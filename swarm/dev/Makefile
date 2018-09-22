.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/arfblocks/goarf/cmd/swarm
	go build -o bin/geth     github.com/arfblocks/goarf/cmd/geth
	go build -o bin/bootnode github.com/arfblocks/goarf/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/arfblocks/goarf/swarm/...
