
.PHONY: api
# generate api
api:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: errors
# generate errors
errors:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) errors'

.PHONY: wire
# generate wire
wire:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: proto
# generate proto
proto:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'

.PHONY: build
# build
build:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) build'

.PHONY: migrate
# migrate
migrate:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) migrate'
