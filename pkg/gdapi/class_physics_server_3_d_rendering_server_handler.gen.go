// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports





  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3DRenderingServerHandler struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3DRenderingServerHandler) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3DRenderingServerHandler) BaseClass() string {
  return "PhysicsServer3DRenderingServerHandler"
}



// Enums

func (me *PhysicsServer3DRenderingServerHandler) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer3DRenderingServerHandler) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer3DRenderingServerHandler) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsServer3DRenderingServerHandler) XSetVertex(vertex_id int, vertices unsafe.Pointer, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DRenderingServerHandler) XSetNormal(vertex_id int, normals unsafe.Pointer, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DRenderingServerHandler) XSetAabb(aabb AABB, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
