// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *VisualInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualInstance3D) BaseClass() string {
  return "VisualInstance3D"
}



// Enums

func (me *VisualInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualInstance3D) XGetAabb()  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) SetBase(base RID, )  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) GetBase()  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) GetInstance()  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) SetLayerMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) GetLayerMask()  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) SetLayerMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) GetLayerMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) SetSortingOffset(offset float32, )  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) GetSortingOffset()  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) SetSortingUseAabbCenter(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) IsSortingUseAabbCenter()  {
  panic("TODO: implement")
}

func  (me *VisualInstance3D) GetAabb()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
