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

type ptrsForSubViewportList struct {
  fnSetSize gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetSize2DOverride gdc.MethodBindPtr
  fnGetSize2DOverride gdc.MethodBindPtr
  fnSetSize2DOverrideStretch gdc.MethodBindPtr
  fnIsSize2DOverrideStretchEnabled gdc.MethodBindPtr
  fnSetUpdateMode gdc.MethodBindPtr
  fnGetUpdateMode gdc.MethodBindPtr
  fnSetClearMode gdc.MethodBindPtr
  fnGetClearMode gdc.MethodBindPtr
}

var ptrsForSubViewport ptrsForSubViewportList

func initSubViewportPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SubViewport")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForSubViewport.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForSubViewport.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_size_2d_override")
    defer methodName.Destroy()
    ptrsForSubViewport.fnSetSize2DOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_size_2d_override")
    defer methodName.Destroy()
    ptrsForSubViewport.fnGetSize2DOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_size_2d_override_stretch")
    defer methodName.Destroy()
    ptrsForSubViewport.fnSetSize2DOverrideStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_size_2d_override_stretch_enabled")
    defer methodName.Destroy()
    ptrsForSubViewport.fnIsSize2DOverrideStretchEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_update_mode")
    defer methodName.Destroy()
    ptrsForSubViewport.fnSetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1295690030))
  }
  {
    methodName := StringNameFromStr("get_update_mode")
    defer methodName.Destroy()
    ptrsForSubViewport.fnGetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2980171553))
  }
  {
    methodName := StringNameFromStr("set_clear_mode")
    defer methodName.Destroy()
    ptrsForSubViewport.fnSetClearMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2834454712))
  }
  {
    methodName := StringNameFromStr("get_clear_mode")
    defer methodName.Destroy()
    ptrsForSubViewport.fnGetClearMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 331324495))
  }
}

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
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) GetSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SubViewport) SetSize2DOverride(size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnSetSize2DOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) GetSize2DOverride() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnGetSize2DOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SubViewport) SetSize2DOverrideStretch(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnSetSize2DOverrideStretch), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) IsSize2DOverrideStretchEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnIsSize2DOverrideStretchEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SubViewport) SetUpdateMode(mode SubViewportUpdateMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnSetUpdateMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) GetUpdateMode() SubViewportUpdateMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SubViewportUpdateMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnGetUpdateMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SubViewport) SetClearMode(mode SubViewportClearMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnSetClearMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewport) GetClearMode() SubViewportClearMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SubViewportClearMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewport.fnGetClearMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
