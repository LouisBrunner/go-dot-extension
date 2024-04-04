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

type PhysicsServer2DManager struct {
  Object
}

func (me *PhysicsServer2DManager) BaseClass() string {
  return "PhysicsServer2DManager"
}

func NewPhysicsServer2DManager() *PhysicsServer2DManager {
  str := StringNameFromStr("PhysicsServer2DManager") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsServer2DManager{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), create_callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsServer2DManager) SetDefaultServer(name String, priority int64, )  {
  classNameV := StringNameFromStr("PhysicsServer2DManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956805083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
