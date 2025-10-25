package nutrition

import "github.com/brianvoe/gofakeit/v7"

func FakeProduct() Product {
	p := Product{}

	if err := gofakeit.Struct(&p); err != nil {
		panic(err)
	}

	return p
}
