test:
	go test ./... --cover

run:
	go run ./cmd -order=$(order) -plugin=$(plugin) -path=$(path)