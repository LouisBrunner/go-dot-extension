// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *ImmediateMesh) SurfaceSetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *ImmediateMesh) SurfaceSetNormal(normal Vector3, )  {
  panic("TODO: implement")
}

func  (me *ImmediateMesh) SurfaceSetTangent(tangent Plane, )  {
  panic("TODO: implement")
}

func  (me *ImmediateMesh) SurfaceSetUv(uv Vector2, )  {
  panic("TODO: implement")
}

func  (me *ImmediateMesh) SurfaceSetUv2(uv2 Vector2, )  {
  panic("TODO: implement")
}

func  (me *ImmediateMesh) SurfaceAddVertex(vertex Vector3, )  {
  panic("TODO: implement")
}

func  (me *ImmediateMesh) SurfaceAddVertex2D(vertex Vector2, )  {
  panic("TODO: implement")
}

func  (me *ImmediateMesh) SurfaceEnd()  {
  panic("TODO: implement")
}

func  (me *ImmediateMesh) ClearSurfaces()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
