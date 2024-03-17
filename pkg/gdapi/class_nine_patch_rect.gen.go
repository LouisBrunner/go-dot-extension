// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NinePatchRect struct {
  Control
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
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NinePatchRectTextureChangedSignalFn func()

func (me *NinePatchRect) ConnectTextureChanged(subs SignalSubscribers, fn NinePatchRectTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NinePatchRect) DisconnectTextureChanged(subs SignalSubscribers, fn NinePatchRectTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
