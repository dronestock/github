package main

import (
	"context"
	"fmt"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (r *release) check(plugin *plugin) (err error) {
	uri := fmt.Sprintf("repos/%s/%s/releases/tags/%s", plugin.Owner, plugin.Repo, r.Tag)
	rsp := new(githubRelease)
	if he := plugin.call(context.Background(), uri, nil, rsp, gox.HttpMethodGet); nil != he {
		err = he
		plugin.Warn("检索发布出错", field.New("release", r), field.Error(he))
	} else if de := r.delete(plugin, rsp.Id); nil != de {
		err = de
		plugin.Warn("删除发布出错", field.New("release", r), field.Error(de))
	}

	return
}
