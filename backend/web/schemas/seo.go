package schemas


type CreateIndexSeoRequest struct {
	Title string `json:"title"`
	Logo string `json:"logo"`
	Image string `json:"image"`
	About string `json:"about"`
}


type EditIndexSeoRequest struct {
	Title *string `json:"title"`
	Logo *string `json:"logo"`
	Image *string `json:"image"`
	About *string `json:"about"`
}
