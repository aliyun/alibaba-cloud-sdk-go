package paifeaturestore

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

// DeleteFeatureView invokes the paifeaturestore.DeleteFeatureView API synchronously
func (client *Client) DeleteFeatureView(request *DeleteFeatureViewRequest) (response *DeleteFeatureViewResponse, err error) {
	response = CreateDeleteFeatureViewResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFeatureViewWithChan invokes the paifeaturestore.DeleteFeatureView API asynchronously
func (client *Client) DeleteFeatureViewWithChan(request *DeleteFeatureViewRequest) (<-chan *DeleteFeatureViewResponse, <-chan error) {
	responseChan := make(chan *DeleteFeatureViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFeatureView(request)
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

// DeleteFeatureViewWithCallback invokes the paifeaturestore.DeleteFeatureView API asynchronously
func (client *Client) DeleteFeatureViewWithCallback(request *DeleteFeatureViewRequest, callback func(response *DeleteFeatureViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFeatureViewResponse
		var err error
		defer close(result)
		response, err = client.DeleteFeatureView(request)
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

// DeleteFeatureViewRequest is the request struct for api DeleteFeatureView
type DeleteFeatureViewRequest struct {
	*requests.RoaRequest
	InstanceId    string `position:"Path" name:"InstanceId"`
	FeatureViewId string `position:"Path" name:"FeatureViewId"`
}

// DeleteFeatureViewResponse is the response struct for api DeleteFeatureView
type DeleteFeatureViewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFeatureViewRequest creates a request to invoke DeleteFeatureView API
func CreateDeleteFeatureViewRequest() (request *DeleteFeatureViewRequest) {
	request = &DeleteFeatureViewRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "DeleteFeatureView", "/api/v1/instances/[InstanceId]/featureviews/[FeatureViewId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteFeatureViewResponse creates a response to parse from DeleteFeatureView response
func CreateDeleteFeatureViewResponse() (response *DeleteFeatureViewResponse) {
	response = &DeleteFeatureViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
