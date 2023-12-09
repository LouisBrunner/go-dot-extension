// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightOccluder2D struct {
  obj gdc.ObjectPtr
}

func (me *LightOccluder2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LightOccluder2D) BaseClass() string {
  return "LightOccluder2D"
}



// Enums

func (me *LightOccluder2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LightOccluder2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *LightOccluder2D) SetOccluderPolygon(polygon OccluderPolygon2D, )  {
  panic("TODO: implement")
}

func  (me *LightOccluder2D) GetOccluderPolygon()  {
  panic("TODO: implement")
}

func  (me *LightOccluder2D) SetOccluderLightMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *LightOccluder2D) GetOccluderLightMask()  {
  panic("TODO: implement")
}

func  (me *LightOccluder2D) SetAsSdfCollision(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LightOccluder2D) IsSetAsSdfCollision()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
