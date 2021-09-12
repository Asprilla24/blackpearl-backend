SVC=search-movies
export REVISION_ID=latest

build:
	docker build --tag="$(SVC):$(REVISION_ID)" --tag="$(SVC):latest" -f Dockerfile .

run:
	docker-compose up

COMPOSE_TEST=docker-compose -f docker-compose-test.yaml
RUN_TEST=run -e CGO_ENABLED=0 -e PWD=$(CURDIR) \
	-v $(GOPATH)/pkg/mod:/opt/go/pkg/mod \
	-e GO111MODULE=on
test:
	$(COMPOSE_TEST) down --remove-orphans
	$(COMPOSE_TEST) up -d
	$(COMPOSE_TEST) $(RUN_TEST) test go test -p=1 -coverprofile=profile.out ./...
	$(COMPOSE_TEST) down --remove-orphans