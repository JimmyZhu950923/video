package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 直播拉流
type LivePullStreamController struct {
	beego.Controller
}

// @Title 添加拉流配置
// @Description CreatePullStreamConfig
// @Param FromUrl body String false "源Url"
// @Param ToUrl body String false "目的Url，目前限制该目标地址为腾讯域名"
// @Param AreaId body Integer false "区域id,1-深圳,2-上海，3-天津,4-香港"
// @Param Ispld body Integer false "运营商id,1-电信,2-移动,3-联通,4-其他,AreaId为4的时候,IspId只能为其他"
// @Param StartTime body String false "开始时间。 使用UTC格式时间， 例如：2019-01-08T10:00:00Z"
// @Param EndTime body String false "结束时间，注意： 1. 结束时间必须大于开始时间； 2. 结束时间和开始时间必须大于当前时间； 3. 结束时间 和 开始时间 间隔必须小于七天。 使用UTC格式时间， 例如：2019-01-08T10:00:00Z"
// @Success 200 {object} models.User
// @router /create [get]
func (l *LiveCallbackController) CreatePullStreamConfig() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreatePullStreamConfigRequest()
	FromUrl := l.Input().Get("FromUrl")
	ToUrl := l.Input().Get("ToUrl")
	AreaId := l.Input().Get("AreaId")
	IspId := l.Input().Get("IspId")
	StartTime := l.Input().Get("StartTime")
	EndTime := l.Input().Get("EndTime")
	params := `{\"FromUrl\":`+ FromUrl +`,\"ToUrl\":`+ ToUrl +`,\"AreaId\":`+ AreaId +`,\"IspId\":`+ IspId +`,\"StartTime\":`+ StartTime +`,\"EndTime\":`+ EndTime +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreatePullStreamConfig(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除拉流配置
// @Description DeletePullStreamConfig
// @Param ConfigId body String false "配置Id"
// @Success 200 {object} models.User
// @router /delete [get]
func (l *LiveCallbackController) DeletePullStreamConfig() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeletePullStreamConfigRequest()
	ConfigId := l.Input().Get("ConfigId")
	params := `{\"ConfigId\":`+ ConfigId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeletePullStreamConfig(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询拉流配置
// @Description DescribePullStreamConfigs
// @Success 200 {object} models.User
// @router /de [get]
func (l *LiveCallbackController) DescribePullStreamConfigs() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribePullStreamConfigsRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribePullStreamConfigs(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 更新拉流配置
// @Description ModifyPullStreamConfig
// @Param ConfigId body String false "配置Id"
// @Success 200 {object} models.User
// @router /modify [get]
func (l *LiveCallbackController) ModifyPullStreamConfig() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyPullStreamConfigRequest()
	ConfigId := l.Input().Get("ConfigId")
	params := `{\"ConfigId\":`+ ConfigId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyPullStreamConfig(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改拉流配置状态
// @Description ModifyPullStreamConfig
// @Param ConfigIds body String false "配置Id列表"
// @Param Status body String false "目标状态。0无效，2正在运行，4暂停"
// @Success 200 {object} models.User
// @router /ms [get]
func (l *LiveCallbackController) ModifyPullStreamStatus() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyPullStreamStatusRequest()
	ConfigIds := l.Input().Get("ConfigIds")
	Status := l.Input().Get("Status")
	params := `{\"ConfigIds\":[`+ ConfigIds +`],\"Status\":`+ Status +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyPullStreamStatus(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
