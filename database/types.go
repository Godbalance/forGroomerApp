package database

type Product struct {
	Image string `json:"image"`
	Brand string `json:"name"`
	Size  string `json:"size"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
