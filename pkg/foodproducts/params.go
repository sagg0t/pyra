package foodproducts

const (
	errNoName      = "can't be blank"
	errNegative    = "can't be less than 0"
	errNotPositive = "must be greater than 0"
)

type ProductForm struct {
	FoodProduct
	Per    float32
	Errors map[string]string
}

func (f *ProductForm) Validate() bool {
	if len(f.Name) == 0 {
		f.Errors["name"] = errNoName
	}

	if f.Calories < 0 {
		f.Errors["calories"] = errNegative
	}

	if f.Per <= 0 {
		f.Errors["per"] = errNotPositive
	}

	if f.Proteins < 0 {
		f.Errors["proteins"] = errNegative
	}

	if f.Fats < 0 {
		f.Errors["fats"] = errNegative
	}

	if f.Carbs < 0 {
		f.Errors["carbs"] = errNegative
	}

	return len(f.Errors) == 0
}

func (f *ProductForm) NormalizedProduct() FoodProduct {
	clone := f.FoodProduct
	clone.Normalize(f.Per)

	return clone
}
