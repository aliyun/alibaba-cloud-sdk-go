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

// DeleteVpcEndpoint invokes the elasticsearch.DeleteVpcEndpoint API synchronously
func (client *Client) DeleteVpcEndpoint(request *DeleteVpcEndpointRequest) (response *DeleteVpcEndpointResponse, err error) {
	response = CreateDeleteVpcEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVpcEndpointWithChan invokes the elasticsearch.DeleteVpcEndpoint API asynchronously
func (client *Client) DeleteVpcEndpointWithChan(request *DeleteVpcEndpointRequest) (<-chan *DeleteVpcEndpointResponse, <-chan error) {
	responseChan := make(chan *DeleteVpcEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVpcEndpoint(request)
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

// DeleteVpcEndpointWithCallback invokes the elasticsearch.DeleteVpcEndpoint API asynchronously
func (client *Client) DeleteVpcEndpointWithCallback(request *DeleteVpcEndpointRequest, callback func(response *DeleteVpcEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVpcEndpointResponse
		var err error
		defer close(result)
		response, err = client.DeleteVpcEndpoint(request)
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

// DeleteVpcEndpointRequest is the request struct for api DeleteVpcEndpoint
type DeleteVpcEndpointRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	EndpointId  string `position:"Path" name:"EndpointId"`
}

// DeleteVpcEndpointResponse is the response struct for api DeleteVpcEndpoint
type DeleteVpcEndpointResponse struct {
	*responses.BaseResponse
	Result    bool   `json:"Result" xml:"Result"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVpcEndpointRequest creates a request to invoke DeleteVpcEndpoint API
func CreateDeleteVpcEndpointRequest() (request *DeleteVpcEndpointRequest) {
	request = &DeleteVpcEndpointRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DeleteVpcEndpoint", "/openapi/instances/[InstanceId]/vpc-endpoints/[EndpointId]", "elasticsearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteVpcEndpointResponse creates a response to parse from DeleteVpcEndpoint response
func CreateDeleteVpcEndpointResponse() (response *DeleteVpcEndpointResponse) {
	response = &DeleteVpcEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
