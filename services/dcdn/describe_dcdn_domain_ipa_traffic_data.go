package dcdn

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

// DescribeDcdnDomainIpaTrafficData invokes the dcdn.DescribeDcdnDomainIpaTrafficData API synchronously
func (client *Client) DescribeDcdnDomainIpaTrafficData(request *DescribeDcdnDomainIpaTrafficDataRequest) (response *DescribeDcdnDomainIpaTrafficDataResponse, err error) {
	response = CreateDescribeDcdnDomainIpaTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainIpaTrafficDataWithChan invokes the dcdn.DescribeDcdnDomainIpaTrafficData API asynchronously
func (client *Client) DescribeDcdnDomainIpaTrafficDataWithChan(request *DescribeDcdnDomainIpaTrafficDataRequest) (<-chan *DescribeDcdnDomainIpaTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainIpaTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainIpaTrafficData(request)
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

// DescribeDcdnDomainIpaTrafficDataWithCallback invokes the dcdn.DescribeDcdnDomainIpaTrafficData API asynchronously
func (client *Client) DescribeDcdnDomainIpaTrafficDataWithCallback(request *DescribeDcdnDomainIpaTrafficDataRequest, callback func(response *DescribeDcdnDomainIpaTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainIpaTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainIpaTrafficData(request)
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

// DescribeDcdnDomainIpaTrafficDataRequest is the request struct for api DescribeDcdnDomainIpaTrafficData
type DescribeDcdnDomainIpaTrafficDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	FixTimeGap     string           `position:"Query" name:"FixTimeGap"`
	TimeMerge      string           `position:"Query" name:"TimeMerge"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeDcdnDomainIpaTrafficDataResponse is the response struct for api DescribeDcdnDomainIpaTrafficData
type DescribeDcdnDomainIpaTrafficDataResponse struct {
	*responses.BaseResponse
	EndTime                string                                                   `json:"EndTime" xml:"EndTime"`
	StartTime              string                                                   `json:"StartTime" xml:"StartTime"`
	RequestId              string                                                   `json:"RequestId" xml:"RequestId"`
	DomainName             string                                                   `json:"DomainName" xml:"DomainName"`
	DataInterval           string                                                   `json:"DataInterval" xml:"DataInterval"`
	TrafficDataPerInterval TrafficDataPerIntervalInDescribeDcdnDomainIpaTrafficData `json:"TrafficDataPerInterval" xml:"TrafficDataPerInterval"`
}

// CreateDescribeDcdnDomainIpaTrafficDataRequest creates a request to invoke DescribeDcdnDomainIpaTrafficData API
func CreateDescribeDcdnDomainIpaTrafficDataRequest() (request *DescribeDcdnDomainIpaTrafficDataRequest) {
	request = &DescribeDcdnDomainIpaTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainIpaTrafficData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainIpaTrafficDataResponse creates a response to parse from DescribeDcdnDomainIpaTrafficData response
func CreateDescribeDcdnDomainIpaTrafficDataResponse() (response *DescribeDcdnDomainIpaTrafficDataResponse) {
	response = &DescribeDcdnDomainIpaTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
