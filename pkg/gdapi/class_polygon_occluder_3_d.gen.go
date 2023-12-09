// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PolygonOccluder3D struct {
  obj gdc.ObjectPtr
}

func (me *PolygonOccluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PolygonOccluder3D) BaseClass() string {
  return "PolygonOccluder3D"
}



// Enums

func (me *PolygonOccluder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PolygonOccluder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PolygonOccluder3D) SetPolygon(polygon PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *PolygonOccluder3D) GetPolygon()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
