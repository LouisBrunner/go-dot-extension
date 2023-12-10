// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisibleOnScreenEnabler3D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenEnabler3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenEnabler3D) BaseClass() string {
  return "VisibleOnScreenEnabler3D"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisibleOnScreenEnabler3D) GetEnableMode() VisibleOnScreenEnabler3DEnableMode {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3352990031) // FIXME: should cache?
  var ret VisibleOnScreenEnabler3DEnableMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisibleOnScreenEnabler3D) SetEnableNodePath(path NodePath, )  {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_node_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisibleOnScreenEnabler3D) GetEnableNodePath() NodePath {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_node_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 277076166) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisibleOnScreenEnabler3D) GetPropEnableMode() int {
  panic("TODO: implement")
}

func (me *VisibleOnScreenEnabler3D) SetPropEnableMode(value int) {
  panic("TODO: implement")
}

func (me *VisibleOnScreenEnabler3D) GetPropEnableNodePath() NodePath {
  panic("TODO: implement")
}

func (me *VisibleOnScreenEnabler3D) SetPropEnableNodePath(value NodePath) {
  panic("TODO: implement")
}