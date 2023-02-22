package main

import (
	"log"

	diagram2 "github.com/mmierzwa/go-diagrams-v2/pkg/diagram"
	gcp2 "github.com/mmierzwa/go-diagrams-v2/pkg/nodes/gcp"
)

func main() {
	d, err := diagram2.New(diagram2.WithFilename("app"), diagram2.WithLabel("App"), diagram2.WithDirection("LR"))
	if err != nil {
		log.Fatal(err)
	}

	dns := gcp2.Network.Dns(diagram2.NodeLabel("DNS"))
	lb := gcp2.Network.LoadBalancing(diagram2.NodeLabel("NLB"))
	cache := gcp2.Database.Memorystore(diagram2.NodeLabel("Cache"))
	db := gcp2.Database.Sql(diagram2.NodeLabel("Database"))

	dc := diagram2.NewGroup("GCP")
	dc.NewGroup("services").
		Label("Service Layer").
		Add(
			gcp2.Compute.ComputeEngine(diagram2.NodeLabel("Server 1")),
			gcp2.Compute.ComputeEngine(diagram2.NodeLabel("Server 2")),
			gcp2.Compute.ComputeEngine(diagram2.NodeLabel("Server 3")),
		).
		ConnectAllFrom(lb.ID(), diagram2.Forward()).
		ConnectAllTo(cache.ID(), diagram2.Forward())

	dc.NewGroup("data").Label("Data Layer").Add(cache, db).Connect(cache, db)

	d.Connect(dns, lb, diagram2.Forward()).Group(dc)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
