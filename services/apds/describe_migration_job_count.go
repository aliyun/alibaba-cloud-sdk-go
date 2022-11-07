package apds

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

// DescribeMigrationJobCount invokes the apds.DescribeMigrationJobCount API synchronously
func (client *Client) DescribeMigrationJobCount(request *DescribeMigrationJobCountRequest) (response *DescribeMigrationJobCountResponse, err error) {
	response = CreateDescribeMigrationJobCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMigrationJobCountWithChan invokes the apds.DescribeMigrationJobCount API asynchronously
func (client *Client) DescribeMigrationJobCountWithChan(request *DescribeMigrationJobCountRequest) (<-chan *DescribeMigrationJobCountResponse, <-chan error) {
	responseChan := make(chan *DescribeMigrationJobCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMigrationJobCount(request)
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

// DescribeMigrationJobCountWithCallback invokes the apds.DescribeMigrationJobCount API asynchronously
func (client *Client) DescribeMigrationJobCountWithCallback(request *DescribeMigrationJobCountRequest, callback func(response *DescribeMigrationJobCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMigrationJobCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeMigrationJobCount(request)
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

// DescribeMigrationJobCountRequest is the request struct for api DescribeMigrationJobCount
type DescribeMigrationJobCountRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// DescribeMigrationJobCountResponse is the response struct for api DescribeMigrationJobCount
type DescribeMigrationJobCountResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateDescribeMigrationJobCountRequest creates a request to invoke DescribeMigrationJobCount API
func CreateDescribeMigrationJobCountRequest() (request *DescribeMigrationJobCountRequest) {
	request = &DescribeMigrationJobCountRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "DescribeMigrationJobCount", "/okss-services/migration-job/count-migration-jobs", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeMigrationJobCountResponse creates a response to parse from DescribeMigrationJobCount response
func CreateDescribeMigrationJobCountResponse() (response *DescribeMigrationJobCountResponse) {
	response = &DescribeMigrationJobCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
