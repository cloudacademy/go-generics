all: collections constraints functions interfaces

collections: list set

constraints: basic comparable custom inlined

functions: printer filter

interfaces: worker

# ======================
# collections

list:
	go run ./cmd/collections/list/main.go

set:
	go run ./cmd/collections/set/main.go

# ======================
# constraints

basic:
	go run ./cmd/constraints/basic/main.go

comparable:
	go run ./cmd/constraints/comparable/main.go

custom:
	go run ./cmd/constraints/custom/main.go

inlined:
	go run ./cmd/constraints/inlined/main.go

# ======================
# functions

printer:
	go run ./cmd/functions/printer/main.go

filter:
	go run ./cmd/functions/filter/main.go

# ======================
# interfaces:

worker:
	go run ./cmd/interfaces/worker/main.go

