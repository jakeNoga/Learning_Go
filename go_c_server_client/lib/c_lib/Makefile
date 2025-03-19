root := ../..

lib := $(root)/lib/bin/
include := $(root)/include/

all: libc_serv_calls.so

libc_serv_calls.so:
	@echo "Building the library"
	gcc -c -I$(include) c_serv_calls.c -o c_serv_calls.o
	gcc -shared -o $(lib)/libcServCalls.so c_serv_calls.o -L$(lib) -lclientMakeSvrCall

# Not tested yet need to debug and add to the all
libc_serv_calls.a:
	@echo "Building the library"
	gcc -c -I$(include) c_serv_calls.c -o c_serv_calls.o
	ar rcs libcServCalls.a c_serv_calls.o -L$(lib) -lclientMakeSvrCall