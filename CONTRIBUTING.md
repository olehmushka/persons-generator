# Contributing

Thanks for considering a contribution.

## Dev environment

```sh
git clone https://github.com/olehmushka/persons-generator.git
cd persons-generator
make install       # go mod download
```

No infrastructure is needed to work on the pure generation logic
(`engine/entities/*`, `engine/probability_machine`, etc.) — `go run main.go
generate_religion` runs entirely in memory. For the HTTP API / persistent
worlds, see the [README's full-demo section](README.md#full-demo--one-command).

## Before opening a PR

```sh
make test    # go test ./...
make lint    # golangci-lint run
make fmt     # go fmt ./...
```

CI runs the same checks (build, vet, `test -race -cover`, lint) against your
diff — lint only fails on issues introduced by the change itself
(`--new-from-rev`), not the repo's pre-existing findings, so you're not on the
hook for cleaning up code you didn't touch.

## PR expectations

- Keep PRs small and focused on one thing — easier to review, easier to revert
  if something's wrong.
- Add or update tests for behavior changes. Not everything needs a test (pure
  refactors, docs, config), but new logic and bug fixes should come with one.
- Explain the *why* in the PR description, not just the *what* — especially for
  anything that isn't obvious from the diff alone.

## Reporting bugs / requesting features

Use the issue templates. For anything that isn't quite a bug report — questions,
ideas, discussion — use [Discussions](https://github.com/olehmushka/persons-generator/discussions)
instead.

## Reporting a security issue

Please don't open a public issue — see [SECURITY.md](SECURITY.md).

## Code of Conduct

This project follows the [Code of Conduct](CODE_OF_CONDUCT.md).
