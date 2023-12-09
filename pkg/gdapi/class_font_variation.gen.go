// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FontVariation struct {
  obj gdc.ObjectPtr
}

func (me *FontVariation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FontVariation) BaseClass() string {
  return "FontVariation"
}



// Enums

func (me *FontVariation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FontVariation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *FontVariation) SetBaseFont(font Font, )  {
  panic("TODO: implement")
}

func  (me *FontVariation) GetBaseFont()  {
  panic("TODO: implement")
}

func  (me *FontVariation) SetVariationOpentype(coords Dictionary, )  {
  panic("TODO: implement")
}

func  (me *FontVariation) GetVariationOpentype()  {
  panic("TODO: implement")
}

func  (me *FontVariation) SetVariationEmbolden(strength float32, )  {
  panic("TODO: implement")
}

func  (me *FontVariation) GetVariationEmbolden()  {
  panic("TODO: implement")
}

func  (me *FontVariation) SetVariationFaceIndex(face_index int, )  {
  panic("TODO: implement")
}

func  (me *FontVariation) GetVariationFaceIndex()  {
  panic("TODO: implement")
}

func  (me *FontVariation) SetVariationTransform(transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *FontVariation) GetVariationTransform()  {
  panic("TODO: implement")
}

func  (me *FontVariation) SetOpentypeFeatures(features Dictionary, )  {
  panic("TODO: implement")
}

func  (me *FontVariation) SetSpacing(spacing TextServerSpacingType, value int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
