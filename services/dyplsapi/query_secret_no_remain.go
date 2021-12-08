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

// QuerySecretNoRemain invokes the dyplsapi.QuerySecretNoRemain API synchronously
func (client *Client) QuerySecretNoRemain(request *QuerySecretNoRemainRequest) (response *QuerySecretNoRemainResponse, err error) {
	response = CreateQuerySecretNoRemainResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySecretNoRemainWithChan invokes the dyplsapi.QuerySecretNoRemain API asynchronously
func (client *Client) QuerySecretNoRemainWithChan(request *QuerySecretNoRemainRequest) (<-chan *QuerySecretNoRemainResponse, <-chan error) {
	responseChan := make(chan *QuerySecretNoRemainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySecretNoRemain(request)
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

// QuerySecretNoRemainWithCallback invokes the dyplsapi.QuerySecretNoRemain API asynchronously
func (client *Client) QuerySecretNoRemainWithCallback(request *QuerySecretNoRemainRequest, callback func(response *QuerySecretNoRemainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySecretNoRemainResponse
		var err error
		defer close(result)
		response, err = client.QuerySecretNoRemain(request)
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

// QuerySecretNoRemainRequest is the request struct for api QuerySecretNoRemain
type QuerySecretNoRemainRequest struct {
	*requests.RpcRequest
	SpecId               requests.Integer `position:"Query" name:"SpecId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	City                 string           `position:"Query" name:"City"`
	SecretNo             string           `position:"Query" name:"SecretNo"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QuerySecretNoRemainResponse is the response struct for api QuerySecretNoRemain
type QuerySecretNoRemainResponse struct {
	*responses.BaseResponse
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	SecretRemainDTO SecretRemainDTO `json:"SecretRemainDTO" xml:"SecretRemainDTO"`
}

// CreateQuerySecretNoRemainRequest creates a request to invoke QuerySecretNoRemain API
func CreateQuerySecretNoRemainRequest() (request *QuerySecretNoRemainRequest) {
	request = &QuerySecretNoRemainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "QuerySecretNoRemain", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySecretNoRemainResponse creates a response to parse from QuerySecretNoRemain response
func CreateQuerySecretNoRemainResponse() (response *QuerySecretNoRemainResponse) {
	response = &QuerySecretNoRemainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
