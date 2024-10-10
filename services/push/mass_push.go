package push

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

// MassPush invokes the push.MassPush API synchronously
func (client *Client) MassPush(request *MassPushRequest) (response *MassPushResponse, err error) {
	response = CreateMassPushResponse()
	err = client.DoAction(request, response)
	return
}

// MassPushWithChan invokes the push.MassPush API asynchronously
func (client *Client) MassPushWithChan(request *MassPushRequest) (<-chan *MassPushResponse, <-chan error) {
	responseChan := make(chan *MassPushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MassPush(request)
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

// MassPushWithCallback invokes the push.MassPush API asynchronously
func (client *Client) MassPushWithCallback(request *MassPushRequest, callback func(response *MassPushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MassPushResponse
		var err error
		defer close(result)
		response, err = client.MassPush(request)
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

// MassPushRequest is the request struct for api MassPush
type MassPushRequest struct {
	*requests.RpcRequest
	PushTask *[]MassPushPushTask `position:"Body" name:"PushTask"  type:"Repeated"`
	AppKey   requests.Integer    `position:"Query" name:"AppKey"`
}

// MassPushPushTask is a repeated param struct in MassPushRequest
type MassPushPushTask struct {
	AndroidNotificationBarType       string `name:"AndroidNotificationBarType"`
	Body                             string `name:"Body"`
	DeviceType                       string `name:"DeviceType"`
	PushTime                         string `name:"PushTime"`
	SendSpeed                        string `name:"SendSpeed"`
	AndroidNotificationHuaweiChannel string `name:"AndroidNotificationHuaweiChannel"`
	AndroidPopupActivity             string `name:"AndroidPopupActivity"`
	HarmonyNotifyId                  string `name:"HarmonyNotifyId"`
	HarmonyRenderStyle               string `name:"HarmonyRenderStyle"`
	IOSRemindBody                    string `name:"iOSRemindBody"`
	Trim                             string `name:"Trim"`
	AndroidMessageVivoCategory       string `name:"AndroidMessageVivoCategory"`
	AndroidNotifyType                string `name:"AndroidNotifyType"`
	AndroidPopupTitle                string `name:"AndroidPopupTitle"`
	AndroidMessageHuaweiCategory     string `name:"AndroidMessageHuaweiCategory"`
	IOSMusic                         string `name:"iOSMusic"`
	IOSApnsEnv                       string `name:"iOSApnsEnv"`
	IOSMutableContent                string `name:"iOSMutableContent"`
	AndroidNotificationBarPriority   string `name:"AndroidNotificationBarPriority"`
	ExpireTime                       string `name:"ExpireTime"`
	AndroidImageUrl                  string `name:"AndroidImageUrl"`
	AndroidHonorTargetUserType       string `name:"AndroidHonorTargetUserType"`
	HarmonyRemindBody                string `name:"HarmonyRemindBody"`
	AndroidNotificationVivoChannel   string `name:"AndroidNotificationVivoChannel"`
	AndroidVivoReceiptId             string `name:"AndroidVivoReceiptId"`
	IOSNotificationCategory          string `name:"iOSNotificationCategory"`
	AndroidNotificationXiaomiChannel string `name:"AndroidNotificationXiaomiChannel"`
	HarmonyAction                    string `name:"HarmonyAction"`
	StoreOffline                     string `name:"StoreOffline"`
	IOSRelevanceScore                string `name:"iOSRelevanceScore"`
	AndroidVivoPushMode              string `name:"AndroidVivoPushMode"`
	AndroidInboxBody                 string `name:"AndroidInboxBody"`
	JobKey                           string `name:"JobKey"`
	HarmonyReceiptId                 string `name:"HarmonyReceiptId"`
	AndroidOpenUrl                   string `name:"AndroidOpenUrl"`
	AndroidBadgeSetNum               string `name:"AndroidBadgeSetNum"`
	AndroidXiaoMiNotifyBody          string `name:"AndroidXiaoMiNotifyBody"`
	IOSSubtitle                      string `name:"iOSSubtitle"`
	AndroidXiaomiBigPictureUrl       string `name:"AndroidXiaomiBigPictureUrl"`
	HarmonyCategory                  string `name:"HarmonyCategory"`
	IOSRemind                        string `name:"iOSRemind"`
	IOSNotificationThreadId          string `name:"iOSNotificationThreadId"`
	AndroidHuaweiTargetUserType      string `name:"AndroidHuaweiTargetUserType"`
	HarmonyRemind                    string `name:"HarmonyRemind"`
	AndroidMusic                     string `name:"AndroidMusic"`
	HarmonyExtensionPush             string `name:"HarmonyExtensionPush"`
	IOSNotificationCollapseId        string `name:"iOSNotificationCollapseId"`
	AndroidMessageHuaweiUrgency      string `name:"AndroidMessageHuaweiUrgency"`
	PushType                         string `name:"PushType"`
	IOSInterruptionLevel             string `name:"iOSInterruptionLevel"`
	HarmonyExtensionExtraData        string `name:"HarmonyExtensionExtraData"`
	AndroidExtParameters             string `name:"AndroidExtParameters"`
	HarmonyImageUrl                  string `name:"HarmonyImageUrl"`
	IOSBadge                         string `name:"iOSBadge"`
	AndroidBigBody                   string `name:"AndroidBigBody"`
	IOSBadgeAutoIncrement            string `name:"iOSBadgeAutoIncrement"`
	AndroidOpenType                  string `name:"AndroidOpenType"`
	HarmonyRemindTitle               string `name:"HarmonyRemindTitle"`
	Title                            string `name:"Title"`
	AndroidBadgeClass                string `name:"AndroidBadgeClass"`
	HarmonyBadgeAddNum               string `name:"HarmonyBadgeAddNum"`
	HarmonyTestMessage               string `name:"HarmonyTestMessage"`
	AndroidRenderStyle               string `name:"AndroidRenderStyle"`
	IOSExtParameters                 string `name:"iOSExtParameters"`
	AndroidBadgeAddNum               string `name:"AndroidBadgeAddNum"`
	AndroidHuaweiReceiptId           string `name:"AndroidHuaweiReceiptId"`
	AndroidNotificationHonorChannel  string `name:"AndroidNotificationHonorChannel"`
	AndroidXiaomiImageUrl            string `name:"AndroidXiaomiImageUrl"`
	AndroidTargetUserType            string `name:"AndroidTargetUserType"`
	HarmonyUri                       string `name:"HarmonyUri"`
	AndroidPopupBody                 string `name:"AndroidPopupBody"`
	HarmonyExtParameters             string `name:"HarmonyExtParameters"`
	AndroidBigPictureUrl             string `name:"AndroidBigPictureUrl"`
	IOSSilentNotification            string `name:"iOSSilentNotification"`
	AndroidNotificationGroup         string `name:"AndroidNotificationGroup"`
	SendChannels                     string `name:"SendChannels"`
	HarmonyActionType                string `name:"HarmonyActionType"`
	Target                           string `name:"Target"`
	HarmonyNotificationSlotType      string `name:"HarmonyNotificationSlotType"`
	AndroidBigTitle                  string `name:"AndroidBigTitle"`
	AndroidNotificationChannel       string `name:"AndroidNotificationChannel"`
	AndroidRemind                    string `name:"AndroidRemind"`
	HarmonyInboxContent              string `name:"HarmonyInboxContent"`
	AndroidActivity                  string `name:"AndroidActivity"`
	AndroidNotificationNotifyId      string `name:"AndroidNotificationNotifyId"`
	TargetValue                      string `name:"TargetValue"`
	HarmonyBadgeSetNum               string `name:"HarmonyBadgeSetNum"`
	AndroidXiaoMiNotifyTitle         string `name:"AndroidXiaoMiNotifyTitle"`
	AndroidXiaoMiActivity            string `name:"AndroidXiaoMiActivity"`
}

// MassPushResponse is the response struct for api MassPush
type MassPushResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	MessageIds MessageIds `json:"MessageIds" xml:"MessageIds"`
}

// CreateMassPushRequest creates a request to invoke MassPush API
func CreateMassPushRequest() (request *MassPushRequest) {
	request = &MassPushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "MassPush", "", "")
	request.Method = requests.POST
	return
}

// CreateMassPushResponse creates a response to parse from MassPush response
func CreateMassPushResponse() (response *MassPushResponse) {
	response = &MassPushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
