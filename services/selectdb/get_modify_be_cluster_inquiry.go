package selectdb

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

// GetModifyBEClusterInquiry invokes the selectdb.GetModifyBEClusterInquiry API synchronously
func (client *Client) GetModifyBEClusterInquiry(request *GetModifyBEClusterInquiryRequest) (response *GetModifyBEClusterInquiryResponse, err error) {
	response = CreateGetModifyBEClusterInquiryResponse()
	err = client.DoAction(request, response)
	return
}

// GetModifyBEClusterInquiryWithChan invokes the selectdb.GetModifyBEClusterInquiry API asynchronously
func (client *Client) GetModifyBEClusterInquiryWithChan(request *GetModifyBEClusterInquiryRequest) (<-chan *GetModifyBEClusterInquiryResponse, <-chan error) {
	responseChan := make(chan *GetModifyBEClusterInquiryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetModifyBEClusterInquiry(request)
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

// GetModifyBEClusterInquiryWithCallback invokes the selectdb.GetModifyBEClusterInquiry API asynchronously
func (client *Client) GetModifyBEClusterInquiryWithCallback(request *GetModifyBEClusterInquiryRequest, callback func(response *GetModifyBEClusterInquiryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetModifyBEClusterInquiryResponse
		var err error
		defer close(result)
		response, err = client.GetModifyBEClusterInquiry(request)
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

// GetModifyBEClusterInquiryRequest is the request struct for api GetModifyBEClusterInquiry
type GetModifyBEClusterInquiryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CacheSize       requests.Integer `position:"Query" name:"CacheSize"`
	PreCacheSize    requests.Integer `position:"Query" name:"PreCacheSize"`
	ComputeSize     requests.Integer `position:"Query" name:"ComputeSize"`
	DbInstanceId    string           `position:"Query" name:"DbInstanceId"`
	PreComputeSize  requests.Integer `position:"Query" name:"PreComputeSize"`
	Quantity        requests.Integer `position:"Query" name:"Quantity"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	CommodityCode   string           `position:"Query" name:"CommodityCode"`
	ChargeType      string           `position:"Query" name:"ChargeType"`
	PricingCycle    string           `position:"Query" name:"PricingCycle"`
}

// GetModifyBEClusterInquiryResponse is the response struct for api GetModifyBEClusterInquiry
type GetModifyBEClusterInquiryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetModifyBEClusterInquiryRequest creates a request to invoke GetModifyBEClusterInquiry API
func CreateGetModifyBEClusterInquiryRequest() (request *GetModifyBEClusterInquiryRequest) {
	request = &GetModifyBEClusterInquiryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("selectdb", "2023-05-22", "GetModifyBEClusterInquiry", "", "")
	request.Method = requests.GET
	return
}

// CreateGetModifyBEClusterInquiryResponse creates a response to parse from GetModifyBEClusterInquiry response
func CreateGetModifyBEClusterInquiryResponse() (response *GetModifyBEClusterInquiryResponse) {
	response = &GetModifyBEClusterInquiryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
