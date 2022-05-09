package rest

import (
	"net/http"
	"os"

	"facette.io/natsort"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
)

const envDocRoot = "SWAGGER_PATH"

func bootstrapSwagger(container *restful.Container) error {
	docRoot := os.Getenv(envDocRoot)

	if docRoot == "" {
		docRoot = "./docs"
	}

	routePath := "/internship-api/docs/"
	specPath := "/internship-api/docs/spec.json"
	servePath := http.Dir(docRoot)

	seedRouteTags(container)

	swgConfig := restfulspec.Config{
		WebServices:                   container.RegisteredWebServices(),
		APIPath:                       specPath,
		PostBuildSwaggerObjectHandler: postBuildSwaggerObjectHandler(container),
	}

	container.Add(restfulspec.NewOpenAPIService(swgConfig))
	container.ServeMux.Handle(routePath, http.StripPrefix(routePath, http.FileServer(servePath)))

	return nil
}

func seedRouteTags(container *restful.Container) {
	for _, ws := range container.RegisteredWebServices() {
		// get local reference to route slice
		routes := ws.Routes()
		for i := range routes {
			// test to see if route already has tag defined in metadata...
			if _, ok := routes[i].Metadata[restfulspec.KeyOpenAPITags]; !ok {
				if routes[i].Metadata == nil {
					// if no metadata is defined on the route the map will be nil by the time we get here, so make on.
					routes[i].Metadata = make(map[string]interface{})
				}
				// set route tag to its containing webservice's path.
				routes[i].Metadata[restfulspec.KeyOpenAPITags] = []string{ws.RootPath()}
			}
		}
	}
}

func postBuildSwaggerObjectHandler(container *restful.Container) restfulspec.PostBuildSwaggerObjectFunc {
	return func(swo *spec.Swagger) {
		swo.Swagger = "2.0"
		swo.ID = "internship"

		swo.Schemes = []string{"https", "http"}

		swo.Info = &spec.Info{
			InfoProps: spec.InfoProps{
				Title:       "internship",
				Description: "internship api",
				Contact: &spec.ContactInfo{
					ContactInfoProps: spec.ContactInfoProps{
						Name:  "Miron George",
						Email: "george.miron2003@gmail.com",
					},
				},
				License: &spec.License{
					LicenseProps: spec.LicenseProps{
						Name: "Proprietary",
					},
				},
				Version: "1.0.0",
			},
		}

		swo.Definitions["error"] = spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"string"},
			},
		}

		annotateSwaggerTags(container, swo)
	}
}

// annotateSwagger tags does a few things:
// 		1. Adds any tags and descriptions created when creating a webservice using one of the helpers
//		2. Loops through all registered webservices and their associated routes and locates any defined tags
//		3. If tag(s) defined, takes first (if more than one) defined and creates a new Swagger Tag definition using the
//			defined tag as the name of the tag and the route's containing webservice's description as the swagger tag
//			description
//		4. Appends tags not defined by webservice helper to the swagger object.
func annotateSwaggerTags(container *restful.Container, swo *spec.Swagger) {
	// used to keep track of tags defined in routes so we don't repeat ourselves.
	routeTags := make(map[string]spec.Tag)

	// loop over all webservices...
	for _, ws := range container.RegisteredWebServices() {
		// localize routes so we aren't shallow copying shit all over the place.
		routes := ws.Routes()
		for i := range routes {
			if v, ok := routes[i].Metadata[restfulspec.KeyOpenAPITags]; ok {
				// for now, only look at the first one
				if tagList, ok := v.([]string); ok && len(tagList) > 0 {
					tagName := tagList[0]
					// check to see if the webservice creator already annotated the tag
					if _, ok = routeTags[tagName]; !ok {
						// check to see if we've already seen a route with this tag.
						if _, ok = routeTags[tagName]; !ok {
							routeTags[tagName] = spec.NewTag(tagName, ws.Documentation(), nil)
						}
					}
				}
			}
		}
	}

	// all this nonsense sorts the tags alphabetically

	tagNames := make([]string, 0)

routeTagOuter:
	for _, tag := range routeTags {
		for _, curr := range tagNames {
			if tag.Name == curr {
				continue routeTagOuter
			}
		}
		tagNames = append(tagNames, tag.Name)
	}

swoTagOuter:
	for _, tag := range swo.Tags {
		for _, curr := range tagNames {
			if tag.Name == curr {
				continue swoTagOuter
			}
		}
		tagNames = append(tagNames, tag.Name)
	}

	var cswo []spec.Tag
	if cl := len(cswo); cl != 0 {
		cswo = make([]spec.Tag, cl)
		copy(cswo, swo.Tags)
	}

	// reset source tags
	swo.Tags = make([]spec.Tag, 0)

	natsort.Sort(tagNames)

finalOuter:
	for _, tagName := range tagNames {
		for _, tag := range routeTags {
			if tag.Name == tagName {
				swo.Tags = append(swo.Tags, tag)
				continue finalOuter
			}
		}
		for _, tag := range cswo {
			if tag.Name == tagName {
				swo.Tags = append(swo.Tags, tag)
			}
		}
	}
}
