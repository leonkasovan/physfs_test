# Set source c files from directory packages/physfs
srcFiles=packages/physfs/*.c

# Set CFLAGS
CFLAGS= -Ipackages/physfs

all: c go

c: $(srcFiles) src/test_physfs.c
	gcc -o c_test_physfs src/test_physfs.c $(srcFiles) $(CFLAGS)

go: src/test_physfs.go
	export CGO_ENABLED=1 && go build -trimpath -ldflags="-s -w" -v -o go_test_physfs src/test_physfs.go
