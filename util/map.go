// This file contains util functions dealing with map[string]interface{}
// Useful when dealing with configs readed from viper

package util

// HasKey is a shortcut for checking if certain key exists in map
func HasKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}

// GetWithDefault return default vaule if certain key does not exists
func GetWithDefault(m map[string]interface{}, key string, d interface{}) interface{} {
	v, ok := m[key]
	if ok {
		return v
	}
	return d
}
