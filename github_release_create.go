package main

type githubReleaseCreateReq struct {
	TagName                string `json:"tag_name,omitempty" struct:"tag"`
	TargetCommitish        string `json:"target_commitish,omitempty"  struct:"target"`
	Name                   string `json:"name,omitempty" struct:"name"`
	Body                   string `json:"body,omitempty" struct:"body"`
	Draft                  bool   `json:"draft,omitempty" struct:"draft"`
	Prerelease             bool   `json:"prerelease,omitempty" struct:"prerelease"`
	DiscussionCategoryName string `json:"discussion_category_name,omitempty" struct:"discussion"`
	GenerateReleaseNotes   bool   `json:"generate_release_notes,omitempty" struct:"notes"`
	MakeLatest             string `json:"make_latest,omitempty" struct:"latest" validate:"oneof=true false legacy"`
}
