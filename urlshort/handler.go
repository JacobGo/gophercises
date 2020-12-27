package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if url,contains := pathsToUrls[req.URL.Path]; contains {
			http.Redirect(w, req, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, req)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Unmarshal the YAML data into an array of pathUrls
	var urlMaps []pathUrl
	err := yaml.Unmarshal(yml, &urlMaps)
	if err != nil {
		return nil, err
	}
	// Convert the array of pathUrls into a map
	pathsToUrls := make(map[string]string)
	for _, urlMap := range urlMaps {
		pathsToUrls[urlMap.Path] = urlMap.URL
	}
	// Return a MapHandler with the converted map
	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}