// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorCommandPalette struct {
  obj gdc.ObjectPtr
}

func (me *EditorCommandPalette) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorCommandPalette) BaseClass() string {
  return "EditorCommandPalette"
}



// Enums

func (me *EditorCommandPalette) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorCommandPalette) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorCommandPalette) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorCommandPalette) AddCommand(command_name String, key_name String, binded_callable Callable, shortcut_text String, )  {
  panic("TODO: implement")
}

func  (me *EditorCommandPalette) RemoveCommand(key_name String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
