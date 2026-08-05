# Dependency review: go-sql-driver/mysql

- Decision: approved for the isolated
  `github.com/spice-framework/starter-mysql` module. The dependency never
  enters the Spice core module graph.
- Version: `github.com/go-sql-driver/mysql` v1.10.0.
- Upstream: <https://github.com/go-sql-driver/mysql>.
- License: MPL-2.0; the unchanged vendored dependency retains its license and
  Spice does not modify its files.
- Transitive cryptography: `filippo.io/edwards25519` v1.2.0 is pulled by the
  driver for Ed25519 authentication. It uses the BSD-3-Clause license, remains
  unmodified in vendor, and is covered by the same vulnerability gate.
- Maintenance: active stable driver for Go's `database/sql` API with context
  cancellation and current MySQL support.
- Security: Spice constructs the driver configuration without parsing an
  arbitrary DSN, verifies TLS certificates and hostnames by default, rejects
  unsafe driver options, disables multi-statements and file loading, and never
  includes the URL or password in errors. Insecure local connections require
  an option and an explicit URL marker.
- Cancellation: connection establishment, ping, queries, and transactions use
  caller-owned contexts. Driver cancellation closes an in-flight connection,
  matching `database/sql` cancellation semantics.
- Observability: a bounded application name is sent as a connection attribute;
  the application owns tracing wrappers and the pool lifecycle.
- Configuration: Spice validates identities, pool bounds, timeouts, and TLS
  policy before constructing a pool. `Open` performs no network I/O and
  registers no global TLS configuration.
- Migration limitation: MySQL DDL is atomic per supported InnoDB statement but
  is not transactional and implicitly commits. This starter deliberately
  claims only pool and standard SQL capabilities. Applications such as
  Petclinic own locked, idempotent, resumable migration plans instead of
  claiming the transactional migration backend contract.
- Real-system acceptance: the official MySQL 8.4.11 multi-platform image is
  pinned by index digest. Tests prove independent pool ownership, caller-owned
  cancellation, connection recovery, cleanup isolation, and credential-safe
  failures.

Primary references:

- <https://pkg.go.dev/github.com/go-sql-driver/mysql>
- <https://github.com/go-sql-driver/mysql>
- <https://pkg.go.dev/filippo.io/edwards25519>
- <https://dev.mysql.com/doc/refman/8.4/en/atomic-ddl.html>
- <https://hub.docker.com/_/mysql>
