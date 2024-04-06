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

type ptrsForQuadMeshList struct {
}

var ptrsForQuadMesh ptrsForQuadMeshList

func initQuadMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("QuadMesh")
	defer className.Destroy()

}

type QuadMesh struct {
	PlaneMesh
}

func (me *QuadMesh) BaseClass() string {
	return "QuadMesh"
}

func NewQuadMesh() *QuadMesh {
	str := StringNameFromStr("QuadMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &QuadMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *QuadMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *QuadMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *QuadMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
