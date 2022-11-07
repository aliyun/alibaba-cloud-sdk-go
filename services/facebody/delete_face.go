package facebody

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

// DeleteFace invokes the facebody.DeleteFace API synchronously
func (client *Client) DeleteFace(request *DeleteFaceRequest) (response *DeleteFaceResponse, err error) {
	response = CreateDeleteFaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceWithChan invokes the facebody.DeleteFace API asynchronously
func (client *Client) DeleteFaceWithChan(request *DeleteFaceRequest) (<-chan *DeleteFaceResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFace(request)
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

// DeleteFaceWithCallback invokes the facebody.DeleteFace API asynchronously
func (client *Client) DeleteFaceWithCallback(request *DeleteFaceRequest, callback func(response *DeleteFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceResponse
		var err error
		defer close(result)
		response, err = client.DeleteFace(request)
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

// DeleteFaceRequest is the request struct for api DeleteFace
type DeleteFaceRequest struct {
	*requests.RpcRequest
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	OssFile            string           `position:"Query" name:"OssFile"`
	FaceId             string           `position:"Body" name:"FaceId"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
	DbName             string           `position:"Body" name:"DbName"`
}

// DeleteFaceResponse is the response struct for api DeleteFace
type DeleteFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFaceRequest creates a request to invoke DeleteFace API
func CreateDeleteFaceRequest() (request *DeleteFaceRequest) {
	request = &DeleteFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DeleteFace", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFaceResponse creates a response to parse from DeleteFace response
func CreateDeleteFaceResponse() (response *DeleteFaceResponse) {
	response = &DeleteFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
