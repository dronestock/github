package github

type release struct {
	// 标签
	Tag string `json:"tag" validate:"required"`
	// 目标
	Target string `json:"target"`
	// 名称
	Name string `json:"name"`
	// 描述
	Body string `json:"body"`
	// 是否是草稿
	Draft bool `json:"draft"`
	// 是否是预发布
	Prerelease bool `json:"prerelease"`
	// 如果指定，将创建指定类别的讨论并将其链接到发布
	Discussion string `json:"discussion"`
	// 是否为本次发布自动生成名称和正文
	Notes bool `json:"notes"`
	// 指定是否应将此版本设置为存储库的最新版本
	Latest string `json:"latest" validate:"oneof=true false legacy"`

	// 附件
	Asset *asset `json:"asset"`
	// 附件列表
	Assets []*asset `json:"assets"`
}

func (p *plugin) release() (undo bool, err error) {
	if undo = nil == p.Release; undo {
		return
	}

	return
}
