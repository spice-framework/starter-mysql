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
| Release trust anchor | [`security/release/ed25519-public.pem`](../security/release/ed25519-public.pem), SHA-256 fingerprint `1c3d374cfe4465f6c04ad7afe22acad2bb05d51bc6d249745c92a817265d08a9` |

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

The reviewed public release anchor is committed and its matching private key is
stored only as the repository Actions secret
`SPICE_LIBRARY_RELEASE_SIGNING_KEY`. The caller forwards that one named secret;
the protected `release-signing` environment contains no secret copy and still
gates the signing job through required review. These controls do not mean a
signed release exists: publication still requires an approved canonical tag and
the complete documented ceremony.
