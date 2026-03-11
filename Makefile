COVERAGE_DIR ?= .coverage.out

test:
	@if [ -d $(COVERAGE_DIR) ]; then rm -r $(COVERAGE_DIR); fi
	@mkdir -p $(COVERAGE_DIR)
	$(MAKE) test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
