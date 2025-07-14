package schemas

type CreateIndexSeoRequest struct {
	Title    string `json:"title"`
	Logo     string `json:"logo"`
	Image    string `json:"image"`
	About    string `json:"about"`
	FipiLink string `json:"fipi_link"`
}

type EditIndexSeoRequest struct {
	Title    *string `json:"title"`
	Logo     *string `json:"logo"`
	Image    *string `json:"image"`
	About    *string `json:"about"`
	FipiLink *string `json:"fipi_link"`
}
