// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HeightMapShape3D struct {
  obj gdc.ObjectPtr
}

func (me *HeightMapShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HeightMapShape3D) BaseClass() string {
  return "HeightMapShape3D"
}



// Enums

func (me *HeightMapShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HeightMapShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *HeightMapShape3D) SetMapWidth(width int, )  {
  panic("TODO: implement")
}

func  (me *HeightMapShape3D) GetMapWidth()  {
  panic("TODO: implement")
}

func  (me *HeightMapShape3D) SetMapDepth(height int, )  {
  panic("TODO: implement")
}

func  (me *HeightMapShape3D) GetMapDepth()  {
  panic("TODO: implement")
}

func  (me *HeightMapShape3D) SetMapData(data PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *HeightMapShape3D) GetMapData()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
