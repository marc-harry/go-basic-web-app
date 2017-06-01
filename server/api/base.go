package api

// Action defines a standard function signature for us to use when creating
import (
	"gopkg.in/unrolled/render.v1"
)

// This is our Base Controller
type AppController struct{}

type BaseApiController struct {
	AppController
	*render.Render
}
