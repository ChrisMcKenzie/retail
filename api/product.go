package api

type Product struct {
	Data struct {
		Product struct {
			TCIN string `json:"tcin"`
			Item struct {
				ProductDescription struct {
					Title string `json:"title"`
				} `json:"product_description"`
			} `json:"item"`
		} `json:"product"`
	} `json:"data"`
}
