package scdn

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

// DescribeScdnDomainTopUrlVisit invokes the scdn.DescribeScdnDomainTopUrlVisit API synchronously
func (client *Client) DescribeScdnDomainTopUrlVisit(request *DescribeScdnDomainTopUrlVisitRequest) (response *DescribeScdnDomainTopUrlVisitResponse, err error) {
	response = CreateDescribeScdnDomainTopUrlVisitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainTopUrlVisitWithChan invokes the scdn.DescribeScdnDomainTopUrlVisit API asynchronously
func (client *Client) DescribeScdnDomainTopUrlVisitWithChan(request *DescribeScdnDomainTopUrlVisitRequest) (<-chan *DescribeScdnDomainTopUrlVisitResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainTopUrlVisitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainTopUrlVisit(request)
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

// DescribeScdnDomainTopUrlVisitWithCallback invokes the scdn.DescribeScdnDomainTopUrlVisit API asynchronously
func (client *Client) DescribeScdnDomainTopUrlVisitWithCallback(request *DescribeScdnDomainTopUrlVisitRequest, callback func(response *DescribeScdnDomainTopUrlVisitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainTopUrlVisitResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainTopUrlVisit(request)
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

// DescribeScdnDomainTopUrlVisitRequest is the request struct for api DescribeScdnDomainTopUrlVisit
type DescribeScdnDomainTopUrlVisitRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	SortBy     string `position:"Query" name:"SortBy"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeScdnDomainTopUrlVisitResponse is the response struct for api DescribeScdnDomainTopUrlVisit
type DescribeScdnDomainTopUrlVisitResponse struct {
	*responses.BaseResponse
	StartTime  string     `json:"StartTime" xml:"StartTime"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	DomainName string     `json:"DomainName" xml:"DomainName"`
	AllUrlList AllUrlList `json:"AllUrlList" xml:"AllUrlList"`
	Url200List Url200List `json:"Url200List" xml:"Url200List"`
	Url300List Url300List `json:"Url300List" xml:"Url300List"`
	Url400List Url400List `json:"Url400List" xml:"Url400List"`
	Url500List Url500List `json:"Url500List" xml:"Url500List"`
}

// CreateDescribeScdnDomainTopUrlVisitRequest creates a request to invoke DescribeScdnDomainTopUrlVisit API
func CreateDescribeScdnDomainTopUrlVisitRequest() (request *DescribeScdnDomainTopUrlVisitRequest) {
	request = &DescribeScdnDomainTopUrlVisitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainTopUrlVisit", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainTopUrlVisitResponse creates a response to parse from DescribeScdnDomainTopUrlVisit response
func CreateDescribeScdnDomainTopUrlVisitResponse() (response *DescribeScdnDomainTopUrlVisitResponse) {
	response = &DescribeScdnDomainTopUrlVisitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
