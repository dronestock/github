package main

import (
	"github.com/goexl/gox/field"
)

func (r *release) publish(plugin *plugin) (err error) {
	if ce := r.check(plugin); nil != ce { // 检查是否已经存在相应的发布
		err = ce
		plugin.Warn("检索发布出错", field.New("release", r), field.Error(ce))
	} else if ae := r.add(plugin); nil != ae { // 创建发布
		err = ae
		plugin.Warn("创建发布出错", field.New("release", r), field.Error(ae))
	}

	return
}
