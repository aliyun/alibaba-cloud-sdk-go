package dms_enterprise

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

// ListDataLakeCatalog invokes the dms_enterprise.ListDataLakeCatalog API synchronously
func (client *Client) ListDataLakeCatalog(request *ListDataLakeCatalogRequest) (response *ListDataLakeCatalogResponse, err error) {
	response = CreateListDataLakeCatalogResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataLakeCatalogWithChan invokes the dms_enterprise.ListDataLakeCatalog API asynchronously
func (client *Client) ListDataLakeCatalogWithChan(request *ListDataLakeCatalogRequest) (<-chan *ListDataLakeCatalogResponse, <-chan error) {
	responseChan := make(chan *ListDataLakeCatalogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataLakeCatalog(request)
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

// ListDataLakeCatalogWithCallback invokes the dms_enterprise.ListDataLakeCatalog API asynchronously
func (client *Client) ListDataLakeCatalogWithCallback(request *ListDataLakeCatalogRequest, callback func(response *ListDataLakeCatalogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataLakeCatalogResponse
		var err error
		defer close(result)
		response, err = client.ListDataLakeCatalog(request)
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

// ListDataLakeCatalogRequest is the request struct for api ListDataLakeCatalog
type ListDataLakeCatalogRequest struct {
	*requests.RpcRequest
	SearchKey  string           `position:"Query" name:"SearchKey"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	DataRegion string           `position:"Query" name:"DataRegion"`
}

// ListDataLakeCatalogResponse is the response struct for api ListDataLakeCatalog
type ListDataLakeCatalogResponse struct {
	*responses.BaseResponse
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	Success      bool        `json:"Success" xml:"Success"`
	ErrorCode    string      `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string      `json:"ErrorMessage" xml:"ErrorMessage"`
	CataLogList  CataLogList `json:"CataLogList" xml:"CataLogList"`
}

// CreateListDataLakeCatalogRequest creates a request to invoke ListDataLakeCatalog API
func CreateListDataLakeCatalogRequest() (request *ListDataLakeCatalogRequest) {
	request = &ListDataLakeCatalogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDataLakeCatalog", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDataLakeCatalogResponse creates a response to parse from ListDataLakeCatalog response
func CreateListDataLakeCatalogResponse() (response *ListDataLakeCatalogResponse) {
	response = &ListDataLakeCatalogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
