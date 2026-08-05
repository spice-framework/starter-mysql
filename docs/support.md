# Support and compatibility

| Contract | Current development support |
|---|---|
| Go | Exactly 1.26.5 for development and release verification |
| Spice | `v0.0.0-20260805185924-ee45e0aa386e` |
| MySQL | 8.4.11 real-system acceptance; versions supported by go-sql-driver/mysql v1.10.0 remain integration targets, not yet release claims |
| Operating systems | Windows, Linux, and macOS; Linux container acceptance |
| Architectures | amd64 and arm64 compilation through the public core API |
| Transport security | TLS 1.2 hostname verification by default; insecure mode requires explicit test-only opt-in |
| Real-system artifact | `mysql:8.4.11` index digest `sha256:b3b90af2a6552ae30c266fdb7d5dd55f3afb72404bb78d37fe8a23eb857fd3fb` |

The first preview tag will define the minimum supported Spice version. Until
then, development commits intentionally declare one exact compatible Spice
commit and fail closed outside that tested combination. Future releases will
test both the published minimum and current supported Spice lines before
raising that floor.
