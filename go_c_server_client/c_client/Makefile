root := ..
include := $(root)/include
lib := $(root)/lib/bin
bin := $(root)/bin

all: c_client

c_client:
	gcc -o $(bin)/c_client c_client.c -I$(include) -L$(lib) -lcServCalls -lclientMakeSvrCall