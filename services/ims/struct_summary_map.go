package ims

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

// SummaryMap is a nested struct in ims response
type SummaryMap struct {
	MFADevices                          int `json:"MFADevices" xml:"MFADevices"`
	AccessKeysPerUserQuota              int `json:"AccessKeysPerUserQuota" xml:"AccessKeysPerUserQuota"`
	AttachedPoliciesPerGroupQuota       int `json:"AttachedPoliciesPerGroupQuota" xml:"AttachedPoliciesPerGroupQuota"`
	AttachedSystemPoliciesPerRoleQuota  int `json:"AttachedSystemPoliciesPerRoleQuota" xml:"AttachedSystemPoliciesPerRoleQuota"`
	AttachedPoliciesPerRoleQuota        int `json:"AttachedPoliciesPerRoleQuota" xml:"AttachedPoliciesPerRoleQuota"`
	GroupsPerUserQuota                  int `json:"GroupsPerUserQuota" xml:"GroupsPerUserQuota"`
	Roles                               int `json:"Roles" xml:"Roles"`
	PolicySizeQuota                     int `json:"PolicySizeQuota" xml:"PolicySizeQuota"`
	AttachedSystemPoliciesPerGroupQuota int `json:"AttachedSystemPoliciesPerGroupQuota" xml:"AttachedSystemPoliciesPerGroupQuota"`
	AttachedSystemPoliciesPerUserQuota  int `json:"AttachedSystemPoliciesPerUserQuota" xml:"AttachedSystemPoliciesPerUserQuota"`
	AttachedPoliciesPerUserQuota        int `json:"AttachedPoliciesPerUserQuota" xml:"AttachedPoliciesPerUserQuota"`
	GroupsQuota                         int `json:"GroupsQuota" xml:"GroupsQuota"`
	Groups                              int `json:"Groups" xml:"Groups"`
	PoliciesQuota                       int `json:"PoliciesQuota" xml:"PoliciesQuota"`
	VirtualMFADevicesQuota              int `json:"VirtualMFADevicesQuota" xml:"VirtualMFADevicesQuota"`
	VersionsPerPolicyQuota              int `json:"VersionsPerPolicyQuota" xml:"VersionsPerPolicyQuota"`
	RolesQuota                          int `json:"RolesQuota" xml:"RolesQuota"`
	UsersQuota                          int `json:"UsersQuota" xml:"UsersQuota"`
	Policies                            int `json:"Policies" xml:"Policies"`
	Users                               int `json:"Users" xml:"Users"`
	MFADevicesInUse                     int `json:"MFADevicesInUse" xml:"MFADevicesInUse"`
}
