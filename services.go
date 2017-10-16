package gounity

import types "github.com/equelin/gounity/types/v1"

//GetDNSServer purpose is to get the DNS configuration of the Unity array.
func (session *Session) GetDNSServer() (resp *types.DNSServer, err error) {

	fields := "id,domain,origin,addresses"

	err = session.Request("GET", "/api/types/dnsServer/instances", fields, "", nil, &resp)
	return resp, err
}
