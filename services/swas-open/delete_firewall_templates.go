package swas_open

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

// DeleteFirewallTemplates invokes the swas_open.DeleteFirewallTemplates API synchronously
func (client *Client) DeleteFirewallTemplates(request *DeleteFirewallTemplatesRequest) (response *DeleteFirewallTemplatesResponse, err error) {
	response = CreateDeleteFirewallTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFirewallTemplatesWithChan invokes the swas_open.DeleteFirewallTemplates API asynchronously
func (client *Client) DeleteFirewallTemplatesWithChan(request *DeleteFirewallTemplatesRequest) (<-chan *DeleteFirewallTemplatesResponse, <-chan error) {
	responseChan := make(chan *DeleteFirewallTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFirewallTemplates(request)
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

// DeleteFirewallTemplatesWithCallback invokes the swas_open.DeleteFirewallTemplates API asynchronously
func (client *Client) DeleteFirewallTemplatesWithCallback(request *DeleteFirewallTemplatesRequest, callback func(response *DeleteFirewallTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFirewallTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DeleteFirewallTemplates(request)
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

// DeleteFirewallTemplatesRequest is the request struct for api DeleteFirewallTemplates
type DeleteFirewallTemplatesRequest struct {
	*requests.RpcRequest
	FirewallTemplateId *[]string `position:"Query" name:"FirewallTemplateId"  type:"Repeated"`
	ClientToken        string    `position:"Query" name:"ClientToken"`
	InstanceId         string    `position:"Query" name:"InstanceId"`
}

// DeleteFirewallTemplatesResponse is the response struct for api DeleteFirewallTemplates
type DeleteFirewallTemplatesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFirewallTemplatesRequest creates a request to invoke DeleteFirewallTemplates API
func CreateDeleteFirewallTemplatesRequest() (request *DeleteFirewallTemplatesRequest) {
	request = &DeleteFirewallTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DeleteFirewallTemplates", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFirewallTemplatesResponse creates a response to parse from DeleteFirewallTemplates response
func CreateDeleteFirewallTemplatesResponse() (response *DeleteFirewallTemplatesResponse) {
	response = &DeleteFirewallTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
