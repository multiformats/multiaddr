all: go-multiaddr #js-multiaddr

go-multiaddr:
	go get -d -v -t .
	go build -o ./go-multiaddr github.com/multiformats/go-multiaddr/multiaddr
	MULTIADDR_BIN="./go-multiaddr" go test

js-multiaddr:
	npm install multiaddr
	MULTIADDR_BIN="node_modules/.bin/multiaddr" go test

clean:
	rm -rf ./go-multiaddr ./node_modules

.PHONY: go-multiaddr js-multiaddr clean
