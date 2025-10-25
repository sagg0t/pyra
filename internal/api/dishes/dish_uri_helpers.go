package dishes

import (
	"fmt"
	"html/template"
)

var URIHelpers = template.FuncMap{
	"dishURI": DishURI,
}

func DishURI(id uint64) string {
	return fmt.Sprintf("/dishes/%d", id)
}
