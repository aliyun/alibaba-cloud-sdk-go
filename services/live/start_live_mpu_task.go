package live

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

// StartLiveMPUTask invokes the live.StartLiveMPUTask API synchronously
func (client *Client) StartLiveMPUTask(request *StartLiveMPUTaskRequest) (response *StartLiveMPUTaskResponse, err error) {
	response = CreateStartLiveMPUTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StartLiveMPUTaskWithChan invokes the live.StartLiveMPUTask API asynchronously
func (client *Client) StartLiveMPUTaskWithChan(request *StartLiveMPUTaskRequest) (<-chan *StartLiveMPUTaskResponse, <-chan error) {
	responseChan := make(chan *StartLiveMPUTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartLiveMPUTask(request)
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

// StartLiveMPUTaskWithCallback invokes the live.StartLiveMPUTask API asynchronously
func (client *Client) StartLiveMPUTaskWithCallback(request *StartLiveMPUTaskRequest, callback func(response *StartLiveMPUTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartLiveMPUTaskResponse
		var err error
		defer close(result)
		response, err = client.StartLiveMPUTask(request)
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

// StartLiveMPUTaskRequest is the request struct for api StartLiveMPUTask
type StartLiveMPUTaskRequest struct {
	*requests.RpcRequest
	SingleSubParams StartLiveMPUTaskSingleSubParams `position:"Query" name:"SingleSubParams"  type:"Struct"`
	SeiParams       StartLiveMPUTaskSeiParams       `position:"Query" name:"SeiParams"  type:"Struct"`
	TranscodeParams StartLiveMPUTaskTranscodeParams `position:"Query" name:"TranscodeParams"  type:"Struct"`
	AppId           string                          `position:"Query" name:"AppId"`
	Region          string                          `position:"Query" name:"Region"`
	MixMode         string                          `position:"Query" name:"MixMode"`
	ChannelId       string                          `position:"Query" name:"ChannelId"`
	TaskId          string                          `position:"Query" name:"TaskId"`
	StreamURL       string                          `position:"Query" name:"StreamURL"`
}

// StartLiveMPUTaskSingleSubParams is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskSingleSubParams struct {
	StreamType string `name:"StreamType"`
	SourceType string `name:"SourceType"`
	UserId     string `name:"UserId"`
}

// StartLiveMPUTaskSeiParams is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskSeiParams struct {
	PayloadType  string                                `name:"PayloadType"`
	LayoutVolume StartLiveMPUTaskSeiParamsLayoutVolume `name:"LayoutVolume" type:"Struct"`
	PassThrough  StartLiveMPUTaskSeiParamsPassThrough  `name:"PassThrough" type:"Struct"`
}

// StartLiveMPUTaskTranscodeParams is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskTranscodeParams struct {
	Layout       StartLiveMPUTaskTranscodeParamsLayout           `name:"Layout" type:"Struct"`
	Background   StartLiveMPUTaskTranscodeParamsBackground       `name:"Background" type:"Struct"`
	UserInfos    *[]StartLiveMPUTaskTranscodeParamsUserInfosItem `name:"UserInfos" type:"Repeated"`
	EncodeParams StartLiveMPUTaskTranscodeParamsEncodeParams     `name:"EncodeParams" type:"Struct"`
}

// StartLiveMPUTaskSeiParamsLayoutVolume is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskSeiParamsLayoutVolume struct {
	FollowIdr string `name:"FollowIdr"`
	Interval  string `name:"Interval"`
}

// StartLiveMPUTaskSeiParamsPassThrough is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskSeiParamsPassThrough struct {
	FollowIdr         string `name:"FollowIdr"`
	PayloadContentKey string `name:"PayloadContentKey"`
	PayloadContent    string `name:"PayloadContent"`
	Interval          string `name:"Interval"`
}

// StartLiveMPUTaskTranscodeParamsLayout is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskTranscodeParamsLayout struct {
	UserPanes    *[]StartLiveMPUTaskTranscodeParamsLayoutUserPanesItem `name:"UserPanes" type:"Repeated"`
	LayoutMode   string                                                `name:"LayoutMode"`
	MaxVideoUser StartLiveMPUTaskTranscodeParamsLayoutMaxVideoUser     `name:"MaxVideoUser" type:"Struct"`
}

// StartLiveMPUTaskTranscodeParamsBackground is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskTranscodeParamsBackground struct {
	URL        string `name:"URL"`
	RenderMode string `name:"RenderMode"`
}

// StartLiveMPUTaskTranscodeParamsUserInfosItem is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskTranscodeParamsUserInfosItem struct {
	StreamType string `name:"StreamType"`
	SourceType string `name:"SourceType"`
	UserId     string `name:"UserId"`
}

// StartLiveMPUTaskTranscodeParamsEncodeParams is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskTranscodeParamsEncodeParams struct {
	AudioOnly       string `name:"AudioOnly"`
	VideoWidth      string `name:"VideoWidth"`
	AudioBitrate    string `name:"AudioBitrate"`
	VideoFramerate  string `name:"VideoFramerate"`
	VIdeoHeight     string `name:"VIdeoHeight"`
	VideoBitrate    string `name:"VideoBitrate"`
	AudioSampleRate string `name:"AudioSampleRate"`
	VideoGop        string `name:"VideoGop"`
	AudioChannels   string `name:"AudioChannels"`
}

// StartLiveMPUTaskTranscodeParamsLayoutUserPanesItem is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskTranscodeParamsLayoutUserPanesItem struct {
	BackgroundImageUrl string                                                     `name:"BackgroundImageUrl"`
	ZOrder             string                                                     `name:"ZOrder"`
	X                  string                                                     `name:"X"`
	Width              string                                                     `name:"Width"`
	Y                  string                                                     `name:"Y"`
	UserInfo           StartLiveMPUTaskTranscodeParamsLayoutUserPanesItemUserInfo `name:"UserInfo" type:"Struct"`
	RenderMode         string                                                     `name:"RenderMode"`
	Height             string                                                     `name:"Height"`
}

// StartLiveMPUTaskTranscodeParamsLayoutMaxVideoUser is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskTranscodeParamsLayoutMaxVideoUser struct {
	StreamType string `name:"StreamType"`
	SourceType string `name:"SourceType"`
	UserId     string `name:"UserId"`
}

// StartLiveMPUTaskTranscodeParamsLayoutUserPanesItemUserInfo is a repeated param struct in StartLiveMPUTaskRequest
type StartLiveMPUTaskTranscodeParamsLayoutUserPanesItemUserInfo struct {
	SourceType string `name:"SourceType"`
	UserId     string `name:"UserId"`
}

// StartLiveMPUTaskResponse is the response struct for api StartLiveMPUTask
type StartLiveMPUTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartLiveMPUTaskRequest creates a request to invoke StartLiveMPUTask API
func CreateStartLiveMPUTaskRequest() (request *StartLiveMPUTaskRequest) {
	request = &StartLiveMPUTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "StartLiveMPUTask", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartLiveMPUTaskResponse creates a response to parse from StartLiveMPUTask response
func CreateStartLiveMPUTaskResponse() (response *StartLiveMPUTaskResponse) {
	response = &StartLiveMPUTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
