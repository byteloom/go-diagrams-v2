package main

import (
	"log"

	diagram2 "github.com/mmierzwa/go-diagrams-v2/pkg/diagram"
	apps2 "github.com/mmierzwa/go-diagrams-v2/pkg/nodes/apps"
)

func main() {
	d, err := diagram2.New(diagram2.WithLabel("Web Service"), diagram2.WithDirection("LR"))
	if err != nil {
		log.Fatal(err)
	}

	inet := apps2.Network.Internet().Label("Internet")
	proxy := apps2.Network.Caddy().Label("Caddy")

	d.Connect(inet, proxy)

	ss := apps2.Inmemory.Redis().Label("session")
	rs := apps2.Inmemory.Redis().Label("replica")

	cache := diagram2.NewGroup("cache").Label("Sessions").
		Connect(ss, rs, diagram2.Bidirectional())

	dbmain := apps2.Database.Postgresql().Label("Master DB")
	repls := []*diagram2.Node{
		apps2.Database.Postgresql().Label("DB Replica 1"),
		apps2.Database.Postgresql().Label("DB Replica 2"),
	}

	db := diagram2.NewGroup("database").Label("Database").
		Add(dbmain).Add(repls...)

	svcs := diagram2.NewGroup("services").Label("Services").
		Add(
			apps2.Container.Docker().Label("Replica 1"),
			apps2.Container.Docker().Label("Replica 2"),
		).
		ConnectAllFrom(proxy.ID()).
		ConnectAllTo(dbmain.ID()).
		ConnectAllTo(ss.ID())

	for _, r := range repls {
		db.Connect(dbmain, r)
	}

	d.Group(svcs).
		Group(cache).
		Group(db)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}

}
