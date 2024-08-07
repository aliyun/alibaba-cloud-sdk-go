package csas

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

// ListNacUserCert invokes the csas.ListNacUserCert API synchronously
func (client *Client) ListNacUserCert(request *ListNacUserCertRequest) (response *ListNacUserCertResponse, err error) {
	response = CreateListNacUserCertResponse()
	err = client.DoAction(request, response)
	return
}

// ListNacUserCertWithChan invokes the csas.ListNacUserCert API asynchronously
func (client *Client) ListNacUserCertWithChan(request *ListNacUserCertRequest) (<-chan *ListNacUserCertResponse, <-chan error) {
	responseChan := make(chan *ListNacUserCertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNacUserCert(request)
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

// ListNacUserCertWithCallback invokes the csas.ListNacUserCert API asynchronously
func (client *Client) ListNacUserCertWithCallback(request *ListNacUserCertRequest, callback func(response *ListNacUserCertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNacUserCertResponse
		var err error
		defer close(result)
		response, err = client.ListNacUserCert(request)
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

// ListNacUserCertRequest is the request struct for api ListNacUserCert
type ListNacUserCertRequest struct {
	*requests.RpcRequest
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage string           `position:"Query" name:"CurrentPage"`
	StartTime   requests.Integer `position:"Query" name:"StartTime"`
	DeviceType  string           `position:"Query" name:"DeviceType"`
	PageSize    string           `position:"Query" name:"PageSize"`
	Department  string           `position:"Query" name:"Department"`
	Username    string           `position:"Query" name:"Username"`
	Status      string           `position:"Query" name:"Status"`
}

// ListNacUserCertResponse is the response struct for api ListNacUserCert
type ListNacUserCertResponse struct {
	*responses.BaseResponse
	RequestId string                          `json:"RequestId" xml:"RequestId"`
	Code      int64                           `json:"Code" xml:"Code"`
	Message   string                          `json:"Message" xml:"Message"`
	TotalNum  int64                           `json:"TotalNum" xml:"TotalNum"`
	DataList  []DataListItemInListNacUserCert `json:"DataList" xml:"DataList"`
}

// CreateListNacUserCertRequest creates a request to invoke ListNacUserCert API
func CreateListNacUserCertRequest() (request *ListNacUserCertRequest) {
	request = &ListNacUserCertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListNacUserCert", "", "")
	request.Method = requests.POST
	return
}

// CreateListNacUserCertResponse creates a response to parse from ListNacUserCert response
func CreateListNacUserCertResponse() (response *ListNacUserCertResponse) {
	response = &ListNacUserCertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
