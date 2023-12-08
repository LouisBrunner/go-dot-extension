package gdapi

import (
	"context"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Class interface {
	BaseClass() string
	setInternalObject(obj gdc.ObjectPtr)
}

type Context struct {
	context.Context
	iface gdc.Interface
}

func callUtilityFunc(ctx Context, ret gdc.TypePtr, ptr gdc.PtrUtilityFunction, args ...gdc.ConstTypePtr) gdc.TypePtr {
	ctx.iface.CallPtrUtilityFunction(ptr, ret, unsafe.SliceData(args), len(args))
	return ret
}
