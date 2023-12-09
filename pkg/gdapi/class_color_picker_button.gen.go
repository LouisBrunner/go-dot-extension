// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ColorPickerButton struct {
  obj gdc.ObjectPtr
}

func (me *ColorPickerButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ColorPickerButton) BaseClass() string {
  return "ColorPickerButton"
}



// Enums

func (me *ColorPickerButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ColorPickerButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ColorPickerButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ColorPickerButton) SetPickColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *ColorPickerButton) GetPickColor()  {
  panic("TODO: implement")
}

func  (me *ColorPickerButton) GetPicker()  {
  panic("TODO: implement")
}

func  (me *ColorPickerButton) GetPopup()  {
  panic("TODO: implement")
}

func  (me *ColorPickerButton) SetEditAlpha(show bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPickerButton) IsEditingAlpha()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
