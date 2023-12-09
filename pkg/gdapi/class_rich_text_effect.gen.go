// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RichTextEffect struct {
  obj gdc.ObjectPtr
}

func (me *RichTextEffect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RichTextEffect) BaseClass() string {
  return "RichTextEffect"
}



// Enums

func (me *RichTextEffect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RichTextEffect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RichTextEffect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RichTextEffect) XProcessCustomFx(char_fx CharFXTransform, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
