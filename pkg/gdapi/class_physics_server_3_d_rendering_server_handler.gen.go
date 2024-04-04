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

type PhysicsServer3DRenderingServerHandler struct {
  Object
}

func (me *PhysicsServer3DRenderingServerHandler) BaseClass() string {
  return "PhysicsServer3DRenderingServerHandler"
}

func NewPhysicsServer3DRenderingServerHandler() *PhysicsServer3DRenderingServerHandler {
  str := StringNameFromStr("PhysicsServer3DRenderingServerHandler") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsServer3DRenderingServerHandler{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *PhysicsServer3DRenderingServerHandler) SetVertex(vertex_id int64, vertex Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3DRenderingServerHandler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertex_id) , vertex.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsServer3DRenderingServerHandler) SetNormal(vertex_id int64, normal Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3DRenderingServerHandler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertex_id) , normal.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsServer3DRenderingServerHandler) SetAabb(aabb AABB, )  {
  classNameV := StringNameFromStr("PhysicsServer3DRenderingServerHandler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
