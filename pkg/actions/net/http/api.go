package http

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"gopkg.in/yaml.v3"
)

type openapi struct {
	Paths map[string]map[string]struct {
		Summary string
	}
}

type OpenApiPathOpts struct {
	Spec   string
	Method string
}

// ActionOpenApiPaths completes api paths
//
//	  http.ActionOpenApiPaths(http.OpenApiPathOpts{Spec: "petstore.json"}).
//			MultiPartsP("/", "{.*}", func(placeholder string, matches map[string]string) carapace.Action {
//				switch placeholder {
//				case "{petId}":
//					return carapace.ActionValues("pet1", "pet2", "pet3")
//				case "{orderId}":
//					return carapace.ActionValues("order1", "order2", "order3")
//				case "{username}":
//					return carapace.ActionValues("user1", "user2", "user3")
//				default:
//					return carapace.ActionValues()
//				}
//			})
func ActionOpenApiPaths(opts OpenApiPathOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if splitted := strings.SplitN(opts.Spec, "\n", 2); len(splitted) == 1 {
			switch filepath.Ext(opts.Spec) {
			case ".json", ".yaml", ".yml":
				content, err := os.ReadFile(opts.Spec)
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				opts.Spec = string(content)
			}
		}
		if opts.Method == "" {
			opts.Method = "GET"
		}

		var api openapi
		switch {
		case strings.HasPrefix(opts.Spec, "{"):
			if err := json.Unmarshal([]byte(opts.Spec), &api); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		default:
			if err := yaml.Unmarshal([]byte(opts.Spec), &api); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		pathsDescribed := make([]string, 0)
		for path, methods := range api.Paths {
			for _method, details := range methods {
				if strings.EqualFold(_method, opts.Method) {
					pathsDescribed = append(pathsDescribed, path, details.Summary)
				}
			}
		}
		return carapace.ActionValuesDescribed(pathsDescribed...)
	})
}
