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

type ptrsForTextureRectList struct {
	fnSetTexture     gdc.MethodBindPtr
	fnGetTexture     gdc.MethodBindPtr
	fnSetExpandMode  gdc.MethodBindPtr
	fnGetExpandMode  gdc.MethodBindPtr
	fnSetFlipH       gdc.MethodBindPtr
	fnIsFlippedH     gdc.MethodBindPtr
	fnSetFlipV       gdc.MethodBindPtr
	fnIsFlippedV     gdc.MethodBindPtr
	fnSetStretchMode gdc.MethodBindPtr
	fnGetStretchMode gdc.MethodBindPtr
}

var ptrsForTextureRect ptrsForTextureRectList

func initTextureRectPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextureRect")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForTextureRect.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForTextureRect.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_expand_mode")
		defer methodName.Destroy()
		ptrsForTextureRect.fnSetExpandMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1870766882))
	}
	{
		methodName := StringNameFromStr("get_expand_mode")
		defer methodName.Destroy()
		ptrsForTextureRect.fnGetExpandMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3863824733))
	}
	{
		methodName := StringNameFromStr("set_flip_h")
		defer methodName.Destroy()
		ptrsForTextureRect.fnSetFlipH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_flipped_h")
		defer methodName.Destroy()
		ptrsForTextureRect.fnIsFlippedH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_flip_v")
		defer methodName.Destroy()
		ptrsForTextureRect.fnSetFlipV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_flipped_v")
		defer methodName.Destroy()
		ptrsForTextureRect.fnIsFlippedV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_stretch_mode")
		defer methodName.Destroy()
		ptrsForTextureRect.fnSetStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 58788729))
	}
	{
		methodName := StringNameFromStr("get_stretch_mode")
		defer methodName.Destroy()
		ptrsForTextureRect.fnGetStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 346396079))
	}

}

type TextureRect struct {
	Control
}

func (me *TextureRect) BaseClass() string {
	return "TextureRect"
}

func NewTextureRect() *TextureRect {
	str := StringNameFromStr("TextureRect") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextureRect{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TextureRectExpandMode int

const (
	TextureRectExpandModeExpandKeepSize              TextureRectExpandMode = 0
	TextureRectExpandModeExpandIgnoreSize            TextureRectExpandMode = 1
	TextureRectExpandModeExpandFitWidth              TextureRectExpandMode = 2
	TextureRectExpandModeExpandFitWidthProportional  TextureRectExpandMode = 3
	TextureRectExpandModeExpandFitHeight             TextureRectExpandMode = 4
	TextureRectExpandModeExpandFitHeightProportional TextureRectExpandMode = 5
)

type TextureRectStretchMode int

const (
	TextureRectStretchModeStretchScale              TextureRectStretchMode = 0
	TextureRectStretchModeStretchTile               TextureRectStretchMode = 1
	TextureRectStretchModeStretchKeep               TextureRectStretchMode = 2
	TextureRectStretchModeStretchKeepCentered       TextureRectStretchMode = 3
	TextureRectStretchModeStretchKeepAspect         TextureRectStretchMode = 4
	TextureRectStretchModeStretchKeepAspectCentered TextureRectStretchMode = 5
	TextureRectStretchModeStretchKeepAspectCovered  TextureRectStretchMode = 6
)

func (me *TextureRect) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextureRect) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextureRect) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TextureRect) SetTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureRect) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextureRect) SetExpandMode(expand_mode TextureRectExpandMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&expand_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnSetExpandMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureRect) GetExpandMode() TextureRectExpandMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextureRectExpandMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnGetExpandMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextureRect) SetFlipH(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnSetFlipH), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureRect) IsFlippedH() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnIsFlippedH), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureRect) SetFlipV(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnSetFlipV), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureRect) IsFlippedV() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnIsFlippedV), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureRect) SetStretchMode(stretch_mode TextureRectStretchMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnSetStretchMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureRect) GetStretchMode() TextureRectStretchMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextureRectStretchMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureRect.fnGetStretchMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
