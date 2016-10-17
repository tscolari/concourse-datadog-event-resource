image: linux-bins
	docker build --rm -t tscolari/datadog-event-resource .

linux-bins:
	mkdir -p output
	GOOS=linux go build -o output/check ./check/main.go
	GOOS=linux go build -o output/in ./in/main.go
	GOOS=linux go build -o output/out ./out/main.go
