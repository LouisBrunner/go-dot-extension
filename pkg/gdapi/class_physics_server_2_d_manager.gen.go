// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsServer2DManager struct {
  Object
}

func (me *PhysicsServer2DManager) BaseClass() string {
  return "PhysicsServer2DManager"
}



// Enums

func (me *PhysicsServer2DManager) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer2DManager) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer2DManager) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsServer2DManager) RegisterServer(name String, create_callback Callable, )  {
  classNameV := StringNameFromStr("PhysicsServer2DManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2137474292) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(create_callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2DManager) SetDefaultServer(name String, priority int, )  {
  classNameV := StringNameFromStr("PhysicsServer2DManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956805083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
