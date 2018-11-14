package obs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryConsumeDetailRequest struct {
	*requests.RpcRequest
	CommodityCode string `position:"Query" name:"CommodityCode" required:"false"`
	ConsumePeriod string `position:"Query" name:"ConsumePeriod" required:"true"`
	StartTime string `position:"Query" name:"StartTime" required:"false"`
	EndTime string `position:"Query" name:"EndTime" required:"false"`
	ConsumeCategory string `position:"Query" name:"ConsumeCategory" required:"false"`
	OwnerAccount string `position:"Query" name:"OwnerAccount" required:"false"`
	NeedDetail requests.Boolean `position:"Query" name:"NeedDetail" required:"false"`
	PageSize requests.Integer `position:"Query" name:"PageSize" required:"false"`
	PageMinId requests.Integer `position:"Query" name:"PageMinId" required:"false"`

}

type QueryConsumeDetailResponse struct {

	*responses.BaseResponse
	TotalCount int `json:"TotalCount" xml:"TotalCount"`
	NextPageMinId int `json:"NextPageMinId" xml:"NextPageMinId"`
	PageSize int `json:"PageSize" xml:"PageSize"`
	ConsumeList ConsumeInstance `json:"ConsumeList" xml:"ConsumeList"`

}

func CreateQueryConsumeDetailRequest() (request *QueryConsumeDetailRequest)  {
	request = &QueryConsumeDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}

	request.InitWithApiInfo("Obs", "2016-11-11", "QueryConsumeDetail", "obs", "openAPI")

	return
}

func CreateQueryConsumeDetailResponse()  (response *QueryConsumeDetailResponse){
	response = &QueryConsumeDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}


func (client *Client) QueryConsumeDetail(request *QueryConsumeDetailRequest) (response *QueryConsumeDetailResponse, err error){
	response = CreateQueryConsumeDetailResponse()
	err = client.DoAction(request, response)
	return
}
