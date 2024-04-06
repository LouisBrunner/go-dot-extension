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

type ptrsForPlaceholderMeshList struct {
	fnSetAabb gdc.MethodBindPtr
}

var ptrsForPlaceholderMesh ptrsForPlaceholderMeshList

func initPlaceholderMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PlaceholderMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_aabb")
		defer methodName.Destroy()
		ptrsForPlaceholderMesh.fnSetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259215842))
	}

}

type PlaceholderMesh struct {
	Mesh
}

func (me *PlaceholderMesh) BaseClass() string {
	return "PlaceholderMesh"
}

func NewPlaceholderMesh() *PlaceholderMesh {
	str := StringNameFromStr("PlaceholderMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PlaceholderMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PlaceholderMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PlaceholderMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PlaceholderMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PlaceholderMesh) SetAabb(aabb AABB) {
	cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaceholderMesh.fnSetAabb), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
