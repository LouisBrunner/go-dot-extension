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

type ptrsForTextureButtonList struct {
	fnSetTextureNormal     gdc.MethodBindPtr
	fnSetTexturePressed    gdc.MethodBindPtr
	fnSetTextureHover      gdc.MethodBindPtr
	fnSetTextureDisabled   gdc.MethodBindPtr
	fnSetTextureFocused    gdc.MethodBindPtr
	fnSetClickMask         gdc.MethodBindPtr
	fnSetIgnoreTextureSize gdc.MethodBindPtr
	fnSetStretchMode       gdc.MethodBindPtr
	fnSetFlipH             gdc.MethodBindPtr
	fnIsFlippedH           gdc.MethodBindPtr
	fnSetFlipV             gdc.MethodBindPtr
	fnIsFlippedV           gdc.MethodBindPtr
	fnGetTextureNormal     gdc.MethodBindPtr
	fnGetTexturePressed    gdc.MethodBindPtr
	fnGetTextureHover      gdc.MethodBindPtr
	fnGetTextureDisabled   gdc.MethodBindPtr
	fnGetTextureFocused    gdc.MethodBindPtr
	fnGetClickMask         gdc.MethodBindPtr
	fnGetIgnoreTextureSize gdc.MethodBindPtr
	fnGetStretchMode       gdc.MethodBindPtr
}

var ptrsForTextureButton ptrsForTextureButtonList

func initTextureButtonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextureButton")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture_normal")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetTextureNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("set_texture_pressed")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetTexturePressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("set_texture_hover")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetTextureHover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("set_texture_disabled")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetTextureDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("set_texture_focused")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetTextureFocused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("set_click_mask")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetClickMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 698588216))
	}
	{
		methodName := StringNameFromStr("set_ignore_texture_size")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetIgnoreTextureSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_stretch_mode")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 252530840))
	}
	{
		methodName := StringNameFromStr("set_flip_h")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetFlipH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_flipped_h")
		defer methodName.Destroy()
		ptrsForTextureButton.fnIsFlippedH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_flip_v")
		defer methodName.Destroy()
		ptrsForTextureButton.fnSetFlipV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_flipped_v")
		defer methodName.Destroy()
		ptrsForTextureButton.fnIsFlippedV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_texture_normal")
		defer methodName.Destroy()
		ptrsForTextureButton.fnGetTextureNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("get_texture_pressed")
		defer methodName.Destroy()
		ptrsForTextureButton.fnGetTexturePressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("get_texture_hover")
		defer methodName.Destroy()
		ptrsForTextureButton.fnGetTextureHover = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("get_texture_disabled")
		defer methodName.Destroy()
		ptrsForTextureButton.fnGetTextureDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("get_texture_focused")
		defer methodName.Destroy()
		ptrsForTextureButton.fnGetTextureFocused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("get_click_mask")
		defer methodName.Destroy()
		ptrsForTextureButton.fnGetClickMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2459671998))
	}
	{
		methodName := StringNameFromStr("get_ignore_texture_size")
		defer methodName.Destroy()
		ptrsForTextureButton.fnGetIgnoreTextureSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_stretch_mode")
		defer methodName.Destroy()
		ptrsForTextureButton.fnGetStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 33815122))
	}

}

type TextureButton struct {
	BaseButton
}

func (me *TextureButton) BaseClass() string {
	return "TextureButton"
}

func NewTextureButton() *TextureButton {
	str := StringNameFromStr("TextureButton") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextureButton{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TextureButtonStretchMode int

const (
	TextureButtonStretchModeStretchScale              TextureButtonStretchMode = 0
	TextureButtonStretchModeStretchTile               TextureButtonStretchMode = 1
	TextureButtonStretchModeStretchKeep               TextureButtonStretchMode = 2
	TextureButtonStretchModeStretchKeepCentered       TextureButtonStretchMode = 3
	TextureButtonStretchModeStretchKeepAspect         TextureButtonStretchMode = 4
	TextureButtonStretchModeStretchKeepAspectCentered TextureButtonStretchMode = 5
	TextureButtonStretchModeStretchKeepAspectCovered  TextureButtonStretchMode = 6
)

func (me *TextureButton) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextureButton) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextureButton) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TextureButton) SetTextureNormal(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetTextureNormal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) SetTexturePressed(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetTexturePressed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) SetTextureHover(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetTextureHover), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) SetTextureDisabled(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetTextureDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) SetTextureFocused(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetTextureFocused), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) SetClickMask(mask BitMap) {
	cargs := []gdc.ConstTypePtr{mask.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetClickMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) SetIgnoreTextureSize(ignore bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetIgnoreTextureSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) SetStretchMode(mode TextureButtonStretchMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetStretchMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) SetFlipH(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetFlipH), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) IsFlippedH() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnIsFlippedH), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureButton) SetFlipV(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnSetFlipV), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureButton) IsFlippedV() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnIsFlippedV), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureButton) GetTextureNormal() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnGetTextureNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextureButton) GetTexturePressed() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnGetTexturePressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextureButton) GetTextureHover() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnGetTextureHover), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextureButton) GetTextureDisabled() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnGetTextureDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextureButton) GetTextureFocused() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnGetTextureFocused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextureButton) GetClickMask() BitMap {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBitMap()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnGetClickMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextureButton) GetIgnoreTextureSize() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnGetIgnoreTextureSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureButton) GetStretchMode() TextureButtonStretchMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextureButtonStretchMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureButton.fnGetStretchMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
