package products

const (
	ProductPATH = "/products/{uid}/{version}"
	ProductsPATH = "/products"
	NewProductPATH = ProductsPATH + "/new"
	EditProductPATH = ProductPATH + "/edit"

	ListProductsEP = "GET " + ProductsPATH
	ShowProductEP = "GET " + ProductPATH
	NewProductEP = "GET " + NewProductPATH
	CreateProductEP = "POST " + ProductsPATH
	DeleteProductEP = "DELETE " + ProductPATH
	EditProductEP = "GET " + EditProductPATH
	UpdateProductEP = "PUT " + ProductPATH
)
