.PHONY: dylib so
build-dynamic:
	gcc -c -Wall -Werror -fPIC hello.c

dylib: build-dynamic
	gcc -dynamiclib -o libhello.dylib hello.o
	rm -rf *.o
	mv libhello.dylib ../libhello.dylib

so: build-dynamic
	gcc -shared -W1 -o libhello.so hello.o
	rm -rf *.o
	mv libhello.so ../libhello.so

dll: build-dynamic
	gcc -shared -W1 -o hello.dll hello.o
	rm -rf *.o
	mv hello.dll ../hello.dll