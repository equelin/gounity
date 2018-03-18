package gounity

import (
	"time"
)

// Type declarations

//DNSServer struct
type DNSServer struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Links   []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Entries []struct {
		Content struct {
			Origin    int      `json:"origin"`
			ID        string   `json:"id"`
			Addresses []string `json:"addresses"`
		} `json:"content"`
	} `json:"entries"`
}

//MetricRealTimeQueryResponse struct
type MetricRealTimeQueryResponse struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Links   []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Content struct {
		Interval       int       `json:"interval"`
		ID             int       `json:"id"`
		Paths          []string  `json:"paths"`
		MaximumSamples int       `json:"maximumSamples"`
		Expiration     time.Time `json:"expiration"`
	} `json:"content"`
}

//MetricRealTimeQueryResult struct
type MetricRealTimeQueryResult struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Links   []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Entries []struct {
		Base    string    `json:"@base"`
		Updated time.Time `json:"updated"`
		Links   []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		Content struct {
			QueryID   int         `json:"queryId"`
			Path      string      `json:"path"`
			Timestamp time.Time   `json:"timestamp"`
			Values    interface{} `json:"values"`
		} `json:"content"`
	} `json:"entries"`
}

//BasicSystemInfo struct
type BasicSystemInfo struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Links   []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Entries []struct {
		Base    string    `json:"@base"`
		Updated time.Time `json:"updated"`
		Links   []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		Content struct {
			ID                 string `json:"id"`
			Model              string `json:"model"`
			Name               string `json:"name"`
			SoftwareVersion    string `json:"softwareVersion"`
			APIVersion         string `json:"apiVersion"`
			EarliestAPIVersion string `json:"earliestApiVersion"`
		} `json:"content"`
	} `json:"entries"`
}
