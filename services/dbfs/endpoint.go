package dbfs

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "dbfs.aliyuncs.com",
			"cn-beijing-gov-1":            "dbfs.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "dbfs.aliyuncs.com",
			"cn-wulanchabu":               "dbfs.aliyuncs.com",
			"cn-shanghai-inner":           "dbfs.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "dbfs.aliyuncs.com",
			"cn-haidian-cm12-c01":         "dbfs.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "dbfs.aliyuncs.com",
			"cn-north-2-gov-1":            "dbfs.aliyuncs.com",
			"cn-yushanfang":               "dbfs.aliyuncs.com",
			"cn-hongkong-finance-pop":     "dbfs.aliyuncs.com",
			"cn-qingdao-nebula":           "dbfs.aliyuncs.com",
			"cn-shanghai-finance-1":       "dbfs.aliyuncs.com",
			"cn-beijing-finance-pop":      "dbfs.aliyuncs.com",
			"cn-wuhan":                    "dbfs.aliyuncs.com",
			"cn-zhangbei":                 "dbfs.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "dbfs.aliyuncs.com",
			"rus-west-1-pop":              "dbfs.aliyuncs.com",
			"cn-shanghai-et15-b01":        "dbfs.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "dbfs.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "dbfs.aliyuncs.com",
			"eu-west-1-oxs":               "dbfs.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "dbfs.aliyuncs.com",
			"cn-beijing-finance-1":        "dbfs.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "dbfs.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "dbfs.aliyuncs.com",
			"cn-shenzhen-finance-1":       "dbfs.aliyuncs.com",
			"cn-hangzhou-test-306":        "dbfs.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "dbfs.aliyuncs.com",
			"cn-shanghai-et2-b01":         "dbfs.aliyuncs.com",
			"cn-hangzhou-finance":         "dbfs.aliyuncs.com",
			"cn-beijing-nu16-b01":         "dbfs.aliyuncs.com",
			"cn-edge-1":                   "dbfs.aliyuncs.com",
			"cn-fujian":                   "dbfs.aliyuncs.com",
			"ap-northeast-2-pop":          "dbfs.aliyuncs.com",
			"cn-shenzhen-inner":           "dbfs.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "dbfs.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
