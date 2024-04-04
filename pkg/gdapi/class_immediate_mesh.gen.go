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

type ImmediateMesh struct {
  Mesh
}

func (me *ImmediateMesh) BaseClass() string {
  return "ImmediateMesh"
}

func NewImmediateMesh() *ImmediateMesh {
  str := StringNameFromStr("ImmediateMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ImmediateMesh{}
  obj.SetBaseObject(objPtr)
  return obj
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
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2794442543) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive) , material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) SurfaceSetColor(color Color, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) SurfaceSetNormal(normal Vector3, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{normal.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) SurfaceSetTangent(tangent Plane, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_tangent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3505987427) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{tangent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) SurfaceSetUv(uv Vector2, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{uv.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) SurfaceSetUv2(uv2 Vector2, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_uv2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{uv2.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) SurfaceAddVertex(vertex Vector3, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_add_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{vertex.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) SurfaceAddVertex2D(vertex Vector2, )  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_add_vertex_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{vertex.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) SurfaceEnd()  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImmediateMesh) ClearSurfaces()  {
  classNameV := StringNameFromStr("ImmediateMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_surfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
