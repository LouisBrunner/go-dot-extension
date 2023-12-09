// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FogVolume struct {
  obj gdc.ObjectPtr
}

func (me *FogVolume) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FogVolume) BaseClass() string {
  return "FogVolume"
}



// Enums

func (me *FogVolume) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FogVolume) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *FogVolume) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *FogVolume) GetSize()  {
  panic("TODO: implement")
}

func  (me *FogVolume) SetShape(shape RenderingServerFogVolumeShape, )  {
  panic("TODO: implement")
}

func  (me *FogVolume) GetShape()  {
  panic("TODO: implement")
}

func  (me *FogVolume) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *FogVolume) GetMaterial()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
