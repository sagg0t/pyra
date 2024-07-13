package foodproducts

import "errors"

type Validator interface {
	Validate()
	Err() error
}

var (
	ErrNoName      = errors.New("can't be blank")
	ErrNegative    = errors.New("can't be less than 0")
	ErrNotPositive = errors.New("must be greater than 0")
)

type validator struct {
	params Form
	err    error
}

func NewCreateValidator(params Form) Validator {
	return &validator{
		params: params,
		err:    nil,
	}
}

func (v *validator) Err() error {
	return v.err
}

func (v *validator) Validate() {
	p := &v.params

	if len(p.Name) == 0 {
		v.err = ErrNoName
		return
	}

	if p.Calories < 0 {
		v.err = ErrNegative
		return
	}

	if p.Per <= 0 {
		v.err = ErrNotPositive
		return
	}

	if p.Proteins < 0 {
		v.err = ErrNegative
		return
	}

	if p.Fats < 0 {
		v.err = ErrNegative
		return
	}

	if p.Carbs < 0 {
		v.err = ErrNegative
		return
	}
}
