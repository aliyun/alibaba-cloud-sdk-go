package pvtz

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

// DescribeZoneInfo invokes the pvtz.DescribeZoneInfo API synchronously
func (client *Client) DescribeZoneInfo(request *DescribeZoneInfoRequest) (response *DescribeZoneInfoResponse, err error) {
	response = CreateDescribeZoneInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeZoneInfoWithChan invokes the pvtz.DescribeZoneInfo API asynchronously
func (client *Client) DescribeZoneInfoWithChan(request *DescribeZoneInfoRequest) (<-chan *DescribeZoneInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeZoneInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeZoneInfo(request)
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

// DescribeZoneInfoWithCallback invokes the pvtz.DescribeZoneInfo API asynchronously
func (client *Client) DescribeZoneInfoWithCallback(request *DescribeZoneInfoRequest, callback func(response *DescribeZoneInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeZoneInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeZoneInfo(request)
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

// DescribeZoneInfoRequest is the request struct for api DescribeZoneInfo
type DescribeZoneInfoRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeZoneInfoResponse is the response struct for api DescribeZoneInfo
type DescribeZoneInfoResponse struct {
	*responses.BaseResponse
	RequestId        string                     `json:"RequestId" xml:"RequestId"`
	SlaveDns         bool                       `json:"SlaveDns" xml:"SlaveDns"`
	ResourceGroupId  string                     `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ZoneId           string                     `json:"ZoneId" xml:"ZoneId"`
	ProxyPattern     string                     `json:"ProxyPattern" xml:"ProxyPattern"`
	CreateTime       string                     `json:"CreateTime" xml:"CreateTime"`
	ZoneType         string                     `json:"ZoneType" xml:"ZoneType"`
	Remark           string                     `json:"Remark" xml:"Remark"`
	ZoneName         string                     `json:"ZoneName" xml:"ZoneName"`
	ZoneTag          string                     `json:"ZoneTag" xml:"ZoneTag"`
	UpdateTime       string                     `json:"UpdateTime" xml:"UpdateTime"`
	UpdateTimestamp  int64                      `json:"UpdateTimestamp" xml:"UpdateTimestamp"`
	RecordCount      int                        `json:"RecordCount" xml:"RecordCount"`
	CreateTimestamp  int64                      `json:"CreateTimestamp" xml:"CreateTimestamp"`
	IsPtr            bool                       `json:"IsPtr" xml:"IsPtr"`
	DnsGroup         string                     `json:"DnsGroup" xml:"DnsGroup"`
	Creator          string                     `json:"Creator" xml:"Creator"`
	CreatorType      string                     `json:"CreatorType" xml:"CreatorType"`
	DnsGroupChanging bool                       `json:"DnsGroupChanging" xml:"DnsGroupChanging"`
	BindVpcs         BindVpcsInDescribeZoneInfo `json:"BindVpcs" xml:"BindVpcs"`
}

// CreateDescribeZoneInfoRequest creates a request to invoke DescribeZoneInfo API
func CreateDescribeZoneInfoRequest() (request *DescribeZoneInfoRequest) {
	request = &DescribeZoneInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "DescribeZoneInfo", "pvtz", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeZoneInfoResponse creates a response to parse from DescribeZoneInfo response
func CreateDescribeZoneInfoResponse() (response *DescribeZoneInfoResponse) {
	response = &DescribeZoneInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
