# oort — c12s policy engine

The authorization brain of the control plane: a **RHABAC** engine (resource
hierarchy + attributes + relationships) stored in **Neo4j**, exposed as two
gRPC services — an *evaluator* (`Authorize`,
`GetGrantedPermissions`) and an *administrator* (resource/relationship/
attribute/policy management) — plus an **async administration** endpoint over
NATS (`oort.administration`) used by the other services.

## Role in c12s

```
apollo  register          ──▶ create org→user inheritance + one policy per permission
apollo  VerifyToken (every protected request) ──▶ GetGrantedPermissions → JWT `permissions` claim
kuiper  namespace/config  ──▶ async CreateInheritanceRel (org→resource, config→namespace, ...)
quasar  schema create     ──▶ async CreateInheritanceRel (schema→namespace)
```

The effective per-request authorization is done by services against the
JWT `permissions` claim (magnetar, kuiper, quasar parse `perm|kind|id`
strings); Oort is what *produces* those permissions (via Apollo) and what the
other services *write* inheritance relationships to. Apollo is the only live
caller of the evaluator in the request path — which makes Oort availability
part of API availability (see apollo README).

## API

Evaluator (`pkg/api/proto/evaluator.proto`):

| RPC | Behavior |
|---|---|
| `Authorize` | evaluates one `subject/permission/object` triple against policies + attributes + hierarchy |
| `GetGrantedPermissions` | lists every policy applicable to a subject that currently evaluates to allowed |

Administrator (`pkg/api/proto/administrator.proto`), also available async
via `AdministrationAsyncClient` (NATS subject `oort.administration`, queue
group `oort`, request/reply with callbacks):

`CreateResource`, `DeleteResource`, `CreateInheritanceRel`,
`DeleteInheritanceRel`, `PutAttribute`, `DeleteAttribute`, `CreatePolicy`,
`DeletePolicy`.

## Running

Compose service `oort` (`tools/docker-compose.yml`), gRPC on `OORT_PORT`
(default `8000`), reflection enabled. Depends on Neo4j (bolt) and NATS.

## Configuration (env)

| Variable | Used for |
|---|---|
| `OORT_PORT` | gRPC listen port |
| `NEO4J_HOSTNAME`, `NEO4J_BOLT_PORT` | graph store (`bolt://host:port`) |
| `NEO4J_USERNAME`, `NEO4J_PASSWORD` | optional bolt credentials (lab runs with auth disabled) |
| `NEO4J_DBNAME` | Neo4j database name |
| `NATS_HOSTNAME`, `NATS_PORT`, `NATS_USERNAME`, `NATS_PASSWORD` | async administration |

## Data and dependencies

- **Neo4j** (`neo4jdata` volume, image `neo4j:4.4.12`, APOC enabled, auth
  disabled in the lab): resources, attributes, inheritance relationships and
  policies. Wiping the volume wipes all authorization state — every existing
  user keeps its JWT until it expires but loses all permissions.
- **NATS** for the async administration path (same broker as everything else).
- Nothing else persists locally.

## Status

Stable enough to carry the lab's authorization (registration plus
per-request verification). It is the least-hardened control-plane service:
error swallowing, a subscription leak and no tests (below).

## Bugs, disabled and unfinished features

### Confirmed bugs

| Where | Effect |
|---|---|
| `internal/services/evaluation.go` — `GetGrantedPermissions` | When fetching the **subject's attributes** fails, the function returns `resp.Error` (nil at that point) instead of the actual error: the caller gets an **empty permission list with a success status**. Apollo then mints a JWT without permissions — a transient Neo4j error silently strips a user of all rights instead of surfacing an error. |
| `pkg/api/administrator_async.go` — `SendRequest` | The per-request **reply subscription is leaked**: `subscriber.Unsubscribe()` runs only when the publish fails, never after a successful reply. Every async administration call (each policy created at registration!) leaves a live NATS subscription behind; some callers (apollo) also create a brand-new NATS connection per request. |
| `internal/startup/nats.go`, `pkg/api/administrator_async.go` | NATS connections use **default nats.go options** (no `RetryOnFailedConnect`, reconnects stop after ≈2 min) — same class as review R16 for star/magnetar/gravity, still open for oort. A broker blip can leave the async path deaf while the process looks healthy. |
| Async error paths | Unmarshal/validation failures in the async handler are logged and the request is dropped **without a reply**; callers waiting on the callback never fire (no timeouts on either side). |

### Disabled / unfinished features

- **Permission caching is designed but not implemented**: `services/cache.go`
  defines a `Cache` interface with no implementation, and
  `internal/repos/rhabac/neo4j/cypher_factory.go` contains a large
  commented-out `cachedPermsCypherFactory`. Every evaluation hits Neo4j.
- **Tests are commented out**: `test/store_test.go` and `test/bench_test.go`
  (testcontainers-based Neo4j suite) have no active test functions —
  `go test ./...` reports `no tests to run`. The RHABAC hierarchy/eval logic
  is effectively untested.
- No Neo4j schema/migration management; the graph is whatever the services
  happened to write.

### PoC boundaries (by design, not bugs)

- **No authentication on either gRPC service** — anything that can reach
  oort:8000 can create/delete policies (the gRPC administrator is a god-mode
  API). The lab relies on the compose network boundary (same class as S1).
- Attribute conditions are evaluated in-process; condition expressions are
  generally empty in the lab (policies are plain allows).
- The evaluator recomputes everything per call (see cache gap) — fine at lab
  scale, a known scalability limit.

References: `../plan/architecture.md` (auth chain), `../plan/review.md`
(R16 class), `../apollo/README.md` (the main consumer).
