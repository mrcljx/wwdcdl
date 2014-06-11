PRODUCT=wwdcdl
RELEASES_FOLDER=bin

all: bindata.go ${RELEASES_FOLDER}/${PRODUCT}-linux.tar.gz ${RELEASES_FOLDER}/${PRODUCT}-darwin.tar.gz

.PHONY: install
	
${RELEASES_FOLDER}/linux/${PRODUCT}: *.go
	GOARCH=amd64 GOOS=linux go build -o $@

${RELEASES_FOLDER}/${PRODUCT}-linux.tar.gz: ${RELEASES_FOLDER}/linux/${PRODUCT}
	tar czf $@ -C ${RELEASES_FOLDER}/linux ${PRODUCT}

${RELEASES_FOLDER}/darwin/${PRODUCT}: *.go
	GOARCH=amd64 GOOS=darwin go build -o $@
	
${RELEASES_FOLDER}/${PRODUCT}-darwin.tar.gz: ${RELEASES_FOLDER}/darwin/${PRODUCT}
	tar czf $@ -C ${RELEASES_FOLDER}/darwin ${PRODUCT}

bindata.go: data/*
	go-bindata data
