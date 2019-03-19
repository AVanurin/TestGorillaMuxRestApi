package APIs

import (
	"encoding/json"
	"net/http"
)

func handleVersionCall(w http.ResponseWriter, r *http.Request) {
	var responseMap = make(map[string]string)
	responseMap["version"] = "v1"

	json.NewEncoder(w).Encode(responseMap)
}
