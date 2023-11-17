package ens

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

// GetBucketLifecycle invokes the ens.GetBucketLifecycle API synchronously
func (client *Client) GetBucketLifecycle(request *GetBucketLifecycleRequest) (response *GetBucketLifecycleResponse, err error) {
	response = CreateGetBucketLifecycleResponse()
	err = client.DoAction(request, response)
	return
}

// GetBucketLifecycleWithChan invokes the ens.GetBucketLifecycle API asynchronously
func (client *Client) GetBucketLifecycleWithChan(request *GetBucketLifecycleRequest) (<-chan *GetBucketLifecycleResponse, <-chan error) {
	responseChan := make(chan *GetBucketLifecycleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBucketLifecycle(request)
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

// GetBucketLifecycleWithCallback invokes the ens.GetBucketLifecycle API asynchronously
func (client *Client) GetBucketLifecycleWithCallback(request *GetBucketLifecycleRequest, callback func(response *GetBucketLifecycleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBucketLifecycleResponse
		var err error
		defer close(result)
		response, err = client.GetBucketLifecycle(request)
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

// GetBucketLifecycleRequest is the request struct for api GetBucketLifecycle
type GetBucketLifecycleRequest struct {
	*requests.RpcRequest
	BucketName string `position:"Query" name:"BucketName"`
	RuleId     string `position:"Query" name:"RuleId"`
}

// GetBucketLifecycleResponse is the response struct for api GetBucketLifecycle
type GetBucketLifecycleResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Rule      []RuleItem `json:"Rule" xml:"Rule"`
}

// CreateGetBucketLifecycleRequest creates a request to invoke GetBucketLifecycle API
func CreateGetBucketLifecycleRequest() (request *GetBucketLifecycleRequest) {
	request = &GetBucketLifecycleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "GetBucketLifecycle", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetBucketLifecycleResponse creates a response to parse from GetBucketLifecycle response
func CreateGetBucketLifecycleResponse() (response *GetBucketLifecycleResponse) {
	response = &GetBucketLifecycleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
