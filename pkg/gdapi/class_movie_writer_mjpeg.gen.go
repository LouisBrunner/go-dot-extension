// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MovieWriterMJPEG struct {
	obj gdc.ObjectPtr
}

func (me *MovieWriterMJPEG) SetBaseObject(obj gdc.ObjectPtr) {
	me.obj = obj
}

func (me *MovieWriterMJPEG) BaseClass() string {
	return "MovieWriterMJPEG"
}

// Enums

func (me *MovieWriterMJPEG) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MovieWriterMJPEG) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MovieWriterMJPEG) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
