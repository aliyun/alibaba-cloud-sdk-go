package cbn

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

// AssociateCenBandwidthPackage invokes the cbn.AssociateCenBandwidthPackage API synchronously
func (client *Client) AssociateCenBandwidthPackage(request *AssociateCenBandwidthPackageRequest) (response *AssociateCenBandwidthPackageResponse, err error) {
	response = CreateAssociateCenBandwidthPackageResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateCenBandwidthPackageWithChan invokes the cbn.AssociateCenBandwidthPackage API asynchronously
func (client *Client) AssociateCenBandwidthPackageWithChan(request *AssociateCenBandwidthPackageRequest) (<-chan *AssociateCenBandwidthPackageResponse, <-chan error) {
	responseChan := make(chan *AssociateCenBandwidthPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateCenBandwidthPackage(request)
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

// AssociateCenBandwidthPackageWithCallback invokes the cbn.AssociateCenBandwidthPackage API asynchronously
func (client *Client) AssociateCenBandwidthPackageWithCallback(request *AssociateCenBandwidthPackageRequest, callback func(response *AssociateCenBandwidthPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateCenBandwidthPackageResponse
		var err error
		defer close(result)
		response, err = client.AssociateCenBandwidthPackage(request)
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

// AssociateCenBandwidthPackageRequest is the request struct for api AssociateCenBandwidthPackage
type AssociateCenBandwidthPackageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                 string           `position:"Query" name:"CenId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	Version               string           `position:"Query" name:"Version"`
	CenBandwidthPackageId string           `position:"Query" name:"CenBandwidthPackageId"`
}

// AssociateCenBandwidthPackageResponse is the response struct for api AssociateCenBandwidthPackage
type AssociateCenBandwidthPackageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateCenBandwidthPackageRequest creates a request to invoke AssociateCenBandwidthPackage API
func CreateAssociateCenBandwidthPackageRequest() (request *AssociateCenBandwidthPackageRequest) {
	request = &AssociateCenBandwidthPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "AssociateCenBandwidthPackage", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateCenBandwidthPackageResponse creates a response to parse from AssociateCenBandwidthPackage response
func CreateAssociateCenBandwidthPackageResponse() (response *AssociateCenBandwidthPackageResponse) {
	response = &AssociateCenBandwidthPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
