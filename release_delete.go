package main

import (
	"context"
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/goexl/gox/http"
)

func (r *release) delete(plugin *plugin, id int64) (err error) {
	if 0 == id {
		return
	}

	uri := fmt.Sprintf("repos/%s/%s/releases/%d", plugin.Owner, plugin.Repo, id)
	if he := plugin.call(context.Background(), uri, nil, nil, http.MethodDelete); nil != he {
		err = he
		plugin.Warn("删除发布出错", field.New("release", r), field.Error(he))
	} else {
		plugin.Info("删除发布成功", field.New("release", r), field.New("id", id))
	}

	return
}
