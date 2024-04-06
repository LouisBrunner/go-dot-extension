// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForXRCamera3DList struct {
}

var ptrsForXRCamera3D ptrsForXRCamera3DList

func initXRCamera3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRCamera3D")
	defer className.Destroy()

}

type XRCamera3D struct {
	Camera3D
}

func (me *XRCamera3D) BaseClass() string {
	return "XRCamera3D"
}

func NewXRCamera3D() *XRCamera3D {
	str := StringNameFromStr("XRCamera3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRCamera3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *XRCamera3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRCamera3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRCamera3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
