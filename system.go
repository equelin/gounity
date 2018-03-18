package gounity

import types "github.com/equelin/gounity/types/v1"

//GetbasicSystemInfo Provides access to system model, system name, software version, and API version information.
func (session *Session) GetbasicSystemInfo() (resp *types.BasicSystemInfo, err error) {

	fields := "id,model,name,softwareVersion,apiVersion,earliestApiVersion"

	err = session.Request("GET", "/api/types/basicSystemInfo/instances", fields, "", false, nil, &resp)
	return resp, err
}
