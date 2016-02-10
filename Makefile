default: build

guard-%:
	@ if [ "${${*}}" == "" ]; then \
		echo "Environment variable $* not set"; \
		exit 1; \
	fi

build:
	docker run -v $(CURDIR):/src -e BUILD_GOARCH="amd64" centurylink/golang-builder-cross

release: guard-GITHUB_TOKEN guard-VERSION build
	git tag v$(VERSION) && git push origin v$(VERSION)
	docker run -e GITHUB_TOKEN=$(GITHUB_TOKEN) jimdo/github-release release --user Jimdo --repo docker-machine-fs-notify --tag v$(VERSION)
	docker run -e GITHUB_TOKEN=$(GITHUB_TOKEN) -v $(CURDIR):/src -w /src jimdo/github-release upload --user Jimdo --repo docker-machine-fs-notify --tag v$(VERSION) --name "docker-machine-fs-notify-$(VERSION)-darwin-amd64" --file docker-machine-fs-notify-darwin-amd64
	docker run -e GITHUB_TOKEN=$(GITHUB_TOKEN) -v $(CURDIR):/src -w /src jimdo/github-release upload --user Jimdo --repo docker-machine-fs-notify --tag v$(VERSION) --name "docker-machine-fs-notify-$(VERSION)-linux-amd64" --file docker-machine-fs-notify-linux-amd64
