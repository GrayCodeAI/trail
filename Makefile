.PHONY: check fixtures fmt simulate test test-race vet

check: fmt test vet fixtures

fmt:
	gofmt -w $$(find cmd internal -name '*.go' -type f)

test:
	go test ./...

test-race:
	go test -race ./...

vet:
	go vet ./...

fixtures:
	@set -e; \
	events_tmp=$$(mktemp); \
	commands_tmp=$$(mktemp); \
	trap 'rm -f "$$events_tmp" "$$commands_tmp"' EXIT; \
	go run ./cmd/trail-sim -mode events -output "$$events_tmp"; \
	go run ./cmd/trail-sim -mode scenario -output "$$commands_tmp"; \
	cmp spec/fixtures/v0/valid/safe-recovery-events.json "$$events_tmp"; \
	cmp spec/fixtures/v0/valid/safe-recovery-commands.json "$$commands_tmp"

simulate:
	go run ./cmd/trail-sim
