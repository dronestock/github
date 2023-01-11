package github

import (
	"context"
	"fmt"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/structer"
)

func (r *release) create(plugin *plugin) (err error) {
	// 创建发布
	req := new(releaseReq)
	rsp := new(releaseRsp)
	uri := fmt.Sprintf("repos/%s/%s/releases", plugin.Owner, plugin.Repo)
	if ce := structer.New().Map().From(r).To(req).Tag(copyTag).Convert(); nil != ce {
		err = ce
		plugin.Warn("复制结构体出错", field.New("from", r), field.Error(ce))
	} else if he := plugin.call(context.Background(), uri, req, rsp, gox.HttpMethodPost); nil != he {
		err = he
		plugin.Warn("创建发布出错", field.New("release", r), field.Error(he))
	}
	if nil != err {
		return
	}

	// 上传附件
	if "" != r.Asset {
		r.Assets = append(r.Assets, r.Asset)
	}
	err = r.uploads(plugin, r.Assets)

	return
}
