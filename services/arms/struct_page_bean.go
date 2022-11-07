package arms

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

// PageBean is a nested struct in arms response
type PageBean struct {
	Page                    int64                `json:"Page" xml:"Page"`
	PageSize                int                  `json:"PageSize" xml:"PageSize"`
	TotalCount              int                  `json:"TotalCount" xml:"TotalCount"`
	Size                    int64                `json:"Size" xml:"Size"`
	PageNumber              int                  `json:"PageNumber" xml:"PageNumber"`
	Total                   int64                `json:"Total" xml:"Total"`
	EscalationPolicies      []EscalationPolicies `json:"EscalationPolicies" xml:"EscalationPolicies"`
	RetcodeApps             []RetcodeApp         `json:"RetcodeApps" xml:"RetcodeApps"`
	ListAlerts              []ListAlerts         `json:"ListAlerts" xml:"ListAlerts"`
	Event                   []EventItem          `json:"Event" xml:"Event"`
	NotificationPolicies    []Policies           `json:"NotificationPolicies" xml:"NotificationPolicies"`
	AlertContacts           []Contacts           `json:"AlertContacts" xml:"AlertContacts"`
	AlertRules              []AlertRules         `json:"AlertRules" xml:"AlertRules"`
	SilencePolicies         []Policies           `json:"SilencePolicies" xml:"SilencePolicies"`
	AlarmHistories          []AlarmHistory       `json:"AlarmHistories" xml:"AlarmHistories"`
	WebhookContacts         []WebhookContacts    `json:"WebhookContacts" xml:"WebhookContacts"`
	EventBridgeIntegrations []Integrations       `json:"EventBridgeIntegrations" xml:"EventBridgeIntegrations"`
	OnCallSchedules         []OnCallSchedules    `json:"OnCallSchedules" xml:"OnCallSchedules"`
	AlertContactGroups      []ContactGroups      `json:"AlertContactGroups" xml:"AlertContactGroups"`
	TraceApps               []TraceApp           `json:"TraceApps" xml:"TraceApps"`
	AlertIMRobots           []IMRobots           `json:"AlertIMRobots" xml:"AlertIMRobots"`
	Events                  []Events             `json:"Events" xml:"Events"`
	Contacts                []Contact            `json:"Contacts" xml:"Contacts"`
}
