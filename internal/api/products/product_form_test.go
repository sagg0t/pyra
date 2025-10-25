package products

import (
	"testing"

	"pyra/pkg/nutrition"

	"github.com/stretchr/testify/assert"
)

type formData map[string]string

func (fd formData) fetch(key string) string {
	return fd[key]
}

func Test_NewProductForm(t *testing.T) {
	examples := []struct {
		in  formData
		out ProductForm
	}{
		{
			in:  formData{"name": "  asdf  "},
			out: ProductForm{Name: "asdf"},
		},
		{
			in:  formData{"calories": "  123  "},
			out: ProductForm{Calories: "123"},
		},
		{
			in:  formData{"proteins": "  123  "},
			out: ProductForm{Proteins: "123"},
		},
		{
			in:  formData{"fats": "  123  "},
			out: ProductForm{Fats: "123"},
		},
		{
			in:  formData{"carbs": "  123  "},
			out: ProductForm{Carbs: "123"},
		},
		{
			in:  formData{"per": "  123  "},
			out: ProductForm{Per: "123"},
		},
	}

	for _, example := range examples {
		form := NewProductForm(example.in.fetch)

		assert.Equal(t, form, example.out)
	}
}

func Test_BuildProduct(t *testing.T) {
	t.Run("valid minimal form", func(t *testing.T) {
		form := ProductForm{
			Name: "name",
			Per:  "100",
		}

		product := form.BuildProduct()

		assert.False(t, form.HasErrors(), "expected no errors")

		assert.EqualValues(t, product.Name, form.Name)
		assert.EqualValues(t, product.Calories, 0)
		assert.EqualValues(t, product.Proteins, 0)
		assert.EqualValues(t, product.Fats, 0)
		assert.EqualValues(t, product.Carbs, 0)
	})

	t.Run("non-standard portion", func(t *testing.T) {
		form := ProductForm{
			Name:     "name",
			Per:      "50",
			Calories: "1",
			Proteins: "1",
			Fats:     "1",
			Carbs:    "1",
		}

		product := form.BuildProduct()

		assert.False(t, form.HasErrors(), "expected no errors")

		assert.Equal(t, form.Name, string(product.Name))
		assert.Equal(t, nutrition.NewMeasurement(2), product.Calories)
		assert.Equal(t, nutrition.NewMeasurement(2), product.Proteins)
		assert.Equal(t, nutrition.NewMeasurement(2), product.Fats)
		assert.Equal(t, nutrition.NewMeasurement(2), product.Carbs)
	})

	t.Run("validates ProductForm.Name", func(t *testing.T) {
		form := ProductForm{
			Name: "",
			Per:  "100",
		}
		_ = form.BuildProduct()

		assert.True(t, form.HasErrors(), "expected errors")

		if assert.Error(t, form.Errors.Name) {
			assert.Equal(t, form.Errors.Name, nutrition.ErrBlank)
		}
	})

	t.Run("validates ProductForm.Per", func(t *testing.T) {
		t.Run("when not a number", func(t *testing.T) {
			form := ProductForm{
				Name: "name",
				Per:  "asdf",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Per) {
				assert.Equal(t, form.Errors.Per, ErrInvalidNumber)
			}
		})

		t.Run("when 0", func(t *testing.T) {
			form := ProductForm{
				Name: "name",
				Per:  "",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Per) {
				assert.Equal(t, form.Errors.Per, ErrNotPositive)
			}
		})

		t.Run("when negative", func(t *testing.T) {
			form := ProductForm{
				Name: "name",
				Per:  "-1.5",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Per) {
				assert.Equal(t, form.Errors.Per, ErrNotPositive)
			}
		})
	})

	t.Run("validates ProductForm.Calories", func(t *testing.T) {
		t.Run("when not a number", func(t *testing.T) {
			form := ProductForm{
				Name:     "name",
				Calories: "asdf",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Calories) {
				assert.Equal(t, form.Errors.Calories, ErrInvalidNumber)
			}
		})

		t.Run("when negative", func(t *testing.T) {
			form := ProductForm{
				Name:     "name",
				Calories: "-1.5",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Calories) {
				assert.Equal(t, form.Errors.Calories, nutrition.ErrNegative)
			}
		})
	})

	t.Run("validates ProductForm.Proteins", func(t *testing.T) {
		t.Run("when not a number", func(t *testing.T) {
			form := ProductForm{
				Name:     "name",
				Proteins: "asdf",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Proteins) {
				assert.Equal(t, form.Errors.Proteins, ErrInvalidNumber)
			}
		})

		t.Run("when negative", func(t *testing.T) {
			form := ProductForm{
				Name:     "name",
				Proteins: "-1.5",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Proteins) {
				assert.Equal(t, form.Errors.Proteins, nutrition.ErrNegative)
			}
		})
	})

	t.Run("validates ProductForm.Fats", func(t *testing.T) {
		t.Run("when not a number", func(t *testing.T) {
			form := ProductForm{
				Name: "name",
				Fats: "asdf",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Fats) {
				assert.Equal(t, form.Errors.Fats, ErrInvalidNumber)
			}
		})

		t.Run("when negative", func(t *testing.T) {
			form := ProductForm{
				Name: "name",
				Fats: "-1.5",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Fats) {
				assert.Equal(t, form.Errors.Fats, nutrition.ErrNegative)
			}
		})
	})

	t.Run("validates ProductForm.Carbs", func(t *testing.T) {
		t.Run("when not a number", func(t *testing.T) {
			form := ProductForm{
				Name:  "name",
				Carbs: "asdf",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Carbs) {
				assert.Equal(t, form.Errors.Carbs, ErrInvalidNumber)
			}
		})

		t.Run("when negative", func(t *testing.T) {
			form := ProductForm{
				Name:  "name",
				Carbs: "-1.5",
			}
			_ = form.BuildProduct()

			assert.True(t, form.HasErrors(), "expected errors")

			if assert.Error(t, form.Errors.Carbs) {
				assert.Equal(t, form.Errors.Carbs, nutrition.ErrNegative)
			}
		})
	})
}
