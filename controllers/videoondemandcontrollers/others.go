package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// 其他
type OthersController struct {
	beego.Controller
}

// @Title 直播即时剪辑
// @Description LiveRealTimeClip
// @Param StreamId body string false "推流[直播码]"
// @Param StartTime body string false "流剪辑的开始时间，格式参照 [ISO 日期格式说明]"
// @Param EndTime body string false "流剪辑的结束时间，格式参照 [ISO 日期格式说明]"
// @Success 200 {object} models.User
// @router /lrtc [get]
func (o *OthersController) LiveRealTimeClip() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewLiveRealTimeClipRequest()
	StreamId := o.Input().Get("StreamId")
	StartTime := o.Input().Get("StartTime")
	EndTime := o.Input().Get("EndTime")
	params := `{\"StreamId\":`+ StreamId +`,\"StartTime\":`+ StartTime +`,\"EndTime\":`+ EndTime +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.LiveRealTimeClip(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 简单HLS剪辑
// @Description SimpleHlsClip
// @Param Url body string false "需要裁剪的腾讯云点播 HLS 视频 URL"
// @Success 200 {object} models.User
// @router /shc [get]
func (o *OthersController) SimpleHlsClip() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewSimpleHlsClipRequest()
	Url := o.Input().Get("Url")
	params := `{\"Url\":`+ Url +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.SimpleHlsClip(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}