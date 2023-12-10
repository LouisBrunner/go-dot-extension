// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Tweener struct {
  obj gdc.ObjectPtr
}

func (me *Tweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Tweener) BaseClass() string {
  return "Tweener"
}



// Enums

func (me *Tweener) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Tweener) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Tweener) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API