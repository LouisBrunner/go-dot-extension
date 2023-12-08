package gdapi

import (
	"context"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Context struct {
	context.Context
	iface gdc.Interface
}

func callUtilityFunc(ctx Context, ret gdc.TypePtr, ptr gdc.PtrUtilityFunction, args ...gdc.ConstTypePtr) gdc.TypePtr {
	ctx.iface.CallPtrUtilityFunction(ptr, ret, &args[0], len(args))
	return ret
}
