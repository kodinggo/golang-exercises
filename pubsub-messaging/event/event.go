package event

const (
	ProductsTopic = "PRODUCTS"

	ProductsEvent      = "PRODUCTS.*"
	CreateProductEvent = "PRODUCTS.create"
	UpdateProductEvent = "PRODUCTS.update"
	DeleteProductEvent = "PRODUCTS.delete"
)

type Product struct {
	ID   int
	Name string
}
