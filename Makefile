configure:
	go mod tidy
	rm -rf vendor
	rm -rf monigo

local_run:
	go mod tidy && go mod vendor
	go run main.go
