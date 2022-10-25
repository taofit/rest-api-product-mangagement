package entity

type Product struct {
	Id         string  `json:id`
	Name       string  `json:name`
	Stock      int     `json:stock`
	Created_at string  `json:created_at`
	Prices     []Price `json:prices`
}

type Product_price struct {
	Product_id string
	Price      float32
}

// Good post saying that never use float for money,
// https://husobee.github.io/money/float/2016/09/23/never-use-floats-for-currency.html
type Price struct {
	Amount   string `json:amount`
	Currency string `json:currency`
}
