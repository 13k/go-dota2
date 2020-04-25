SUBDIRS = apigen cso socache
GO_GENERATE = apigen cso socache

.PHONY: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C "$@" $(MAKECMDGOALS)

.PHONY: go-generate
go-generate: $(GO_GENERATE)

.PHONY: generate-api
generate-api: apigen

.PHONY: lint
lint:
	golangci-lint run
