package diagrams

import (
	diagram2 "github.com/mmierzwa/go-diagrams-v2/pkg/diagram"
)

func New(opts ...diagram2.Option) (*diagram2.Diagram, error) {
	return diagram2.New(opts...)
}
