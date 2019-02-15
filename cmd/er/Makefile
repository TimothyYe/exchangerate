# Binary name
BINARY=er
# Builds the project
build:
		GO111MODULE=on go build -ldflags "-X main.Version=${VERSION}" -o ${BINARY}
		GO111MODULE=on go test -v
# Installs our project: copies binaries
install:
		GO111MODULE=on go install
release:
		# Clean
		go clean
		rm -rf *.gz
		# Build for mac
		GO111MODULE=on go build -ldflags "-X main.Version=${VERSION}"
		tar czvf ${BINARY}-mac64-${VERSION}.tar.gz ./${BINARY}
		# Build for arm
		go clean
		GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-X main.Version=${VERSION}"
		tar czvf ${BINARY}-arm64-${VERSION}.tar.gz ./${BINARY}
		# Build for linux
		go clean
		GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=${VERSION}"
		tar czvf ${BINARY}-linux64-${VERSION}.tar.gz ./${BINARY}
		go clean
# Cleans our projects: deletes binaries
clean:
		go clean
		rm -rf *.gz

.PHONY:  clean build
