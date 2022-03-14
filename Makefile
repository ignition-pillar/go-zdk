.PHONY: all clean examples version

GO ?= latest

BUILDDIR = $(shell pwd)/build

examples:
	go build -o $(BUILDDIR)/fake-znnd examples/fake-znnd.go
	go build -o $(BUILDDIR)/account-history examples/account-history.go
	@echo "Build examples done. See ${BUILDDIR} directory."

clean:
	rm -r $(BUILDDIR)/

all: examples
