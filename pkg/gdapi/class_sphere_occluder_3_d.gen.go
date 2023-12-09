// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SphereOccluder3D struct {
  obj gdc.ObjectPtr
}

func (me *SphereOccluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SphereOccluder3D) BaseClass() string {
  return "SphereOccluder3D"
}



// Enums

func (me *SphereOccluder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SphereOccluder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SphereOccluder3D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *SphereOccluder3D) GetRadius()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
