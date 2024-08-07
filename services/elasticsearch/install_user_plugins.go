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

// InstallUserPlugins invokes the elasticsearch.InstallUserPlugins API synchronously
func (client *Client) InstallUserPlugins(request *InstallUserPluginsRequest) (response *InstallUserPluginsResponse, err error) {
	response = CreateInstallUserPluginsResponse()
	err = client.DoAction(request, response)
	return
}

// InstallUserPluginsWithChan invokes the elasticsearch.InstallUserPlugins API asynchronously
func (client *Client) InstallUserPluginsWithChan(request *InstallUserPluginsRequest) (<-chan *InstallUserPluginsResponse, <-chan error) {
	responseChan := make(chan *InstallUserPluginsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallUserPlugins(request)
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

// InstallUserPluginsWithCallback invokes the elasticsearch.InstallUserPlugins API asynchronously
func (client *Client) InstallUserPluginsWithCallback(request *InstallUserPluginsRequest, callback func(response *InstallUserPluginsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallUserPluginsResponse
		var err error
		defer close(result)
		response, err = client.InstallUserPlugins(request)
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

// InstallUserPluginsRequest is the request struct for api InstallUserPlugins
type InstallUserPluginsRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	Body       string `position:"Body" name:"body"`
}

// InstallUserPluginsResponse is the response struct for api InstallUserPlugins
type InstallUserPluginsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Result    []string `json:"Result" xml:"Result"`
}

// CreateInstallUserPluginsRequest creates a request to invoke InstallUserPlugins API
func CreateInstallUserPluginsRequest() (request *InstallUserPluginsRequest) {
	request = &InstallUserPluginsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "InstallUserPlugins", "/openapi/instances/[InstanceId]/plugins/user/actions/install", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInstallUserPluginsResponse creates a response to parse from InstallUserPlugins response
func CreateInstallUserPluginsResponse() (response *InstallUserPluginsResponse) {
	response = &InstallUserPluginsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
