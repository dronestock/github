package github

import (
	"context"
	"fmt"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/structer"
)

type asset struct {
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Label    string `json:"label"`
}

func (a *asset) upload(plugin *plugin, id int64) (err error) {
	// 创建附件
	req := new(assetUploadReq)
	rsp := new(assetUploadRsp)
	ctx := context.Background()
	uri := fmt.Sprintf("/repos/%s/%s/releases/%d/assets", plugin.Owner, plugin.Repo, id)
	if ce := structer.New().Map().From(a).To(req).Convert(); nil != ce {
		err = ce
		plugin.Warn("复制结构体出错", field.New("from", a), field.Error(ce))
	} else if he := plugin.call(ctx, uri, req, rsp, gox.HttpMethodPost); nil != he {
		err = he
		plugin.Warn("创建附件出错", field.New("asset", a), field.Error(he))
	} else if se := plugin.sendfile(ctx, rsp.Url, a.Filename); nil != se {
		err = se
		plugin.Warn("上传附件出错", field.New("asset", a), field.Error(he))
	}

	return
}
