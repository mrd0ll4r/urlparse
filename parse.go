package urlparse
import (
	"strings"
	"net/url"
)

type Params map[string][]string

func ParseStdlib(path string) (params Params, err error) {
	idx := strings.Index(path, "?")
	if idx == -1 {
		return
	}

	s := path[idx + 1:]

	v, err := url.ParseQuery(s)
	if err != nil {
		return
	}

	params = Params(v)
	return
}

func ParseNoUrldecode(path string) (params Params) {
	idx := strings.Index(path, "?")
	params = make(Params)
	if idx == -1 {
		return
	}

	// implementation of url.ParseQuery(string) without unescaping
	for path != "" {
		key := path
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			key, path = key[:i], key[i+1:]
		} else {
			path = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}

		params[key] = append(params[key], value)
	}
	return
}

func ParseFull(path string) (*url.URL,error) {
	return url.Parse(path)
}