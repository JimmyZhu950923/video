package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 直播流管理
type LiveStreamManagementController struct {
	beego.Controller
}

// @Title 获取禁推流列表
// @Description DescribeLiveForbidStreamList
// @Success 200 {object} models.User
// @router /dlfsl [get]
func (l *LiveStreamManagementController) DescribeLiveForbidStreamList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveForbidStreamListRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveForbidStreamList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询推断流事件
// @Description DescribeLiveForbidStreamList
// @Param StartTime body String false "起始时间。 UTC 格式，例如：2018-12-29T19:00:00Z。 支持查询60天内的历史记录"
// @Param EndTime body String false "结束时间。 UTC 格式，例如：2018-12-29T20:00:00Z。 不超过当前时间，且和起始时间相差不得超过30天"
// @Success 200 {object} models.User
// @router /dlsel [get]
func (l *LiveStreamManagementController) DescribeLiveStreamEventList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveStreamEventListRequest()
	StartTime := l.Input().Get("StartTime")
	EndTime := l.Input().Get("EndTime")
	params := `{\"StartTime\":`+ StartTime +`,\"EndTime\":`+ EndTime +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveStreamEventList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询在线推流信息列表
// @Description DescribeLiveForbidStreamList
// @Param PageNum body String false "取得第几页"
// @Param PageSize body String false "分页大小。 最大值：100。 取值范围：1~100 之前的任意整数。 默认值：10"
// @Success 200 {object} models.User
// @router /dlsoi [get]
func (l *LiveStreamManagementController) DescribeLiveStreamOnlineInfo() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveStreamOnlineInfoRequest()
	PageNum := l.Input().Get("PageNum")
	PageSize := l.Input().Get("PageSize")
	params := `{\"PageNum\":`+ PageNum +`,\"PageSize\":`+ PageSize +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveStreamOnlineInfo(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询直播中的流
// @Description DescribeLiveForbidStreamList
// @Success 200 {object} models.User
// @router /dlsol [get]
func (l *LiveStreamManagementController) DescribeLiveStreamOnlineList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveStreamOnlineListRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveStreamOnlineList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询推过流的流列表
// @Description DescribeLiveStreamPublishedList
// @Param DomainName body String false "您的域名"
// @Param StartTime body String false "起始时间。 UTC 格式，例如：2016-06-29T19:00:00Z。 和当前时间相隔不超过7天"
// @Param EndTime body String false "结束时间。 UTC 格式，例如：2016-06-30T19:00:00Z。 不超过当前时间"
// @Success 200 {object} models.User
// @router /dlspi [get]
func (l *LiveStreamManagementController) DescribeLiveStreamPublishedList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveStreamPublishedListRequest()
	DomainName := l.Input().Get("DomainName")
	StartTime := l.Input().Get("StartTime")
	EndTime := l.Input().Get("EndTime")
	params := `{\"DomainName\":`+ DomainName +`,\"StartTime\":`+ StartTime +`,\"EndTime\":`+ EndTime +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveStreamPublishedList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询流状态
// @Description DescribeLiveStreamState
// @Param AppName body String false "应用名称"
// @Param DomainName body String false "您的推流域名"
// @Param StreamName body String false "流名称"
// @Success 200 {object} models.User
// @router /dlss [get]
func (l *LiveStreamManagementController) DescribeLiveStreamState() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveStreamStateRequest()
	AppName := l.Input().Get("AppName")
	DomainName := l.Input().Get("DomainName")
	StreamName := l.Input().Get("StreamName")
	params := `{\"AppName\":`+ AppName +`,\"DomainName\":`+ DomainName +`,\"StreamName\":`+ StreamName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveStreamState(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 断开直播流
// @Description DropLiveStream
// @Param AppName body String false "应用名称"
// @Param DomainName body String false "您的加速域名"
// @Param StreamName body String false "流名称"
// @Success 200 {object} models.User
// @router /drop [get]
func (l *LiveStreamManagementController) DropLiveStream() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDropLiveStreamRequest()
	AppName := l.Input().Get("AppName")
	DomainName := l.Input().Get("DomainName")
	StreamName := l.Input().Get("StreamName")
	params := `{\"AppName\":`+ AppName +`,\"DomainName\":`+ DomainName +`,\"StreamName\":`+ StreamName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DropLiveStream(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 禁播直播流
// @Description DropLiveStream
// @Param AppName body String false "应用名称"
// @Param DomainName body String false "您的加速域名"
// @Param StreamName body String false "流名称"
// @Success 200 {object} models.User
// @router /forbid [get]
func (l *LiveStreamManagementController) ForbidLiveStream() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewForbidLiveStreamRequest()
	AppName := l.Input().Get("AppName")
	DomainName := l.Input().Get("DomainName")
	StreamName := l.Input().Get("StreamName")
	params := `{\"AppName\":`+ AppName +`,\"DomainName\":`+ DomainName +`,\"StreamName\":`+ StreamName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ForbidLiveStream(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 恢复直播流
// @Description ResumeLiveStream
// @Param AppName body String false "应用名称"
// @Param DomainName body String false "您的加速域名"
// @Param StreamName body String false "流名称"
// @Success 200 {object} models.User
// @router /resume [get]
func (l *LiveStreamManagementController) ResumeLiveStream() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewResumeLiveStreamRequest()
	AppName := l.Input().Get("AppName")
	DomainName := l.Input().Get("DomainName")
	StreamName := l.Input().Get("StreamName")
	params := `{\"AppName\":`+ AppName +`,\"DomainName\":`+ DomainName +`,\"StreamName\":`+ StreamName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ResumeLiveStream(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}