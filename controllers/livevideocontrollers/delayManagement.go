package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 延播管理
type DelayManagementController struct {
	beego.Controller
}

// @Title 延迟播放
// @Description AddDelayLiveStream
// @Param AppName body String false "应用名称"
// @Param DomainName body String false "您的加速域名"
// @Param StreamName body String false "流名称"
// @Param DelayTime body String false "延播时间，单位：秒，上限：600秒"
// @Success 200 {object} models.User
// @router /add [get]
func (d *DelayManagementController) AddDelayLiveStream() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewAddDelayLiveStreamRequest()
	AppName := d.Input().Get("AppName")
	DomainName := d.Input().Get("DomainName")
	StreamName := d.Input().Get("StreamName")
	DelayTime := d.Input().Get("DelayTime")
	params := `{\"AppName\":`+ AppName +`,\"DomainName\":`+ DomainName +`,\"StreamName\":`+ StreamName +`,\"DelayTime\":`+ DelayTime +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.AddDelayLiveStream(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 恢复延播
// @Description ResumeDelayLiveStream
// @Param AppName body String false "应用名称"
// @Param DomainName body String false "您的加速域名"
// @Param StreamName body String false "流名称"
// @Success 200 {object} models.User
// @router /resume [get]
func (d *DelayManagementController) ResumeDelayLiveStream() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewResumeDelayLiveStreamRequest()
	AppName := d.Input().Get("AppName")
	DomainName := d.Input().Get("DomainName")
	StreamName := d.Input().Get("StreamName")
	params := `{\"AppName\":`+ AppName +`,\"DomainName\":`+ DomainName +`,\"StreamName\":`+ StreamName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ResumeDelayLiveStream(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}