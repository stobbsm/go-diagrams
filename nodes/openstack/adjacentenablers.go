package openstack

import "github.com/stobbsm/go-diagrams/diagram"

type adjacentenablersContainer struct {
	path string
	opts []diagram.NodeOption
}

var Adjacentenablers = &adjacentenablersContainer{
	opts: diagram.OptionSet{diagram.Provider("openstack"), diagram.NodeShape("none")},
	path: "assets/openstack/adjacentenablers",
}
