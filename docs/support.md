# Support and compatibility

| Contract | Current development support |
|---|---|
| Go | Exactly 1.26.5 for development and release verification |
| Spice minimum | `v0.0.0-20260805222830-a2ecd56df246` |
| Spice current | `v0.0.0-20260806053623-2ec6f862073f` |
| MySQL | 8.4.11 real-system acceptance; versions supported by go-sql-driver/mysql v1.10.0 remain integration targets, not yet release claims |
| Operating systems | Windows, Linux, and macOS; Linux container acceptance |
| Architectures | amd64 and arm64 compilation through the public core API |
| Transport security | TLS 1.2 hostname verification by default; insecure mode requires explicit test-only opt-in |
| Real-system artifact | `mysql:8.4.11` index digest `sha256:b3b90af2a6552ae30c266fdb7d5dd55f3afb72404bb78d37fe8a23eb857fd3fb` |
| Release parity tool | `github.com/spice-framework/development/cmd/spice-dev` at `v0.0.0-20260806121906-963bb6676069` |
| Release verifier tool | `github.com/spice-framework/toolchain/cmd/spice-library-release-verify` at `v0.0.0-20260806054457-a83d9b58034c` |

The first preview tag will define the minimum supported Spice version. Until
then, development commits fail closed outside the exact minimum and current
Spice commits recorded above. Future releases will preserve both tests before
raising that floor.

The pinned central signer and independent verifier are the protected production
path. Windows and Linux CI still compare the central renderer with the retained
builder under vendor-only offline resolution; the retained command is a parity
oracle only.
