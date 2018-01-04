.PHONY: benchmark generate

benchmark:
	go test -run=none -tags bench -bench . -benchmem

generate:
	go generate