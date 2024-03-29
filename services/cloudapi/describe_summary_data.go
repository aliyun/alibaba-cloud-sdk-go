package cloudapi

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

// DescribeSummaryData invokes the cloudapi.DescribeSummaryData API synchronously
func (client *Client) DescribeSummaryData(request *DescribeSummaryDataRequest) (response *DescribeSummaryDataResponse, err error) {
	response = CreateDescribeSummaryDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSummaryDataWithChan invokes the cloudapi.DescribeSummaryData API asynchronously
func (client *Client) DescribeSummaryDataWithChan(request *DescribeSummaryDataRequest) (<-chan *DescribeSummaryDataResponse, <-chan error) {
	responseChan := make(chan *DescribeSummaryDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSummaryData(request)
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

// DescribeSummaryDataWithCallback invokes the cloudapi.DescribeSummaryData API asynchronously
func (client *Client) DescribeSummaryDataWithCallback(request *DescribeSummaryDataRequest, callback func(response *DescribeSummaryDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSummaryDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeSummaryData(request)
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

// DescribeSummaryDataRequest is the request struct for api DescribeSummaryData
type DescribeSummaryDataRequest struct {
	*requests.RpcRequest
	Language      string `position:"Query" name:"Language"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeSummaryDataResponse is the response struct for api DescribeSummaryData
type DescribeSummaryDataResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	UsageGroupNum          int    `json:"UsageGroupNum" xml:"UsageGroupNum"`
	UsageInstanceNum       int    `json:"UsageInstanceNum" xml:"UsageInstanceNum"`
	UsageApiNum            int    `json:"UsageApiNum" xml:"UsageApiNum"`
	Region                 string `json:"Region" xml:"Region"`
	ComeingSoonInstanceNum int    `json:"ComeingSoonInstanceNum" xml:"ComeingSoonInstanceNum"`
	ExpireInstanceNum      int    `json:"ExpireInstanceNum" xml:"ExpireInstanceNum"`
}

// CreateDescribeSummaryDataRequest creates a request to invoke DescribeSummaryData API
func CreateDescribeSummaryDataRequest() (request *DescribeSummaryDataRequest) {
	request = &DescribeSummaryDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeSummaryData", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSummaryDataResponse creates a response to parse from DescribeSummaryData response
func CreateDescribeSummaryDataResponse() (response *DescribeSummaryDataResponse) {
	response = &DescribeSummaryDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
