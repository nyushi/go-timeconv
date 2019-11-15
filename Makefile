SOURCES := $(shell find . -name '*.go')
COVER_HTML := cover.html
COVER_FILE := cov.out

test: $(SOURCES)
	go test ./... -v -cover

$(COVER_HTML): $(COVER_FILE)
	go tool cover -html=cov.out -o cover.html

$(COVER_FILE): $(SOURCES)
	go test ./... -v -coverprofile=cov.out

clean:
	rm -rf $(COVER_FILE) $(COVER_HTML)

.PHONY: test clean
