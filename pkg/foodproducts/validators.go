package foodproducts

type Validator interface {
	Validate()
	Err() map[string]string
}

var (
	errNoName      = "can't be blank"
	errNegative    = "can't be less than 0"
	errNotPositive = "must be greater than 0"
)

type validator struct {
	params CreateRequest
	err    map[string]string
}

func NewCreateValidator(params CreateRequest) Validator {
	return &validator{
		params: params,
		err:    make(map[string]string),
	}
}

func (v *validator) Err() map[string]string {
	return v.err
}

func (v *validator) Validate() {
	p := &v.params

	if len(p.Name) == 0 {
		v.err["name"] = errNoName
	}

	if p.Calories < 0 {
		v.err["calories"] = errNegative
	}

	if p.Per <= 0 {
		v.err["per"] = errNotPositive
	}

	if p.Proteins < 0 {
		v.err["proteins"] = errNegative
	}

	if p.Fats < 0 {
		v.err["fats"] = errNegative
	}

	if p.Carbs < 0 {
		v.err["carbs"] = errNegative
	}
}
