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

// DescribeAlertsWithEntity invokes the cloud_siem.DescribeAlertsWithEntity API synchronously
func (client *Client) DescribeAlertsWithEntity(request *DescribeAlertsWithEntityRequest) (response *DescribeAlertsWithEntityResponse, err error) {
	response = CreateDescribeAlertsWithEntityResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlertsWithEntityWithChan invokes the cloud_siem.DescribeAlertsWithEntity API asynchronously
func (client *Client) DescribeAlertsWithEntityWithChan(request *DescribeAlertsWithEntityRequest) (<-chan *DescribeAlertsWithEntityResponse, <-chan error) {
	responseChan := make(chan *DescribeAlertsWithEntityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlertsWithEntity(request)
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

// DescribeAlertsWithEntityWithCallback invokes the cloud_siem.DescribeAlertsWithEntity API asynchronously
func (client *Client) DescribeAlertsWithEntityWithCallback(request *DescribeAlertsWithEntityRequest, callback func(response *DescribeAlertsWithEntityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlertsWithEntityResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlertsWithEntity(request)
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

// DescribeAlertsWithEntityRequest is the request struct for api DescribeAlertsWithEntity
type DescribeAlertsWithEntityRequest struct {
	*requests.RpcRequest
	EntityId     requests.Integer `position:"Body" name:"EntityId"`
	PageSize     requests.Integer `position:"Body" name:"PageSize"`
	SophonTaskId string           `position:"Body" name:"SophonTaskId"`
	CurrentPage  requests.Integer `position:"Body" name:"CurrentPage"`
	IncidentUuid string           `position:"Body" name:"IncidentUuid"`
}

// DescribeAlertsWithEntityResponse is the response struct for api DescribeAlertsWithEntity
type DescribeAlertsWithEntityResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAlertsWithEntityRequest creates a request to invoke DescribeAlertsWithEntity API
func CreateDescribeAlertsWithEntityRequest() (request *DescribeAlertsWithEntityRequest) {
	request = &DescribeAlertsWithEntityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeAlertsWithEntity", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAlertsWithEntityResponse creates a response to parse from DescribeAlertsWithEntity response
func CreateDescribeAlertsWithEntityResponse() (response *DescribeAlertsWithEntityResponse) {
	response = &DescribeAlertsWithEntityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
