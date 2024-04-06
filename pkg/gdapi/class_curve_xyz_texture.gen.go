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

type ptrsForCurveXYZTextureList struct {
	fnSetWidth  gdc.MethodBindPtr
	fnSetCurveX gdc.MethodBindPtr
	fnGetCurveX gdc.MethodBindPtr
	fnSetCurveY gdc.MethodBindPtr
	fnGetCurveY gdc.MethodBindPtr
	fnSetCurveZ gdc.MethodBindPtr
	fnGetCurveZ gdc.MethodBindPtr
}

var ptrsForCurveXYZTexture ptrsForCurveXYZTextureList

func initCurveXYZTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("CurveXYZTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForCurveXYZTexture.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_curve_x")
		defer methodName.Destroy()
		ptrsForCurveXYZTexture.fnSetCurveX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_curve_x")
		defer methodName.Destroy()
		ptrsForCurveXYZTexture.fnGetCurveX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_curve_y")
		defer methodName.Destroy()
		ptrsForCurveXYZTexture.fnSetCurveY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_curve_y")
		defer methodName.Destroy()
		ptrsForCurveXYZTexture.fnGetCurveY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_curve_z")
		defer methodName.Destroy()
		ptrsForCurveXYZTexture.fnSetCurveZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_curve_z")
		defer methodName.Destroy()
		ptrsForCurveXYZTexture.fnGetCurveZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
}

type CurveXYZTexture struct {
	Texture2D
}

func (me *CurveXYZTexture) BaseClass() string {
	return "CurveXYZTexture"
}

func NewCurveXYZTexture() *CurveXYZTexture {
	str := StringNameFromStr("CurveXYZTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CurveXYZTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CurveXYZTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CurveXYZTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CurveXYZTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CurveXYZTexture) SetWidth(width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveXYZTexture.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CurveXYZTexture) SetCurveX(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveXYZTexture.fnSetCurveX), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CurveXYZTexture) GetCurveX() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveXYZTexture.fnGetCurveX), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CurveXYZTexture) SetCurveY(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveXYZTexture.fnSetCurveY), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CurveXYZTexture) GetCurveY() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveXYZTexture.fnGetCurveY), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CurveXYZTexture) SetCurveZ(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveXYZTexture.fnSetCurveZ), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CurveXYZTexture) GetCurveZ() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveXYZTexture.fnGetCurveZ), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
