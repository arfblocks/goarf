.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/tayfunakcay/goArf/cmd/swarm
	go build -o bin/geth     github.com/tayfunakcay/goArf/cmd/geth
	go build -o bin/bootnode github.com/tayfunakcay/goArf/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/tayfunakcay/goArf/swarm/...
