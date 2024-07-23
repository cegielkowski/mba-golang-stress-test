.PHONY: build run

build:
	docker build -t load-tester .

run:
	docker run --rm load-tester --url=http://example.com --requests=1000 --concurrency=10
