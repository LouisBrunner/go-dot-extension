// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PanoramaSkyMaterial struct {
  obj gdc.ObjectPtr
}

func (me *PanoramaSkyMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PanoramaSkyMaterial) BaseClass() string {
  return "PanoramaSkyMaterial"
}



// Enums

func (me *PanoramaSkyMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PanoramaSkyMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PanoramaSkyMaterial) SetPanorama(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *PanoramaSkyMaterial) GetPanorama()  {
  panic("TODO: implement")
}

func  (me *PanoramaSkyMaterial) SetFilteringEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PanoramaSkyMaterial) IsFilteringEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
