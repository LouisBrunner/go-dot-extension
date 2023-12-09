// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxMesh struct {
  obj gdc.ObjectPtr
}

func (me *BoxMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoxMesh) BaseClass() string {
  return "BoxMesh"
}



// Enums

func (me *BoxMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BoxMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BoxMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *BoxMesh) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *BoxMesh) GetSize()  {
  panic("TODO: implement")
}

func  (me *BoxMesh) SetSubdivideWidth(subdivide int, )  {
  panic("TODO: implement")
}

func  (me *BoxMesh) GetSubdivideWidth()  {
  panic("TODO: implement")
}

func  (me *BoxMesh) SetSubdivideHeight(divisions int, )  {
  panic("TODO: implement")
}

func  (me *BoxMesh) GetSubdivideHeight()  {
  panic("TODO: implement")
}

func  (me *BoxMesh) SetSubdivideDepth(divisions int, )  {
  panic("TODO: implement")
}

func  (me *BoxMesh) GetSubdivideDepth()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
