package dyplsapi

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

// GetFaceVerify invokes the dyplsapi.GetFaceVerify API synchronously
func (client *Client) GetFaceVerify(request *GetFaceVerifyRequest) (response *GetFaceVerifyResponse, err error) {
	response = CreateGetFaceVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// GetFaceVerifyWithChan invokes the dyplsapi.GetFaceVerify API asynchronously
func (client *Client) GetFaceVerifyWithChan(request *GetFaceVerifyRequest) (<-chan *GetFaceVerifyResponse, <-chan error) {
	responseChan := make(chan *GetFaceVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFaceVerify(request)
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

// GetFaceVerifyWithCallback invokes the dyplsapi.GetFaceVerify API asynchronously
func (client *Client) GetFaceVerifyWithCallback(request *GetFaceVerifyRequest, callback func(response *GetFaceVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFaceVerifyResponse
		var err error
		defer close(result)
		response, err = client.GetFaceVerify(request)
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

// GetFaceVerifyRequest is the request struct for api GetFaceVerify
type GetFaceVerifyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	VerifyToken          string           `position:"Query" name:"VerifyToken"`
	ProdCode             string           `position:"Query" name:"ProdCode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// GetFaceVerifyResponse is the response struct for api GetFaceVerify
type GetFaceVerifyResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetFaceVerifyRequest creates a request to invoke GetFaceVerify API
func CreateGetFaceVerifyRequest() (request *GetFaceVerifyRequest) {
	request = &GetFaceVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "GetFaceVerify", "", "")
	request.Method = requests.POST
	return
}

// CreateGetFaceVerifyResponse creates a response to parse from GetFaceVerify response
func CreateGetFaceVerifyResponse() (response *GetFaceVerifyResponse) {
	response = &GetFaceVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
