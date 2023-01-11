package github

type (
	releaseReq struct {
		TagName                string `json:"tag_name" structer:"tag"`
		TargetCommitish        string `json:"target_commitish"  structer:"target"`
		Name                   string `json:"name"  structer:"name"`
		Body                   string `json:"body"  structer:"body"`
		Draft                  bool   `json:"draft"  structer:"draft"`
		Prerelease             bool   `json:"prerelease"  structer:"prerelease"`
		DiscussionCategoryName string `json:"discussion_category_name" structer:"discussion"`
		GenerateReleaseNotes   bool   `json:"generate_release_notes" structer:"notes"`
		MakeLatest             string `json:"make_latest" structer:"latest" validate:"oneof=true false legacy"`
	}

	releaseRsp struct {
		Id int64 `json:"id"`
	}
)
