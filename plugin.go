package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dronestock/drone"
	"github.com/go-resty/resty/v2"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/structer"
)

type plugin struct {
	drone.Base

	// 密钥
	Token string `default:"${TOKEN}"`
	// 拥有者
	Owner string `default:"${OWNER}"`
	// 仓库
	Repo string `default:"${REPO}"`

	// 发布
	Release *release `default:"${RELEASE}"`

	// 版本
	Version string `default:"2022-11-28"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(p.release, drone.Name("发布")),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("release", p.Release),
	}
}

func (p *plugin) call(ctx context.Context, uri string, req any, rsp any, method gox.HttpMethod) (err error) {
	fields := gox.Fields[any]{
		field.New("uri", uri),
		field.New("req", req),
		field.New("rsp", rsp),
	}
	p.Debug("调用Github接口", fields...)

	http := p.http(ctx)
	http.SetBody(req).SetResult(rsp)

	response := new(resty.Response)
	url := p.apiUrl(uri)
	switch method {
	case gox.HttpMethodGet:
		response, err = http.Get(url)
	case gox.HttpMethodPost:
		response, err = http.Post(url)
	}
	if nil != err {
		p.Warn("调用Github出错", fields.Connect(field.Error(err))...)
	} else if response.IsError() {
		err = exc.NewException(response.StatusCode(), "调用Github返回错误", fields...)
		p.Warn("Github返回错误", fields.Connect(field.Error(err))...)
	}

	return
}

func (p *plugin) sendfile(ctx context.Context, uri string, req any, filepath string) (err error) {
	fields := gox.Fields[any]{
		field.New("uri", uri),
		field.New("filepath", filepath),
	}

	http := p.http(ctx)
	queries := make(map[string]string)
	if ce := structer.New().Map().From(req).To(&queries).Convert(); nil != ce {
		err = ce
		p.Warn("复制结构体出错", field.New("from", req), field.Error(ce))
	} else {
		http.SetQueryParams(queries)
	}
	if nil != err {
		return
	}

	if bytes, oe := os.ReadFile(filepath); nil != oe {
		err = oe
		p.Warn("打开文件出错", fields.Connect(field.Error(oe))...)
	} else {
		http.SetBody(bytes)
		p.Info("准备上传文件",fields.Connect(field.New("size", len(bytes)))...)
	}
	if nil != err {
		return
	}

	if hr, he := http.Post(p.uploadUrl(uri)); nil != he {
		err = he
		p.Warn("向Github上传文件出错", fields.Connect(field.Error(err))...)
	} else if hr.IsError() {
		err = exc.NewException(hr.StatusCode(), "Github返回错误", fields...)
		p.Warn("Github返回错误", fields.Connect(field.Error(err))...)
	} else {
		p.Debug("向Github上传文件成功", fields...)
	}

	return
}

func (p *plugin) http(ctx context.Context) (http *resty.Request) {
	http = p.Http()
	http.SetHeader("Accept", "application/vnd.github+json")
	http.SetHeader("Authorization", fmt.Sprintf("Bearer %s", p.Token))
	http.SetHeader("X-GitHub-Api-Version", p.Version)
	http.SetContext(ctx)

	return
}

func (p *plugin) apiUrl(uri string) string {
	return fmt.Sprintf("https://api.github.com/%s", uri)
}

func (p *plugin) uploadUrl(uri string) string {
	return fmt.Sprintf("https://uploads.github.com/%s", uri)
}
