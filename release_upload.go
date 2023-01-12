package main

import (
	"github.com/goexl/gox/field"
)

func (r *release) upload(plugin *plugin, id int64, assets []*asset) (err error) {
	count := 0
	for _, _asset := range assets {
		if ue := _asset.upload(plugin, id); nil != ue {
			err = ue
			count += 1
			plugin.Warn("上传附件出错", field.New("id", id), field.New("asset", _asset), field.Error(ue))
		}
	}
	if count != len(assets) {
		err = nil
	}

	return
}
