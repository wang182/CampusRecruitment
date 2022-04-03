GOBIN=go

GIT_TAG=$(shell git describe --tags --abbrev=0 --always)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GOLDFLAGS="-X CampusRecruitment/pkg/consts.VERSION=$(GIT_TAG) -X CampusRecruitment/pkg/consts.COMMIT=$(GIT_COMMIT)"

GOBUILD=$(GOBIN) build -tags "jsoniter,$(BUILD_TAGS)" -ldflags $(GOLDFLAGS)
GORUN=$(GOBIN) run -tags "jsoniter,$(BUILD_TAGS)"

APP_NAME=CampusRecruitment


.PHONY: all build run-serve image

all: build
build: swag-docs
	$(GOBUILD) -o bin/$(APP_NAME) main.go

run-serve: 
	$(GORUN) ./main.go serve 

swag-docs:
	GOOS="" go get github.com/swaggo/swag/cmd/swag && \
	swag init -g pkg/web/apiv1/route.go

image:
	docker build -t $(APP_NAME):$(GIT_TAG) .