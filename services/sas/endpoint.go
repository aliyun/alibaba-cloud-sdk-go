package sas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "tds.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "tds.aliyuncs.com",
			"cn-beijing":                  "tds.aliyuncs.com",
			"cn-shanghai-inner":           "tds.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "tds.aliyuncs.com",
			"cn-north-2-gov-1":            "tds.aliyuncs.com",
			"cn-yushanfang":               "tds.aliyuncs.com",
			"cn-qingdao-nebula":           "tds.aliyuncs.com",
			"cn-beijing-finance-pop":      "tds.aliyuncs.com",
			"cn-wuhan":                    "tds.aliyuncs.com",
			"cn-zhangjiakou":              "tds.aliyuncs.com",
			"us-west-1":                   "tds.aliyuncs.com",
			"cn-zhangbei":                 "tds.aliyuncs.com",
			"rus-west-1-pop":              "tds.aliyuncs.com",
			"cn-shanghai-et15-b01":        "tds.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "tds.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "tds.aliyuncs.com",
			"ap-northeast-1":              "tds.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "tds.aliyuncs.com",
			"cn-shanghai-et2-b01":         "tds.aliyuncs.com",
			"ap-southeast-1":              "tds.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":              "tds.aliyuncs.com",
			"ap-southeast-3":              "tds.ap-southeast-1.aliyuncs.com",
			"ap-southeast-5":              "tds.ap-southeast-1.aliyuncs.com",
			"ap-southeast-6":              "tds.ap-southeast-1.aliyuncs.com",
			"ap-southeast-7":              "tds.ap-southeast-1.aliyuncs.com",
			"me-central-1":                "tds.ap-southeast-1.aliyuncs.com",
			"us-east-1":                   "tds.aliyuncs.com",
			"cn-shenzhen-inner":           "tds.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "tds.aliyuncs.com",
			"cn-beijing-gov-1":            "tds.aliyuncs.com",
			"cn-wulanchabu":               "tds.aliyuncs.com",
			"ap-south-1":                  "tds.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "tds.aliyuncs.com",
			"cn-haidian-cm12-c01":         "tds.aliyuncs.com",
			"cn-qingdao":                  "tds.aliyuncs.com",
			"cn-hongkong-finance-pop":     "tds.aliyuncs.com",
			"cn-shanghai":                 "tds.aliyuncs.com",
			"cn-shanghai-finance-1":       "tds.aliyuncs.com",
			"cn-hongkong":                 "tds.aliyuncs.com",
			"eu-central-1":                "tds.aliyuncs.com",
			"cn-shenzhen":                 "tds.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "tds.aliyuncs.com",
			"eu-west-1":                   "tds.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "tds.aliyuncs.com",
			"eu-west-1-oxs":               "tds.aliyuncs.com",
			"cn-beijing-finance-1":        "tds.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "tds.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "tds.aliyuncs.com",
			"cn-shenzhen-finance-1":       "tds.aliyuncs.com",
			"me-east-1":                   "tds.aliyuncs.com",
			"cn-chengdu":                  "tds.aliyuncs.com",
			"cn-hangzhou-test-306":        "tds.aliyuncs.com",
			"cn-hangzhou-finance":         "tds.aliyuncs.com",
			"cn-beijing-nu16-b01":         "tds.aliyuncs.com",
			"cn-edge-1":                   "tds.aliyuncs.com",
			"cn-huhehaote":                "tds.aliyuncs.com",
			"cn-fujian":                   "tds.aliyuncs.com",
			"ap-northeast-2-pop":          "tds.ap-southeast-1.aliyuncs.com",
			"ap-northeast-2":              "tds.ap-southeast-1.aliyuncs.com",
			"cn-hangzhou":                 "tds.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
