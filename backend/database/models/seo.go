package models


type IndexSeo struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Logo  string `json:"logo"`
	Image string `json:"image"`
	About string `json:"about"`
}
