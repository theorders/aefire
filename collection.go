package aefire

import "encoding/json"

func MapOf(pairs ...interface{}) map[string]interface{} {
	if len(pairs)%2 != 0 {
		panic("MapOf: key-value pair cannot be odd")
	}

	m := map[string]interface{}{}

	for i, kv := range pairs {
		if i%2 == 1 {
			k := pairs[i-1].(string)
			m[k] = kv
		}
	}

	return m
}

func StringMapOf(pairs ...string) map[string]string {
	if len(pairs)%2 != 0 {
		panic("MapOf: key-value pair cannot be odd")
	}

	m := map[string]string{}

	for i, kv := range pairs {
		if i%2 == 1 {
			k := pairs[i-1]
			m[k] = kv
		}
	}

	return m
}

func ToMap(v interface{}) (m map[string]interface{}) {
	b, err := json.Marshal(v)

	PanicIfError(err)

	PanicIfError(json.Unmarshal(b, &m))

	return
}

func ToJson(v interface{}, indent ...string) string {
	var b []byte
	var err error

	if len(indent) > 0 {
		b, err = json.MarshalIndent(v, "", indent[0])
	} else {
		b, err = json.Marshal(v)
	}

	PanicIfError(err)

	return string(b)
}

func StringMapValuesToSlice(m map[string]string) (s []string) {
	for _, v := range m {
		s = append(s, v)
	}
	return
}

func LastOf(arr []string) string {
	if len(arr) == 0 {
		return ""
	} else {
		return arr[len(arr)-1]
	}
}

func StringMapContains(m map[string]string, key string) bool {
	_, ok := m[key]

	return ok
}

func MapContains(m map[string]interface{}, key string) bool {
	_, ok := m[key]

	return ok
}
