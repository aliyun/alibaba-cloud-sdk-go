package ocr

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

// DataInRecognizeMalaysiaIdentityCard is a nested struct in ocr response
type DataInRecognizeMalaysiaIdentityCard struct {
	AddressFirstLine  AddressFirstLine  `json:"AddressFirstLine" xml:"AddressFirstLine"`
	AddressSecondLine AddressSecondLine `json:"AddressSecondLine" xml:"AddressSecondLine"`
	AddressThirdLine  AddressThirdLine  `json:"AddressThirdLine" xml:"AddressThirdLine"`
	AddressFourthLine AddressFourthLine `json:"AddressFourthLine" xml:"AddressFourthLine"`
	AddressFifthLine  AddressFifthLine  `json:"AddressFifthLine" xml:"AddressFifthLine"`
	CardBox           CardBox           `json:"CardBox" xml:"CardBox"`
	DriveClass        DriveClass        `json:"DriveClass" xml:"DriveClass"`
	ExpiryDate        ExpiryDate        `json:"ExpiryDate" xml:"ExpiryDate"`
	IdNumber          IdNumber          `json:"IdNumber" xml:"IdNumber"`
	IssueDate         IssueDate         `json:"IssueDate" xml:"IssueDate"`
	NameFirstLine     NameFirstLine     `json:"NameFirstLine" xml:"NameFirstLine"`
	NameSecondLine    NameSecondLine    `json:"NameSecondLine" xml:"NameSecondLine"`
	Nationality       Nationality       `json:"Nationality" xml:"Nationality"`
	PortraitBox       PortraitBox       `json:"PortraitBox" xml:"PortraitBox"`
	Sex               Sex               `json:"Sex" xml:"Sex"`
}
