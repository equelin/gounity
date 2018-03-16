package gounity

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

//Session struct
type Session struct {
	server   string
	insecure bool
	token    string
	http     http.Client
}

// EncodeCredentials purpose is to encode the Username and Password with base64 encoding which is
// required for UNITY
func EncodeCredentials(username string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

// URL purpose is to return the main entry point for the UNITY API
func URL(server string, URI string) string {
	return "https://" + server + URI
}

//NewSession purpose is to build a session object and authenticate to the Unity array
func NewSession(server string, insecure bool, username string, password string) (*Session, error) {
	if server == "" || username == "" || password == "" {
		return nil, errors.New("Missing server (Unity IP/FQDN), username, or password")
	}

	var httpClient http.Client
	cookieJar, _ := cookiejar.New(nil)

	// Create http client
	if insecure {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = http.Client{Transport: tr, Jar: cookieJar}
	} else {
		httpClient = http.Client{Jar: cookieJar}
	}

	// Creates a http Request pointer
	var req *http.Request

	// Creates a GET HTTP Request to the UNITY API
	req, _ = http.NewRequest("GET", URL(server, "/api/types/system/instances"), nil)

	// Add the Authorization header with the value of base64 encoded Username and Password
	// and X-EMC-REST-CLIENT header that is required if using basic authentication
	req.Header.Set("Authorization", "Basic "+EncodeCredentials(username, password))
	req.Header.Set("X-EMC-REST-CLIENT", "true")

	// Execute request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Get the CSRF Token
	token := resp.Header.Get("Emc-Csrf-Token")

	return &Session{server, insecure, token, httpClient}, nil
}

//Request purpose is to send a Rest API request to the Unity array.
func (session *Session) Request(method string, URI string, fields string, filter string, engineering bool, body, resp interface{}) error {

	if method == "" || URI == "" || resp == nil {
		return errors.New("Missing method, URI or response interface")
	}

	// build endpoint URL
	endpoint := URL(session.server, URI)

	// create a http Request pointer
	var req *http.Request

	if body != nil {
		// Parse out body struct into JSON
		bodyBytes, _ := json.Marshal(body)

		// Create new http request
		req, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(bodyBytes))

	} else {

		// Create new http request
		req, _ = http.NewRequest(method, endpoint, nil)
	}

	// In case of POST or DELETE method, add the CSRF token in the request header
	if method == "POST" || method == "DELETE" {
		req.Header.Set("Emc-Csrf-Token", session.token)
	}

	// Add the mandatory headers to the request
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-EMC-REST-CLIENT", "true")

	if method != "DELETE" {
		// Create an URL query object
		a := req.URL.Query()

		if fields != "" {
			a.Add("fields", fields)
		}

		if filter != "" {
			a.Add("filter", filter)
		}

		if engineering == true {
			a.Add("visibility", "Engineering")
		}

		a.Add("compact", "true")

		req.URL.RawQuery = a.Encode()
	}

	// Perform request
	httpResp, err := session.http.Do(req)
	if err != nil {
		return err
	}

	// Cleanup Response
	defer httpResp.Body.Close()

	switch {
	case httpResp.StatusCode == 200 || httpResp.StatusCode == 201 || httpResp.StatusCode == 202:
		// Decode JSON of response into our interface defined for the specific request sent
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}

		// Unmarshal the body into a struct
		bodyByte := json.Unmarshal(body, resp)

		return bodyByte
	case httpResp.StatusCode == 204:
		return nil

	case httpResp.StatusCode == 422:
		return fmt.Errorf("HTTP status codes: %d, detail: %v", httpResp.StatusCode, httpResp.Body)

	default:
		return fmt.Errorf("HTTP status codes: %d", httpResp.StatusCode)
	}
}

//CloseSession purpose is to end the session with the Unity array.
func (session *Session) CloseSession() (err error) {
	err = session.Request("POST", "/api/types/loginSessionInfo/action/logout", "", "", false, nil, nil)
	return err
}
