package dts

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "dts.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "dts.aliyuncs.com",
			"cn-beijing":                  "dts.aliyuncs.com",
			"cn-shanghai-inner":           "dts.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "dts.aliyuncs.com",
			"cn-north-2-gov-1":            "dts.aliyuncs.com",
			"cn-yushanfang":               "dts.aliyuncs.com",
			"cn-qingdao-nebula":           "dts.aliyuncs.com",
			"cn-beijing-finance-pop":      "dts.aliyuncs.com",
			"cn-wuhan":                    "dts.aliyuncs.com",
			"cn-zhangjiakou":              "dts.aliyuncs.com",
			"us-west-1":                   "dts.aliyuncs.com",
			"rus-west-1-pop":              "dts.aliyuncs.com",
			"cn-shanghai-et15-b01":        "dts.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "dts.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "dts.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "dts.aliyuncs.com",
			"cn-shanghai-et2-b01":         "dts.aliyuncs.com",
			"ap-southeast-1":              "dts.aliyuncs.com",
			"ap-southeast-2":              "dts.aliyuncs.com",
			"ap-southeast-3":              "dts.aliyuncs.com",
			"ap-southeast-5":              "dts.aliyuncs.com",
			"us-east-1":                   "dts.aliyuncs.com",
			"cn-shenzhen-inner":           "dts.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "dts.aliyuncs.com",
			"cn-beijing-gov-1":            "dts.aliyuncs.com",
			"cn-wulanchabu":               "dts.aliyuncs.com",
			"ap-south-1":                  "dts.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "dts.aliyuncs.com",
			"cn-haidian-cm12-c01":         "dts.aliyuncs.com",
			"cn-qingdao":                  "dts.aliyuncs.com",
			"cn-hongkong-finance-pop":     "dts.aliyuncs.com",
			"cn-shanghai":                 "dts.aliyuncs.com",
			"cn-shanghai-finance-1":       "dts.aliyuncs.com",
			"cn-hongkong":                 "dts.aliyuncs.com",
			"eu-central-1":                "dts.aliyuncs.com",
			"cn-shenzhen":                 "dts.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "dts.aliyuncs.com",
			"eu-west-1":                   "dts.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "dts.aliyuncs.com",
			"eu-west-1-oxs":               "dts.aliyuncs.com",
			"cn-beijing-finance-1":        "dts.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "dts.aliyuncs.com",
			"cn-shenzhen-finance-1":       "dts.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "dts.aliyuncs.com",
			"me-east-1":                   "dts.aliyuncs.com",
			"cn-chengdu":                  "dts.aliyuncs.com",
			"cn-hangzhou-test-306":        "dts.aliyuncs.com",
			"cn-hangzhou-finance":         "dts.aliyuncs.com",
			"cn-beijing-nu16-b01":         "dts.aliyuncs.com",
			"cn-edge-1":                   "dts.aliyuncs.com",
			"cn-huhehaote":                "dts.aliyuncs.com",
			"cn-fujian":                   "dts.aliyuncs.com",
			"ap-northeast-2-pop":          "dts.aliyuncs.com",
			"cn-hangzhou":                 "dts.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
