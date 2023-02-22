package main

import (
	"log"

	diagram2 "github.com/mmierzwa/go-diagrams-v2/pkg/diagram"
	k8s2 "github.com/mmierzwa/go-diagrams-v2/pkg/nodes/k8s"
)

func main() {
	d, err := diagram2.New(
		diagram2.WithLabel("Kubernetes"),
		diagram2.WithFilename("k8s"),
		diagram2.WithDirection("TB"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ingress := k8s2.Network.Ing(diagram2.NodeLabel("nginx"))
	svc := k8s2.Network.Svc(diagram2.NodeLabel("http"))

	d.Connect(ingress, svc)

	g := diagram2.NewGroup("pods").Label("Deployment").Connect(svc, k8s2.Compute.Pod(diagram2.NodeLabel("web server")))

	d.Group(g)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
