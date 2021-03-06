package live

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

// AddDRMCertificate invokes the live.AddDRMCertificate API synchronously
func (client *Client) AddDRMCertificate(request *AddDRMCertificateRequest) (response *AddDRMCertificateResponse, err error) {
	response = CreateAddDRMCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// AddDRMCertificateWithChan invokes the live.AddDRMCertificate API asynchronously
func (client *Client) AddDRMCertificateWithChan(request *AddDRMCertificateRequest) (<-chan *AddDRMCertificateResponse, <-chan error) {
	responseChan := make(chan *AddDRMCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDRMCertificate(request)
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

// AddDRMCertificateWithCallback invokes the live.AddDRMCertificate API asynchronously
func (client *Client) AddDRMCertificateWithCallback(request *AddDRMCertificateRequest, callback func(response *AddDRMCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDRMCertificateResponse
		var err error
		defer close(result)
		response, err = client.AddDRMCertificate(request)
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

// AddDRMCertificateRequest is the request struct for api AddDRMCertificate
type AddDRMCertificateRequest struct {
	*requests.RpcRequest
	ServCert    string           `position:"Query" name:"ServCert"`
	Description string           `position:"Query" name:"Description"`
	PrivateKey  string           `position:"Query" name:"PrivateKey"`
	CertName    string           `position:"Query" name:"CertName"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	Ask         string           `position:"Query" name:"Ask"`
	Passphrase  string           `position:"Query" name:"Passphrase"`
}

// AddDRMCertificateResponse is the response struct for api AddDRMCertificate
type AddDRMCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CertId    string `json:"CertId" xml:"CertId"`
}

// CreateAddDRMCertificateRequest creates a request to invoke AddDRMCertificate API
func CreateAddDRMCertificateRequest() (request *AddDRMCertificateRequest) {
	request = &AddDRMCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddDRMCertificate", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDRMCertificateResponse creates a response to parse from AddDRMCertificate response
func CreateAddDRMCertificateResponse() (response *AddDRMCertificateResponse) {
	response = &AddDRMCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
