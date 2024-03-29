package vod

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

// DeleteVodSpecificConfig invokes the vod.DeleteVodSpecificConfig API synchronously
func (client *Client) DeleteVodSpecificConfig(request *DeleteVodSpecificConfigRequest) (response *DeleteVodSpecificConfigResponse, err error) {
	response = CreateDeleteVodSpecificConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVodSpecificConfigWithChan invokes the vod.DeleteVodSpecificConfig API asynchronously
func (client *Client) DeleteVodSpecificConfigWithChan(request *DeleteVodSpecificConfigRequest) (<-chan *DeleteVodSpecificConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteVodSpecificConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVodSpecificConfig(request)
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

// DeleteVodSpecificConfigWithCallback invokes the vod.DeleteVodSpecificConfig API asynchronously
func (client *Client) DeleteVodSpecificConfigWithCallback(request *DeleteVodSpecificConfigRequest, callback func(response *DeleteVodSpecificConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVodSpecificConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteVodSpecificConfig(request)
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

// DeleteVodSpecificConfigRequest is the request struct for api DeleteVodSpecificConfig
type DeleteVodSpecificConfigRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Env           string           `position:"Query" name:"Env"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	ConfigId      string           `position:"Query" name:"ConfigId"`
}

// DeleteVodSpecificConfigResponse is the response struct for api DeleteVodSpecificConfig
type DeleteVodSpecificConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVodSpecificConfigRequest creates a request to invoke DeleteVodSpecificConfig API
func CreateDeleteVodSpecificConfigRequest() (request *DeleteVodSpecificConfigRequest) {
	request = &DeleteVodSpecificConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteVodSpecificConfig", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVodSpecificConfigResponse creates a response to parse from DeleteVodSpecificConfig response
func CreateDeleteVodSpecificConfigResponse() (response *DeleteVodSpecificConfigResponse) {
	response = &DeleteVodSpecificConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
