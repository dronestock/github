package main

import (
	"context"
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/goexl/structer"
)

type asset struct {
	File  string `json:"file"`
	Name  string `json:"name" validate:"required,ascii"`
	Label string `json:"label"`
}

func (a *asset) upload(plugin *plugin, id int64) (err error) {
	if "" == a.Label {
		a.Label = a.Name
	}

	// 创建附件
	req := new(assetUploadReq)
	ctx := context.Background()
	uri := fmt.Sprintf("repos/%s/%s/releases/%d/assets", plugin.Owner, plugin.Repo, id)

	if ce := structer.New().Map().From(a).To(req).Convert(); nil != ce {
		err = ce
		plugin.Warn("复制结构体出错", field.New("from", a), field.Error(ce))
	} else if se := plugin.sendfile(ctx, uri, req, a.File); nil != se {
		err = se
		plugin.Warn("上传附件出错", field.New("asset", a), field.Error(se))
	}

	return
}
