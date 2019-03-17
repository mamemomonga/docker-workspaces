.PHONY: all

PROGNAME := workspace
SRCDIR   := ./workspace
BINDIR   := bin

SRC = $(shell find $(SRCDIR) -type f)
BUILDER = go build -o $@ $(SRCDIR)

all: \
	$(BINDIR)/$(PROGNAME)-darwin-amd64 \
	$(BINDIR)/$(PROGNAME)-linux-amd64 \
	$(BINDIR)/$(PROGNAME)-windows-amd64.exe

clean:
	rm -rf $(BINDIR)

deps:
	go get -v $(SRCDIR)

$(BINDIR)/$(PROGNAME)-darwin-amd64: $(SRC)
	GOOS=darwin GOARCH=amd64 $(BUILDER)

$(BINDIR)/$(PROGNAME)-linux-amd64: $(SRC)
	GOOS=linux GOARCH=amd64 $(BUILDER)

$(BINDIR)/$(PROGNAME)-windows-amd64.exe: $(SRC)
	GOOS=windows GOARCH=amd64 $(BUILDER)


