package metadata

import (
	"strings"
)

// Metadata is our way of representing request headers internally.
// They're used at the RPC level and translate back and forth
// from Transport headers.
type Metadata map[string]string

// New creates an Metadata from a given key-value map.
func New(m map[string]string) Metadata {
	md := Metadata{}
	for key, val := range m {
		md[key] = val
	}
	return md
}

// Len returns the number of items in md.
func (md Metadata) Len() int {
	return len(md)
}

func (md Metadata) Get(key string) (string, bool) {
	// attempt to get as is
	val, ok := md[key]
	if ok {
		return val, ok
	}

	// attempt to get lower case
	val, ok = md[strings.Title(key)]
	return val, ok
}

func (md Metadata) Set(key, val string) {
	md[key] = val
}

func (md Metadata) Delete(keys ...string) {
	for _, key := range keys {
		// delete key as-is
		delete(md, key)
		// delete also Title key
		delete(md, strings.Title(key))
	}
}

// Copy makes a copy of the metadata
func (md Metadata) Copy() Metadata {
	return New(md)
}

func (md Metadata) Patch(patches []Metadata, overwrite bool) {
	for _, patch := range patches {
		for k, v := range patch {
			if _, ok := md[k]; ok && !overwrite {
				// skip
			} else if v != "" {
				md[k] = v
			} else {
				delete(md, k)
			}
		}
	}
}
