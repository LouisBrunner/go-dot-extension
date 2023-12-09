// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMeshInstance2D struct {
  obj gdc.ObjectPtr
}

func (me *MultiMeshInstance2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiMeshInstance2D) BaseClass() string {
  return "MultiMeshInstance2D"
}



// Enums

func (me *MultiMeshInstance2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiMeshInstance2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MultiMeshInstance2D) SetMultimesh(multimesh MultiMesh, )  {
  panic("TODO: implement")
}

func  (me *MultiMeshInstance2D) GetMultimesh()  {
  panic("TODO: implement")
}

func  (me *MultiMeshInstance2D) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *MultiMeshInstance2D) GetTexture()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
