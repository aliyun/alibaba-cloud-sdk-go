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

// DeleteMigrationJob invokes the apds.DeleteMigrationJob API synchronously
func (client *Client) DeleteMigrationJob(request *DeleteMigrationJobRequest) (response *DeleteMigrationJobResponse, err error) {
	response = CreateDeleteMigrationJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMigrationJobWithChan invokes the apds.DeleteMigrationJob API asynchronously
func (client *Client) DeleteMigrationJobWithChan(request *DeleteMigrationJobRequest) (<-chan *DeleteMigrationJobResponse, <-chan error) {
	responseChan := make(chan *DeleteMigrationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMigrationJob(request)
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

// DeleteMigrationJobWithCallback invokes the apds.DeleteMigrationJob API asynchronously
func (client *Client) DeleteMigrationJobWithCallback(request *DeleteMigrationJobRequest, callback func(response *DeleteMigrationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMigrationJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteMigrationJob(request)
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

// DeleteMigrationJobRequest is the request struct for api DeleteMigrationJob
type DeleteMigrationJobRequest struct {
	*requests.RoaRequest
	Id string `position:"Query" name:"id"`
}

// DeleteMigrationJobResponse is the response struct for api DeleteMigrationJob
type DeleteMigrationJobResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateDeleteMigrationJobRequest creates a request to invoke DeleteMigrationJob API
func CreateDeleteMigrationJobRequest() (request *DeleteMigrationJobRequest) {
	request = &DeleteMigrationJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "DeleteMigrationJob", "/okss-services/migration-job/remove-migration-job", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMigrationJobResponse creates a response to parse from DeleteMigrationJob response
func CreateDeleteMigrationJobResponse() (response *DeleteMigrationJobResponse) {
	response = &DeleteMigrationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
