package csas

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

// Data is a nested struct in csas response
type Data struct {
	MfaConfigType       string           `json:"MfaConfigType" xml:"MfaConfigType"`
	FileHashMd5         string           `json:"FileHashMd5" xml:"FileHashMd5"`
	GetGroupUrl         string           `json:"GetGroupUrl" xml:"GetGroupUrl"`
	IdpMetadata         string           `json:"IdpMetadata" xml:"IdpMetadata"`
	ImageId             string           `json:"ImageId" xml:"ImageId"`
	WmInfoBytesB64      string           `json:"WmInfoBytesB64" xml:"WmInfoBytesB64"`
	PcLoginType         string           `json:"PcLoginType" xml:"PcLoginType"`
	MobileMfaConfigType string           `json:"MobileMfaConfigType" xml:"MobileMfaConfigType"`
	TaskStatus          string           `json:"TaskStatus" xml:"TaskStatus"`
	ImageUrlExp         int64            `json:"ImageUrlExp" xml:"ImageUrlExp"`
	InFileSize          string           `json:"InFileSize" xml:"InFileSize"`
	TaskId              string           `json:"TaskId" xml:"TaskId"`
	PolicyId            string           `json:"PolicyId" xml:"PolicyId"`
	ImageUrl            string           `json:"ImageUrl" xml:"ImageUrl"`
	Id                  string           `json:"Id" xml:"Id"`
	AccessKey           string           `json:"AccessKey" xml:"AccessKey"`
	MobileLoginType     string           `json:"MobileLoginType" xml:"MobileLoginType"`
	Type                string           `json:"Type" xml:"Type"`
	Email               string           `json:"Email" xml:"Email"`
	UserGroupId         string           `json:"UserGroupId" xml:"UserGroupId"`
	MultiIdpInfo        string           `json:"MultiIdpInfo" xml:"MultiIdpInfo"`
	WmType              string           `json:"WmType" xml:"WmType"`
	DepartmentId        string           `json:"DepartmentId" xml:"DepartmentId"`
	FileUrl             string           `json:"FileUrl" xml:"FileUrl"`
	Status              string           `json:"Status" xml:"Status"`
	WmInfoSize          int64            `json:"WmInfoSize" xml:"WmInfoSize"`
	WmInfoUint          int64            `json:"WmInfoUint" xml:"WmInfoUint"`
	Name                string           `json:"Name" xml:"Name"`
	OutFileHashMd5      string           `json:"OutFileHashMd5" xml:"OutFileHashMd5"`
	CreateTime          string           `json:"CreateTime" xml:"CreateTime"`
	MobileNumber        string           `json:"MobileNumber" xml:"MobileNumber"`
	FileUrlExp          string           `json:"FileUrlExp" xml:"FileUrlExp"`
	VerifyAesKey        string           `json:"VerifyAesKey" xml:"VerifyAesKey"`
	InFileHashMd5       string           `json:"InFileHashMd5" xml:"InFileHashMd5"`
	UserId              string           `json:"UserId" xml:"UserId"`
	IdpConfigId         string           `json:"IdpConfigId" xml:"IdpConfigId"`
	TotalNum            int64            `json:"TotalNum" xml:"TotalNum"`
	VerifyToken         string           `json:"VerifyToken" xml:"VerifyToken"`
	Description         string           `json:"Description" xml:"Description"`
	Username            string           `json:"Username" xml:"Username"`
	UpdateTime          string           `json:"UpdateTime" xml:"UpdateTime"`
	Filename            string           `json:"Filename" xml:"Filename"`
	VerifyUrl           string           `json:"VerifyUrl" xml:"VerifyUrl"`
	AccessKeySecret     string           `json:"AccessKeySecret" xml:"AccessKeySecret"`
	FileSize            string           `json:"FileSize" xml:"FileSize"`
	MetricName          string           `json:"MetricName" xml:"MetricName"`
	OutFileSize         int64            `json:"OutFileSize" xml:"OutFileSize"`
	Department          Department       `json:"Department" xml:"Department"`
	Datapoints          []Point          `json:"Datapoints" xml:"Datapoints"`
	DataList            []DataListItem   `json:"DataList" xml:"DataList"`
	Policies            []PoliciesItem   `json:"Policies" xml:"Policies"`
	UserGroups          []UserGroupsItem `json:"UserGroups" xml:"UserGroups"`
}
