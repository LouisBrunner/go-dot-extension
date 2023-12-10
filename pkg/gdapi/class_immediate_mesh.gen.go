// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ImmediateMesh struct {
  obj gdc.ObjectPtr
}

func (me *ImmediateMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImmediateMesh) BaseClass() string {
  return "ImmediateMesh"
}



// Enums

func (me *ImmediateMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImmediateMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImmediateMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ImmediateMesh) SurfaceBegin(primitive MeshPrimitiveType, material Material, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3716480242) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive), gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) SurfaceSetColor(color Color, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) SurfaceSetNormal(normal Vector3, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(normal.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) SurfaceSetTangent(tangent Plane, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_tangent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3505987427) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tangent.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) SurfaceSetUv(uv Vector2, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(uv.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) SurfaceSetUv2(uv2 Vector2, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_uv2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(uv2.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) SurfaceAddVertex(vertex Vector3, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_add_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vertex.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) SurfaceAddVertex2D(vertex Vector2, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_add_vertex_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vertex.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) SurfaceEnd()  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImmediateMesh) ClearSurfaces()  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_surfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties