package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// 参数模板
type ParamTemplateController struct {
	beego.Controller
}

// @Title 创建视屏内容分析模板
// @Description CreateAIAnalysisTemplate
// @Success 200 {object} models.User
// @router /caat [get]
func (p *ParamTemplateController) CreateAIAnalysisTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCreateAIAnalysisTemplateRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateAIAnalysisTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建视屏内容识别模板
// @Description CreateAIRecognitionTemplate
// @Success 200 {object} models.User
// @router /cart [get]
func (p *ParamTemplateController) CreateAIRecognitionTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCreateAIRecognitionTemplateRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateAIRecognitionTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建视屏内容审核模板
// @Description CreateContentReviewTemplate
// @Param ReviewWallSwitch body string false "审核结果是否进入审核墙（对审核结果进行人工复核）的开关"
// @Success 200 {object} models.User
// @router /ccrt [get]
func (p *ParamTemplateController) CreateContentReviewTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCreateContentReviewTemplateRequest()
	ReviewWallSwitch := p.Input().Get("ReviewWallSwitch")
	params := `{\"ReviewWallSwitch\":`+ ReviewWallSwitch +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateContentReviewTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建任务流模板
// @Description CreateProcedureTemplate
// @Param Name body string false "任务流名字（支持中文，不超过20个字）"
// @Success 200 {object} models.User
// @router /cpt [get]
func (p *ParamTemplateController) CreateProcedureTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCreateProcedureTemplateRequest()
	Name := p.Input().Get("Name")
	params := `{\"Name\":`+ Name +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateProcedureTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建转码模板
// @Description CreateTranscodeTemplate
// @Param Container body string false "封装格式"
// @Success 200 {object} models.User
// @router /ctt [get]
func (p *ParamTemplateController) CreateTranscodeTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCreateTranscodeTemplateRequest()
	Container := p.Input().Get("Container")
	params := `{\"Container\":`+ Container +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateTranscodeTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建水印模板
// @Description CreateWatermarkTemplate
// @Param Type body string false "水印类型 "
// @Success 200 {object} models.User
// @router /cwt [get]
func (p *ParamTemplateController) CreateWatermarkTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCreateWatermarkTemplateRequest()
	Type := p.Input().Get("Type")
	params := `{\"Type\":`+ Type +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateWatermarkTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除视屏内容分析模板
// @Description DeleteAIAnalysisTemplate
// @Param Definition body integer false "视频内容分析模板唯一标识"
// @Success 200 {object} models.User
// @router /daat [get]
func (p *ParamTemplateController) DeleteAIAnalysisTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteAIAnalysisTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteAIAnalysisTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除视屏内容识别模板
// @Description DeleteAIRecognitionTemplate
// @Param Definition body integer false "视频内容分析模板唯一标识"
// @Success 200 {object} models.User
// @router /dart [get]
func (p *ParamTemplateController) DeleteAIRecognitionTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteAIRecognitionTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteAIRecognitionTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除视屏内容审核模板
// @Description DeleteContentReviewTemplate
// @Param Definition body integer false "视频内容分析模板唯一标识"
// @Success 200 {object} models.User
// @router /dcrt [get]
func (p *ParamTemplateController) DeleteContentReviewTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteContentReviewTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteContentReviewTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除任务流模板
// @Description DeleteProcedureTemplate
// @Param Name body string false "任务流名称"
// @Success 200 {object} models.User
// @router /dpt [get]
func (p *ParamTemplateController) DeleteProcedureTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteProcedureTemplateRequest()
	Name := p.Input().Get("Name")
	params := `{\"Name\":`+ Name +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteProcedureTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除转码模板
// @Description DeleteTranscodeTemplate
// @Param Definition body integer false "转码模板唯一标识"
// @Success 200 {object} models.User
// @router /dtt [get]
func (p *ParamTemplateController) DeleteTranscodeTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteTranscodeTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteTranscodeTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除水印模板
// @Description DeleteWatermarkTemplate
// @Param Definition body integer false "水印模板唯一标识"
// @Success 200 {object} models.User
// @router /dwt [get]
func (p *ParamTemplateController) DeleteWatermarkTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteWatermarkTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteWatermarkTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取视屏内容分析模板列表
// @Description DescribeAIAnalysisTemplates
// @Success 200 {object} models.User
// @router /deaat [get]
func (p *ParamTemplateController) DescribeAIAnalysisTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeAIAnalysisTemplatesRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeAIAnalysisTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取视屏内容识别模板列表
// @Description DescribeAIRecognitionTemplates
// @Success 200 {object} models.User
// @router /deart [get]
func (p *ParamTemplateController) DescribeAIRecognitionTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeAIRecognitionTemplatesRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeAIRecognitionTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取视屏内容审核模板列表
// @Description DescribeContentReviewTemplates
// @Success 200 {object} models.User
// @router /decrt [get]
func (p *ParamTemplateController) DescribeContentReviewTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeContentReviewTemplatesRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeContentReviewTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取任务流模板列表
// @Description DescribeProcedureTemplates
// @Success 200 {object} models.User
// @router /dept [get]
func (p *ParamTemplateController) DescribeProcedureTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeProcedureTemplatesRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeProcedureTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取转码模板列表
// @Description DescribeTranscodeTemplates
// @Success 200 {object} models.User
// @router /dett [get]
func (p *ParamTemplateController) DescribeTranscodeTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeTranscodeTemplatesRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeTranscodeTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取水印模板列表
// @Description DescribeWatermarkTemplates
// @Success 200 {object} models.User
// @router /dett [get]
func (p *ParamTemplateController) DescribeWatermarkTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeWatermarkTemplatesRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeWatermarkTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改视频内容分析模板
// @Description ModifyAIAnalysisTemplate
// @Param Definition body integer false "视频内容分析模板唯一标识"
// @Success 200 {object} models.User
// @router /maat [get]
func (p *ParamTemplateController) ModifyAIAnalysisTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyAIAnalysisTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyAIAnalysisTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改视频内容识别模板
// @Description ModifyAIRecognitionTemplate
// @Param Definition body integer false "视频内容识别模板唯一标识"
// @Success 200 {object} models.User
// @router /mart [get]
func (p *ParamTemplateController) ModifyAIRecognitionTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyAIRecognitionTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyAIRecognitionTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改视屏内容审核模板
// @Description ModifyContentReviewTemplate
// @Param Definition body integer false "内容审核模板唯一标识"
// @Success 200 {object} models.User
// @router /mcrt [get]
func (p *ParamTemplateController) ModifyContentReviewTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyContentReviewTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyContentReviewTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改转码模板
// @Description ModifyTranscodeTemplate
// @Param Definition body integer false "转码模板唯一标识"
// @Success 200 {object} models.User
// @router /mtt [get]
func (p *ParamTemplateController) ModifyTranscodeTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyTranscodeTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyTranscodeTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改水印模板
// @Description ModifyWatermarkTemplate
// @Param Definition body integer false "水印模板唯一标识"
// @Success 200 {object} models.User
// @router /mwt [get]
func (p *ParamTemplateController) ModifyWatermarkTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyWatermarkTemplateRequest()
	Definition := p.Input().Get("Definition")
	params := `{\"Definition\":`+ Definition +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyWatermarkTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 重设任务流模板
// @Description ResetProcedureTemplate
// @Param Name body integer false "人物流名字"
// @Success 200 {object} models.User
// @router /rpt [get]
func (p *ParamTemplateController) ResetProcedureTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewResetProcedureTemplateRequest()
	Name := p.Input().Get("Name")
	params := `{\"Name\":`+ Name +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ResetProcedureTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}