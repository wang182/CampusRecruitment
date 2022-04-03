package utils

import (
	"encoding/json"
)

func JSONMarshal(v interface{}) []byte {
	bs, _ := json.Marshal(v)
	return bs
}
