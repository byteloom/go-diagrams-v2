package openstack

import (
	"github.com/mmierzwa/go-diagrams-v2/pkg/diagram"
)

type adjacentenablersContainer struct {
	path string
	opts []diagram.NodeOption
}

var Adjacentenablers = &adjacentenablersContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/adjacentenablers",
}
