# CONSTANTS
NAME=javawrapper
MAJOR=2
MINOR=0

#Check for build number
ifeq ($(BUILD_NUMBER), )
BUILD=99999999
else
BUILD=$(BUILD_NUMBER)
endif


#Create version
VERSION=${MAJOR}.${MINOR}.${BUILD}
BIN=bin
BINDIR="${BIN}/${VERSION}"


build: compileLinux compileOsx compileWin

prepare:
	mkdir -p ${BINDIR}

compileLinux: prepare
	GOOS=linux GOARCH=386 go build -o ${BINDIR}/javawrapper-v${VERSION}-linux



compileWin: prepare
	cp resources/versioninfo.json .
	cp resources/manifest.template javawrapper.exe.manifest
	sed -i 's/@MAJOR@/${MAJOR}/' versioninfo.json
	sed -i 's/@MINOR@/${MINOR}/' versioninfo.json
	sed -i 's/@BUILD@/${BUILD}/' versioninfo.json
	sed -i 's/@VERSION@/${VERSION}/' versioninfo.json
	sed -i 's/@VERSION@/${VERSION}/' javawrapper.exe.manifest
	go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	go generate
	GOOS=windows G0ARCH=386 go build -o ${BINDIR}/javawrapper-v${VERSION}-win
	rm versioninfo.json
	rm javawrapper.exe.manifest
	rm *.syso

compileOsx: prepare
	GOOS=darwin GOARCH=amd64 go build -o ${BINDIR}/javawrapper-v${VERSION}-osx


clean:
	rm -rf ${BIN}

