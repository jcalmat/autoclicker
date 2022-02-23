PACKAGE     = autoclicker
VERSION    ?= $(shell git describe --tags --always)

GO          = go

V           = 0
Q           = $(if $(filter 1,$V),,@)
M           = $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all
all: tidy build

# Executables
build: tidy ## Build autoclicker in bin
	$(info $(M) building autoclicker…) @
	$Q CGO_ENABLED=1 \
	$(GO) build \
		-ldflags '-X main.version=$(VERSION)' \
		-o $(PACKAGE)

build_windows: tidy
	$(info $(M) building autoclicker for windows…) @
	$Q GOOS=windows \
	GOARCH=386 \
	CC=i686-w64-mingw32-gcc \
	CXX=i686-w64-mingw32-g++ \
	CGO_ENABLED=1 \
	CGO_CXXFLAGS="-static-libgcc -static-libstdc++ -Wl,-Bstatic -lstdc++ -lpthread -Wl,-Bdynamic" \
	$(GO) build -ldflags "-s -w -X main.version=$(VERSION)" \
	-o $(PACKAGE).exe

# Tidy
.PHONY: tidy
tidy: ## Update go.sum with go.mod
	$(info $(M) running mod tidy…) @
	$Q $(GO) mod tidy