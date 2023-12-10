// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Sprite3D struct {
  obj gdc.ObjectPtr
}

func (me *Sprite3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Sprite3D) BaseClass() string {
  return "Sprite3D"
}



// Enums

func (me *Sprite3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Sprite3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Sprite3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Sprite3D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite3D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite3D) SetRegionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite3D) IsRegionEnabled() bool {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_region_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite3D) SetRegionRect(rect Rect2, )  {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite3D) GetRegionRect() Rect2 {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite3D) SetFrame(frame int, )  {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite3D) GetFrame() int {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite3D) SetFrameCoords(coords Vector2i, )  {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(coords.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite3D) GetFrameCoords() Vector2i {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame_coords")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite3D) SetVframes(vframes int, )  {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vframes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vframes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite3D) GetVframes() int {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vframes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Sprite3D) SetHframes(hframes int, )  {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hframes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hframes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Sprite3D) GetHframes() int {
  classNameV := StringNameFromStr("Sprite3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hframes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Sprite3D) GetPropTexture() Texture {
  panic("TODO: implement")
}

func (me *Sprite3D) SetPropTexture(value Texture) {
  panic("TODO: implement")
}

func (me *Sprite3D) GetPropHframes() int {
  panic("TODO: implement")
}

func (me *Sprite3D) SetPropHframes(value int) {
  panic("TODO: implement")
}

func (me *Sprite3D) GetPropVframes() int {
  panic("TODO: implement")
}

func (me *Sprite3D) SetPropVframes(value int) {
  panic("TODO: implement")
}

func (me *Sprite3D) GetPropFrame() int {
  panic("TODO: implement")
}

func (me *Sprite3D) SetPropFrame(value int) {
  panic("TODO: implement")
}

func (me *Sprite3D) GetPropFrameCoords() Vector2 {
  panic("TODO: implement")
}

func (me *Sprite3D) SetPropFrameCoords(value Vector2) {
  panic("TODO: implement")
}

func (me *Sprite3D) GetPropRegionEnabled() bool {
  panic("TODO: implement")
}

func (me *Sprite3D) SetPropRegionEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Sprite3D) GetPropRegionRect() Rect2 {
  panic("TODO: implement")
}

func (me *Sprite3D) SetPropRegionRect(value Rect2) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API