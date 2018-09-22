.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/arfblocks/goArf/cmd/swarm
	go build -o bin/geth     github.com/arfblocks/goArf/cmd/geth
	go build -o bin/bootnode github.com/arfblocks/goArf/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/arfblocks/goArf/swarm/...
