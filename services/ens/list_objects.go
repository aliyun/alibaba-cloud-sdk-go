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

// ListObjects invokes the ens.ListObjects API synchronously
func (client *Client) ListObjects(request *ListObjectsRequest) (response *ListObjectsResponse, err error) {
	response = CreateListObjectsResponse()
	err = client.DoAction(request, response)
	return
}

// ListObjectsWithChan invokes the ens.ListObjects API asynchronously
func (client *Client) ListObjectsWithChan(request *ListObjectsRequest) (<-chan *ListObjectsResponse, <-chan error) {
	responseChan := make(chan *ListObjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListObjects(request)
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

// ListObjectsWithCallback invokes the ens.ListObjects API asynchronously
func (client *Client) ListObjectsWithCallback(request *ListObjectsRequest, callback func(response *ListObjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListObjectsResponse
		var err error
		defer close(result)
		response, err = client.ListObjects(request)
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

// ListObjectsRequest is the request struct for api ListObjects
type ListObjectsRequest struct {
	*requests.RpcRequest
	MaxKeys           requests.Integer `position:"Query" name:"MaxKeys"`
	ContinuationToken string           `position:"Query" name:"ContinuationToken"`
	Prefix            string           `position:"Query" name:"Prefix"`
	Delimiter         string           `position:"Query" name:"Delimiter"`
	Marker            string           `position:"Query" name:"Marker"`
	BucketName        string           `position:"Query" name:"BucketName"`
	EncodingType      string           `position:"Query" name:"EncodingType"`
	StartAfter        string           `position:"Query" name:"StartAfter"`
}

// ListObjectsResponse is the response struct for api ListObjects
type ListObjectsResponse struct {
	*responses.BaseResponse
	NextContinuationToken string             `json:"NextContinuationToken" xml:"NextContinuationToken"`
	ContinuationToken     string             `json:"ContinuationToken" xml:"ContinuationToken"`
	Delimiter             string             `json:"Delimiter" xml:"Delimiter"`
	EncodingType          string             `json:"EncodingType" xml:"EncodingType"`
	Prefix                string             `json:"Prefix" xml:"Prefix"`
	Marker                string             `json:"Marker" xml:"Marker"`
	BucketName            string             `json:"BucketName" xml:"BucketName"`
	IsTruncated           bool               `json:"IsTruncated" xml:"IsTruncated"`
	KeyCount              int64              `json:"KeyCount" xml:"KeyCount"`
	NextMarker            string             `json:"NextMarker" xml:"NextMarker"`
	MaxKeys               int64              `json:"MaxKeys" xml:"MaxKeys"`
	RequestId             string             `json:"RequestId" xml:"RequestId"`
	CommonPrefixes        []string           `json:"CommonPrefixes" xml:"CommonPrefixes"`
	Contents              []Content          `json:"Contents" xml:"Contents"`
	CommonPrefixInfos     []CommonPrefixInfo `json:"CommonPrefixInfos" xml:"CommonPrefixInfos"`
}

// CreateListObjectsRequest creates a request to invoke ListObjects API
func CreateListObjectsRequest() (request *ListObjectsRequest) {
	request = &ListObjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ListObjects", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListObjectsResponse creates a response to parse from ListObjects response
func CreateListObjectsResponse() (response *ListObjectsResponse) {
	response = &ListObjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
