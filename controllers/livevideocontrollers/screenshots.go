package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 截图鉴黄
type ScreenshotsController struct {
	beego.Controller
}

// @Title 创建截图规则
// @Description CreateLiveSnapshotRule
// @Param DomainName body String false "推流域名"
// @Param AppName body String false "推流路径"
// @Param StreamName body String false "流名称"
// @Param TemplateId body Integer false "模板ID"
// @Success 200 {object} models.User
// @router /clsr [get]
func (s *ScreenshotsController) CreateLiveSnapshotRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveSnapshotRuleRequest()
	DomainName := s.Input().Get("DomainName")
	AppName := s.Input().Get("AppName")
	StreamName := s.Input().Get("StreamName")
	TemplateId := s.Input().Get("TemplateId")
	params := `{\"DomainName\":`+ DomainName +`,\"AppName\":`+ AppName +`,\"StreamName\":`+ StreamName +`,\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveSnapshotRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建截图模板
// @Description CreateLiveSnapshotTemplate
// @Param TemplateName body String false "模板域名"
// @Param CosAppId body Integer false "Cos AppId"
// @Param CosRegion body String false "Cos地区"
// @Param CosBucker body String false "Cos Bucket名称"
// @Success 200 {object} models.User
// @router /clst [get]
func (s *ScreenshotsController) CreateLiveSnapshotTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveSnapshotTemplateRequest()
	TemplateName := s.Input().Get("TemplateName")
	CosAppId := s.Input().Get("CosAppId")
	CosBucket := s.Input().Get("CosBucket")
	CosRegion := s.Input().Get("CosRegion")
	params := `{\"TemplateName\":`+ TemplateName +`,\"CosAppId\":`+ CosAppId +`,\"CosBucket\":`+ CosBucket +`,\"CosRegion\":`+ CosRegion +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveSnapshotTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除截图规则
// @Description CreateLiveSnapshotTemplate
// @Param DomainName body String false "推流域名"
// @Param AppName body String false "推流路径"
// @Param StreamName body String false "流名称"
// @Success 200 {object} models.User
// @router /dlsr [get]
func (s *ScreenshotsController) DeleteLiveSnapshotRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveSnapshotRuleRequest()
	DomainName := s.Input().Get("DomainName")
	AppName := s.Input().Get("AppName")
	StreamName := s.Input().Get("StreamName")
	params := `{\"DomainName\":`+ DomainName +`,\"AppName\":`+ AppName +`,\"StreamName\":`+ StreamName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveSnapshotRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除截图模板
// @Description CreateLiveSnapshotTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /dlst [get]
func (s *ScreenshotsController) DeleteLiveSnapshotTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveSnapshotTemplateRequest()
	TemplateId := s.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveSnapshotTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取截图规则列表
// @Description DescribeLiveSnapshotRules
// @Success 200 {object} models.User
// @router /delsr [get]
func (s *ScreenshotsController) DescribeLiveSnapshotRules() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveSnapshotRulesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveSnapshotRules(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取单个截图模板
// @Description DescribeLiveSnapshotTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /delst [get]
func (s *ScreenshotsController) DescribeLiveSnapshotTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveSnapshotTemplateRequest()
	TemplateId := s.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveSnapshotTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取截图模板列表
// @Description DescribeLiveSnapshotTemplate
// @Success 200 {object} models.User
// @router /delsts [get]
func (s *ScreenshotsController) DescribeLiveSnapshotTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveSnapshotTemplatesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveSnapshotTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改截图模板
// @Description ModifyLiveSnapshotTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /modify [get]
func (s *ScreenshotsController) ModifyLiveSnapshotTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLiveSnapshotTemplateRequest()
	TemplateId := s.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLiveSnapshotTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}