package eais

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

// CreateEaiJupyter invokes the eais.CreateEaiJupyter API synchronously
func (client *Client) CreateEaiJupyter(request *CreateEaiJupyterRequest) (response *CreateEaiJupyterResponse, err error) {
	response = CreateCreateEaiJupyterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEaiJupyterWithChan invokes the eais.CreateEaiJupyter API asynchronously
func (client *Client) CreateEaiJupyterWithChan(request *CreateEaiJupyterRequest) (<-chan *CreateEaiJupyterResponse, <-chan error) {
	responseChan := make(chan *CreateEaiJupyterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEaiJupyter(request)
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

// CreateEaiJupyterWithCallback invokes the eais.CreateEaiJupyter API asynchronously
func (client *Client) CreateEaiJupyterWithCallback(request *CreateEaiJupyterRequest, callback func(response *CreateEaiJupyterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEaiJupyterResponse
		var err error
		defer close(result)
		response, err = client.CreateEaiJupyter(request)
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

// CreateEaiJupyterRequest is the request struct for api CreateEaiJupyter
type CreateEaiJupyterRequest struct {
	*requests.RpcRequest
	ClientToken     string                            `position:"Query" name:"ClientToken"`
	SecurityGroupId string                            `position:"Query" name:"SecurityGroupId"`
	EaisType        string                            `position:"Query" name:"EaisType"`
	ResourceGroupId string                            `position:"Query" name:"ResourceGroupId"`
	EaisName        string                            `position:"Query" name:"EaisName"`
	Tag             *[]CreateEaiJupyterTag            `position:"Query" name:"Tag"  type:"Repeated"`
	VSwitchId       string                            `position:"Query" name:"VSwitchId"`
	EnvironmentVar  *[]CreateEaiJupyterEnvironmentVar `position:"Query" name:"EnvironmentVar"  type:"Json"`
}

// CreateEaiJupyterEnvironmentVar is a repeated param struct in CreateEaiJupyterRequest
type CreateEaiJupyterEnvironmentVar struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateEaiJupyterTag is a repeated param struct in CreateEaiJupyterRequest
type CreateEaiJupyterTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateEaiJupyterResponse is the response struct for api CreateEaiJupyter
type CreateEaiJupyterResponse struct {
	*responses.BaseResponse
	RequestId                    string `json:"RequestId" xml:"RequestId"`
	ElasticAcceleratedInstanceId string `json:"ElasticAcceleratedInstanceId" xml:"ElasticAcceleratedInstanceId"`
}

// CreateCreateEaiJupyterRequest creates a request to invoke CreateEaiJupyter API
func CreateCreateEaiJupyterRequest() (request *CreateEaiJupyterRequest) {
	request = &CreateEaiJupyterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eais", "2019-06-24", "CreateEaiJupyter", "eais", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateEaiJupyterResponse creates a response to parse from CreateEaiJupyter response
func CreateCreateEaiJupyterResponse() (response *CreateEaiJupyterResponse) {
	response = &CreateEaiJupyterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
