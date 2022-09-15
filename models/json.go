package models

import (
	"encoding/json"
	"strings"
)

// Structs for json data

type JSONData struct {
	Id           string
	Type         string
	TotalItems   uint
	OrderedItems []struct {
		Id        string
		Published string
		To        []string
		Object    ObjectWrapper
	}
}

type ObjectWrapper struct {
	Object
}

type Object struct {
	Type    string
	Content string
}

// Method that overrides default UnmarshalJSON method for ObjectWrapper because
// sometimes it's a plain string and some other times it's a "dict"
func (w *ObjectWrapper) UnmarshalJSON(data []byte) error {
	datas := string(data)
	if strings.HasPrefix(datas, "\"") {
		w.Content = datas
		return nil
	}
	return json.Unmarshal(data, &w.Object)
}
