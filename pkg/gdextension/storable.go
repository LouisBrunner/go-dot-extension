package gdextension

import (
	"fmt"
	"unsafe"

	"github.com/mattn/go-pointer"
)

type isStorable any

func store[T isStorable](me T) unsafe.Pointer {
	return pointer.Save(me)
}

func restore[T isStorable](p unsafe.Pointer) (T, error) {
	v, cast := pointer.Restore(p).(T)
	if !cast {
		return v, fmt.Errorf("could not restore pointer for %T", v)
	}
	return v, nil
}
