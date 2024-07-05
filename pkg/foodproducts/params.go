package foodproducts

type Params struct {
	Name string

	Calories float32
	Per      uint32

	Proteins float32
	Fats     float32
	Carbs    float32
}

func (p *Params) Normalize() {
	// Normal values are per 100g
	ratio := 100.0 / float32(p.Per)
	p.Calories *= ratio
	p.Proteins *= ratio
	p.Fats *= ratio
	p.Carbs *= ratio
}
