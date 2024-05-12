package productRequest

type ProductGetFilter struct {
	Id          string
	Limit       int
	Offset      int
	Name        string
	IsAvailable interface{}
	Category    string
	Sku         string
	Price       string
	InStock     interface{}
	CreatedAt   string
}
