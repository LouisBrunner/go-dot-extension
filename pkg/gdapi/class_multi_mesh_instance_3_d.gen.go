// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMeshInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *MultiMeshInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiMeshInstance3D) BaseClass() string {
  return "MultiMeshInstance3D"
}



// Enums

func (me *MultiMeshInstance3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiMeshInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiMeshInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MultiMeshInstance3D) SetMultimesh(multimesh MultiMesh, )  {
  panic("TODO: implement")
}

func  (me *MultiMeshInstance3D) GetMultimesh()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
