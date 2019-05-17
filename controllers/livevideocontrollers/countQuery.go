package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 统计查询
type CountQueryController struct {
	beego.Controller
}

// @Title 获取在线流的推流数据
// @Description DescribeLiveStreamPushInfoList
// @Success 200 {object} models.User
// @router /delspil [get]
func (c *CountQueryController) DescribeLiveStreamPushInfoList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveStreamPushInfoListRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveStreamPushInfoList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询直播转码统计信息
// @Description DescribeLiveTranscodeDetailInfo
// @Param DayTime body String false "起始时间，北京时间， 格式：yyyymmdd。 注意：当前只支持查询近30天内某天的详细数据"
// @Success 200 {object} models.User
// @router /deltdi [get]
func (c *CountQueryController) DescribeLiveTranscodeDetailInfo() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveTranscodeDetailInfoRequest()
	DayTime := c.Input().Get("DayTime")
	params := `{\"DayTime\":`+ DayTime +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveTranscodeDetailInfo(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询分身份运营商播放汇总总数据
// @Description DescribeProIspPlaySumInfoList
// @Param StartTime body String false "起始时间，北京时间， 格式：yyyy-mm-dd HH:MM:SS"
// @Param EndTime body String false "结束时间，北京时间， 格式：yyyy-mm-dd HH:MM:SS。 注：EndTime 和 StartTime 只支持最近1天的数据查询"
// @Param StatType body String false "统计的类型，可选值包括”Province”，”Isp”"
// @Success 200 {object} models.User
// @router /deltdi [get]
func (c *CountQueryController) DescribeProIspPlaySumInfoList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeProIspPlaySumInfoListRequest()
	StartTime := c.Input().Get("StartTime")
	EndTime := c.Input().Get("EndTime")
	StatType := c.Input().Get("StatType")
	params := `{\"StartTime\":`+ StartTime +`,\"EndTime\":`+ EndTime +`,\"StatType\":`+ StatType +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeProIspPlaySumInfoList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 按省份运营商查询播放信息
// @Description DescribeProvinceIspPlayInfoList
// @Param StartTime body String false "起始时间，北京时间， 格式：yyyy-mm-dd HH:MM:SS"
// @Param EndTime body String false "结束时间，北京时间， 格式：yyyy-mm-dd HH:MM:SS。 注：EndTime 和 StartTime 只支持最近1天的数据查询"
// @Param Granularity body Integer false "支持如下粒度： 1：1分钟粒度（跨度不支持超过1天）"
// @Param StatType body String false "统计指标类型： “Bandwidth”：带宽 “FluxPerSecond”：平均流量 “Flux”：流量 “Request”：请求数 “Online”：并发连接数"
// @Success 200 {object} models.User
// @router /depipil [get]
func (c *CountQueryController) DescribeProvinceIspPlayInfoList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeProvinceIspPlayInfoListRequest()
	StartTime := c.Input().Get("StartTime")
	EndTime := c.Input().Get("EndTime")
	Granularity := c.Input().Get("Granularity")
	StatType := c.Input().Get("StatType")
	params := `{\"StartTime\":`+ StartTime +`,\"EndTime\":`+ EndTime +`,\"Granularity\":`+ Granularity +`,\"StatType\":`+ StatType +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeProvinceIspPlayInfoList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询所有流的流量数据
// @Description DescribeStreamDayPlayInfoList
// @Param StartTime body String false "日期， 格式：YYYY-mm-dd"
// @Success 200 {object} models.User
// @router /desdpil [get]
func (c *CountQueryController) DescribeStreamDayPlayInfoList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeStreamDayPlayInfoListRequest()
	DayTime := c.Input().Get("DayTime")
	params := `{\"DayTime\":`+ DayTime +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeStreamDayPlayInfoList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询留的播放信息列表
// @Description DescribeStreamPlayInfoList
// @Param StartTime body String false "开始时间，北京时间，格式为yyyy-mm-dd HH:MM:SS， 当前时间 和 开始时间 间隔不超过30天"
// @Param EndTime body String false "结束时间，北京时间，格式为yyyy-mm-dd HH:MM:SS， 结束时间 和 开始时间 必须在同一天内"
// @Success 200 {object} models.User
// @router /despil [get]
func (c *CountQueryController) DescribeStreamPlayInfoList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeStreamPlayInfoListRequest()
	StartTime := c.Input().Get("StartTime")
	EndTime := c.Input().Get("EndTime")
	params := `{\"StartTime\":`+ StartTime +`,\"EndTime\":`+ EndTime +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeStreamPlayInfoList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}