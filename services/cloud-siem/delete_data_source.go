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

// DeleteDataSource invokes the cloud_siem.DeleteDataSource API synchronously
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (response *DeleteDataSourceResponse, err error) {
	response = CreateDeleteDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDataSourceWithChan invokes the cloud_siem.DeleteDataSource API asynchronously
func (client *Client) DeleteDataSourceWithChan(request *DeleteDataSourceRequest) (<-chan *DeleteDataSourceResponse, <-chan error) {
	responseChan := make(chan *DeleteDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDataSource(request)
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

// DeleteDataSourceWithCallback invokes the cloud_siem.DeleteDataSource API asynchronously
func (client *Client) DeleteDataSourceWithCallback(request *DeleteDataSourceRequest, callback func(response *DeleteDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDataSourceResponse
		var err error
		defer close(result)
		response, err = client.DeleteDataSource(request)
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

// DeleteDataSourceRequest is the request struct for api DeleteDataSource
type DeleteDataSourceRequest struct {
	*requests.RpcRequest
	CloudCode            string `position:"Body" name:"CloudCode"`
	AccountId            string `position:"Body" name:"AccountId"`
	DataSourceInstanceId string `position:"Body" name:"DataSourceInstanceId"`
}

// DeleteDataSourceResponse is the response struct for api DeleteDataSource
type DeleteDataSourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteDataSourceRequest creates a request to invoke DeleteDataSource API
func CreateDeleteDataSourceRequest() (request *DeleteDataSourceRequest) {
	request = &DeleteDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DeleteDataSource", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDataSourceResponse creates a response to parse from DeleteDataSource response
func CreateDeleteDataSourceResponse() (response *DeleteDataSourceResponse) {
	response = &DeleteDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
