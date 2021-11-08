package tools

import "encoding/json"

func ToJSON(obj interface{}) []byte {
	j, _ := json.Marshal(obj)
	return j
}
