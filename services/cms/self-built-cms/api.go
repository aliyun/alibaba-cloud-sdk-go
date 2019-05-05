// This file is auto-generated, don't edit it
package self_cms

import "github.com/alibabacloud-go/tea/tea"

type CustomMetric struct {
	GroupId    int                    `json:"groupId"`
	MetricName string                 `json:"metricName"`
	Dimensions map[string]interface{} `json:"dimensions"`
	Time       string                 `json:"time"`
	Type       int                    `json:"type"`
	Period     int                    `json:"period"`
	Values     map[string]interface{} `json:"values"`
}

type CustomMetricResponse struct {
	RequestId string                 `json:"requestId"`
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data"`
}

type CustomEvent struct {
	GroupId int    `json:"groupId"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Name    string `json:"name"`
}

type CustomEventResponse struct {
	RequestId string                 `json:"requestId"`
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data"`
}

func (client *Client) UploadCustomMetric(metrics *[]CustomMetric) (result_ CustomMetricResponse, err error) {
	request_ := tea.NewRequest()
	request_.Protocol = client.Protocol__
	request_.Method = "POST"
	request_.Pathname = "/metric/custom/upload"
	request_.Headers = map[string]string{
		"content-type":      "application/json",
		"date":              client.GetGMTDateString__(),
		"host":              client.Endpoint__,
		"x-cms-signature":   "hmac-sha1",
		"x-cms-api-version": "1.0",
		"x-cms-ip":          client.GetIp__(),
		"user-agent":        client.UserAgent__,
	}
	request_.Body = client.Stringify__(metrics)
	request_.Headers["content-length"] = client.GetContentLength__(request_)
	request_.Headers["content-md5"] = client.GetContentMD5__(request_)
	request_.Headers["authorization"] = client.GetAuthorization__(request_)
	response_, err := tea.DoRequest(request_)
	if err != nil {
		return
	}

	if client.IsNotOk__(response_) {
		err = tea.NewSDKError(map[string]interface{}{
			"code": response_.StatusCode,
		})
		return
	}

	body, err := client.ReadJSON__(response_)
	if err != nil {
		return
	}
	if client.Is4xx__(body) {
		err = tea.NewSDKError(map[string]interface{}{
			"code":    body["code"],
			"message": body["msg"],
		})
		return
	}

	tea.Convert(map[string]interface{}{
		"requestId": body["requestId"],
		"code":      body["code"],
		"message":   body["msg"],
	}, &result_)
	return
}

func (client *Client) UploadCustomEvent(events *[]CustomEvent) (result_ CustomEventResponse, err error) {
	request_ := tea.NewRequest()
	request_.Protocol = client.Protocol__
	request_.Method = "POST"
	request_.Pathname = "/event/custom/upload"
	request_.Headers = map[string]string{
		"content-type":      "application/json",
		"date":              client.GetGMTDateString__(),
		"host":              client.Endpoint__,
		"x-cms-signature":   "hmac-sha1",
		"x-cms-api-version": "1.0",
		"x-cms-ip":          client.GetIp__(),
		"user-agent":        client.UserAgent__,
	}
	request_.Body = client.Stringify__(events)
	request_.Headers["content-length"] = client.GetContentLength__(request_)
	request_.Headers["content-md5"] = client.GetContentMD5__(request_)
	request_.Headers["authorization"] = client.GetAuthorization__(request_)
	response_, err := tea.DoRequest(request_)
	if err != nil {
		return
	}

	if client.IsNotOk__(response_) {
		err = tea.NewSDKError(map[string]interface{}{
			"code": response_.StatusCode,
		})
		return
	}

	body, err := client.ReadJSON__(response_)
	if err != nil {
		return
	}
	if client.Is4xx__(body) {
		err = tea.NewSDKError(map[string]interface{}{
			"code":    body["code"],
			"message": body["msg"],
		})
		return
	}

	tea.Convert(map[string]interface{}{
		"requestId": body["requestId"],
		"code":      body["code"],
		"message":   body["msg"],
	}, &result_)
	return
}
