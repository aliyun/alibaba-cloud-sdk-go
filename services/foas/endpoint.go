package foas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "foas.aliyuncs.com",
			"cn-beijing-gov-1":            "foas.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "foas.aliyuncs.com",
			"ap-south-1":                  "foas.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-inner":           "foas.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "foas.aliyuncs.com",
			"cn-haidian-cm12-c01":         "foas.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "foas.aliyuncs.com",
			"cn-yushanfang":               "foas.aliyuncs.com",
			"cn-hongkong-finance-pop":     "foas.aliyuncs.com",
			"cn-qingdao-nebula":           "foas.aliyuncs.com",
			"cn-hongkong":                 "foas.aliyuncs.com",
			"cn-beijing-finance-pop":      "foas.aliyuncs.com",
			"cn-wuhan":                    "foas.aliyuncs.com",
			"eu-central-1":                "foas.ap-northeast-1.aliyuncs.com",
			"us-west-1":                   "foas.ap-northeast-1.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "foas.aliyuncs.com",
			"rus-west-1-pop":              "foas.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "foas.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "foas.aliyuncs.com",
			"eu-west-1":                   "foas.ap-northeast-1.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "foas.aliyuncs.com",
			"eu-west-1-oxs":               "foas.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "foas.aliyuncs.com",
			"cn-beijing-finance-1":        "foas.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "foas.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "foas.aliyuncs.com",
			"cn-shenzhen-finance-1":       "foas.aliyuncs.com",
			"me-east-1":                   "foas.ap-northeast-1.aliyuncs.com",
			"cn-chengdu":                  "foas.aliyuncs.com",
			"cn-hangzhou-test-306":        "foas.aliyuncs.com",
			"cn-shanghai-et2-b01":         "foas.aliyuncs.com",
			"cn-beijing-nu16-b01":         "foas.aliyuncs.com",
			"cn-edge-1":                   "foas.aliyuncs.com",
			"ap-southeast-2":              "foas.ap-northeast-1.aliyuncs.com",
			"cn-huhehaote":                "foas.aliyuncs.com",
			"ap-southeast-5":              "foas.ap-northeast-1.aliyuncs.com",
			"cn-fujian":                   "foas.aliyuncs.com",
			"us-east-1":                   "foas.ap-northeast-1.aliyuncs.com",
			"ap-northeast-2-pop":          "foas.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "foas.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "foas.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
