package dishes

import (
	"net/http"
	"strconv"
)

func dishID(r *http.Request) (uint64, error) {
	paramID := r.PathValue("id")
	return strconv.ParseUint(paramID, 10, 64)
}
