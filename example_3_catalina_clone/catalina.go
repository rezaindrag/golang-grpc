package catalina

type Image struct {
	URL     string `json:"id"`
	Height  int32  `json:"height"`
	Width   int32  `json:"width"`
	Mime    string `json:"mime"`
	Caption string `json:"caption"`
}
