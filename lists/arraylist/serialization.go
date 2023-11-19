package arraylist

import (
	"encoding/json"
	"gods/containers"
)

var _ containers.JSONDeserializer = (*List)(nil)
var _ containers.JSONSerializer = (*List)(nil)

// ToJSON 将数据转换为json并输出
// ToJSON outputs the JSON representation of list's elements.
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.elements[:list.size])
}

// FromJSON 将json数据转换为list
func (list *List) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (list *List) UnmarshalJSON(bytes []byte) error {
	return list.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (list *List) MarshalJSON() ([]byte, error) {
	return list.ToJSON()
}
