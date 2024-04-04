// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type ParallaxBackground struct {
  CanvasLayer
}

func (me *ParallaxBackground) BaseClass() string {
  return "ParallaxBackground"
}

func NewParallaxBackground() *ParallaxBackground {
  str := StringNameFromStr("ParallaxBackground") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ParallaxBackground{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxBackground) GetScrollOffset() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParallaxBackground) SetScrollBaseOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_base_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxBackground) GetScrollBaseOffset() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_base_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParallaxBackground) SetScrollBaseScale(scale Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxBackground) GetScrollBaseScale() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_base_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParallaxBackground) SetLimitBegin(offset Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_limit_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxBackground) GetLimitBegin() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_limit_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParallaxBackground) SetLimitEnd(offset Vector2, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_limit_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxBackground) GetLimitEnd() Vector2 {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_limit_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ParallaxBackground) SetIgnoreCameraZoom(ignore bool, )  {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ignore_camera_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ParallaxBackground) IsIgnoreCameraZoom() bool {
  classNameV := StringNameFromStr("ParallaxBackground")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ignore_camera_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
