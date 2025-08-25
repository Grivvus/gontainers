.PHONY: build
build:
	go build -v .

.PHONY: test
test:
	go test -p 4 -v ./... 

.PHONY: bench
bench:
	go test -v -bench=. -benchmem -count 3 ./... -run=^#
