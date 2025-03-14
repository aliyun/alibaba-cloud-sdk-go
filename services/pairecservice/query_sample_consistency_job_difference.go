package pairecservice

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

// QuerySampleConsistencyJobDifference invokes the pairecservice.QuerySampleConsistencyJobDifference API synchronously
func (client *Client) QuerySampleConsistencyJobDifference(request *QuerySampleConsistencyJobDifferenceRequest) (response *QuerySampleConsistencyJobDifferenceResponse, err error) {
	response = CreateQuerySampleConsistencyJobDifferenceResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySampleConsistencyJobDifferenceWithChan invokes the pairecservice.QuerySampleConsistencyJobDifference API asynchronously
func (client *Client) QuerySampleConsistencyJobDifferenceWithChan(request *QuerySampleConsistencyJobDifferenceRequest) (<-chan *QuerySampleConsistencyJobDifferenceResponse, <-chan error) {
	responseChan := make(chan *QuerySampleConsistencyJobDifferenceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySampleConsistencyJobDifference(request)
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

// QuerySampleConsistencyJobDifferenceWithCallback invokes the pairecservice.QuerySampleConsistencyJobDifference API asynchronously
func (client *Client) QuerySampleConsistencyJobDifferenceWithCallback(request *QuerySampleConsistencyJobDifferenceRequest, callback func(response *QuerySampleConsistencyJobDifferenceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySampleConsistencyJobDifferenceResponse
		var err error
		defer close(result)
		response, err = client.QuerySampleConsistencyJobDifference(request)
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

// QuerySampleConsistencyJobDifferenceRequest is the request struct for api QuerySampleConsistencyJobDifference
type QuerySampleConsistencyJobDifferenceRequest struct {
	*requests.RoaRequest
	FeatureType            string `position:"Query" name:"FeatureType"`
	SampleConsistencyJobId string `position:"Path" name:"SampleConsistencyJobId"`
	FeatureName            string `position:"Query" name:"FeatureName"`
	InstanceId             string `position:"Query" name:"InstanceId"`
}

// QuerySampleConsistencyJobDifferenceResponse is the response struct for api QuerySampleConsistencyJobDifference
type QuerySampleConsistencyJobDifferenceResponse struct {
	*responses.BaseResponse
	RequestId                string                         `json:"RequestId" xml:"RequestId"`
	StringFeatureDifferences []StringFeatureDifferencesItem `json:"StringFeatureDifferences" xml:"StringFeatureDifferences"`
	NumberFeatureDifferences []NumberFeatureDifferencesItem `json:"NumberFeatureDifferences" xml:"NumberFeatureDifferences"`
	DifferenceHistogram      []DifferenceHistogramItem      `json:"DifferenceHistogram" xml:"DifferenceHistogram"`
}

// CreateQuerySampleConsistencyJobDifferenceRequest creates a request to invoke QuerySampleConsistencyJobDifference API
func CreateQuerySampleConsistencyJobDifferenceRequest() (request *QuerySampleConsistencyJobDifferenceRequest) {
	request = &QuerySampleConsistencyJobDifferenceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "QuerySampleConsistencyJobDifference", "/api/v1/sampleconsistencyjobs/[SampleConsistencyJobId]/action/querydifference", "", "")
	request.Method = requests.GET
	return
}

// CreateQuerySampleConsistencyJobDifferenceResponse creates a response to parse from QuerySampleConsistencyJobDifference response
func CreateQuerySampleConsistencyJobDifferenceResponse() (response *QuerySampleConsistencyJobDifferenceResponse) {
	response = &QuerySampleConsistencyJobDifferenceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
