package gounity

import (
	"encoding/base64"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestURL(t *testing.T) {

	url := "unity.exemple.com"
	uri := "/api/instance"

	t.Log("Given the need to test URL function.")
	{
		t.Logf("\tWhen checking url \"%s\" with uri \"%s\"", url, uri)

		actualResult := URL(url, uri)
		var expectedResult = "https://unity.exemple.com/api/instance"

		if actualResult != expectedResult {
			t.Fatalf("\t\tShould receive %s but got %s %v", expectedResult, actualResult, ballotX)
		}

		t.Logf("\t\tShould receive %s %v", expectedResult, checkMark)

	}
}

func TestEncodeCredentials(t *testing.T) {

	user := "user"
	password := "password"

	t.Log("Given the need to test base64 credential's encoding.")
	{

		t.Logf("\tWhen encoding user \"%s\" with password \"%s\"", user, password)

		actualResult := EncodeCredentials(user, password)
		var expectedResult = base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

		if actualResult != expectedResult {
			t.Fatalf("Expected %s but got %s %v", expectedResult, actualResult, ballotX)
		}

		t.Logf("\t\tShould receive %s %v", expectedResult, checkMark)
	}

}

func TestNewSession(t *testing.T) {
	t.Log("Given the need to test calling NewSession with missing parameters")
	{

		server := ""
		insecure := true
		username := "user"
		password := "password"

		t.Logf("\tWhen parameters are server: \"%s\", insecure: \"%t\", username \"%s\", password: \"%s\"", server, insecure, username, password)
		{
			_, err := NewSession(server, insecure, username, password)
			if err != nil {
				t.Log("\t\tShould receive an error", checkMark, err)
			} else {
				t.Error("\t\tShould receive an error", ballotX, err)
			}
		}

		server = "unity.exemple.com"
		insecure = true
		username = ""
		password = "password"

		t.Logf("\tWhen parameters are server: \"%s\", insecure: \"%t\", username \"%s\", password: \"%s\"", server, insecure, username, password)
		{
			_, err := NewSession(server, insecure, username, password)
			if err != nil {
				t.Log("\t\tShould receive an error", checkMark, err)
			} else {
				t.Error("\t\tShould receive an error", ballotX, err)
			}
		}

		server = "unity.exemple.com"
		insecure = true
		username = "user"
		password = ""

		t.Logf("\tWhen parameters are server: \"%s\", insecure: \"%t\", username \"%s\", password: \"%s\"", server, insecure, username, password)
		{
			_, err := NewSession(server, insecure, username, password)
			if err != nil {
				t.Log("\t\tShould receive an error", checkMark, err)
			} else {
				t.Error("\t\tShould receive an error", ballotX, err)
			}
		}
	}
}
