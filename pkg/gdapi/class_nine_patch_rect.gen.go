// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NinePatchRect struct {
  obj gdc.ObjectPtr
}

func (me *NinePatchRect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NinePatchRect) BaseClass() string {
  return "NinePatchRect"
}



// Enums

type NinePatchRectAxisStretchMode int
const (
  NinePatchRectAxisStretchModeAxisStretchModeStretch NinePatchRectAxisStretchMode = 0
  NinePatchRectAxisStretchModeAxisStretchModeTile NinePatchRectAxisStretchMode = 1
  NinePatchRectAxisStretchModeAxisStretchModeTileFit NinePatchRectAxisStretchMode = 2
)

func (me *NinePatchRect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NinePatchRect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NinePatchRect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NinePatchRect) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NinePatchRect) GetTexture() Texture2D {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NinePatchRect) SetPatchMargin(margin Side, value int, )  {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_patch_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 437707142) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NinePatchRect) GetPatchMargin(margin Side, ) int {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_patch_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1983885014) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NinePatchRect) SetRegionRect(rect Rect2, )  {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NinePatchRect) GetRegionRect() Rect2 {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NinePatchRect) SetDrawCenter(draw_center bool, )  {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_center), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NinePatchRect) IsDrawCenterEnabled() bool {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draw_center_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NinePatchRect) SetHAxisStretchMode(mode NinePatchRectAxisStretchMode, )  {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_axis_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3219608417) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NinePatchRect) GetHAxisStretchMode() NinePatchRectAxisStretchMode {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_axis_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3317113799) // FIXME: should cache?
  var ret NinePatchRectAxisStretchMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NinePatchRect) SetVAxisStretchMode(mode NinePatchRectAxisStretchMode, )  {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_axis_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3219608417) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NinePatchRect) GetVAxisStretchMode() NinePatchRectAxisStretchMode {
  classNameV := StringNameFromStr("NinePatchRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_axis_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3317113799) // FIXME: should cache?
  var ret NinePatchRectAxisStretchMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *NinePatchRect) GetPropTexture() Texture2D {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *NinePatchRect) GetPropDrawCenter() bool {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropDrawCenter(value bool) {
  panic("TODO: implement")
}

func (me *NinePatchRect) GetPropRegionRect() Rect2 {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropRegionRect(value Rect2) {
  panic("TODO: implement")
}

func (me *NinePatchRect) GetPropPatchMarginLeft() int {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropPatchMarginLeft(value int) {
  panic("TODO: implement")
}

func (me *NinePatchRect) GetPropPatchMarginTop() int {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropPatchMarginTop(value int) {
  panic("TODO: implement")
}

func (me *NinePatchRect) GetPropPatchMarginRight() int {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropPatchMarginRight(value int) {
  panic("TODO: implement")
}

func (me *NinePatchRect) GetPropPatchMarginBottom() int {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropPatchMarginBottom(value int) {
  panic("TODO: implement")
}

func (me *NinePatchRect) GetPropAxisStretchHorizontal() int {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropAxisStretchHorizontal(value int) {
  panic("TODO: implement")
}

func (me *NinePatchRect) GetPropAxisStretchVertical() int {
  panic("TODO: implement")
}

func (me *NinePatchRect) SetPropAxisStretchVertical(value int) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API