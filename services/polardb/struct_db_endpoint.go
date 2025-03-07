package polardb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DBEndpoint is a nested struct in polardb response
type DBEndpoint struct {
	NodeWithRoles         string    `json:"NodeWithRoles" xml:"NodeWithRoles"`
	Nodes                 string    `json:"Nodes" xml:"Nodes"`
	ReadWriteMode         string    `json:"ReadWriteMode" xml:"ReadWriteMode"`
	DBEndpointId          string    `json:"DBEndpointId" xml:"DBEndpointId"`
	EndpointConfig        string    `json:"EndpointConfig" xml:"EndpointConfig"`
	DBEndpointDescription string    `json:"DBEndpointDescription" xml:"DBEndpointDescription"`
	EndpointType          string    `json:"EndpointType" xml:"EndpointType"`
	AutoAddNewNodes       string    `json:"AutoAddNewNodes" xml:"AutoAddNewNodes"`
	DBClusterId           string    `json:"DBClusterId" xml:"DBClusterId"`
	SccMode               string    `json:"SccMode" xml:"SccMode"`
	PolarSccTimeoutAction string    `json:"PolarSccTimeoutAction" xml:"PolarSccTimeoutAction"`
	PolarSccWaitTimeout   string    `json:"PolarSccWaitTimeout" xml:"PolarSccWaitTimeout"`
	AddressItems          []Address `json:"AddressItems" xml:"AddressItems"`
}
