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

type ptrsForCanvasGroupList struct {
	fnSetFitMargin   gdc.MethodBindPtr
	fnGetFitMargin   gdc.MethodBindPtr
	fnSetClearMargin gdc.MethodBindPtr
	fnGetClearMargin gdc.MethodBindPtr
	fnSetUseMipmaps  gdc.MethodBindPtr
	fnIsUsingMipmaps gdc.MethodBindPtr
}

var ptrsForCanvasGroup ptrsForCanvasGroupList

func initCanvasGroupPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CanvasGroup")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_fit_margin")
		defer methodName.Destroy()
		ptrsForCanvasGroup.fnSetFitMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fit_margin")
		defer methodName.Destroy()
		ptrsForCanvasGroup.fnGetFitMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_clear_margin")
		defer methodName.Destroy()
		ptrsForCanvasGroup.fnSetClearMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_clear_margin")
		defer methodName.Destroy()
		ptrsForCanvasGroup.fnGetClearMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_use_mipmaps")
		defer methodName.Destroy()
		ptrsForCanvasGroup.fnSetUseMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_mipmaps")
		defer methodName.Destroy()
		ptrsForCanvasGroup.fnIsUsingMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type CanvasGroup struct {
	Node2D
}

func (me *CanvasGroup) BaseClass() string {
	return "CanvasGroup"
}

func NewCanvasGroup() *CanvasGroup {
	str := StringNameFromStr("CanvasGroup") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CanvasGroup{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CanvasGroup) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CanvasGroup) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CanvasGroup) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CanvasGroup) SetFitMargin(fit_margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fit_margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasGroup.fnSetFitMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasGroup) GetFitMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasGroup.fnGetFitMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasGroup) SetClearMargin(clear_margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clear_margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasGroup.fnSetClearMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasGroup) GetClearMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasGroup.fnGetClearMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasGroup) SetUseMipmaps(use_mipmaps bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_mipmaps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasGroup.fnSetUseMipmaps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasGroup) IsUsingMipmaps() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasGroup.fnIsUsingMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
