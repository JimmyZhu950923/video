package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// 视频分类
type VideoClassificationController struct {
	beego.Controller
}

// @Title 创建分类
// @Description CreateClass
// @Param ParentId body Integer false "父类ID，一级分类填写 -1"
// @Param ClassName body String false "分类名称"
// @Success 200 {object} models.User
// @router /cc [get]
func (v *VideoClassificationController) CreateClass() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCreateClassRequest()
	ParentId := v.Input().Get("ParentId")
	ClassName := v.Input().Get("ClassName")
	params := `{\"ParentId\":`+ ParentId +`,\"ClassName\":`+ ClassName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateClass(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除分类
// @Description DeleteClass
// @Param ClassId body Integer false "分类ID"
// @Success 200 {object} models.User
// @router /dc [get]
func (v *VideoClassificationController) DeleteClass() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteClassRequest()
	ClassId := v.Input().Get("ClassId")
	params := `{\"ClassId\":`+ ClassId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteClass(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取所有分类
// @Description DescribeAllClass
// @Success 200 {object} models.User
// @router /deac [get]
func (v *VideoClassificationController) DescribeAllClass() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeAllClassRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeAllClass(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改分类
// @Description ModifyClass
// @Param ClassId body Integer false "分类ID"
// @Param ClassName body String false "分类名称"
// @Success 200 {object} models.User
// @router /mc [get]
func (v *VideoClassificationController) ModifyClass() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyClassRequest()
	ClassId := v.Input().Get("ClassId")
	ClassName := v.Input().Get("ClassName")
	params := `{\"ClassId\":`+ ClassId +`,\"ClassName\":`+ ClassName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyClass(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}