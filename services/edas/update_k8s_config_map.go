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

// UpdateK8sConfigMap invokes the edas.UpdateK8sConfigMap API synchronously
func (client *Client) UpdateK8sConfigMap(request *UpdateK8sConfigMapRequest) (response *UpdateK8sConfigMapResponse, err error) {
	response = CreateUpdateK8sConfigMapResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateK8sConfigMapWithChan invokes the edas.UpdateK8sConfigMap API asynchronously
func (client *Client) UpdateK8sConfigMapWithChan(request *UpdateK8sConfigMapRequest) (<-chan *UpdateK8sConfigMapResponse, <-chan error) {
	responseChan := make(chan *UpdateK8sConfigMapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateK8sConfigMap(request)
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

// UpdateK8sConfigMapWithCallback invokes the edas.UpdateK8sConfigMap API asynchronously
func (client *Client) UpdateK8sConfigMapWithCallback(request *UpdateK8sConfigMapRequest, callback func(response *UpdateK8sConfigMapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateK8sConfigMapResponse
		var err error
		defer close(result)
		response, err = client.UpdateK8sConfigMap(request)
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

// UpdateK8sConfigMapRequest is the request struct for api UpdateK8sConfigMap
type UpdateK8sConfigMapRequest struct {
	*requests.RoaRequest
	Data      string `position:"Body" name:"Data"`
	Namespace string `position:"Body" name:"Namespace"`
	Name      string `position:"Body" name:"Name"`
	ClusterId string `position:"Body" name:"ClusterId"`
}

// UpdateK8sConfigMapResponse is the response struct for api UpdateK8sConfigMap
type UpdateK8sConfigMapResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateK8sConfigMapRequest creates a request to invoke UpdateK8sConfigMap API
func CreateUpdateK8sConfigMapRequest() (request *UpdateK8sConfigMapRequest) {
	request = &UpdateK8sConfigMapRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateK8sConfigMap", "/pop/v5/k8s/acs/k8s_config_map", "edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateK8sConfigMapResponse creates a response to parse from UpdateK8sConfigMap response
func CreateUpdateK8sConfigMapResponse() (response *UpdateK8sConfigMapResponse) {
	response = &UpdateK8sConfigMapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
