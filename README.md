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
| `NEO4J_USERNAME`, `NEO4J_PASSWORD` | optional bolt credentials (the compose stack runs with auth disabled) |
| `NEO4J_DBNAME` | Neo4j database name |
| `NATS_HOSTNAME`, `NATS_PORT`, `NATS_USERNAME`, `NATS_PASSWORD` | async administration |

## Data and dependencies

- **Neo4j** (`neo4jdata` volume, image `neo4j:4.4.12`, APOC enabled, auth
  disabled in the compose stack): resources, attributes, inheritance relationships and
  policies. Wiping the volume wipes all authorization state — every existing
  user keeps its JWT until it expires but loses all permissions.
- **NATS** for the async administration path (same broker as everything else).
- Nothing else persists locally.
