package main

import (
	"fmt"
	"log"

	diagram2 "github.com/mmierzwa/go-diagrams-v2/pkg/diagram"
	gcp2 "github.com/mmierzwa/go-diagrams-v2/pkg/nodes/gcp"
)

func main() {
	workerCount := 5

	d, err := diagram2.New(
		diagram2.WithLabel("Workers"),
		diagram2.WithFilename("workers"),
		diagram2.WithDirection("TB"),
	)

	if err != nil {
		log.Fatal(err)
	}

	lb := gcp2.Network.LoadBalancing(diagram2.NodeLabel("nlb"))
	d.Add(lb)

	db := gcp2.Database.Sql(diagram2.NodeLabel("db"))
	d.Add(db)

	workers := make([]*diagram2.Node, workerCount)

	for i := 0; i < workerCount; i++ {
		label := fmt.Sprintf("worker %d", i+1)
		workers[i] = gcp2.Compute.ComputeEngine(diagram2.NodeLabel(label))
	}

	d.Group(diagram2.NewGroup("workers").
		Add(workers...).
		ConnectAllTo(db.ID()).
		ConnectAllFrom(lb.ID()),
	)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
