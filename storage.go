package gounity

import types "github.com/equelin/gounity/types/v1"

//GetPool purpose is to get information about pool in the storage system. .
func (session *Session) GetPool() (resp *types.Pool, err error) {

	fields := "id,name,sizeFree,sizeTotal,sizeUsed,sizeSubscribed"

	err = session.Request("GET", "/api/types/pool/instances", fields, "", false, nil, &resp)
	return resp, err
}

//GetStorageResource purpose is to get information about storage resources in the storage system.
func (session *Session) GetStorageResource() (resp *types.StorageResource, err error) {

	fields := "id,name,sizeAllocated,sizeTotal,sizeUsed"

	err = session.Request("GET", "/api/types/storageResource/instances", fields, "", false, nil, &resp)
	return resp, err
}
