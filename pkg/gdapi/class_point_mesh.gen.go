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

type ptrsForPointMeshList struct {
}

var ptrsForPointMesh ptrsForPointMeshList

func initPointMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PointMesh")
	defer className.Destroy()

}

type PointMesh struct {
	PrimitiveMesh
}

func (me *PointMesh) BaseClass() string {
	return "PointMesh"
}

func NewPointMesh() *PointMesh {
	str := StringNameFromStr("PointMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PointMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PointMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PointMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PointMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
