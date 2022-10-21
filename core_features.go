package of

import (
	"github.com/tikafog/of/feature/bootstrap"
	"github.com/tikafog/of/feature/resource"
)

type CoreFeature interface {
	Module(name Name) Module
	API() API
	Event() Event
	Node() Node
	Resource() resource.Resource
	Bootstrap() bootstrap.Bootstrap
}
