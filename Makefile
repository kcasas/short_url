# What is .PHONY? https://unix.stackexchange.com/a/321947
.PHONY: run_local

#####################################
# Build, test & lint
#####################################

run_dev:
	go run cmd/web/server.go