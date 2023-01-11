package github

import (
	"github.com/goexl/gox"
)

func (r *release) uploads(plugin *plugin, id int64, assets []*asset) (err error) {
	// 创建附件
	req := new(gox.Object)
	rsp := new(assetUploadRsp)

	return
}
