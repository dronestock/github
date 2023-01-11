package github

type (
	assetUploadReq struct {
		Name string `json:"name"`
		Label string `json:"label"`
	}

	assetUploadRsp struct {
		Url string `json:"url"`
	}
)
