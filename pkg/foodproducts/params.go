package foodproducts

type CreateRequest struct {
	Name string

	Calories float32
	Per      float32

	Proteins float32
	Fats     float32
	Carbs    float32
}

func (p *CreateRequest) Normalize() {
	// Normal values are per 100g
	ratio := 100.0 / p.Per
	p.Calories *= ratio
	p.Proteins *= ratio
	p.Fats *= ratio
	p.Carbs *= ratio
}

type CreateResponse struct {
	CreateRequest
	Errors map[string]string
}
