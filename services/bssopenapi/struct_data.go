package bssopenapi

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

// Data is a nested struct in bssopenapi response
type Data struct {
	RefundAmount                string                                                  `json:"RefundAmount" xml:"RefundAmount"`
	MaxResults                  int                                                     `json:"MaxResults" xml:"MaxResults"`
	CreditLimitStatus           string                                                  `json:"CreditLimitStatus" xml:"CreditLimitStatus"`
	Type                        string                                                  `json:"Type" xml:"Type"`
	ReleaseResult               bool                                                    `json:"ReleaseResult" xml:"ReleaseResult"`
	Id                          int64                                                   `json:"Id" xml:"Id"`
	Quantity                    int                                                     `json:"Quantity" xml:"Quantity"`
	MemberNickName              string                                                  `json:"MemberNickName" xml:"MemberNickName"`
	HostName                    string                                                  `json:"HostName" xml:"HostName"`
	TotalOutstandingAmount      float64                                                 `json:"TotalOutstandingAmount" xml:"TotalOutstandingAmount"`
	Status                      string                                                  `json:"Status" xml:"Status"`
	NewInvoiceAmount            float64                                                 `json:"NewInvoiceAmount" xml:"NewInvoiceAmount"`
	Marker                      string                                                  `json:"Marker" xml:"Marker"`
	PageSize                    int                                                     `json:"PageSize" xml:"PageSize"`
	MybankCreditAmount          string                                                  `json:"MybankCreditAmount" xml:"MybankCreditAmount"`
	InstanceId                  string                                                  `json:"InstanceId" xml:"InstanceId"`
	OwnerUid                    int64                                                   `json:"OwnerUid" xml:"OwnerUid"`
	InvoiceApplyId              int64                                                   `json:"InvoiceApplyId" xml:"InvoiceApplyId"`
	RelationId                  int64                                                   `json:"RelationId" xml:"RelationId"`
	TotalInvoiceAmount          int64                                                   `json:"TotalInvoiceAmount" xml:"TotalInvoiceAmount"`
	PageNum                     int                                                     `json:"PageNum" xml:"PageNum"`
	LoginEmail                  string                                                  `json:"LoginEmail" xml:"LoginEmail"`
	DiscountPrice               float64                                                 `json:"DiscountPrice" xml:"DiscountPrice"`
	NextToken                   string                                                  `json:"NextToken" xml:"NextToken"`
	AccountName                 string                                                  `json:"AccountName" xml:"AccountName"`
	MemberGroupName             string                                                  `json:"MemberGroupName" xml:"MemberGroupName"`
	ToUnitUserId                int64                                                   `json:"ToUnitUserId" xml:"ToUnitUserId"`
	TradePrice                  float64                                                 `json:"TradePrice" xml:"TradePrice"`
	HostId                      string                                                  `json:"HostId" xml:"HostId"`
	TotalUnAppliedInvoiceAmount int64                                                   `json:"TotalUnAppliedInvoiceAmount" xml:"TotalUnAppliedInvoiceAmount"`
	OriginalPrice               float64                                                 `json:"OriginalPrice" xml:"OriginalPrice"`
	ToUnitId                    int64                                                   `json:"ToUnitId" xml:"ToUnitId"`
	EndTime                     int64                                                   `json:"EndTime" xml:"EndTime"`
	OrderId                     string                                                  `json:"OrderId" xml:"OrderId"`
	StartTime                   int64                                                   `json:"StartTime" xml:"StartTime"`
	HostingStatus               string                                                  `json:"HostingStatus" xml:"HostingStatus"`
	TotalCount                  int                                                     `json:"TotalCount" xml:"TotalCount"`
	BillingCycle                string                                                  `json:"BillingCycle" xml:"BillingCycle"`
	OutstandingAmount           float64                                                 `json:"OutstandingAmount" xml:"OutstandingAmount"`
	AvailableCashAmount         string                                                  `json:"AvailableCashAmount" xml:"AvailableCashAmount"`
	GmtModified                 int64                                                   `json:"GmtModified" xml:"GmtModified"`
	ParentUserId                int64                                                   `json:"ParentUserId" xml:"ParentUserId"`
	AccountType                 string                                                  `json:"AccountType" xml:"AccountType"`
	AvailableAmount             string                                                  `json:"AvailableAmount" xml:"AvailableAmount"`
	UnitId                      int64                                                   `json:"UnitId" xml:"UnitId"`
	CreditAmount                string                                                  `json:"CreditAmount" xml:"CreditAmount"`
	AccountID                   string                                                  `json:"AccountID" xml:"AccountID"`
	IsFinancialAccount          bool                                                    `json:"IsFinancialAccount" xml:"IsFinancialAccount"`
	ChildUserId                 int64                                                   `json:"ChildUserId" xml:"ChildUserId"`
	Boolean                     bool                                                    `json:"Boolean" xml:"Boolean"`
	IsCertified                 bool                                                    `json:"IsCertified" xml:"IsCertified"`
	IsSuccess                   bool                                                    `json:"IsSuccess" xml:"IsSuccess"`
	Currency                    string                                                  `json:"Currency" xml:"Currency"`
	MemberGroupId               int64                                                   `json:"MemberGroupId" xml:"MemberGroupId"`
	UserName                    string                                                  `json:"UserName" xml:"UserName"`
	Mpk                         int64                                                   `json:"Mpk" xml:"Mpk"`
	OmsData                     []map[string]interface{}                                `json:"OmsData" xml:"OmsData"`
	UidList                     []string                                                `json:"UidList" xml:"UidList"`
	TotalUsage                  TotalUsage                                              `json:"TotalUsage" xml:"TotalUsage"`
	TotalCoverage               TotalCoverage                                           `json:"TotalCoverage" xml:"TotalCoverage"`
	CostUnitStatisInfo          CostUnitStatisInfo                                      `json:"CostUnitStatisInfo" xml:"CostUnitStatisInfo"`
	CostUnit                    CostUnit                                                `json:"CostUnit" xml:"CostUnit"`
	Modules                     Modules                                                 `json:"Modules" xml:"Modules"`
	Items                       []Item                                                  `json:"Items" xml:"Items"`
	PeriodCoverage              []Item                                                  `json:"PeriodCoverage" xml:"PeriodCoverage"`
	CustomerInvoiceAddressList  CustomerInvoiceAddressList                              `json:"CustomerInvoiceAddressList" xml:"CustomerInvoiceAddressList"`
	ModuleDetails               ModuleDetailsInGetPayAsYouGoPrice                       `json:"ModuleDetails" xml:"ModuleDetails"`
	FinancialRelationInfoList   []FinancialRelationInfoListItem                         `json:"FinancialRelationInfoList" xml:"FinancialRelationInfoList"`
	AttributeList               AttributeList                                           `json:"AttributeList" xml:"AttributeList"`
	ResourceInstanceDtoList     []ResourceInstanceList                                  `json:"ResourceInstanceDtoList" xml:"ResourceInstanceDtoList"`
	DetailList                  DetailListInQueryDPUtilizationDetail                    `json:"DetailList" xml:"DetailList"`
	OrderList                   OrderListInQueryOrders                                  `json:"OrderList" xml:"OrderList"`
	AccountTransactionsList     AccountTransactionsListInQueryAccountTransactionDetails `json:"AccountTransactionsList" xml:"AccountTransactionsList"`
	CostUnitDtoList             []CostUnitDtoListItem                                   `json:"CostUnitDtoList" xml:"CostUnitDtoList"`
	EvaluateList                EvaluateList                                            `json:"EvaluateList" xml:"EvaluateList"`
	CustomerInvoiceList         CustomerInvoiceList                                     `json:"CustomerInvoiceList" xml:"CustomerInvoiceList"`
	InstanceList                []Instance                                              `json:"InstanceList" xml:"InstanceList"`
	Promotions                  Promotions                                              `json:"Promotions" xml:"Promotions"`
	PromotionDetails            PromotionDetailsInGetPayAsYouGoPrice                    `json:"PromotionDetails" xml:"PromotionDetails"`
	ProductList                 ProductList                                             `json:"ProductList" xml:"ProductList"`
	ResourcePackages            ResourcePackages                                        `json:"ResourcePackages" xml:"ResourcePackages"`
	ModuleList                  ModuleList                                              `json:"ModuleList" xml:"ModuleList"`
}
