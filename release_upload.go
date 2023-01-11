package github

import (
	"fmt"

	"github.com/goexl/gox"
)

func (r *release) uploads(plugin *plugin, id int64, assets []string) (err error) {
	// 创建附件
	req := new(gox.Object)
	rsp := new(assetUploadRsp)

	return
}

func (r *release) upload(plugin *plugin, id int64, assets string) (err error) {
	// 创建附件
	req := new(assetUploadReq)
	rsp := new(assetUploadRsp)
	uri := fmt.Sprintf("/repos/%s/%s/releases/%d/assets", plugin.Owner,plugin.Repo, id)

	return
}
