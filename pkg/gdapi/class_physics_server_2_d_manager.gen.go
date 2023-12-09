// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer2DManager struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer2DManager) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer2DManager) BaseClass() string {
  return "PhysicsServer2DManager"
}



// Enums

func (me *PhysicsServer2DManager) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer2DManager) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsServer2DManager) RegisterServer(name String, create_callback Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DManager) SetDefaultServer(name String, priority int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
