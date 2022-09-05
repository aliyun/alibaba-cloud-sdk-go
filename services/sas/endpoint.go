package sas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "sas.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "sas.aliyuncs.com",
			"cn-beijing":                  "sas.aliyuncs.com",
			"cn-shanghai-inner":           "sas.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "sas.aliyuncs.com",
			"cn-north-2-gov-1":            "sas.aliyuncs.com",
			"cn-yushanfang":               "sas.aliyuncs.com",
			"cn-qingdao-nebula":           "sas.aliyuncs.com",
			"cn-beijing-finance-pop":      "sas.aliyuncs.com",
			"cn-wuhan":                    "sas.aliyuncs.com",
			"cn-zhangjiakou":              "sas.aliyuncs.com",
			"us-west-1":                   "sas.aliyuncs.com",
			"cn-zhangbei":                 "sas.aliyuncs.com",
			"rus-west-1-pop":              "sas.aliyuncs.com",
			"cn-shanghai-et15-b01":        "sas.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "sas.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "sas.aliyuncs.com",
			"ap-northeast-1":              "sas.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "sas.aliyuncs.com",
			"cn-shanghai-et2-b01":         "sas.aliyuncs.com",
			"ap-southeast-1":              "tds.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":              "sas.aliyuncs.com",
			"ap-southeast-3":              "tds.ap-southeast-3.aliyuncs.com",
			"ap-southeast-5":              "sas.aliyuncs.com",
			"us-east-1":                   "sas.aliyuncs.com",
			"cn-shenzhen-inner":           "sas.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "sas.aliyuncs.com",
			"cn-beijing-gov-1":            "sas.aliyuncs.com",
			"cn-wulanchabu":               "sas.aliyuncs.com",
			"ap-south-1":                  "sas.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "sas.aliyuncs.com",
			"cn-haidian-cm12-c01":         "sas.aliyuncs.com",
			"cn-qingdao":                  "sas.aliyuncs.com",
			"cn-hongkong-finance-pop":     "sas.aliyuncs.com",
			"cn-shanghai":                 "sas.aliyuncs.com",
			"cn-shanghai-finance-1":       "sas.aliyuncs.com",
			"cn-hongkong":                 "sas.aliyuncs.com",
			"eu-central-1":                "sas.aliyuncs.com",
			"cn-shenzhen":                 "sas.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "sas.aliyuncs.com",
			"eu-west-1":                   "sas.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "sas.aliyuncs.com",
			"eu-west-1-oxs":               "sas.aliyuncs.com",
			"cn-beijing-finance-1":        "sas.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "sas.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "sas.aliyuncs.com",
			"cn-shenzhen-finance-1":       "sas.aliyuncs.com",
			"me-east-1":                   "sas.aliyuncs.com",
			"cn-chengdu":                  "sas.aliyuncs.com",
			"cn-hangzhou-test-306":        "sas.aliyuncs.com",
			"cn-hangzhou-finance":         "sas.aliyuncs.com",
			"cn-beijing-nu16-b01":         "sas.aliyuncs.com",
			"cn-edge-1":                   "sas.aliyuncs.com",
			"cn-huhehaote":                "sas.aliyuncs.com",
			"cn-fujian":                   "sas.aliyuncs.com",
			"ap-northeast-2-pop":          "sas.aliyuncs.com",
			"cn-hangzhou":                 "tds.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
