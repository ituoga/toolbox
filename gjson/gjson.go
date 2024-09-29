package gjson

import "encoding/json"

func Unmarshal[T any](data []byte) (T, error) {
	var v T
	err := json.Unmarshal(data, &v)
	return v, err
}

func MustUnmarshal[T any](data []byte) T {
	v, err := Unmarshal[T](data)
	if err != nil {
		panic(err)
	}
	return v
}
