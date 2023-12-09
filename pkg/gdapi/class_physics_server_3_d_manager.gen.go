// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3DManager struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3DManager) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3DManager) BaseClass() string {
  return "PhysicsServer3DManager"
}



// Enums

func (me *PhysicsServer3DManager) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer3DManager) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer3DManager) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsServer3DManager) RegisterServer(name String, create_callback Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DManager) SetDefaultServer(name String, priority int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
