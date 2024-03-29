package cloud_siem

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

// EnableAccessForCloudSiem invokes the cloud_siem.EnableAccessForCloudSiem API synchronously
func (client *Client) EnableAccessForCloudSiem(request *EnableAccessForCloudSiemRequest) (response *EnableAccessForCloudSiemResponse, err error) {
	response = CreateEnableAccessForCloudSiemResponse()
	err = client.DoAction(request, response)
	return
}

// EnableAccessForCloudSiemWithChan invokes the cloud_siem.EnableAccessForCloudSiem API asynchronously
func (client *Client) EnableAccessForCloudSiemWithChan(request *EnableAccessForCloudSiemRequest) (<-chan *EnableAccessForCloudSiemResponse, <-chan error) {
	responseChan := make(chan *EnableAccessForCloudSiemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableAccessForCloudSiem(request)
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

// EnableAccessForCloudSiemWithCallback invokes the cloud_siem.EnableAccessForCloudSiem API asynchronously
func (client *Client) EnableAccessForCloudSiemWithCallback(request *EnableAccessForCloudSiemRequest, callback func(response *EnableAccessForCloudSiemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableAccessForCloudSiemResponse
		var err error
		defer close(result)
		response, err = client.EnableAccessForCloudSiem(request)
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

// EnableAccessForCloudSiemRequest is the request struct for api EnableAccessForCloudSiem
type EnableAccessForCloudSiemRequest struct {
	*requests.RpcRequest
}

// EnableAccessForCloudSiemResponse is the response struct for api EnableAccessForCloudSiem
type EnableAccessForCloudSiemResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableAccessForCloudSiemRequest creates a request to invoke EnableAccessForCloudSiem API
func CreateEnableAccessForCloudSiemRequest() (request *EnableAccessForCloudSiemRequest) {
	request = &EnableAccessForCloudSiemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "EnableAccessForCloudSiem", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableAccessForCloudSiemResponse creates a response to parse from EnableAccessForCloudSiem response
func CreateEnableAccessForCloudSiemResponse() (response *EnableAccessForCloudSiemResponse) {
	response = &EnableAccessForCloudSiemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
