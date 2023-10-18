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

// DescribeLiveDomainMonitoringUsageData invokes the live.DescribeLiveDomainMonitoringUsageData API synchronously
func (client *Client) DescribeLiveDomainMonitoringUsageData(request *DescribeLiveDomainMonitoringUsageDataRequest) (response *DescribeLiveDomainMonitoringUsageDataResponse, err error) {
	response = CreateDescribeLiveDomainMonitoringUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainMonitoringUsageDataWithChan invokes the live.DescribeLiveDomainMonitoringUsageData API asynchronously
func (client *Client) DescribeLiveDomainMonitoringUsageDataWithChan(request *DescribeLiveDomainMonitoringUsageDataRequest) (<-chan *DescribeLiveDomainMonitoringUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainMonitoringUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainMonitoringUsageData(request)
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

// DescribeLiveDomainMonitoringUsageDataWithCallback invokes the live.DescribeLiveDomainMonitoringUsageData API asynchronously
func (client *Client) DescribeLiveDomainMonitoringUsageDataWithCallback(request *DescribeLiveDomainMonitoringUsageDataRequest, callback func(response *DescribeLiveDomainMonitoringUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainMonitoringUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainMonitoringUsageData(request)
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

// DescribeLiveDomainMonitoringUsageDataRequest is the request struct for api DescribeLiveDomainMonitoringUsageData
type DescribeLiveDomainMonitoringUsageDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	SplitBy    string           `position:"Query" name:"SplitBy"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	Interval   string           `position:"Query" name:"Interval"`
	Region     string           `position:"Query" name:"Region"`
}

// DescribeLiveDomainMonitoringUsageDataResponse is the response struct for api DescribeLiveDomainMonitoringUsageData
type DescribeLiveDomainMonitoringUsageDataResponse struct {
	*responses.BaseResponse
	EndTime        string         `json:"EndTime" xml:"EndTime"`
	StartTime      string         `json:"StartTime" xml:"StartTime"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Region         string         `json:"Region" xml:"Region"`
	InstanceId     string         `json:"InstanceId" xml:"InstanceId"`
	DomainName     string         `json:"DomainName" xml:"DomainName"`
	MonitoringData MonitoringData `json:"MonitoringData" xml:"MonitoringData"`
}

// CreateDescribeLiveDomainMonitoringUsageDataRequest creates a request to invoke DescribeLiveDomainMonitoringUsageData API
func CreateDescribeLiveDomainMonitoringUsageDataRequest() (request *DescribeLiveDomainMonitoringUsageDataRequest) {
	request = &DescribeLiveDomainMonitoringUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainMonitoringUsageData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainMonitoringUsageDataResponse creates a response to parse from DescribeLiveDomainMonitoringUsageData response
func CreateDescribeLiveDomainMonitoringUsageDataResponse() (response *DescribeLiveDomainMonitoringUsageDataResponse) {
	response = &DescribeLiveDomainMonitoringUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
