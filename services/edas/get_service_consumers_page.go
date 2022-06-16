package edas

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

// GetServiceConsumersPage invokes the edas.GetServiceConsumersPage API synchronously
func (client *Client) GetServiceConsumersPage(request *GetServiceConsumersPageRequest) (response *GetServiceConsumersPageResponse, err error) {
	response = CreateGetServiceConsumersPageResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceConsumersPageWithChan invokes the edas.GetServiceConsumersPage API asynchronously
func (client *Client) GetServiceConsumersPageWithChan(request *GetServiceConsumersPageRequest) (<-chan *GetServiceConsumersPageResponse, <-chan error) {
	responseChan := make(chan *GetServiceConsumersPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceConsumersPage(request)
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

// GetServiceConsumersPageWithCallback invokes the edas.GetServiceConsumersPage API asynchronously
func (client *Client) GetServiceConsumersPageWithCallback(request *GetServiceConsumersPageRequest, callback func(response *GetServiceConsumersPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceConsumersPageResponse
		var err error
		defer close(result)
		response, err = client.GetServiceConsumersPage(request)
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

// GetServiceConsumersPageRequest is the request struct for api GetServiceConsumersPage
type GetServiceConsumersPageRequest struct {
	*requests.RoaRequest
	RegistryType   string `position:"Query" name:"registryType"`
	Origin         string `position:"Query" name:"origin"`
	Ip             string `position:"Query" name:"ip"`
	Source         string `position:"Query" name:"source"`
	ServiceType    string `position:"Query" name:"serviceType"`
	Size           string `position:"Query" name:"size"`
	AppId          string `position:"Query" name:"appId"`
	Namespace      string `position:"Query" name:"namespace"`
	ServiceVersion string `position:"Query" name:"serviceVersion"`
	ServiceName    string `position:"Query" name:"serviceName"`
	Page           string `position:"Query" name:"page"`
	Region         string `position:"Query" name:"region"`
	ServiceId      string `position:"Query" name:"serviceId"`
	Group          string `position:"Query" name:"group"`
}

// GetServiceConsumersPageResponse is the response struct for api GetServiceConsumersPage
type GetServiceConsumersPageResponse struct {
	*responses.BaseResponse
	Code    int    `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
	Success bool   `json:"Success" xml:"Success"`
	Data    Data   `json:"Data" xml:"Data"`
}

// CreateGetServiceConsumersPageRequest creates a request to invoke GetServiceConsumersPage API
func CreateGetServiceConsumersPageRequest() (request *GetServiceConsumersPageRequest) {
	request = &GetServiceConsumersPageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetServiceConsumersPage", "/pop/sp/api/mseForOam/getServiceConsumersPage", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetServiceConsumersPageResponse creates a response to parse from GetServiceConsumersPage response
func CreateGetServiceConsumersPageResponse() (response *GetServiceConsumersPageResponse) {
	response = &GetServiceConsumersPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
