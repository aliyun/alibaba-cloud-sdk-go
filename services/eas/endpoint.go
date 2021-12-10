package eas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "eas.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "eas.aliyuncs.com",
			"cn-beijing":                  "pai-eas.cn-beijing.aliyuncs.com",
			"cn-shanghai-inner":           "eas.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "eas.aliyuncs.com",
			"cn-north-2-gov-1":            "pai-eas.cn-north-2-gov-1.aliyuncs.com",
			"cn-yushanfang":               "eas.aliyuncs.com",
			"cn-qingdao-nebula":           "eas.aliyuncs.com",
			"cn-beijing-finance-pop":      "eas.aliyuncs.com",
			"cn-wuhan":                    "eas.aliyuncs.com",
			"cn-zhangjiakou":              "pai-eas.cn-zhangjiakou.aliyuncs.com",
			"us-west-1":                   "pai-eas.us-west-1.aliyuncs.com",
			"cn-zhangbei":                 "eas.aliyuncs.com",
			"rus-west-1-pop":              "eas.aliyuncs.com",
			"cn-shanghai-et15-b01":        "eas.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "eas.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "eas.aliyuncs.com",
			"ap-northeast-1":              "eas.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "eas.aliyuncs.com",
			"cn-shanghai-et2-b01":         "eas.aliyuncs.com",
			"ap-southeast-1":              "pai-eas.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":              "eas.aliyuncs.com",
			"ap-southeast-3":              "eas.aliyuncs.com",
			"ap-southeast-5":              "pai-eas.ap-southeast-5.aliyuncs.com",
			"us-east-1":                   "pai-eas.us-east-1.aliyuncs.com",
			"cn-shenzhen-inner":           "eas.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "eas.aliyuncs.com",
			"cn-beijing-gov-1":            "eas.aliyuncs.com",
			"cn-wulanchabu":               "eas.aliyuncs.com",
			"ap-south-1":                  "pai-eas.ap-south-1.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "eas.aliyuncs.com",
			"cn-haidian-cm12-c01":         "eas.aliyuncs.com",
			"cn-qingdao":                  "eas.aliyuncs.com",
			"cn-hongkong-finance-pop":     "eas.aliyuncs.com",
			"cn-shanghai":                 "pai-eas.cn-shanghai.aliyuncs.com",
			"cn-shanghai-finance-1":       "pai-eas.cn-shanghai-finance-1.aliyuncs.com",
			"cn-hongkong":                 "pai-eas.cn-hongkong.aliyuncs.com",
			"eu-central-1":                "pai-eas.eu-central-1.aliyuncs.com",
			"cn-shenzhen":                 "pai-eas.cn-shenzhen.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "eas.aliyuncs.com",
			"eu-west-1":                   "eas.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "eas.aliyuncs.com",
			"eu-west-1-oxs":               "eas.aliyuncs.com",
			"cn-beijing-finance-1":        "eas.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "eas.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "eas.aliyuncs.com",
			"cn-shenzhen-finance-1":       "eas.aliyuncs.com",
			"me-east-1":                   "eas.aliyuncs.com",
			"cn-chengdu":                  "pai-eas.cn-chengdu.aliyuncs.com",
			"cn-hangzhou-test-306":        "eas.aliyuncs.com",
			"cn-hangzhou-finance":         "eas.aliyuncs.com",
			"cn-beijing-nu16-b01":         "eas.aliyuncs.com",
			"cn-edge-1":                   "eas.aliyuncs.com",
			"cn-huhehaote":                "eas.aliyuncs.com",
			"cn-fujian":                   "eas.aliyuncs.com",
			"ap-northeast-2-pop":          "eas.aliyuncs.com",
			"cn-hangzhou":                 "pai-eas.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
