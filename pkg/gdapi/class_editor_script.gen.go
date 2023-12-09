// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScript struct {
  obj gdc.ObjectPtr
}

func (me *EditorScript) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorScript) BaseClass() string {
  return "EditorScript"
}



// Enums

func (me *EditorScript) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorScript) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorScript) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorScript) XRun()  {
  panic("TODO: implement")
}

func  (me *EditorScript) AddRootNode(node Node, )  {
  panic("TODO: implement")
}

func  (me *EditorScript) GetScene()  {
  panic("TODO: implement")
}

func  (me *EditorScript) GetEditorInterface()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
