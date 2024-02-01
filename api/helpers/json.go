package helpers

import "encoding/json"

func ToJson(data any) []byte {
	str, err := json.Marshal(data)
	if err != nil {
		return []byte(err.Error())
	}
	return str
}
