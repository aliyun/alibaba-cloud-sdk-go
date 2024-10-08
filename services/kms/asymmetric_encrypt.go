package kms

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

// AsymmetricEncrypt invokes the kms.AsymmetricEncrypt API synchronously
func (client *Client) AsymmetricEncrypt(request *AsymmetricEncryptRequest) (response *AsymmetricEncryptResponse, err error) {
	response = CreateAsymmetricEncryptResponse()
	err = client.DoAction(request, response)
	return
}

// AsymmetricEncryptWithChan invokes the kms.AsymmetricEncrypt API asynchronously
func (client *Client) AsymmetricEncryptWithChan(request *AsymmetricEncryptRequest) (<-chan *AsymmetricEncryptResponse, <-chan error) {
	responseChan := make(chan *AsymmetricEncryptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AsymmetricEncrypt(request)
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

// AsymmetricEncryptWithCallback invokes the kms.AsymmetricEncrypt API asynchronously
func (client *Client) AsymmetricEncryptWithCallback(request *AsymmetricEncryptRequest, callback func(response *AsymmetricEncryptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AsymmetricEncryptResponse
		var err error
		defer close(result)
		response, err = client.AsymmetricEncrypt(request)
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

// AsymmetricEncryptRequest is the request struct for api AsymmetricEncrypt
type AsymmetricEncryptRequest struct {
	*requests.RpcRequest
	DryRun       string `position:"Query" name:"DryRun"`
	KeyVersionId string `position:"Query" name:"KeyVersionId"`
	KeyId        string `position:"Query" name:"KeyId"`
	Plaintext    string `position:"Query" name:"Plaintext"`
	Algorithm    string `position:"Query" name:"Algorithm"`
}

// AsymmetricEncryptResponse is the response struct for api AsymmetricEncrypt
type AsymmetricEncryptResponse struct {
	*responses.BaseResponse
	KeyVersionId   string `json:"KeyVersionId" xml:"KeyVersionId"`
	KeyId          string `json:"KeyId" xml:"KeyId"`
	CiphertextBlob string `json:"CiphertextBlob" xml:"CiphertextBlob"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateAsymmetricEncryptRequest creates a request to invoke AsymmetricEncrypt API
func CreateAsymmetricEncryptRequest() (request *AsymmetricEncryptRequest) {
	request = &AsymmetricEncryptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "AsymmetricEncrypt", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAsymmetricEncryptResponse creates a response to parse from AsymmetricEncrypt response
func CreateAsymmetricEncryptResponse() (response *AsymmetricEncryptResponse) {
	response = &AsymmetricEncryptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
