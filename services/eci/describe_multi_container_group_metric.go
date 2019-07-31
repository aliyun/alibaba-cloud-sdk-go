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

package eci

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeMultiContainerGroupMetric invokes the eci.DescribeMultiContainerGroupMetric API synchronously
// api document: https://help.aliyun.com/api/eci/describemulticontainergroupmetric.html
func (client *Client) DescribeMultiContainerGroupMetric(request *DescribeMultiContainerGroupMetricRequest) (response *DescribeMultiContainerGroupMetricResponse, err error) {
	response = CreateDescribeMultiContainerGroupMetricResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMultiContainerGroupMetricWithChan invokes the eci.DescribeMultiContainerGroupMetric API asynchronously
// api document: https://help.aliyun.com/api/eci/describemulticontainergroupmetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMultiContainerGroupMetricWithChan(request *DescribeMultiContainerGroupMetricRequest) (<-chan *DescribeMultiContainerGroupMetricResponse, <-chan error) {
	responseChan := make(chan *DescribeMultiContainerGroupMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMultiContainerGroupMetric(request)
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

// DescribeMultiContainerGroupMetricWithCallback invokes the eci.DescribeMultiContainerGroupMetric API asynchronously
// api document: https://help.aliyun.com/api/eci/describemulticontainergroupmetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMultiContainerGroupMetricWithCallback(request *DescribeMultiContainerGroupMetricRequest, callback func(response *DescribeMultiContainerGroupMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMultiContainerGroupMetricResponse
		var err error
		defer close(result)
		response, err = client.DescribeMultiContainerGroupMetric(request)
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

// DescribeMultiContainerGroupMetricRequest is the request struct for api DescribeMultiContainerGroupMetric
type DescribeMultiContainerGroupMetricRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	RegionId             string           `position:"Query" name:"RegionId"`
	ContainerGroupIds    string           `position:"Query" name:"ContainerGroupIds"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
}

// DescribeMultiContainerGroupMetricResponse is the response struct for api DescribeMultiContainerGroupMetric
type DescribeMultiContainerGroupMetricResponse struct {
	*responses.BaseResponse
	RequestId    string                                           `json:"RequestId" xml:"RequestId"`
	MonitorDatas []DescribeMultiContainerGroupMetricMonitorDatas0 `json:"MonitorDatas" xml:"MonitorDatas"`
}

type DescribeMultiContainerGroupMetricMonitorDatas0 struct {
	Record []DescribeMultiContainerGroupMetricRecord1 `json:"Record" xml:"Record"`
}

type DescribeMultiContainerGroupMetricRecord1 struct {
	ContainerGroupId string                                      `json:"ContainerGroupId" xml:"ContainerGroupId"`
	Records          []DescribeMultiContainerGroupMetricRecords1 `json:"Records" xml:"Records"`
}

type DescribeMultiContainerGroupMetricRecords1 struct {
	PodStat []DescribeMultiContainerGroupMetricPodStat2 `json:"PodStat" xml:"PodStat"`
}

type DescribeMultiContainerGroupMetricPodStat2 struct {
	Timestamp  string                                         `json:"Timestamp" xml:"Timestamp"`
	Containers []DescribeMultiContainerGroupMetricContainers2 `json:"Containers" xml:"Containers"`
	CPU        DescribeMultiContainerGroupMetricCPU2          `json:"CPU" xml:"CPU"`
	Memory     DescribeMultiContainerGroupMetricMemory2       `json:"Memory" xml:"Memory"`
	Network    DescribeMultiContainerGroupMetricNetwork2      `json:"Network" xml:"Network"`
}

type DescribeMultiContainerGroupMetricContainers2 struct {
	Container []DescribeMultiContainerGroupMetricContainer3 `json:"Container" xml:"Container"`
}

type DescribeMultiContainerGroupMetricContainer3 struct {
	Name   string                                   `json:"Name" xml:"Name"`
	CPU    DescribeMultiContainerGroupMetricCPU3    `json:"CPU" xml:"CPU"`
	Memory DescribeMultiContainerGroupMetricMemory3 `json:"Memory" xml:"Memory"`
}

type DescribeMultiContainerGroupMetricCPU3 struct {
	UsageNanoCores       int64 `json:"UsageNanoCores" xml:"UsageNanoCores"`
	UsageCoreNanoSeconds int64 `json:"UsageCoreNanoSeconds" xml:"UsageCoreNanoSeconds"`
	Load                 int64 `json:"Load" xml:"Load"`
	Limit                int64 `json:"Limit" xml:"Limit"`
}

type DescribeMultiContainerGroupMetricMemory3 struct {
	AvailableBytes int64 `json:"AvailableBytes" xml:"AvailableBytes"`
	UsageBytes     int64 `json:"UsageBytes" xml:"UsageBytes"`
	Cache          int64 `json:"Cache" xml:"Cache"`
	WorkingSet     int64 `json:"WorkingSet" xml:"WorkingSet"`
	Rss            int64 `json:"Rss" xml:"Rss"`
}

type DescribeMultiContainerGroupMetricCPU2 struct {
	UsageNanoCores       int64 `json:"UsageNanoCores" xml:"UsageNanoCores"`
	UsageCoreNanoSeconds int64 `json:"UsageCoreNanoSeconds" xml:"UsageCoreNanoSeconds"`
	Load                 int64 `json:"Load" xml:"Load"`
	Limit                int64 `json:"Limit" xml:"Limit"`
}

type DescribeMultiContainerGroupMetricMemory2 struct {
	AvailableBytes int64 `json:"AvailableBytes" xml:"AvailableBytes"`
	UsageBytes     int64 `json:"UsageBytes" xml:"UsageBytes"`
	Cache          int64 `json:"Cache" xml:"Cache"`
	WorkingSet     int64 `json:"WorkingSet" xml:"WorkingSet"`
	Rss            int64 `json:"Rss" xml:"Rss"`
}

type DescribeMultiContainerGroupMetricNetwork2 struct {
	Interfaces []DescribeMultiContainerGroupMetricInterfaces3 `json:"Interfaces" xml:"Interfaces"`
}

type DescribeMultiContainerGroupMetricInterfaces3 struct {
	Interface []DescribeMultiContainerGroupMetricInterface4 `json:"Interface" xml:"Interface"`
}

type DescribeMultiContainerGroupMetricInterface4 struct {
	TxBytes  int64  `json:"TxBytes" xml:"TxBytes"`
	RxBytes  int64  `json:"RxBytes" xml:"RxBytes"`
	TxErrors int64  `json:"TxErrors" xml:"TxErrors"`
	RxErrors int64  `json:"RxErrors" xml:"RxErrors"`
	Name     string `json:"Name" xml:"Name"`
}

// CreateDescribeMultiContainerGroupMetricRequest creates a request to invoke DescribeMultiContainerGroupMetric API
func CreateDescribeMultiContainerGroupMetricRequest() (request *DescribeMultiContainerGroupMetricRequest) {
	request = &DescribeMultiContainerGroupMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eci", "2018-08-08", "DescribeMultiContainerGroupMetric", "eci", "openAPI")
	return
}

// CreateDescribeMultiContainerGroupMetricResponse creates a response to parse from DescribeMultiContainerGroupMetric response
func CreateDescribeMultiContainerGroupMetricResponse() (response *DescribeMultiContainerGroupMetricResponse) {
	response = &DescribeMultiContainerGroupMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
