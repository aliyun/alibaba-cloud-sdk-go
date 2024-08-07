package market

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

// DescribeProduct invokes the market.DescribeProduct API synchronously
func (client *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
	response = CreateDescribeProductResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProductWithChan invokes the market.DescribeProduct API asynchronously
func (client *Client) DescribeProductWithChan(request *DescribeProductRequest) (<-chan *DescribeProductResponse, <-chan error) {
	responseChan := make(chan *DescribeProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProduct(request)
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

// DescribeProductWithCallback invokes the market.DescribeProduct API asynchronously
func (client *Client) DescribeProductWithCallback(request *DescribeProductRequest, callback func(response *DescribeProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProductResponse
		var err error
		defer close(result)
		response, err = client.DescribeProduct(request)
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

// DescribeProductRequest is the request struct for api DescribeProduct
type DescribeProductRequest struct {
	*requests.RpcRequest
	Code       string           `position:"Query" name:"Code"`
	QueryDraft requests.Boolean `position:"Query" name:"QueryDraft"`
	AliUid     string           `position:"Query" name:"AliUid"`
}

// DescribeProductResponse is the response struct for api DescribeProduct
type DescribeProductResponse struct {
	*responses.BaseResponse
	FrontCategoryId  int64         `json:"FrontCategoryId" xml:"FrontCategoryId"`
	Status           string        `json:"Status" xml:"Status"`
	Type             string        `json:"Type" xml:"Type"`
	PicUrl           string        `json:"PicUrl" xml:"PicUrl"`
	Score            float64       `json:"Score" xml:"Score"`
	UseCount         int64         `json:"UseCount" xml:"UseCount"`
	GmtModified      int64         `json:"GmtModified" xml:"GmtModified"`
	SupplierPk       int64         `json:"SupplierPk" xml:"SupplierPk"`
	GmtCreated       int64         `json:"GmtCreated" xml:"GmtCreated"`
	RequestId        string        `json:"RequestId" xml:"RequestId"`
	ShortDescription string        `json:"ShortDescription" xml:"ShortDescription"`
	Description      string        `json:"Description" xml:"Description"`
	Code             string        `json:"Code" xml:"Code"`
	AuditFailMsg     string        `json:"AuditFailMsg" xml:"AuditFailMsg"`
	Name             string        `json:"Name" xml:"Name"`
	AuditTime        int64         `json:"AuditTime" xml:"AuditTime"`
	AuditStatus      string        `json:"AuditStatus" xml:"AuditStatus"`
	ShopInfo         ShopInfo      `json:"ShopInfo" xml:"ShopInfo"`
	ProductSkus      ProductSkus   `json:"ProductSkus" xml:"ProductSkus"`
	ProductExtras    ProductExtras `json:"ProductExtras" xml:"ProductExtras"`
}

// CreateDescribeProductRequest creates a request to invoke DescribeProduct API
func CreateDescribeProductRequest() (request *DescribeProductRequest) {
	request = &DescribeProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "DescribeProduct", "yunmarket", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProductResponse creates a response to parse from DescribeProduct response
func CreateDescribeProductResponse() (response *DescribeProductResponse) {
	response = &DescribeProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
