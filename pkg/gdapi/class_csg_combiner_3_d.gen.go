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

type ptrsForCSGCombiner3DList struct {
}

var ptrsForCSGCombiner3D ptrsForCSGCombiner3DList

func initCSGCombiner3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CSGCombiner3D")
	defer className.Destroy()
}

type CSGCombiner3D struct {
	CSGShape3D
}

func (me *CSGCombiner3D) BaseClass() string {
	return "CSGCombiner3D"
}

func NewCSGCombiner3D() *CSGCombiner3D {
	str := StringNameFromStr("CSGCombiner3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CSGCombiner3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CSGCombiner3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CSGCombiner3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CSGCombiner3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
