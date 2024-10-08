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

// Decrypt invokes the kms.Decrypt API synchronously
func (client *Client) Decrypt(request *DecryptRequest) (response *DecryptResponse, err error) {
	response = CreateDecryptResponse()
	err = client.DoAction(request, response)
	return
}

// DecryptWithChan invokes the kms.Decrypt API asynchronously
func (client *Client) DecryptWithChan(request *DecryptRequest) (<-chan *DecryptResponse, <-chan error) {
	responseChan := make(chan *DecryptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Decrypt(request)
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

// DecryptWithCallback invokes the kms.Decrypt API asynchronously
func (client *Client) DecryptWithCallback(request *DecryptRequest, callback func(response *DecryptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DecryptResponse
		var err error
		defer close(result)
		response, err = client.Decrypt(request)
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

// DecryptRequest is the request struct for api Decrypt
type DecryptRequest struct {
	*requests.RpcRequest
	DryRun            string `position:"Query" name:"DryRun"`
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
	CiphertextBlob    string `position:"Query" name:"CiphertextBlob"`
}

// DecryptResponse is the response struct for api Decrypt
type DecryptResponse struct {
	*responses.BaseResponse
	KeyVersionId string `json:"KeyVersionId" xml:"KeyVersionId"`
	KeyId        string `json:"KeyId" xml:"KeyId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Plaintext    string `json:"Plaintext" xml:"Plaintext"`
}

// CreateDecryptRequest creates a request to invoke Decrypt API
func CreateDecryptRequest() (request *DecryptRequest) {
	request = &DecryptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "Decrypt", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDecryptResponse creates a response to parse from Decrypt response
func CreateDecryptResponse() (response *DecryptResponse) {
	response = &DecryptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
