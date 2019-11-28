# What is .PHONY? https://unix.stackexchange.com/a/321947
.PHONY: run

#####################################
# Build, test & lint
#####################################

run:
	go run cmd/web/server.go