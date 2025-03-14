package quickbi_public

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

// DataItem is a nested struct in quickbi_public response
type DataItem struct {
	WorksName               string     `json:"WorksName" xml:"WorksName"`
	DataSize                float64    `json:"DataSize" xml:"DataSize"`
	ModifyUserAccountName   string     `json:"ModifyUserAccountName" xml:"ModifyUserAccountName"`
	ViewCount               int64      `json:"ViewCount" xml:"ViewCount"`
	Auth3rdFlag             int        `json:"Auth3rdFlag" xml:"Auth3rdFlag"`
	ModifiedTime            string     `json:"ModifiedTime" xml:"ModifiedTime"`
	FavoriteDate            string     `json:"FavoriteDate" xml:"FavoriteDate"`
	GmtModify               string     `json:"GmtModify" xml:"GmtModify"`
	WorkType                string     `json:"WorkType" xml:"WorkType"`
	Favorite                bool       `json:"Favorite" xml:"Favorite"`
	StatusType              int        `json:"StatusType" xml:"StatusType"`
	OrganizationId          string     `json:"OrganizationId" xml:"OrganizationId"`
	Type                    string     `json:"Type" xml:"Type"`
	ShowName                string     `json:"ShowName" xml:"ShowName"`
	PublishStatus           int        `json:"PublishStatus" xml:"PublishStatus"`
	WorksType               string     `json:"WorksType" xml:"WorksType"`
	OwnerNum                string     `json:"OwnerNum" xml:"OwnerNum"`
	Email                   string     `json:"Email" xml:"Email"`
	OwnerName               string     `json:"OwnerName" xml:"OwnerName"`
	ModifyUser              string     `json:"ModifyUser" xml:"ModifyUser"`
	WorkspaceName           string     `json:"WorkspaceName" xml:"WorkspaceName"`
	JoinedDate              int64      `json:"JoinedDate" xml:"JoinedDate"`
	DateUpdateTime          string     `json:"DateUpdateTime" xml:"DateUpdateTime"`
	OpenOfflineAcceleration bool       `json:"OpenOfflineAcceleration" xml:"OpenOfflineAcceleration"`
	WorksId                 string     `json:"WorksId" xml:"WorksId"`
	RowLevel                bool       `json:"RowLevel" xml:"RowLevel"`
	Parameters              string     `json:"Parameters" xml:"Parameters"`
	AllowPublishOperation   bool       `json:"AllowPublishOperation" xml:"AllowPublishOperation"`
	AccountId               string     `json:"AccountId" xml:"AccountId"`
	SecurityLevel           string     `json:"SecurityLevel" xml:"SecurityLevel"`
	IsDeleted               bool       `json:"IsDeleted" xml:"IsDeleted"`
	WorkspaceId             string     `json:"WorkspaceId" xml:"WorkspaceId"`
	WorkspaceDescription    string     `json:"WorkspaceDescription" xml:"WorkspaceDescription"`
	Name                    string     `json:"Name" xml:"Name"`
	ModifyTime              string     `json:"ModifyTime" xml:"ModifyTime"`
	DatasetId               string     `json:"DatasetId" xml:"DatasetId"`
	CreateTime              string     `json:"CreateTime" xml:"CreateTime"`
	AuthAdminUser           bool       `json:"AuthAdminUser" xml:"AuthAdminUser"`
	TreeId                  string     `json:"TreeId" xml:"TreeId"`
	GmtCreate               string     `json:"GmtCreate" xml:"GmtCreate"`
	WorkName                string     `json:"WorkName" xml:"WorkName"`
	UserId                  string     `json:"UserId" xml:"UserId"`
	NickName                string     `json:"NickName" xml:"NickName"`
	UserType                int        `json:"UserType" xml:"UserType"`
	AllowShareOperation     bool       `json:"AllowShareOperation" xml:"AllowShareOperation"`
	CreateUser              string     `json:"CreateUser" xml:"CreateUser"`
	OwnerAccountName        string     `json:"OwnerAccountName" xml:"OwnerAccountName"`
	HasViewAuth             bool       `json:"HasViewAuth" xml:"HasViewAuth"`
	Description             string     `json:"Description" xml:"Description"`
	Phone                   string     `json:"Phone" xml:"Phone"`
	CreateUserAccountName   string     `json:"CreateUserAccountName" xml:"CreateUserAccountName"`
	JobId                   string     `json:"JobId" xml:"JobId"`
	Status                  int        `json:"Status" xml:"Status"`
	ModifyName              string     `json:"ModifyName" xml:"ModifyName"`
	DatasetName             string     `json:"DatasetName" xml:"DatasetName"`
	GmtModified             string     `json:"GmtModified" xml:"GmtModified"`
	Owner                   string     `json:"Owner" xml:"Owner"`
	EmbedTime               string     `json:"EmbedTime" xml:"EmbedTime"`
	Body                    string     `json:"Body" xml:"Body"`
	AdminUser               bool       `json:"AdminUser" xml:"AdminUser"`
	PublicFlag              bool       `json:"PublicFlag" xml:"PublicFlag"`
	ApiId                   string     `json:"ApiId" xml:"ApiId"`
	LatestViewTime          string     `json:"LatestViewTime" xml:"LatestViewTime"`
	AccountName             string     `json:"AccountName" xml:"AccountName"`
	OwnerId                 string     `json:"OwnerId" xml:"OwnerId"`
	PublicInvalidTime       int64      `json:"PublicInvalidTime" xml:"PublicInvalidTime"`
	HasEditAuth             bool       `json:"HasEditAuth" xml:"HasEditAuth"`
	LastLoginTime           int64      `json:"LastLoginTime" xml:"LastLoginTime"`
	RoleIdList              []int64    `json:"RoleIdList" xml:"RoleIdList"`
	DataSource              DataSource `json:"DataSource" xml:"DataSource"`
	Directory               Directory  `json:"Directory" xml:"Directory"`
	Role                    Role       `json:"Role" xml:"Role"`
}
