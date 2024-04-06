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

type ptrsForPointLight2DList struct {
	fnSetTexture       gdc.MethodBindPtr
	fnGetTexture       gdc.MethodBindPtr
	fnSetTextureOffset gdc.MethodBindPtr
	fnGetTextureOffset gdc.MethodBindPtr
	fnSetTextureScale  gdc.MethodBindPtr
	fnGetTextureScale  gdc.MethodBindPtr
}

var ptrsForPointLight2D ptrsForPointLight2DList

func initPointLight2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PointLight2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForPointLight2D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForPointLight2D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_texture_offset")
		defer methodName.Destroy()
		ptrsForPointLight2D.fnSetTextureOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_texture_offset")
		defer methodName.Destroy()
		ptrsForPointLight2D.fnGetTextureOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_texture_scale")
		defer methodName.Destroy()
		ptrsForPointLight2D.fnSetTextureScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_texture_scale")
		defer methodName.Destroy()
		ptrsForPointLight2D.fnGetTextureScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type PointLight2D struct {
	Light2D
}

func (me *PointLight2D) BaseClass() string {
	return "PointLight2D"
}

func NewPointLight2D() *PointLight2D {
	str := StringNameFromStr("PointLight2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PointLight2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PointLight2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PointLight2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PointLight2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PointLight2D) SetTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPointLight2D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PointLight2D) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPointLight2D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PointLight2D) SetTextureOffset(texture_offset Vector2) {
	cargs := []gdc.ConstTypePtr{texture_offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPointLight2D.fnSetTextureOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PointLight2D) GetTextureOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPointLight2D.fnGetTextureOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PointLight2D) SetTextureScale(texture_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPointLight2D.fnSetTextureScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PointLight2D) GetTextureScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPointLight2D.fnGetTextureScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
