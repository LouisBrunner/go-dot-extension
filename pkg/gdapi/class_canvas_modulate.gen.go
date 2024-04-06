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

type ptrsForCanvasModulateList struct {
	fnSetColor gdc.MethodBindPtr
	fnGetColor gdc.MethodBindPtr
}

var ptrsForCanvasModulate ptrsForCanvasModulateList

func initCanvasModulatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("CanvasModulate")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_color")
		defer methodName.Destroy()
		ptrsForCanvasModulate.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_color")
		defer methodName.Destroy()
		ptrsForCanvasModulate.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
}

type CanvasModulate struct {
	Node2D
}

func (me *CanvasModulate) BaseClass() string {
	return "CanvasModulate"
}

func NewCanvasModulate() *CanvasModulate {
	str := StringNameFromStr("CanvasModulate") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CanvasModulate{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CanvasModulate) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CanvasModulate) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CanvasModulate) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CanvasModulate) SetColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasModulate.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasModulate) GetColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasModulate.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
