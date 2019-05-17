package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// 媒资管理
type MediaController struct {
	beego.Controller
}

// @Title 删除媒体
// @Description DeleteMedia
// @Param FileId body string false "媒体文件的唯一标识"
// @Success 200 {object} models.User
// @router /dm [get]
func (m *MediaController) DeleteMedia() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteMediaRequest()
	FileId := m.Input().Get("FileId")
	params := `{\"FileId\":`+FileId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteMedia(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取媒体详细信息
// @Description DescribeMediaInfos
// @Param FileId body string false "媒体文件 ID 列表，N 从 0 开始取值，最大 19"
// @Success 200 {object} models.User
// @router /demi [get]
func (m *MediaController) DescribeMediaInfos() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeMediaInfosRequest()
	FileIds := m.Input().Get("FileIds")
	params := `{\"FileIds\":[`+ FileIds +`]}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeMediaInfos(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改媒体文件属性
// @Description ModifyMediaInfo
// @Param FileId body string false "媒体文件的唯一标识"
// @Success 200 {object} models.User
// @router /mmi [get]
func (m *MediaController) ModifyMediaInfo() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyMediaInfoRequest()
	FileId := m.Input().Get("FileId")
	params := `{\"FileId\":`+ FileId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyMediaInfo(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 搜索媒体信息
// @Description SearchMedia
// @Success 200 {object} models.User
// @router /sm [get]
func (m *MediaController) SearchMedia() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewSearchMediaRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.SearchMedia(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}