package diagrams

import (
	"github.com/mmierzwa/go-diagrams-v2/diagram"
)

func New(opts ...diagram.Option) (*diagram.Diagram, error) {
	return diagram.New(opts...)
}
