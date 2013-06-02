package util

type MapUtil struct{}

/*
 * Copies all keys and entries from the source map to the destination map.
 */
func (MapUtil) AddAll(dest, src map[string]interface{}) {
	for key := range src {
		dest[key] = src[key]
	}
}
