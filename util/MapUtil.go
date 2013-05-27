package util

import "labix.org/v2/mgo/bson"

type MapUtil struct {}

/*
 * Copies all keys and entries from the source map to the destination map.
 */
func (MapUtil) AddAll(src, dest bson.M) {
	for key, value := range src {
		dest[key] = value
	}
}

