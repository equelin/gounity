package gounity

import (
	"fmt"
	"strconv"

	types "github.com/equelin/gounity/types/v1"
)

//NewMetricRealTimeQuery purpose is to create a real time query.
func (session *Session) NewMetricRealTimeQuery(paths []string, interval uint32) (resp *types.MetricRealTimeQueryResponse, err error) {

	type paramNewMetricRealTimeQuery struct {
		Paths    []string `json:"paths"`
		Interval uint32   `json:"interval"`
	}

	jsonData := paramNewMetricRealTimeQuery{
		Paths:    paths,
		Interval: interval,
	}

	err = session.Request("POST", "/api/types/metricRealTimeQuery/instances", "", "", true, jsonData, &resp)

	return resp, err
}

//DeleteMetricRealTimeQuery purpose is to delete a real time query.
func (session *Session) DeleteMetricRealTimeQuery(queryID int) error {

	URI := fmt.Sprintf("/api/instances/metricRealTimeQuery/%s", strconv.Itoa(queryID))

	fmt.Println(URI)

	err := session.Request("DELETE", URI, "", "", false, nil, "")

	return err
}

//GetMetricRealTimeQueryResult tototo
func (session *Session) GetMetricRealTimeQueryResult(queryID int) (resp *types.MetricRealTimeQueryResult, err error) {

	filter := fmt.Sprintf("queryId eq %s", strconv.Itoa(queryID))

	fields := "queryId,path,timestamp,values"

	URI := "/api/types/metricQueryResult/instances"

	err = session.Request("GET", URI, fields, filter, true, nil, &resp)

	return resp, err
}

//GetmetricValue Historical values for requested metrics
func (session *Session) GetmetricValue(path string) (resp *types.MetricValueResponse, err error) {

	filter := fmt.Sprintf("path eq \"%s\"", path)

	//fields := "queryId,path,timestamp,values"

	URI := "/api/types/metricValue/instances"

	err = session.Request("GET", URI, "", filter, true, nil, &resp)

	return resp, err
}

//GetkpiValue Historical values for requested metrics
func (session *Session) GetkpiValue(path string) (resp *types.KpiValueResponse, err error) {

	filter := fmt.Sprintf("path eq \"%s\"", path)

	//fields := "queryId,path,timestamp,values"

	URI := "/api/types/kpiValue/instances"

	err = session.Request("GET", URI, "", filter, true, nil, &resp)

	return resp, err
}
