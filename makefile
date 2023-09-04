# Installation paths
PREFIX = /usr/local
BINDIR = $(PREFIX)/bin
MANDIR = $(PREFIX)/share/man/man1

# Man page source and destination
MAN_SOURCE = man/when.1
MAN_DEST = $(MANDIR)/when.1

# Binary name
BINARY = when

# Define variables
GOCMD = go
GOFLAGS = -v
GOTEST = $(GOCMD) test
PKGS = ./...

install:
	@mkdir -p $(BINDIR) $(MANDIR)
	@$(GOCMD) build -o $(BINDIR)/$(BINARY) ./cmd/when
	@cp $(MAN_SOURCE) $(MAN_DEST)
	@echo "$(BINARY) installed to $(BINDIR)/$(BINARY)"
	@echo "$(MAN_SOURCE) man page installed to $(MAN_DEST)"

uninstall:
	@rm -f $(BINDIR)/$(BINARY) $(MAN_DEST)
	@echo "$(BINARY) and $(MAN_SOURCE) man page removed"

build:
	@$(GOCMD) build -o $(BINARY) ./cmd/when

# Target to run all Go tests
test:
	$(GOTEST) $(GOFLAGS) $(PKGS)

.PHONY: install uninstall build test
