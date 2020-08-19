package smarthosting

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "smarthosting.aliyuncs.com",
			"cn-beijing-gov-1":            "smarthosting.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "smarthosting.aliyuncs.com",
			"cn-beijing":                  "smarthosting.aliyuncs.com",
			"cn-shanghai-inner":           "smarthosting.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "smarthosting.aliyuncs.com",
			"cn-haidian-cm12-c01":         "smarthosting.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "smarthosting.aliyuncs.com",
			"cn-north-2-gov-1":            "smarthosting.aliyuncs.com",
			"cn-yushanfang":               "smarthosting.aliyuncs.com",
			"cn-qingdao":                  "smarthosting.aliyuncs.com",
			"cn-hongkong-finance-pop":     "smarthosting.aliyuncs.com",
			"cn-shanghai":                 "smarthosting.aliyuncs.com",
			"cn-shanghai-finance-1":       "smarthosting.aliyuncs.com",
			"cn-hongkong":                 "smarthosting.aliyuncs.com",
			"cn-beijing-finance-pop":      "smarthosting.aliyuncs.com",
			"cn-wuhan":                    "smarthosting.aliyuncs.com",
			"us-west-1":                   "smarthosting.aliyuncs.com",
			"cn-shenzhen":                 "smarthosting.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "smarthosting.aliyuncs.com",
			"rus-west-1-pop":              "smarthosting.aliyuncs.com",
			"cn-shanghai-et15-b01":        "smarthosting.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "smarthosting.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "smarthosting.aliyuncs.com",
			"eu-west-1-oxs":               "smarthosting.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "smarthosting.aliyuncs.com",
			"cn-beijing-finance-1":        "smarthosting.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "smarthosting.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "smarthosting.aliyuncs.com",
			"cn-shenzhen-finance-1":       "smarthosting.aliyuncs.com",
			"cn-hangzhou-test-306":        "smarthosting.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "smarthosting.aliyuncs.com",
			"cn-shanghai-et2-b01":         "smarthosting.aliyuncs.com",
			"cn-hangzhou-finance":         "smarthosting.aliyuncs.com",
			"ap-southeast-1":              "smarthosting.aliyuncs.com",
			"cn-beijing-nu16-b01":         "smarthosting.aliyuncs.com",
			"cn-edge-1":                   "smarthosting.aliyuncs.com",
			"cn-fujian":                   "smarthosting.aliyuncs.com",
			"us-east-1":                   "smarthosting.aliyuncs.com",
			"ap-northeast-2-pop":          "smarthosting.aliyuncs.com",
			"cn-shenzhen-inner":           "smarthosting.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "smarthosting.aliyuncs.com",
			"cn-hangzhou":                 "smarthosting.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
