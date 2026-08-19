# Persons Generator

A procedural worldbuilding engine written in Go. It generates cultures, religions,
languages, and individual people — complete with phenotype presets, temperament,
family relationships, and a year-by-year population simulation (aging, marriage,
births, deaths) across a coordinate grid ("world").

## What it does

- **Cultures & religions**: generates culture trees rooted in configurable ancient/
  medieval "abstract cultures," and matching religions with doctrine, deities,
  clergy, taboos, rituals, and holy scripture traits.
- **Languages**: procedural word generation per culture, built from a library of
  ~40 real-world word-base sources (for flavor, not for translation).
- **People**: individual persons with phenotype (body/face/hair/size/skin),
  temperament, sexual orientation, and multi-generational family trees.
- **Worlds**: a coordinate grid of locations, each seeded with population, that
  can be simulated year-by-year — people age, marry, have children, and die.
- Exposed two ways: a one-shot **CLI** for quick generation, and a persistent
  **HTTP API** backed by MongoDB (storage) and Redis (async world-generation
  progress, via a pub/sub worker).

## Architecture

```
cli/ , handlers/http/  ──▶  internal/{culture,language,persons,religion,world}/
                             (fx-wired services → adapters → engine)
                                          │
                                          ▼
                             engine/orchestrator
                                          │
                                          ▼
                     engine/entities/{culture,religion,language,person,world}
                                          │
                                          ▼
                          core/{mongodb,redis}  (persistence, pub/sub)
```

`engine/entities/*` holds the actual generation logic (pure, mostly DB-free) and
`engine/orchestrator` is the thin layer that also knows how to persist and read
that data back from MongoDB/Redis. `internal/*` is a ports-and-adapters layer on
top of the orchestrator, wired together with [uber-go/fx](https://github.com/uber-go/fx)
for dependency injection; `handlers/http` and `cli` are the two entrypoints.

## Quick start (no infrastructure required)

`generate_religion` runs entirely in memory — no MongoDB or Redis needed:

```sh
go run main.go generate_religion
```

```
Religion (name=Pekolarinism)
Type=polytheism(monolatry)
Dominated gender=male(moderate)
...
```

## Full demo (HTTP API + persistent worlds)

```sh
cp .env.sample .env
docker-compose up -d          # starts MongoDB
make run_http_server          # go run main.go http_server_run
```

See [`postman_collection/persons-generator.postman_collection.json`](postman_collection/persons-generator.postman_collection.json)
for the full request/response catalogue.

`generate_world` (the CLI equivalent of running a world simulation end to end)
additionally needs Redis for progress tracking — see `core/redis` and
`engine/orchestrator/world_methods.go`'s `saveProgress`.

## Development

```sh
make install       # go mod download
make test          # go test ./...
make test_coverage # go test -cover ./...
make lint          # golangci-lint run
make fmt           # go fmt ./...
```

## Known limitations

- `World.seekPartners` re-scans the whole location grid for every unmarried
  person, every simulated year — fine for small worlds, not optimized for
  scale.
- Test coverage currently favors pure/leaf logic (generation algorithms,
  serialization) over handler- and orchestrator-level coverage.
- The person-generation phenotype presets (`engine/entities/person/*/presets`)
  use codename-based categories (terrain/biome names) rather than real-world
  demographic labels, by design.

## License

[GPL-3.0](LICENSE)
