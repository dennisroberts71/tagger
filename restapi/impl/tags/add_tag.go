package tags

import (
	"github.com/cyverse-de/tagger/restapi/operations/tags"
	"github.com/go-openapi/runtime/middleware"
)

func BuildAddTagHandler() func(tags.AddTagParams) middleware.Responder {

	// Return the handler function.
	return func(tags.AddTagParams) middleware.Responder {
		return tags.NewAddTagOK()
	}
}
