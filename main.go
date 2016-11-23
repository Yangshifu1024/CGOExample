package main

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L. -lhello
void hello();
int sum(int, int);
*/
import "C"
import "fmt"

func main() {
	C.hello()
	fmt.Println(C.sum(C.int(32), C.int(32)))
}
