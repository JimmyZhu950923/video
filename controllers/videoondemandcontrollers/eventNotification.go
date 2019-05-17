package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// AI样本管理
type EventNotificationController struct {
	beego.Controller
}

// @Title 确认时间通知
// @Description ConfirmEvents
// @Param EventHandles body string false "事件句柄，数组长度限制：16"
// @Success 200 {object} models.User
// @router /ce [get]
func (e *EventNotificationController) ConfirmEvents() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewConfirmEventsRequest()
	EventHandles := e.Input().Get("EventHandles")
	params := `{\"EventHandles\":[`+ EventHandles +`]}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ConfirmEvents(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 拉去时间通知
// @Description PullEvents
// @Success 200 {object} models.User
// @router /pe [get]
func (e *EventNotificationController) PullEvents() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewPullEventsRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.PullEvents(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
