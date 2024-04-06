// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type IPUnix struct {
	obj gdc.ObjectPtr
}

func (me *IPUnix) SetBaseObject(obj gdc.ObjectPtr) {
	me.obj = obj
}

func (me *IPUnix) BaseClass() string {
	return "IPUnix"
}

// Enums

func (me *IPUnix) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *IPUnix) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *IPUnix) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
