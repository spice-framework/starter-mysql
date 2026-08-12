module github.com/spice-framework/starter-mysql

go 1.26.0

toolchain go1.26.5

require (
	github.com/go-sql-driver/mysql v1.10.0
	github.com/spice-framework/spice v0.1.0-preview.1
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/spice-framework/development v0.0.0-20260806132124-4c308d1b9fda // indirect
	github.com/spice-framework/toolchain v0.0.0-20260806133530-71211498297c // indirect
	golang.org/x/mod v0.38.0 // indirect
)

tool (
	github.com/spice-framework/development/cmd/spice-dev
	github.com/spice-framework/toolchain/cmd/spice-library-release-verify
)
