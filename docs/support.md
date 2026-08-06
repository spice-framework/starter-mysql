# Support and compatibility

| Contract | Current development support |
|---|---|
| Go | Exactly 1.26.5 for development and release verification |
| Spice minimum/current | Exact versions in [`spice-compatibility.json`](../spice-compatibility.json) |
| MySQL | 8.4.11 real-system acceptance; versions supported by go-sql-driver/mysql v1.10.0 remain integration targets, not yet release claims |
| Operating systems | Windows, Linux, and macOS; Linux container acceptance |
| Architectures | amd64 and arm64 compilation through the public core API |
| Transport security | TLS 1.2 hostname verification by default; insecure mode requires explicit test-only opt-in |
| Real-system artifact | `mysql:8.4.11` index digest `sha256:b3b90af2a6552ae30c266fdb7d5dd55f3afb72404bb78d37fe8a23eb857fd3fb` |
| Release parity tool | `github.com/spice-framework/development/cmd/spice-dev` at `v0.0.0-20260806132124-4c308d1b9fda` |
| Release verifier tool | `github.com/spice-framework/toolchain/cmd/spice-library-release-verify` at `v0.0.0-20260806133530-71211498297c` |
| Release trust anchor | [`security/release/ed25519-public.pem`](../security/release/ed25519-public.pem), SHA-256 fingerprint `db149fbb7d7f20ee2379eb41504202069f9f431df16e22bf84d6561719e0a9d1` |

The first preview tag will define the first published minimum Spice version.
Until then, `spice-compatibility.json` is the sole compatibility boundary
source. Its minimum must equal the exact direct Spice requirement in `go.mod`;
its current value is a forward-compatibility endpoint. The compatibility runner
uses an isolated alternate modfile and never mutates the committed module or
vendor graph.

The pinned central signer and independent verifier are the protected production
path. Windows and Linux CI still compare the central renderer with the retained
builder under vendor-only offline resolution; the retained command is a parity
oracle only.

The reviewed public release anchor is committed and is safe to distribute. It
does not mean a signed release exists: publication remains blocked until the
matching private key and protected release environments are configured and the
documented release ceremony completes.
