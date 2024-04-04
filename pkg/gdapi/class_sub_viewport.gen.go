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

type SubViewport struct {
  Viewport
}

func (me *SubViewport) BaseClass() string {
  return "SubViewport"
}

func NewSubViewport() *SubViewport {
  str := StringNameFromStr("SubViewport") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SubViewport{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) GetSize() Vector2i {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SubViewport) SetSize2DOverride(size Vector2i, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size_2d_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) GetSize2DOverride() Vector2i {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size_2d_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SubViewport) SetSize2DOverrideStretch(enable bool, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size_2d_override_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) IsSize2DOverrideStretchEnabled() bool {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_size_2d_override_stretch_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SubViewport) SetUpdateMode(mode SubViewportUpdateMode, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1295690030) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) GetUpdateMode() SubViewportUpdateMode {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2980171553) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SubViewportUpdateMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SubViewport) SetClearMode(mode SubViewportClearMode, )  {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clear_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2834454712) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) GetClearMode() SubViewportClearMode {
  classNameV := StringNameFromStr("SubViewport")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clear_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 331324495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SubViewportClearMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
