// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFMesh struct {
  obj gdc.ObjectPtr
}

func (me *GLTFMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFMesh) BaseClass() string {
  return "GLTFMesh"
}



// Enums

func (me *GLTFMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GLTFMesh) GetMesh()  {
  panic("TODO: implement")
}

func  (me *GLTFMesh) SetMesh(mesh ImporterMesh, )  {
  panic("TODO: implement")
}

func  (me *GLTFMesh) GetBlendWeights()  {
  panic("TODO: implement")
}

func  (me *GLTFMesh) SetBlendWeights(blend_weights PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *GLTFMesh) GetInstanceMaterials()  {
  panic("TODO: implement")
}

func  (me *GLTFMesh) SetInstanceMaterials(instance_materials Material, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
