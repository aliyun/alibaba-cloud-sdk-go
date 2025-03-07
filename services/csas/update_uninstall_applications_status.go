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

// UpdateUninstallApplicationsStatus invokes the csas.UpdateUninstallApplicationsStatus API synchronously
func (client *Client) UpdateUninstallApplicationsStatus(request *UpdateUninstallApplicationsStatusRequest) (response *UpdateUninstallApplicationsStatusResponse, err error) {
	response = CreateUpdateUninstallApplicationsStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUninstallApplicationsStatusWithChan invokes the csas.UpdateUninstallApplicationsStatus API asynchronously
func (client *Client) UpdateUninstallApplicationsStatusWithChan(request *UpdateUninstallApplicationsStatusRequest) (<-chan *UpdateUninstallApplicationsStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateUninstallApplicationsStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUninstallApplicationsStatus(request)
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

// UpdateUninstallApplicationsStatusWithCallback invokes the csas.UpdateUninstallApplicationsStatus API asynchronously
func (client *Client) UpdateUninstallApplicationsStatusWithCallback(request *UpdateUninstallApplicationsStatusRequest, callback func(response *UpdateUninstallApplicationsStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUninstallApplicationsStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateUninstallApplicationsStatus(request)
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

// UpdateUninstallApplicationsStatusRequest is the request struct for api UpdateUninstallApplicationsStatus
type UpdateUninstallApplicationsStatusRequest struct {
	*requests.RpcRequest
	ApplicationIds *[]string `position:"Body" name:"ApplicationIds"  type:"Repeated"`
	SourceIp       string    `position:"Query" name:"SourceIp"`
	Status         string    `position:"Body" name:"Status"`
}

// UpdateUninstallApplicationsStatusResponse is the response struct for api UpdateUninstallApplicationsStatus
type UpdateUninstallApplicationsStatusResponse struct {
	*responses.BaseResponse
	RequestId    string                                    `json:"RequestId" xml:"RequestId"`
	Applications []DataInUpdateUninstallApplicationsStatus `json:"Applications" xml:"Applications"`
}

// CreateUpdateUninstallApplicationsStatusRequest creates a request to invoke UpdateUninstallApplicationsStatus API
func CreateUpdateUninstallApplicationsStatusRequest() (request *UpdateUninstallApplicationsStatusRequest) {
	request = &UpdateUninstallApplicationsStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "UpdateUninstallApplicationsStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateUninstallApplicationsStatusResponse creates a response to parse from UpdateUninstallApplicationsStatus response
func CreateUpdateUninstallApplicationsStatusResponse() (response *UpdateUninstallApplicationsStatusResponse) {
	response = &UpdateUninstallApplicationsStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
