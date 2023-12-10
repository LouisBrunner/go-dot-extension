// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureButton struct {
  obj gdc.ObjectPtr
}

func (me *TextureButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureButton) BaseClass() string {
  return "TextureButton"
}



// Enums

type TextureButtonStretchMode int
const (
  TextureButtonStretchModeStretchScale TextureButtonStretchMode = 0
  TextureButtonStretchModeStretchTile TextureButtonStretchMode = 1
  TextureButtonStretchModeStretchKeep TextureButtonStretchMode = 2
  TextureButtonStretchModeStretchKeepCentered TextureButtonStretchMode = 3
  TextureButtonStretchModeStretchKeepAspect TextureButtonStretchMode = 4
  TextureButtonStretchModeStretchKeepAspectCentered TextureButtonStretchMode = 5
  TextureButtonStretchModeStretchKeepAspectCovered TextureButtonStretchMode = 6
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

func  (me *TextureButton) SetTextureNormal(texture Texture2D, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) SetTexturePressed(texture Texture2D, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) SetTextureHover(texture Texture2D, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_hover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) SetTextureDisabled(texture Texture2D, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) SetTextureFocused(texture Texture2D, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_focused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) SetClickMask(mask BitMap, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_click_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 698588216) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mask.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) SetIgnoreTextureSize(ignore bool, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ignore_texture_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) SetStretchMode(mode TextureButtonStretchMode, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 252530840) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) SetFlipH(enable bool, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) IsFlippedH() bool {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) SetFlipV(enable bool, )  {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureButton) IsFlippedV() bool {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) GetTextureNormal() Texture2D {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) GetTexturePressed() Texture2D {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) GetTextureHover() Texture2D {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_hover")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) GetTextureDisabled() Texture2D {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) GetTextureFocused() Texture2D {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_focused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) GetClickMask() BitMap {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_click_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2459671998) // FIXME: should cache?
  var ret BitMap
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) GetIgnoreTextureSize() bool {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ignore_texture_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureButton) GetStretchMode() TextureButtonStretchMode {
  classNameV := StringNameFromStr("TextureButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 33815122) // FIXME: should cache?
  var ret TextureButtonStretchMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *TextureButton) GetPropTextureNormal() Texture2D {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropTextureNormal(value Texture2D) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropTexturePressed() Texture2D {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropTexturePressed(value Texture2D) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropTextureHover() Texture2D {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropTextureHover(value Texture2D) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropTextureDisabled() Texture2D {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropTextureDisabled(value Texture2D) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropTextureFocused() Texture2D {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropTextureFocused(value Texture2D) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropTextureClickMask() BitMap {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropTextureClickMask(value BitMap) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropIgnoreTextureSize() bool {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropIgnoreTextureSize(value bool) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropStretchMode() int {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropStretchMode(value int) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropFlipH() bool {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropFlipH(value bool) {
  panic("TODO: implement")
}

func (me *TextureButton) GetPropFlipV() bool {
  panic("TODO: implement")
}

func (me *TextureButton) SetPropFlipV(value bool) {
  panic("TODO: implement")
}