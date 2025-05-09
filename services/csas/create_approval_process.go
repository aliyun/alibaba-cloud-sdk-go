package csas

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateApprovalProcess invokes the csas.CreateApprovalProcess API synchronously
func (client *Client) CreateApprovalProcess(request *CreateApprovalProcessRequest) (response *CreateApprovalProcessResponse, err error) {
	response = CreateCreateApprovalProcessResponse()
	err = client.DoAction(request, response)
	return
}

// CreateApprovalProcessWithChan invokes the csas.CreateApprovalProcess API asynchronously
func (client *Client) CreateApprovalProcessWithChan(request *CreateApprovalProcessRequest) (<-chan *CreateApprovalProcessResponse, <-chan error) {
	responseChan := make(chan *CreateApprovalProcessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApprovalProcess(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateApprovalProcessWithCallback invokes the csas.CreateApprovalProcess API asynchronously
func (client *Client) CreateApprovalProcessWithCallback(request *CreateApprovalProcessRequest, callback func(response *CreateApprovalProcessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateApprovalProcessResponse
		var err error
		defer close(result)
		response, err = client.CreateApprovalProcess(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateApprovalProcessRequest is the request struct for api CreateApprovalProcess
type CreateApprovalProcessRequest struct {
	*requests.RpcRequest
	ProcessNodes *[]*[]string                      `position:"Body" name:"ProcessNodes"  type:"Repeated"`
	MatchSchemas CreateApprovalProcessMatchSchemas `position:"Body" name:"MatchSchemas"  type:"Struct"`
	ProcessName  string                            `position:"Body" name:"ProcessName"`
	Description  string                            `position:"Body" name:"Description"`
	SourceIp     string                            `position:"Query" name:"SourceIp"`
}

// CreateApprovalProcessMatchSchemas is a repeated param struct in CreateApprovalProcessRequest
type CreateApprovalProcessMatchSchemas struct {
	DomainBlacklistSchemaId    string `name:"DomainBlacklistSchemaId"`
	SoftwareBlockSchemaId      string `name:"SoftwareBlockSchemaId"`
	PeripheralBlockSchemaId    string `name:"PeripheralBlockSchemaId"`
	DeviceRegistrationSchemaId string `name:"DeviceRegistrationSchemaId"`
	DlpSendSchemaId            string `name:"DlpSendSchemaId"`
	DomainWhitelistSchemaId    string `name:"DomainWhitelistSchemaId"`
	AppUninstallSchemaId       string `name:"AppUninstallSchemaId"`
}

// CreateApprovalProcessResponse is the response struct for api CreateApprovalProcess
type CreateApprovalProcessResponse struct {
	*responses.BaseResponse
}

// CreateCreateApprovalProcessRequest creates a request to invoke CreateApprovalProcess API
func CreateCreateApprovalProcessRequest() (request *CreateApprovalProcessRequest) {
	request = &CreateApprovalProcessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "CreateApprovalProcess", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateApprovalProcessResponse creates a response to parse from CreateApprovalProcess response
func CreateCreateApprovalProcessResponse() (response *CreateApprovalProcessResponse) {
	response = &CreateApprovalProcessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
