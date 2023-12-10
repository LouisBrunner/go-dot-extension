// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Sprite2D struct {
  obj gdc.ObjectPtr
}

func (me *Sprite2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Sprite2D) BaseClass() string {
  return "Sprite2D"
}



// Enums

func (me *Sprite2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Sprite2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Sprite2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Sprite2D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetCentered(centered bool, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&centered), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) IsCentered() bool {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) GetOffset() Vector2 {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetFlipH(flip_h bool, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_h), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) IsFlippedH() bool {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetFlipV(flip_v bool, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_v), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) IsFlippedV() bool {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetRegionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) IsRegionEnabled() bool {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_region_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) IsPixelOpaque(pos Vector2, ) bool {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_pixel_opaque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 556197845) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pos.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetRegionRect(rect Rect2, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) GetRegionRect() Rect2 {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetRegionFilterClipEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_filter_clip_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) IsRegionFilterClipEnabled() bool {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_region_filter_clip_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetFrame(frame int, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) GetFrame() int {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetFrameCoords(coords Vector2i, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) GetFrameCoords() Vector2i {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetVframes(vframes int, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vframes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vframes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) GetVframes() int {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vframes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) SetHframes(hframes int, )  {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hframes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hframes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite2D) GetHframes() int {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hframes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite2D) GetRect() Rect2 {
  classNameV := StringNameFromStr("Sprite2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Sprite2D) GetPropTexture() Texture2D {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropCentered() bool {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropCentered(value bool) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropOffset() Vector2 {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropFlipH() bool {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropFlipH(value bool) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropFlipV() bool {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropFlipV(value bool) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropHframes() int {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropHframes(value int) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropVframes() int {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropVframes(value int) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropFrame() int {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropFrame(value int) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropFrameCoords() Vector2i {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropFrameCoords(value Vector2i) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropRegionEnabled() bool {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropRegionEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropRegionRect() Rect2 {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropRegionRect(value Rect2) {
  panic("TODO: implement")
}

func (me *Sprite2D) GetPropRegionFilterClipEnabled() bool {
  panic("TODO: implement")
}

func (me *Sprite2D) SetPropRegionFilterClipEnabled(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API