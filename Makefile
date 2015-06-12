PRODUCT=wwdcdl
RELEASES_FOLDER=bin

all: bindata.go darwin linux checksums

checksums: ${RELEASES_FOLDER}/CHECKSUMS

darwin: ${RELEASES_FOLDER}/${PRODUCT}-darwin.tar.gz

linux: ${RELEASES_FOLDER}/${PRODUCT}-linux.tar.gz

.PHONY: install

${RELEASES_FOLDER}/linux/${PRODUCT}: *.go
	GOARCH=amd64 GOOS=linux go build -o $@

${RELEASES_FOLDER}/${PRODUCT}-linux.tar.gz: ${RELEASES_FOLDER}/linux/${PRODUCT}
	tar czf $@ -C ${RELEASES_FOLDER}/linux ${PRODUCT}

${RELEASES_FOLDER}/darwin/${PRODUCT}: *.go
	GOARCH=amd64 GOOS=darwin go build -o $@

${RELEASES_FOLDER}/${PRODUCT}-darwin.tar.gz: ${RELEASES_FOLDER}/darwin/${PRODUCT}
	tar czf $@ -C ${RELEASES_FOLDER}/darwin ${PRODUCT}

${RELEASES_FOLDER}/CHECKSUMS: ${RELEASES_FOLDER}/*.tar.gz
	shasum -a 256 $? > $@

bindata.go: data/*
	go-bindata data
