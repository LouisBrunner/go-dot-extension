// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *PhysicsServer3DRenderingServerHandler) SetVertex(vertex_id int, vertex Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3DRenderingServerHandler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertex_id), gdc.ConstTypePtr(vertex.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3DRenderingServerHandler) SetNormal(vertex_id int, normal Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3DRenderingServerHandler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertex_id), gdc.ConstTypePtr(normal.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3DRenderingServerHandler) SetAabb(aabb AABB, )  {
  classNameV := StringNameFromStr("PhysicsServer3DRenderingServerHandler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(aabb.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
