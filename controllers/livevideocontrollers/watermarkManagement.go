package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 水印管理
type WatermarkManagementController struct {
	beego.Controller
}

// @Title 添加水印
// @Description AddLiveWatermark
// @Param PictureUrl body String false "水印图片url"
// @Param WatermarkName body String false "水印名称"
// @Success 200 {object} models.User
// @router /add [get]
func (w *WatermarkManagementController) AddLiveWatermark() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewAddLiveWatermarkRequest()
	PictureUrl := w.Input().Get("PictureUrl")
	WatermarkName := w.Input().Get("WatermarkName")
	params := `{\"PictureUrl\":`+ PictureUrl +`,\"WatermarkName\":`+ WatermarkName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.AddLiveWatermark(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建水印规则
// @Description CreateLiveWatermarkRule
// @Param DomainName body String false "推流域名"
// @Param AppName body String false "推流路径"
// @Param StreamName body String false "流名称"
// @Param TemplateId body Integer false "水印ID"
// @Success 200 {object} models.User
// @router /create [get]
func (w *WatermarkManagementController) CreateLiveWatermarkRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveWatermarkRuleRequest()
	DomainName := w.Input().Get("DomainName")
	AppName := w.Input().Get("AppName")
	StreamName := w.Input().Get("StreamName")
	TemplateId := w.Input().Get("TemplateId")
	params := `{\"DomainName\":`+ DomainName +`,\"AppName\":`+ AppName +`,\"StreamName\":`+ StreamName +`,\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveWatermarkRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除水印
// @Description DeleteLiveWatermark
// @Param WatermarkId body Integer false "水印ID"
// @Success 200 {object} models.User
// @router /delete [get]
func (w *WatermarkManagementController) DeleteLiveWatermark() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveWatermarkRequest()
	WatermarkId := w.Input().Get("WatermarkId")
	params := `{\"WatermarkId\":`+ WatermarkId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveWatermark(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除水印规则
// @Description DeleteLiveWatermarkRule
// @Param DomainName body String false "推流域名"
// @Param AppName body String false "推流路径"
// @Param StreamName body String false "流名称"
// @Success 200 {object} models.User
// @router /dlwr [get]
func (w *WatermarkManagementController) DeleteLiveWatermarkRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveWatermarkRuleRequest()
	DomainName := w.Input().Get("DomainName")
	AppName := w.Input().Get("AppName")
	StreamName := w.Input().Get("StreamName")
	params := `{\"DomainName\":`+ DomainName +`,\"AppName\":`+ AppName +`,\"StreamName\":`+ StreamName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveWatermarkRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取单个水印
// @Description DescribeLiveWatermark
// @Param WatermarkId body Integer false "水印ID"
// @Success 200 {object} models.User
// @router /de [get]
func (w *WatermarkManagementController) DescribeLiveWatermark() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveWatermarkRequest()
	WatermarkId := w.Input().Get("WatermarkId")
	params := `{\"WatermarkId\":`+ WatermarkId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveWatermark(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取水印规则列表
// @Description DescribeLiveWatermarkRules
// @Success 200 {object} models.User
// @router /delwr [get]
func (w *WatermarkManagementController) DescribeLiveWatermarkRules() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveWatermarkRulesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveWatermarkRules(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询水印列表
// @Description DescribeLiveWatermarks
// @Success 200 {object} models.User
// @router /de [get]
func (w *WatermarkManagementController) DescribeLiveWatermarks() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveWatermarksRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveWatermarks(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 设置水印状态
// @Description SetLiveWatermarkStatus
// @Param WatermarkId body Integer false "水印ID"
// @Param Status body Integer false "状态。0：停用，1:启用"
// @Success 200 {object} models.User
// @router /set [get]
func (w *WatermarkManagementController) SetLiveWatermarkStatus() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewSetLiveWatermarkStatusRequest()
	WatermarkId := w.Input().Get("WatermarkId")
	Status := w.Input().Get("Status")
	params := `{\"WatermarkId\":`+ WatermarkId +`,\"Status\":`+ Status +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.SetLiveWatermarkStatus(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 更新水印
// @Description UpdateLiveWatermark
// @Param WatermarkId body Integer false "水印ID"
// @Param PictureUrl body String false "水印图片url"
// @Param XPosition body Integer false "显示位置，X轴偏移"
// @Param YPosition body Integer false "显示位置，Y轴偏移"
// @Success 200 {object} models.User
// @router /update [get]
func (w *WatermarkManagementController) UpdateLiveWatermark() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewUpdateLiveWatermarkRequest()
	WatermarkId := w.Input().Get("WatermarkId")
	PictureUrl := w.Input().Get("PictureUrl")
	XPosition := w.Input().Get("XPosition")
	YPosition := w.Input().Get("YPosition")
	params := `{\"WatermarkId\":`+ WatermarkId +`,\"PictureUrl\":`+ PictureUrl +`,\"XPosition\":`+ XPosition +`,\"YPosition\":`+ YPosition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.UpdateLiveWatermark(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}