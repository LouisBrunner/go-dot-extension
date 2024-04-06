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

type ptrsForCurveTextureList struct {
	fnSetWidth       gdc.MethodBindPtr
	fnSetCurve       gdc.MethodBindPtr
	fnGetCurve       gdc.MethodBindPtr
	fnSetTextureMode gdc.MethodBindPtr
	fnGetTextureMode gdc.MethodBindPtr
}

var ptrsForCurveTexture ptrsForCurveTextureList

func initCurveTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("CurveTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForCurveTexture.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_curve")
		defer methodName.Destroy()
		ptrsForCurveTexture.fnSetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_curve")
		defer methodName.Destroy()
		ptrsForCurveTexture.fnGetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_texture_mode")
		defer methodName.Destroy()
		ptrsForCurveTexture.fnSetTextureMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321955367))
	}
	{
		methodName := StringNameFromStr("get_texture_mode")
		defer methodName.Destroy()
		ptrsForCurveTexture.fnGetTextureMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 715756376))
	}

}

type CurveTexture struct {
	Texture2D
}

func (me *CurveTexture) BaseClass() string {
	return "CurveTexture"
}

func NewCurveTexture() *CurveTexture {
	str := StringNameFromStr("CurveTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CurveTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CurveTextureTextureMode int

const (
	CurveTextureTextureModeTextureModeRgb CurveTextureTextureMode = 0
	CurveTextureTextureModeTextureModeRed CurveTextureTextureMode = 1
)

func (me *CurveTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CurveTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CurveTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CurveTexture) SetWidth(width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveTexture.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CurveTexture) SetCurve(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveTexture.fnSetCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CurveTexture) GetCurve() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveTexture.fnGetCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CurveTexture) SetTextureMode(texture_mode CurveTextureTextureMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveTexture.fnSetTextureMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CurveTexture) GetTextureMode() CurveTextureTextureMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CurveTextureTextureMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurveTexture.fnGetTextureMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
