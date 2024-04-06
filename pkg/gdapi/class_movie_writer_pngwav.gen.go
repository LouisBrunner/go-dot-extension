// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MovieWriterPNGWAV struct {
	obj gdc.ObjectPtr
}

func (me *MovieWriterPNGWAV) SetBaseObject(obj gdc.ObjectPtr) {
	me.obj = obj
}

func (me *MovieWriterPNGWAV) BaseClass() string {
	return "MovieWriterPNGWAV"
}

// Enums

func (me *MovieWriterPNGWAV) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MovieWriterPNGWAV) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MovieWriterPNGWAV) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
