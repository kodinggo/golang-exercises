package event

const (
	ProductsTopic = "PRODUCTS"

	ProductsEvent = "PRODUCTS.*"
	// ProductsEvent      = "PRODUCTS.>"
	// CreateProductCategoryAEvent = "PRODUCTS.category-a.create"
	CreateProductEvent = "PRODUCTS.create"
	UpdateProductEvent = "PRODUCTS.update"
	DeleteProductEvent = "PRODUCTS.delete"
)

type Product struct {
	ID   int
	Name string
}
