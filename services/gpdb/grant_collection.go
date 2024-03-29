package gpdb

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

// GrantCollection invokes the gpdb.GrantCollection API synchronously
func (client *Client) GrantCollection(request *GrantCollectionRequest) (response *GrantCollectionResponse, err error) {
	response = CreateGrantCollectionResponse()
	err = client.DoAction(request, response)
	return
}

// GrantCollectionWithChan invokes the gpdb.GrantCollection API asynchronously
func (client *Client) GrantCollectionWithChan(request *GrantCollectionRequest) (<-chan *GrantCollectionResponse, <-chan error) {
	responseChan := make(chan *GrantCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrantCollection(request)
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

// GrantCollectionWithCallback invokes the gpdb.GrantCollection API asynchronously
func (client *Client) GrantCollectionWithCallback(request *GrantCollectionRequest, callback func(response *GrantCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrantCollectionResponse
		var err error
		defer close(result)
		response, err = client.GrantCollection(request)
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

// GrantCollectionRequest is the request struct for api GrantCollection
type GrantCollectionRequest struct {
	*requests.RpcRequest
	GrantType              string           `position:"Query" name:"GrantType"`
	ManagerAccount         string           `position:"Query" name:"ManagerAccount"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ManagerAccountPassword string           `position:"Query" name:"ManagerAccountPassword"`
	Collection             string           `position:"Query" name:"Collection"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	GrantToNamespace       string           `position:"Query" name:"GrantToNamespace"`
	Namespace              string           `position:"Query" name:"Namespace"`
}

// GrantCollectionResponse is the response struct for api GrantCollection
type GrantCollectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateGrantCollectionRequest creates a request to invoke GrantCollection API
func CreateGrantCollectionRequest() (request *GrantCollectionRequest) {
	request = &GrantCollectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "GrantCollection", "", "")
	request.Method = requests.POST
	return
}

// CreateGrantCollectionResponse creates a response to parse from GrantCollection response
func CreateGrantCollectionResponse() (response *GrantCollectionResponse) {
	response = &GrantCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
