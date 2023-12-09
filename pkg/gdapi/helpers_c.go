package gdapi

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

func cmalloc(size C.size_t) unsafe.Pointer {
	return C.malloc(size)
}

func cfree(ptr unsafe.Pointer) {
	C.free(ptr)
}
