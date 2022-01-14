package clickhouse

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "clickhouse.aliyuncs.com",
			"cn-beijing-gov-1":            "clickhouse.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "clickhouse.aliyuncs.com",
			"cn-beijing":                  "clickhouse.aliyuncs.com",
			"cn-shanghai-inner":           "clickhouse.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "clickhouse.aliyuncs.com",
			"cn-haidian-cm12-c01":         "clickhouse.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "clickhouse.aliyuncs.com",
			"cn-north-2-gov-1":            "clickhouse.aliyuncs.com",
			"cn-yushanfang":               "clickhouse.aliyuncs.com",
			"cn-qingdao":                  "clickhouse.aliyuncs.com",
			"cn-hongkong-finance-pop":     "clickhouse.aliyuncs.com",
			"cn-qingdao-nebula":           "clickhouse.aliyuncs.com",
			"cn-shanghai":                 "clickhouse.aliyuncs.com",
			"cn-shanghai-finance-1":       "clickhouse.aliyuncs.com",
			"cn-hongkong":                 "clickhouse.aliyuncs.com",
			"cn-beijing-finance-pop":      "clickhouse.aliyuncs.com",
			"cn-wuhan":                    "clickhouse.aliyuncs.com",
			"us-west-1":                   "clickhouse.aliyuncs.com",
			"cn-shenzhen":                 "clickhouse.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "clickhouse.aliyuncs.com",
			"rus-west-1-pop":              "clickhouse.aliyuncs.com",
			"cn-shanghai-et15-b01":        "clickhouse.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "clickhouse.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "clickhouse.aliyuncs.com",
			"eu-west-1-oxs":               "clickhouse.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "clickhouse.aliyuncs.com",
			"cn-beijing-finance-1":        "clickhouse.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "clickhouse.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "clickhouse.aliyuncs.com",
			"cn-shenzhen-finance-1":       "clickhouse.aliyuncs.com",
			"me-east-1":                   "clickhouse.aliyuncs.com",
			"cn-hangzhou-test-306":        "clickhouse.aliyuncs.com",
			"cn-shanghai-et2-b01":         "clickhouse.aliyuncs.com",
			"cn-hangzhou-finance":         "clickhouse.aliyuncs.com",
			"ap-southeast-1":              "clickhouse.aliyuncs.com",
			"cn-beijing-nu16-b01":         "clickhouse.aliyuncs.com",
			"cn-edge-1":                   "clickhouse.aliyuncs.com",
			"cn-fujian":                   "clickhouse.aliyuncs.com",
			"us-east-1":                   "clickhouse.aliyuncs.com",
			"ap-northeast-2-pop":          "clickhouse.aliyuncs.com",
			"cn-shenzhen-inner":           "clickhouse.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "clickhouse.aliyuncs.com",
			"cn-hangzhou":                 "clickhouse.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
