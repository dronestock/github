package main

type githubReleaseCreateReq struct {
	TagName                string `json:"tag_name,omitempty" structer:"tag"`
	TargetCommitish        string `json:"target_commitish,omitempty"  structer:"target"`
	Name                   string `json:"name,omitempty" structer:"name"`
	Body                   string `json:"body,omitempty" structer:"body"`
	Draft                  bool   `json:"draft,omitempty" structer:"draft"`
	Prerelease             bool   `json:"prerelease,omitempty" structer:"prerelease"`
	DiscussionCategoryName string `json:"discussion_category_name,omitempty" structer:"discussion"`
	GenerateReleaseNotes   bool   `json:"generate_release_notes,omitempty" structer:"notes"`
	MakeLatest             string `json:"make_latest,omitempty" structer:"latest" validate:"oneof=true false legacy"`
}
