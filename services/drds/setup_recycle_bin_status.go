package drds

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

// SetupRecycleBinStatus invokes the drds.SetupRecycleBinStatus API synchronously
func (client *Client) SetupRecycleBinStatus(request *SetupRecycleBinStatusRequest) (response *SetupRecycleBinStatusResponse, err error) {
	response = CreateSetupRecycleBinStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SetupRecycleBinStatusWithChan invokes the drds.SetupRecycleBinStatus API asynchronously
func (client *Client) SetupRecycleBinStatusWithChan(request *SetupRecycleBinStatusRequest) (<-chan *SetupRecycleBinStatusResponse, <-chan error) {
	responseChan := make(chan *SetupRecycleBinStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetupRecycleBinStatus(request)
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

// SetupRecycleBinStatusWithCallback invokes the drds.SetupRecycleBinStatus API asynchronously
func (client *Client) SetupRecycleBinStatusWithCallback(request *SetupRecycleBinStatusRequest, callback func(response *SetupRecycleBinStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetupRecycleBinStatusResponse
		var err error
		defer close(result)
		response, err = client.SetupRecycleBinStatus(request)
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

// SetupRecycleBinStatusRequest is the request struct for api SetupRecycleBinStatus
type SetupRecycleBinStatusRequest struct {
	*requests.RpcRequest
	StatusAction   string `position:"Query" name:"StatusAction"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// SetupRecycleBinStatusResponse is the response struct for api SetupRecycleBinStatus
type SetupRecycleBinStatusResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSetupRecycleBinStatusRequest creates a request to invoke SetupRecycleBinStatus API
func CreateSetupRecycleBinStatusRequest() (request *SetupRecycleBinStatusRequest) {
	request = &SetupRecycleBinStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "SetupRecycleBinStatus", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetupRecycleBinStatusResponse creates a response to parse from SetupRecycleBinStatus response
func CreateSetupRecycleBinStatusResponse() (response *SetupRecycleBinStatusResponse) {
	response = &SetupRecycleBinStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
