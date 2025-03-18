package ens

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

// DescribeSnatAttribute invokes the ens.DescribeSnatAttribute API synchronously
func (client *Client) DescribeSnatAttribute(request *DescribeSnatAttributeRequest) (response *DescribeSnatAttributeResponse, err error) {
	response = CreateDescribeSnatAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSnatAttributeWithChan invokes the ens.DescribeSnatAttribute API asynchronously
func (client *Client) DescribeSnatAttributeWithChan(request *DescribeSnatAttributeRequest) (<-chan *DescribeSnatAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeSnatAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSnatAttribute(request)
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

// DescribeSnatAttributeWithCallback invokes the ens.DescribeSnatAttribute API asynchronously
func (client *Client) DescribeSnatAttributeWithCallback(request *DescribeSnatAttributeRequest, callback func(response *DescribeSnatAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSnatAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeSnatAttribute(request)
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

// DescribeSnatAttributeRequest is the request struct for api DescribeSnatAttribute
type DescribeSnatAttributeRequest struct {
	*requests.RpcRequest
	SnatEntryId string `position:"Query" name:"SnatEntryId"`
}

// DescribeSnatAttributeResponse is the response struct for api DescribeSnatAttribute
type DescribeSnatAttributeResponse struct {
	*responses.BaseResponse
	RequestId     string   `json:"RequestId" xml:"RequestId"`
	NatGatewayId  string   `json:"NatGatewayId" xml:"NatGatewayId"`
	SnatEntryId   string   `json:"SnatEntryId" xml:"SnatEntryId"`
	SnatIp        string   `json:"SnatIp" xml:"SnatIp"`
	SourceCIDR    string   `json:"SourceCIDR" xml:"SourceCIDR"`
	SnatEntryName string   `json:"SnatEntryName" xml:"SnatEntryName"`
	Status        string   `json:"Status" xml:"Status"`
	CreationTime  string   `json:"CreationTime" xml:"CreationTime"`
	StandbySnatIp string   `json:"StandbySnatIp" xml:"StandbySnatIp"`
	StandbyStatus string   `json:"StandbyStatus" xml:"StandbyStatus"`
	Type          string   `json:"Type" xml:"Type"`
	IdleTimeout   int      `json:"IdleTimeout" xml:"IdleTimeout"`
	DestCIDR      string   `json:"DestCIDR" xml:"DestCIDR"`
	IspAffinity   bool     `json:"IspAffinity" xml:"IspAffinity"`
	EipAffinity   bool     `json:"EipAffinity" xml:"EipAffinity"`
	SnatIps       []Snatip `json:"SnatIps" xml:"SnatIps"`
}

// CreateDescribeSnatAttributeRequest creates a request to invoke DescribeSnatAttribute API
func CreateDescribeSnatAttributeRequest() (request *DescribeSnatAttributeRequest) {
	request = &DescribeSnatAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeSnatAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSnatAttributeResponse creates a response to parse from DescribeSnatAttribute response
func CreateDescribeSnatAttributeResponse() (response *DescribeSnatAttributeResponse) {
	response = &DescribeSnatAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
