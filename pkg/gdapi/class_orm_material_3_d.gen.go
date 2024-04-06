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

type ptrsForORMMaterial3DList struct {
}

var ptrsForORMMaterial3D ptrsForORMMaterial3DList

func initORMMaterial3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ORMMaterial3D")
	defer className.Destroy()

}

type ORMMaterial3D struct {
	BaseMaterial3D
}

func (me *ORMMaterial3D) BaseClass() string {
	return "ORMMaterial3D"
}

func NewORMMaterial3D() *ORMMaterial3D {
	str := StringNameFromStr("ORMMaterial3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ORMMaterial3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ORMMaterial3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ORMMaterial3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ORMMaterial3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
