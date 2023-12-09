// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type IntervalTweener struct {
  obj gdc.ObjectPtr
}

func (me *IntervalTweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *IntervalTweener) BaseClass() string {
  return "IntervalTweener"
}



// Enums

func (me *IntervalTweener) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *IntervalTweener) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *IntervalTweener) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
