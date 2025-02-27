package urlshort

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathsToUrls[r.URL.Path]

		if ok {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return baseHandler(yml, fallback, yaml.Unmarshal)
}

func JSONHandler(blob []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return baseHandler(blob, fallback, json.Unmarshal)
}

func baseHandler(blob []byte, fallback http.Handler, handler func(data []byte, v any) error) (http.HandlerFunc, error) {
	var out []URLMapping
	err := handler(blob, &out)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(out)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(mappings []URLMapping) map[string]string {
	ret := make(map[string]string)
	for _, v := range mappings {
		ret[v.path] = v.url
	}

	return ret
}

// Technically you can directly map the yaml into a map[string]string
// without defining an entity (only for yaml not for json)
type URLMapping struct {
	path string `yaml:"path"`
	url  string `yaml:"url"`
}

func DBHandler(db *sql.DB, fallback http.Handler) (http.HandlerFunc, error) {
	rows, err := db.Query("SELECT * FROM mappings")
	if err != nil {
		log.Fatal("Error executing query: ", err)
	}
	defer rows.Close()

	var mappings []URLMapping
	for rows.Next() {
		var mapping URLMapping
		err := rows.Scan(&mapping.path, &mapping.url)
		if err != nil {
			return nil, err
		}

		mappings = append(mappings, mapping)
	}

	pathMap := buildMap(mappings)
	return MapHandler(pathMap, fallback), nil
}
