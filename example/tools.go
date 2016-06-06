package main

import (
	"encoding/json"
	"fmt"
)

func toString(v interface{}) string {
	buf, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("unable to marshal %#v: %s", v, err)
	}
	return string(buf)
}
