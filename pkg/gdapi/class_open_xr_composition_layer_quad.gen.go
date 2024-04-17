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

type ptrsForOpenXRCompositionLayerQuadList struct {
	fnSetQuadSize gdc.MethodBindPtr
	fnGetQuadSize gdc.MethodBindPtr
}

var ptrsForOpenXRCompositionLayerQuad ptrsForOpenXRCompositionLayerQuadList

func initOpenXRCompositionLayerQuadPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRCompositionLayerQuad")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_quad_size")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerQuad.fnSetQuadSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_quad_size")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerQuad.fnGetQuadSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}

}

type OpenXRCompositionLayerQuad struct {
	OpenXRCompositionLayer
}

func (me *OpenXRCompositionLayerQuad) BaseClass() string {
	return "OpenXRCompositionLayerQuad"
}

func NewOpenXRCompositionLayerQuad() *OpenXRCompositionLayerQuad {
	str := StringNameFromStr("OpenXRCompositionLayerQuad") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRCompositionLayerQuad{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OpenXRCompositionLayerQuad) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRCompositionLayerQuad) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRCompositionLayerQuad) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRCompositionLayerQuad) SetQuadSize(size Vector2) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerQuad.fnSetQuadSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerQuad) GetQuadSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerQuad.fnGetQuadSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
