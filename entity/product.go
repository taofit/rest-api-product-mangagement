package entity

type Product struct {
	Id         string `json:id`
	Name       string `json:name`
	Stock      int    `json:stock`
	Created_at string `json:created_at`
}

type Product_price struct {
	Product_id string
	Price      float32
}
