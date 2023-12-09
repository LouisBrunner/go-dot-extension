// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OccluderInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *OccluderInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OccluderInstance3D) BaseClass() string {
  return "OccluderInstance3D"
}



// Enums

func (me *OccluderInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OccluderInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *OccluderInstance3D) SetBakeMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *OccluderInstance3D) GetBakeMask()  {
  panic("TODO: implement")
}

func  (me *OccluderInstance3D) SetBakeMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *OccluderInstance3D) GetBakeMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *OccluderInstance3D) SetBakeSimplificationDistance(simplification_distance float32, )  {
  panic("TODO: implement")
}

func  (me *OccluderInstance3D) GetBakeSimplificationDistance()  {
  panic("TODO: implement")
}

func  (me *OccluderInstance3D) SetOccluder(occluder Occluder3D, )  {
  panic("TODO: implement")
}

func  (me *OccluderInstance3D) GetOccluder()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
