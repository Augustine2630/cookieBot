package response

type SpoonacularRecipeResponse struct {
	Results []struct {
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Image     string `json:"image"`
		ImageType string `json:"imageType"`
	} `json:"results"`
	Offset       int `json:"offset"`
	Number       int `json:"number"`
	TotalResults int `json:"totalResults"`
}
