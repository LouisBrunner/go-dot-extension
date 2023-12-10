// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisibleOnScreenEnabler2D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenEnabler2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenEnabler2D) BaseClass() string {
  return "VisibleOnScreenEnabler2D"
}



// Enums

type VisibleOnScreenEnabler2DEnableMode int
const (
  VisibleOnScreenEnabler2DEnableModeEnableModeInherit VisibleOnScreenEnabler2DEnableMode = 0
  VisibleOnScreenEnabler2DEnableModeEnableModeAlways VisibleOnScreenEnabler2DEnableMode = 1
  VisibleOnScreenEnabler2DEnableModeEnableModeWhenPaused VisibleOnScreenEnabler2DEnableMode = 2
)

func (me *VisibleOnScreenEnabler2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisibleOnScreenEnabler2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenEnabler2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisibleOnScreenEnabler2D) SetEnableMode(mode VisibleOnScreenEnabler2DEnableMode, )  {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961788752) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisibleOnScreenEnabler2D) GetEnableMode() VisibleOnScreenEnabler2DEnableMode {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2650445576) // FIXME: should cache?
  var ret VisibleOnScreenEnabler2DEnableMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisibleOnScreenEnabler2D) SetEnableNodePath(path NodePath, )  {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_node_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisibleOnScreenEnabler2D) GetEnableNodePath() NodePath {
  classNameV := StringNameFromStr("VisibleOnScreenEnabler2D")
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

func (me *VisibleOnScreenEnabler2D) GetPropEnableMode() int {
  panic("TODO: implement")
}

func (me *VisibleOnScreenEnabler2D) SetPropEnableMode(value int) {
  panic("TODO: implement")
}

func (me *VisibleOnScreenEnabler2D) GetPropEnableNodePath() NodePath {
  panic("TODO: implement")
}

func (me *VisibleOnScreenEnabler2D) SetPropEnableNodePath(value NodePath) {
  panic("TODO: implement")
}