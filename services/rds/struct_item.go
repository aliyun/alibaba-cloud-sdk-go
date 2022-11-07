package rds

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

// Item is a nested struct in rds response
type Item struct {
	Engine                    string           `json:"Engine" xml:"Engine"`
	ReportTime                string           `json:"ReportTime" xml:"ReportTime"`
	BackupSetStatus           int              `json:"BackupSetStatus" xml:"BackupSetStatus"`
	DBInstanceStatusDesc      string           `json:"DBInstanceStatusDesc" xml:"DBInstanceStatusDesc"`
	BackupEnabled             string           `json:"BackupEnabled" xml:"BackupEnabled"`
	CrossBackupRegion         string           `json:"CrossBackupRegion" xml:"CrossBackupRegion"`
	BackupStartTime           string           `json:"BackupStartTime" xml:"BackupStartTime"`
	DBInstanceId              string           `json:"DBInstanceId" xml:"DBInstanceId"`
	DBInstanceStorageType     string           `json:"DBInstanceStorageType" xml:"DBInstanceStorageType"`
	CrossLogBackupSize        int64            `json:"CrossLogBackupSize" xml:"CrossLogBackupSize"`
	BackupEnabledTime         string           `json:"BackupEnabledTime" xml:"BackupEnabledTime"`
	CrossBackupDownloadLink   string           `json:"CrossBackupDownloadLink" xml:"CrossBackupDownloadLink"`
	LogEndTime                string           `json:"LogEndTime" xml:"LogEndTime"`
	EngineVersion             string           `json:"EngineVersion" xml:"EngineVersion"`
	LogBackupEnabled          string           `json:"LogBackupEnabled" xml:"LogBackupEnabled"`
	HasBackupTableMeta        string           `json:"HasBackupTableMeta" xml:"HasBackupTableMeta"`
	LogBeginTime              string           `json:"LogBeginTime" xml:"LogBeginTime"`
	RetentType                int              `json:"RetentType" xml:"RetentType"`
	ConsistentTime            string           `json:"ConsistentTime" xml:"ConsistentTime"`
	Status                    string           `json:"Status" xml:"Status"`
	RelServiceId              string           `json:"RelServiceId" xml:"RelServiceId"`
	LockMode                  string           `json:"LockMode" xml:"LockMode"`
	CrossBackupSetSize        int64            `json:"CrossBackupSetSize" xml:"CrossBackupSetSize"`
	Category                  string           `json:"Category" xml:"Category"`
	DBInstanceDescription     string           `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	CrossLogBackupId          int              `json:"CrossLogBackupId" xml:"CrossLogBackupId"`
	LogFileName               string           `json:"LogFileName" xml:"LogFileName"`
	CrossDownloadLink         string           `json:"CrossDownloadLink" xml:"CrossDownloadLink"`
	RegionId                  string           `json:"RegionId" xml:"RegionId"`
	Duration                  int              `json:"Duration" xml:"Duration"`
	CrossBackupId             int              `json:"CrossBackupId" xml:"CrossBackupId"`
	LinkExpiredTime           string           `json:"LinkExpiredTime" xml:"LinkExpiredTime"`
	BackupSetScale            int              `json:"BackupSetScale" xml:"BackupSetScale"`
	AutoRenew                 string           `json:"AutoRenew" xml:"AutoRenew"`
	CrossBackupSetFile        string           `json:"CrossBackupSetFile" xml:"CrossBackupSetFile"`
	CrossBackupSetLocation    string           `json:"CrossBackupSetLocation" xml:"CrossBackupSetLocation"`
	LogBackupEnabledTime      string           `json:"LogBackupEnabledTime" xml:"LogBackupEnabledTime"`
	BackupEndTime             string           `json:"BackupEndTime" xml:"BackupEndTime"`
	BackupMethod              string           `json:"BackupMethod" xml:"BackupMethod"`
	RelService                string           `json:"RelService" xml:"RelService"`
	BackupType                string           `json:"BackupType" xml:"BackupType"`
	Retention                 int              `json:"Retention" xml:"Retention"`
	DBInstanceStatus          string           `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	CrossIntranetDownloadLink string           `json:"CrossIntranetDownloadLink" xml:"CrossIntranetDownloadLink"`
	InstanceId                int              `json:"InstanceId" xml:"InstanceId"`
	CrossBackupType           string           `json:"CrossBackupType" xml:"CrossBackupType"`
	RestoreRegions            RestoreRegions   `json:"RestoreRegions" xml:"RestoreRegions"`
	QPSTopNItems              QPSTopNItems     `json:"QPSTopNItems" xml:"QPSTopNItems"`
	LatencyTopNItems          LatencyTopNItems `json:"LatencyTopNItems" xml:"LatencyTopNItems"`
}
