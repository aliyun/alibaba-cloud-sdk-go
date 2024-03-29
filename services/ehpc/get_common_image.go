package ehpc

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

// GetCommonImage invokes the ehpc.GetCommonImage API synchronously
func (client *Client) GetCommonImage(request *GetCommonImageRequest) (response *GetCommonImageResponse, err error) {
	response = CreateGetCommonImageResponse()
	err = client.DoAction(request, response)
	return
}

// GetCommonImageWithChan invokes the ehpc.GetCommonImage API asynchronously
func (client *Client) GetCommonImageWithChan(request *GetCommonImageRequest) (<-chan *GetCommonImageResponse, <-chan error) {
	responseChan := make(chan *GetCommonImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCommonImage(request)
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

// GetCommonImageWithCallback invokes the ehpc.GetCommonImage API asynchronously
func (client *Client) GetCommonImageWithCallback(request *GetCommonImageRequest, callback func(response *GetCommonImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCommonImageResponse
		var err error
		defer close(result)
		response, err = client.GetCommonImage(request)
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

// GetCommonImageRequest is the request struct for api GetCommonImage
type GetCommonImageRequest struct {
	*requests.RpcRequest
	ContainType string `position:"Query" name:"ContainType"`
	ClusterId   string `position:"Query" name:"ClusterId"`
	ImageName   string `position:"Query" name:"ImageName"`
}

// GetCommonImageResponse is the response struct for api GetCommonImage
type GetCommonImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetCommonImageRequest creates a request to invoke GetCommonImage API
func CreateGetCommonImageRequest() (request *GetCommonImageRequest) {
	request = &GetCommonImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "GetCommonImage", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetCommonImageResponse creates a response to parse from GetCommonImage response
func CreateGetCommonImageResponse() (response *GetCommonImageResponse) {
	response = &GetCommonImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
