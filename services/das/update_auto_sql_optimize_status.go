package das

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

// UpdateAutoSqlOptimizeStatus invokes the das.UpdateAutoSqlOptimizeStatus API synchronously
func (client *Client) UpdateAutoSqlOptimizeStatus(request *UpdateAutoSqlOptimizeStatusRequest) (response *UpdateAutoSqlOptimizeStatusResponse, err error) {
	response = CreateUpdateAutoSqlOptimizeStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAutoSqlOptimizeStatusWithChan invokes the das.UpdateAutoSqlOptimizeStatus API asynchronously
func (client *Client) UpdateAutoSqlOptimizeStatusWithChan(request *UpdateAutoSqlOptimizeStatusRequest) (<-chan *UpdateAutoSqlOptimizeStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateAutoSqlOptimizeStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAutoSqlOptimizeStatus(request)
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

// UpdateAutoSqlOptimizeStatusWithCallback invokes the das.UpdateAutoSqlOptimizeStatus API asynchronously
func (client *Client) UpdateAutoSqlOptimizeStatusWithCallback(request *UpdateAutoSqlOptimizeStatusRequest, callback func(response *UpdateAutoSqlOptimizeStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAutoSqlOptimizeStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateAutoSqlOptimizeStatus(request)
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

// UpdateAutoSqlOptimizeStatusRequest is the request struct for api UpdateAutoSqlOptimizeStatus
type UpdateAutoSqlOptimizeStatusRequest struct {
	*requests.RpcRequest
	Instances string           `position:"Query" name:"Instances"`
	Status    requests.Integer `position:"Query" name:"Status"`
}

// UpdateAutoSqlOptimizeStatusResponse is the response struct for api UpdateAutoSqlOptimizeStatus
type UpdateAutoSqlOptimizeStatusResponse struct {
	*responses.BaseResponse
	Code      string                            `json:"Code" xml:"Code"`
	Message   string                            `json:"Message" xml:"Message"`
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Success   string                            `json:"Success" xml:"Success"`
	Data      DataInUpdateAutoSqlOptimizeStatus `json:"Data" xml:"Data"`
}

// CreateUpdateAutoSqlOptimizeStatusRequest creates a request to invoke UpdateAutoSqlOptimizeStatus API
func CreateUpdateAutoSqlOptimizeStatusRequest() (request *UpdateAutoSqlOptimizeStatusRequest) {
	request = &UpdateAutoSqlOptimizeStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "UpdateAutoSqlOptimizeStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateAutoSqlOptimizeStatusResponse creates a response to parse from UpdateAutoSqlOptimizeStatus response
func CreateUpdateAutoSqlOptimizeStatusResponse() (response *UpdateAutoSqlOptimizeStatusResponse) {
	response = &UpdateAutoSqlOptimizeStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
