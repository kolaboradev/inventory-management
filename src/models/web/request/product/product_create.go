package productRequest

type ProductCreate struct {
	Name        string `validate:"required,min=1,max=30"`
	Sku         string `validate:"required,min=1,max=30"`
	Category    string `validate:"required,category"`
	ImageUrl    string `validate:"required"`
	Notes       string `validate:"required,min=1,max=200"`
	Price       int    `validate:"required,min=1"`
	Stock       int    `validate:"required,min=0,max=100000"`
	Location    string `validate:"required,min=1,max=200"`
	IsAvailable bool   `validate:"eq=true|eq=false"`
}
