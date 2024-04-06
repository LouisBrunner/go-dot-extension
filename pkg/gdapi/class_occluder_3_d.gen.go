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

type ptrsForOccluder3DList struct {
	fnGetVertices gdc.MethodBindPtr
	fnGetIndices  gdc.MethodBindPtr
}

var ptrsForOccluder3D ptrsForOccluder3DList

func initOccluder3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Occluder3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_vertices")
		defer methodName.Destroy()
		ptrsForOccluder3D.fnGetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
	}
	{
		methodName := StringNameFromStr("get_indices")
		defer methodName.Destroy()
		ptrsForOccluder3D.fnGetIndices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}

}

type Occluder3D struct {
	Resource
}

func (me *Occluder3D) BaseClass() string {
	return "Occluder3D"
}

func NewOccluder3D() *Occluder3D {
	str := StringNameFromStr("Occluder3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Occluder3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Occluder3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Occluder3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Occluder3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Occluder3D) GetVertices() PackedVector3Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluder3D.fnGetVertices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Occluder3D) GetIndices() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluder3D.fnGetIndices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
