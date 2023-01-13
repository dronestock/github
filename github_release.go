package main

type githubRelease struct {
	Id        int64  `json:"id"`
	UploadUrl string `json:"upload_url"`
}
