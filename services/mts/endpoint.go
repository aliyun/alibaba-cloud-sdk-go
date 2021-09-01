package mts

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "mts.aliyuncs.com",
			"cn-beijing-gov-1":            "mts.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "mts.aliyuncs.com",
			"cn-wulanchabu":               "mts.aliyuncs.com",
			"cn-shanghai-inner":           "mts.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "mts.aliyuncs.com",
			"cn-haidian-cm12-c01":         "mts.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "mts.aliyuncs.com",
			"cn-north-2-gov-1":            "mts.aliyuncs.com",
			"cn-yushanfang":               "mts.aliyuncs.com",
			"cn-hongkong-finance-pop":     "mts.aliyuncs.com",
			"cn-qingdao-nebula":           "mts.aliyuncs.com",
			"cn-shanghai-finance-1":       "mts.aliyuncs.com",
			"cn-beijing-finance-pop":      "mts.aliyuncs.com",
			"cn-wuhan":                    "mts.aliyuncs.com",
			"cn-zhangbei":                 "mts.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "mts.aliyuncs.com",
			"rus-west-1-pop":              "mts.aliyuncs.com",
			"cn-shanghai-et15-b01":        "mts.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "mts.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "mts.aliyuncs.com",
			"eu-west-1-oxs":               "mts.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "mts.aliyuncs.com",
			"cn-beijing-finance-1":        "mts.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "mts.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "mts.aliyuncs.com",
			"cn-shenzhen-finance-1":       "mts.aliyuncs.com",
			"me-east-1":                   "mts.aliyuncs.com",
			"cn-chengdu":                  "mts.aliyuncs.com",
			"cn-hangzhou-test-306":        "mts.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "mts.aliyuncs.com",
			"cn-shanghai-et2-b01":         "mts.aliyuncs.com",
			"cn-hangzhou-finance":         "mts.aliyuncs.com",
			"cn-beijing-nu16-b01":         "mts.aliyuncs.com",
			"cn-edge-1":                   "mts.aliyuncs.com",
			"ap-southeast-2":              "mts.aliyuncs.com",
			"ap-southeast-3":              "mts.aliyuncs.com",
			"cn-fujian":                   "mts.aliyuncs.com",
			"us-east-1":                   "mts.aliyuncs.com",
			"ap-northeast-2-pop":          "mts.aliyuncs.com",
			"cn-shenzhen-inner":           "mts.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "mts.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
