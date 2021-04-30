package main

import (
	"encoding/json"
	"os"

	restful "github.com/emicklei/go-restful/v3"
	restfulspec "github.com/yefengorg/go-restful-openapi"
)

func main() {
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	swagger := restfulspec.BuildSwagger(config)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	enc.Encode(swagger)
}
