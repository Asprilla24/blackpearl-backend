SVC=search-movies
export REVISION_ID=latest

build:
	docker build --tag="$(SVC):$(REVISION_ID)" --tag="$(SVC):latest" -f Dockerfile .

run:
	docker-compose up