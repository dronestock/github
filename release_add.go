package main

import (
	"context"
	"fmt"
	"os"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/structer"
)

func (r *release) add(plugin *plugin) (err error) {
	// 优先假定配置是文件，读取文件中的内容充当描述信息
	if bytes, re := os.ReadFile(r.Body); nil == re {
		r.Body = string(bytes)
	}

	req := new(githubReleaseCreateReq)
	rsp := new(githubRelease)
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
	if nil != r.Asset {
		r.Assets = append(r.Assets, r.Asset)
	}
	err = r.upload(plugin, rsp.Id, r.Assets)

	return
}
