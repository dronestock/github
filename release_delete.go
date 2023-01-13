package main

import (
	"context"
	"fmt"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (r *release) delete(plugin *plugin, id int64) (err error) {
	uri := fmt.Sprintf("repos/%s/%s/releases/%d", plugin.Owner, plugin.Repo, id)
	if he := plugin.call(context.Background(), uri, nil, nil, gox.HttpMethodDelete); nil != he {
		err = he
		plugin.Warn("删除发布出错", field.New("release", r), field.Error(he))
	} else {
		plugin.Info("删除发布成功", field.New("release", r), field.New("id", id))
	}

	return
}
