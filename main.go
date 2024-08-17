package tencent

import (
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	cls "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls/v20201016"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
}

func NewClient(AccessKeyId, AccessKeySecret string) *Client {
	return &Client{
		AccessKeyId:     AccessKeyId,
		AccessKeySecret: AccessKeySecret,
	}
}

func (c *Client) CompleteCredential() *common.Credential {
	return common.NewCredential(c.AccessKeyId, c.AccessKeySecret)
}

func (c *Client) CDB(region string) (*cdb.Client, error) {
	cp := profile.NewClientProfile()
	cp.HttpProfile.Endpoint = "cdb.tencentcloudapi.com"
	return cdb.NewClient(c.CompleteCredential(), region, cp)
}

func (c *Client) CLS(region string) (*cls.Client, error) {
	cp := profile.NewClientProfile()
	cp.HttpProfile.Endpoint = "cls.tencentcloudapi.com"
	return cls.NewClient(c.CompleteCredential(), region, cp)
}

func (c *Client) CVM(region string) (*cvm.Client, error) {
	cp := profile.NewClientProfile()
	cp.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	return cvm.NewClient(c.CompleteCredential(), region, cp)
}

// "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
func (c *Client) DNSPod() (*dnspod.Client, error) {
	cp := profile.NewClientProfile()
	cp.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	return dnspod.NewClient(c.CompleteCredential(), "", cp)
}

func (c *Client) TKE(region string) (*tke.Client, error) {
	cp := profile.NewClientProfile()
	cp.HttpProfile.Endpoint = "tke.tencentcloudapi.com"
	return tke.NewClient(c.CompleteCredential(), region, cp)
}
