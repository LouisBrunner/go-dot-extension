// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ArrayOccluder3D struct {
  obj gdc.ObjectPtr
}

func (me *ArrayOccluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ArrayOccluder3D) BaseClass() string {
  return "ArrayOccluder3D"
}



// Enums

func (me *ArrayOccluder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ArrayOccluder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ArrayOccluder3D) SetArrays(vertices PackedVector3Array, indices PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *ArrayOccluder3D) SetVertices(vertices PackedVector3Array, )  {
  panic("TODO: implement")
}

func  (me *ArrayOccluder3D) SetIndices(indices PackedInt32Array, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
