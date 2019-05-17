package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// 任务管理管理
type TaskManagementController struct {
	beego.Controller
}

// @Title 查询任务详情
// @Description DescribeTaskDetail
// @Param TaskId body string false "视频处理任务的任务 ID"
// @Success 200 {object} models.User
// @router /dtd [get]
func (t *TaskManagementController) DescribeTaskDetail() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeTaskDetailRequest()
	TaskId := t.Input().Get("TaskId")
	params := `{\"TaskId\":`+ TaskId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeTaskDetail(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取任务列表
// @Description DescribeTasks
// @Param Status body string false "过滤条件：任务状态，可选值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）"
// @Success 200 {object} models.User
// @router /dtd [get]
func (t *TaskManagementController) DescribeTasks() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeTasksRequest()
	Status := t.Input().Get("Status")
	params := `{\"Status\":`+ Status +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeTasks(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}