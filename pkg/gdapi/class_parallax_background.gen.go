// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ParallaxBackground struct {
  obj gdc.ObjectPtr
}

func (me *ParallaxBackground) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ParallaxBackground) BaseClass() string {
  return "ParallaxBackground"
}



// Enums

func (me *ParallaxBackground) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ParallaxBackground) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ParallaxBackground) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ParallaxBackground) SetScrollOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParallaxBackground) GetScrollOffset() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParallaxBackground) SetScrollBaseOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_base_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParallaxBackground) GetScrollBaseOffset() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_base_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParallaxBackground) SetScrollBaseScale(scale Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParallaxBackground) GetScrollBaseScale() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParallaxBackground) SetLimitBegin(offset Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_limit_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParallaxBackground) GetLimitBegin() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_limit_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParallaxBackground) SetLimitEnd(offset Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_limit_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParallaxBackground) GetLimitEnd() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_limit_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ParallaxBackground) SetIgnoreCameraZoom(ignore bool, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ignore_camera_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ParallaxBackground) IsIgnoreCameraZoom() bool {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ignore_camera_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
