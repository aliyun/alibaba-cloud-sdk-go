package elasticsearch

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

// ListKibanaPlugins invokes the elasticsearch.ListKibanaPlugins API synchronously
func (client *Client) ListKibanaPlugins(request *ListKibanaPluginsRequest) (response *ListKibanaPluginsResponse, err error) {
	response = CreateListKibanaPluginsResponse()
	err = client.DoAction(request, response)
	return
}

// ListKibanaPluginsWithChan invokes the elasticsearch.ListKibanaPlugins API asynchronously
func (client *Client) ListKibanaPluginsWithChan(request *ListKibanaPluginsRequest) (<-chan *ListKibanaPluginsResponse, <-chan error) {
	responseChan := make(chan *ListKibanaPluginsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListKibanaPlugins(request)
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

// ListKibanaPluginsWithCallback invokes the elasticsearch.ListKibanaPlugins API asynchronously
func (client *Client) ListKibanaPluginsWithCallback(request *ListKibanaPluginsRequest, callback func(response *ListKibanaPluginsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListKibanaPluginsResponse
		var err error
		defer close(result)
		response, err = client.ListKibanaPlugins(request)
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

// ListKibanaPluginsRequest is the request struct for api ListKibanaPlugins
type ListKibanaPluginsRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	Size       requests.Integer `position:"Query" name:"size"`
	Page       string           `position:"Query" name:"page"`
}

// ListKibanaPluginsResponse is the response struct for api ListKibanaPlugins
type ListKibanaPluginsResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Headers   HeadersInListKibanaPlugins `json:"Headers" xml:"Headers"`
	Result    []PluginItem               `json:"Result" xml:"Result"`
}

// CreateListKibanaPluginsRequest creates a request to invoke ListKibanaPlugins API
func CreateListKibanaPluginsRequest() (request *ListKibanaPluginsRequest) {
	request = &ListKibanaPluginsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListKibanaPlugins", "/openapi/instances/[InstanceId]/kibana-plugins", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListKibanaPluginsResponse creates a response to parse from ListKibanaPlugins response
func CreateListKibanaPluginsResponse() (response *ListKibanaPluginsResponse) {
	response = &ListKibanaPluginsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
