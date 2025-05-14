run: ### Run application
	go run .

dev: ### Run application with live debug mode
	air

clean: ### Cleaning
	go mod tidy
	go clean -cache
	go clean -testcache

upgrade: ### Upgrade mods version
	go get -u all

vulncheck: ### Security issues checker
	govulncheck ./...
