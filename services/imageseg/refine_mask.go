package imageseg

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

// RefineMask invokes the imageseg.RefineMask API synchronously
func (client *Client) RefineMask(request *RefineMaskRequest) (response *RefineMaskResponse, err error) {
	response = CreateRefineMaskResponse()
	err = client.DoAction(request, response)
	return
}

// RefineMaskWithChan invokes the imageseg.RefineMask API asynchronously
func (client *Client) RefineMaskWithChan(request *RefineMaskRequest) (<-chan *RefineMaskResponse, <-chan error) {
	responseChan := make(chan *RefineMaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefineMask(request)
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

// RefineMaskWithCallback invokes the imageseg.RefineMask API asynchronously
func (client *Client) RefineMaskWithCallback(request *RefineMaskRequest, callback func(response *RefineMaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefineMaskResponse
		var err error
		defer close(result)
		response, err = client.RefineMask(request)
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

// RefineMaskRequest is the request struct for api RefineMask
type RefineMaskRequest struct {
	*requests.RpcRequest
	MaskImageURL   string `position:"Body" name:"MaskImageURL"`
	OssFile        string `position:"Query" name:"OssFile"`
	RequestProxyBy string `position:"Query" name:"RequestProxyBy"`
	ImageURL       string `position:"Body" name:"ImageURL"`
}

// RefineMaskResponse is the response struct for api RefineMask
type RefineMaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRefineMaskRequest creates a request to invoke RefineMask API
func CreateRefineMaskRequest() (request *RefineMaskRequest) {
	request = &RefineMaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageseg", "2019-12-30", "RefineMask", "imageseg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRefineMaskResponse creates a response to parse from RefineMask response
func CreateRefineMaskResponse() (response *RefineMaskResponse) {
	response = &RefineMaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
