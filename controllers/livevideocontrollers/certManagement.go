package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 证书管理
type CertManagementController struct {
	beego.Controller
}

// @Title 域名绑定证书
// @Description BindLiveDomainCert
// @Param CertId body Integer false "证书Id"
// @Param DomainName body String false "播放域名"
// @Success 200 {object} models.User
// @router /bldc [get]
func (c *CertManagementController) BindLiveDomainCert() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewBindLiveDomainCertRequest()
	CertId := c.Input().Get("CertId")
	DomainName := c.Input().Get("DomainName")
	params := `{\"CertId\":`+ CertId +`,\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.BindLiveDomainCert(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 添加证书
// @Description CreateLiveCert
// @Param CertType body Integer false "证书类型。0-用户添加证书；1-腾讯云托管证书"
// @Param HttpsCrt body String false "证书内容，即公钥"
// @Param HttpsKey body String false "私钥"
// @Success 200 {object} models.User
// @router /clc [get]
func (c *CertManagementController) CreateLiveCert() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveCertRequest()
	CertType := c.Input().Get("CertType")
	HttpsCrt := c.Input().Get("HttpsCrt")
	HttpsKey := c.Input().Get("HttpsKey")
	params := `{\"CertType\":`+ CertType +`,\"HttpsCrt\":`+ HttpsCrt +`,\"HttpsKey\":`+ HttpsKey +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveCert(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除证书
// @Description CreateLiveCert
// @Param CertId body Integer false "证书ID"
// @Success 200 {object} models.User
// @router /dlc [get]
func (c *CertManagementController) DeleteLiveCert() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveCertRequest()
	CertId := c.Input().Get("CertId")
	params := `{\"CertId\":`+ CertId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveCert(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取证书信息
// @Description DescribeLiveCert
// @Param CertId body Integer false "证书ID"
// @Success 200 {object} models.User
// @router /delc [get]
func (c *CertManagementController) DescribeLiveCert() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveCertRequest()
	CertId := c.Input().Get("CertId")
	params := `{\"CertId\":`+ CertId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveCert(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取证书信息列表
// @Description DescribeLiveCerts
// @Success 200 {object} models.User
// @router /delcs [get]
func (c *CertManagementController) DescribeLiveCerts() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveCertsRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveCerts(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取域名证书信息
// @Description DescribeLiveDomainCert
// @Param DomainName body String false "播放域名"
// @Success 200 {object} models.User
// @router /deldc [get]
func (c *CertManagementController) DescribeLiveDomainCert() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveDomainCertRequest()
	DomainName := c.Input().Get("DomainName")
	params := `{\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveDomainCert(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改证书
// @Description ModifyLiveCert
// @Param CertId body Integer false "证书ID"
// @Success 200 {object} models.User
// @router /mlc [get]
func (c *CertManagementController) ModifyLiveCert() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLiveCertRequest()
	CertId := c.Input().Get("CertId")
	params := `{\"CertId\":`+ CertId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLiveCert(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改域名和证书绑定信息
// @Description ModifyLiveDomainCert
// @Param DomainName body String false "播放域名"
// @Success 200 {object} models.User
// @router /mldc [get]
func (c *CertManagementController) ModifyLiveDomainCert() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLiveDomainCertRequest()
	DomainName := c.Input().Get("DomainName")
	params := `{\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLiveDomainCert(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 接触域名证书
// @Description UnBindLiveDomainCert
// @Param DomainName body String false "播放域名"
// @Success 200 {object} models.User
// @router /ubldc [get]
func (c *CertManagementController) UnBindLiveDomainCert() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewUnBindLiveDomainCertRequest()
	DomainName := c.Input().Get("DomainName")
	params := `{\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.UnBindLiveDomainCert(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}