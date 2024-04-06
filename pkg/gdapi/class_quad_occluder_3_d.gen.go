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

type ptrsForQuadOccluder3DList struct {
	fnSetSize gdc.MethodBindPtr
	fnGetSize gdc.MethodBindPtr
}

var ptrsForQuadOccluder3D ptrsForQuadOccluder3DList

func initQuadOccluder3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("QuadOccluder3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForQuadOccluder3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForQuadOccluder3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}

}

type QuadOccluder3D struct {
	Occluder3D
}

func (me *QuadOccluder3D) BaseClass() string {
	return "QuadOccluder3D"
}

func NewQuadOccluder3D() *QuadOccluder3D {
	str := StringNameFromStr("QuadOccluder3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &QuadOccluder3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *QuadOccluder3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *QuadOccluder3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *QuadOccluder3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *QuadOccluder3D) SetSize(size Vector2) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForQuadOccluder3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *QuadOccluder3D) GetSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForQuadOccluder3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
