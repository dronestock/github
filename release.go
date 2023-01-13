package main

type release struct {
	// 标签
	Tag string `json:"tag" validate:"required" structer:"tag"`
	// 目标
	Target string `json:"target" structer:"target"`
	// 名称
	Name string `json:"name" structer:"name"`
	// 描述
	Body string `json:"body" structer:"body"`
	// 是否是草稿
	Draft bool `json:"draft" structer:"draft"`
	// 是否是预发布
	Prerelease bool `json:"prerelease" structer:"prerelease"`
	// 如果指定，将创建指定类别的讨论并将其链接到发布
	Discussion string `json:"discussion" structer:"discussion"`
	// 是否为本次发布自动生成名称和正文
	Notes bool `json:"notes" structer:"notes"`
	// 指定是否应将此版本设置为存储库的最新版本
	Latest string `default:"true" json:"latest" structer:"latest"`

	// 附件
	Asset *asset `json:"asset"`
	// 附件列表
	Assets []*asset `json:"assets"`
}

func (p *plugin) release() (undo bool, err error) {
	if undo = nil == p.Release; undo {
		return
	}

	// 创建发布
	err = p.Release.publish(p)

	return
}
