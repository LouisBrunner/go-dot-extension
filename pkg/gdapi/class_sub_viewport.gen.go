// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SubViewport struct {
  obj gdc.ObjectPtr
}

func (me *SubViewport) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SubViewport) BaseClass() string {
  return "SubViewport"
}



// Enums

type SubViewportClearMode int
const (
  SubViewportClearModeClearModeAlways SubViewportClearMode = 0
  SubViewportClearModeClearModeNever SubViewportClearMode = 1
  SubViewportClearModeClearModeOnce SubViewportClearMode = 2
)

type SubViewportUpdateMode int
const (
  SubViewportUpdateModeUpdateDisabled SubViewportUpdateMode = 0
  SubViewportUpdateModeUpdateOnce SubViewportUpdateMode = 1
  SubViewportUpdateModeUpdateWhenVisible SubViewportUpdateMode = 2
  SubViewportUpdateModeUpdateWhenParentVisible SubViewportUpdateMode = 3
  SubViewportUpdateModeUpdateAlways SubViewportUpdateMode = 4
)

func (me *SubViewport) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SubViewport) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SubViewport) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SubViewport) SetSize(size Vector2i, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SubViewport) GetSize() Vector2i {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SubViewport) SetSize2DOverride(size Vector2i, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size_2d_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SubViewport) GetSize2DOverride() Vector2i {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size_2d_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SubViewport) SetSize2DOverrideStretch(enable bool, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size_2d_override_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SubViewport) IsSize2DOverrideStretchEnabled() bool {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_size_2d_override_stretch_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SubViewport) SetUpdateMode(mode SubViewportUpdateMode, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1295690030) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SubViewport) GetUpdateMode() SubViewportUpdateMode {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2980171553) // FIXME: should cache?
  var ret SubViewportUpdateMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SubViewport) SetClearMode(mode SubViewportClearMode, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clear_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2834454712) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SubViewport) GetClearMode() SubViewportClearMode {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clear_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 331324495) // FIXME: should cache?
  var ret SubViewportClearMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SubViewport) GetPropSize() Vector2i {
  panic("TODO: implement")
}

func (me *SubViewport) SetPropSize(value Vector2i) {
  panic("TODO: implement")
}

func (me *SubViewport) GetPropSize2DOverride() Vector2i {
  panic("TODO: implement")
}

func (me *SubViewport) SetPropSize2DOverride(value Vector2i) {
  panic("TODO: implement")
}

func (me *SubViewport) GetPropSize2DOverrideStretch() bool {
  panic("TODO: implement")
}

func (me *SubViewport) SetPropSize2DOverrideStretch(value bool) {
  panic("TODO: implement")
}

func (me *SubViewport) GetPropRenderTargetClearMode() int {
  panic("TODO: implement")
}

func (me *SubViewport) SetPropRenderTargetClearMode(value int) {
  panic("TODO: implement")
}

func (me *SubViewport) GetPropRenderTargetUpdateMode() int {
  panic("TODO: implement")
}

func (me *SubViewport) SetPropRenderTargetUpdateMode(value int) {
  panic("TODO: implement")
}