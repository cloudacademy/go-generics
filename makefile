GO := GO111MODULE=on go

.PHONY: all
all: collections constraints functions interfaces

.PHONY: collections
collections: list set

.PHONY: constraints
constraints: basic comparable custom inlined

.PHONY: functions
functions: printer filter

.PHONY: interfaces
interfaces: worker

# ======================
# collections

.PHONY: list
list:
	@echo "running list examples..."
	@$(GO) run ./cmd/collections/list/main.go
	@echo

.PHONY: set
set:
	@echo "running set examples..."
	@$(GO) run ./cmd/collections/set/main.go
	@echo

# ======================
# constraints

.PHONY: basic
basic:
	@echo "running basic examples..."
	@$(GO) run ./cmd/constraints/basic/main.go
	@echo

.PHONY: comparable
comparable:
	@echo "running comparable examples..."
	@$(GO) run ./cmd/constraints/comparable/main.go
	@echo

.PHONY: custom
custom:
	@echo "running custom examples..."
	@$(GO) run ./cmd/constraints/custom/main.go
	@echo

.PHONY: inlined
inlined:
	@echo "running inlined examples..."
	@$(GO) run ./cmd/constraints/inlined/main.go
	@echo

# ======================
# functions

.PHONY: printer
printer:
	@echo "running printer examples..."
	@$(GO) run ./cmd/functions/printer/main.go
	@echo

.PHONY: filter
filter:
	@echo "running filter examples..."
	@$(GO) run ./cmd/functions/filter/main.go
	@echo

# ======================
# interfaces:

.PHONY: worker
worker:
	@echo "running worker examples..."
	@$(GO) run ./cmd/interfaces/worker/main.go
	@echo