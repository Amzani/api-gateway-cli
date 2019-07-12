package swagger

//TODO -> class

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

//Parse OAS files
//Return an OpenAPI Spec Object
func Parse(oasPath string) (*spec.Swagger, error) {
	specDoc, err := loads.Spec(oasPath)

	if err != nil {
		return nil, err
	}

	return specDoc.Spec(), nil
}

//PathList function
//Get list of all swagger paths (routes)
func PathList(spec *spec.Swagger) []string {
	keys := make([]string, len(spec.Paths.Paths))

	i := 0
	for k := range spec.Paths.Paths {
		keys[i] = k
		i++
	}
	return keys
}
