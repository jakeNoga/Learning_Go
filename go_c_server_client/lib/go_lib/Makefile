root := ../..

all: clientMakeSvrCall

clientMakeSvrCall:
	@echo "Building the library"
	go build -o libclientMakeSvrCall.a -buildmode=c-archive clientMakeSvrCall.go
	go build -o libclientMakeSvrCall.so -buildmode=c-shared clientMakeSvrCall.go
	chmod 755 libclientMakeSvrCall.so
	# Now copying the library to the verious places it needs to be in order for the rest of the build to find it.
	cp libclientMakeSvrCall.so $(root)/lib/bin/
	cp libclientMakeSvrCall.a $(root)/lib/bin/

	# The h file is auto generated so it needs to be copied over to the tpd-include directory.
	cp libclientMakeSvrCall.h $(root)/include/clientMakeSvrCall.h