// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ViewportTexture struct {
  obj gdc.ObjectPtr
}

func (me *ViewportTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ViewportTexture) BaseClass() string {
  return "ViewportTexture"
}



// Enums

func (me *ViewportTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ViewportTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ViewportTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ViewportTexture) SetViewportPathInScene(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *ViewportTexture) GetViewportPathInScene()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
