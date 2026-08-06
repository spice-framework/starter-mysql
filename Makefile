.PHONY: check fmt integration release-parity verify verify-release

check:
	go run ./internal/qualitygate -mode=check

fmt:
	go run ./internal/qualitygate -mode=fmt

release-parity: export GOWORK := off
release-parity: export GOPROXY := off
release-parity: export GOTOOLCHAIN := local
release-parity: export GOFLAGS := -mod=vendor
release-parity:
	go run ./internal/qualitygate -mode=release-parity

integration:
	go test -tags=integration -race -shuffle=on -count=1 ./...

verify:
	go run ./internal/qualitygate -mode=verify

verify-release:
	go run ./internal/qualitygate -mode=verify-release
