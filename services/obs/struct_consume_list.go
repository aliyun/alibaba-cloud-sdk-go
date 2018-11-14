package obs

type ConsumeInstance struct {
	Instance []CInstance

}

type CInstance struct {
	ConsumrPeriod string
	BuyerId string
	OwnerId string
	CommodityCode string
	ConsumeCategory string
	BizId string
	ConsumeTime int64
	ServiceStartTime string
	ServiceEndTime string
	ServiceDuration string
	InstanceId string
	InstanceNick string
	Tag string
	Region string
	InternetIp string
	intranetIp string
	ConsumeAmount string
	FreeAmount string
	RequireAmount string
	PayStatus string
	CashPayAmount string
	CouponPayAmount string
	UnclearAmount string
	ConsumeDetail ConsumeDetailMap
}
