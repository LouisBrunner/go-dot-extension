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

type ptrsForCanvasTextureList struct {
	fnSetDiffuseTexture    gdc.MethodBindPtr
	fnGetDiffuseTexture    gdc.MethodBindPtr
	fnSetNormalTexture     gdc.MethodBindPtr
	fnGetNormalTexture     gdc.MethodBindPtr
	fnSetSpecularTexture   gdc.MethodBindPtr
	fnGetSpecularTexture   gdc.MethodBindPtr
	fnSetSpecularColor     gdc.MethodBindPtr
	fnGetSpecularColor     gdc.MethodBindPtr
	fnSetSpecularShininess gdc.MethodBindPtr
	fnGetSpecularShininess gdc.MethodBindPtr
	fnSetTextureFilter     gdc.MethodBindPtr
	fnGetTextureFilter     gdc.MethodBindPtr
	fnSetTextureRepeat     gdc.MethodBindPtr
	fnGetTextureRepeat     gdc.MethodBindPtr
}

var ptrsForCanvasTexture ptrsForCanvasTextureList

func initCanvasTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("CanvasTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_diffuse_texture")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnSetDiffuseTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_diffuse_texture")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnGetDiffuseTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_normal_texture")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnSetNormalTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_normal_texture")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnGetNormalTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_specular_texture")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnSetSpecularTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_specular_texture")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnGetSpecularTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_specular_color")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnSetSpecularColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_specular_color")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnGetSpecularColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_specular_shininess")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnSetSpecularShininess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_specular_shininess")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnGetSpecularShininess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_texture_filter")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnSetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1037999706))
	}
	{
		methodName := StringNameFromStr("get_texture_filter")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnGetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121960042))
	}
	{
		methodName := StringNameFromStr("set_texture_repeat")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnSetTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1716472974))
	}
	{
		methodName := StringNameFromStr("get_texture_repeat")
		defer methodName.Destroy()
		ptrsForCanvasTexture.fnGetTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2667158319))
	}
}

type CanvasTexture struct {
	Texture2D
}

func (me *CanvasTexture) BaseClass() string {
	return "CanvasTexture"
}

func NewCanvasTexture() *CanvasTexture {
	str := StringNameFromStr("CanvasTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CanvasTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CanvasTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CanvasTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CanvasTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CanvasTexture) SetDiffuseTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnSetDiffuseTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasTexture) GetDiffuseTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnGetDiffuseTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasTexture) SetNormalTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnSetNormalTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasTexture) GetNormalTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnGetNormalTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasTexture) SetSpecularTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnSetSpecularTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasTexture) GetSpecularTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnGetSpecularTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasTexture) SetSpecularColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnSetSpecularColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasTexture) GetSpecularColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnGetSpecularColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasTexture) SetSpecularShininess(shininess float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shininess)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnSetSpecularShininess), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasTexture) GetSpecularShininess() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnGetSpecularShininess), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasTexture) SetTextureFilter(filter CanvasItemTextureFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnSetTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasTexture) GetTextureFilter() CanvasItemTextureFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CanvasItemTextureFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnGetTextureFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CanvasTexture) SetTextureRepeat(repeat CanvasItemTextureRepeat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&repeat)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnSetTextureRepeat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasTexture) GetTextureRepeat() CanvasItemTextureRepeat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CanvasItemTextureRepeat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasTexture.fnGetTextureRepeat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
