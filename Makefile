BINARY=dictionary-of-chinese

VERSION=0.0.1

BUILD=`date +%FT%T%z`

default:
	go build -o ${BINARY}  -tags=jsoniter

install:
	govendor sync -v

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY:  clean