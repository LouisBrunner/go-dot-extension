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

type VisibleOnScreenEnabler3D struct {
  VisibleOnScreenNotifier3D
}

func (me *VisibleOnScreenEnabler3D) BaseClass() string {
  return "VisibleOnScreenEnabler3D"
}

func NewVisibleOnScreenEnabler3D() *VisibleOnScreenEnabler3D {
  str := StringNameFromStr("VisibleOnScreenEnabler3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisibleOnScreenEnabler3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisibleOnScreenEnabler3DEnableMode int
const (
  VisibleOnScreenEnabler3DEnableModeEnableModeInherit VisibleOnScreenEnabler3DEnableMode = 0
  VisibleOnScreenEnabler3DEnableModeEnableModeAlways VisibleOnScreenEnabler3DEnableMode = 1
  VisibleOnScreenEnabler3DEnableModeEnableModeWhenPaused VisibleOnScreenEnabler3DEnableMode = 2
)

func (me *VisibleOnScreenEnabler3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisibleOnScreenEnabler3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenEnabler3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisibleOnScreenEnabler3D) SetEnableMode(mode VisibleOnScreenEnabler3DEnableMode, )  {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 320303646) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisibleOnScreenEnabler3D) GetEnableMode() VisibleOnScreenEnabler3DEnableMode {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3352990031) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisibleOnScreenEnabler3DEnableMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisibleOnScreenEnabler3D) SetEnableNodePath(path NodePath, )  {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_node_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisibleOnScreenEnabler3D) GetEnableNodePath() NodePath {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_node_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 277076166) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
