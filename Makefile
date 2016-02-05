default: build

build:
	docker run -v $(CURDIR):/src -e BUILD_GOARCH="amd64" centurylink/golang-builder-cross
