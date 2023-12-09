// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MethodTweener struct {
  obj gdc.ObjectPtr
}

func (me *MethodTweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MethodTweener) BaseClass() string {
  return "MethodTweener"
}



// Enums

func (me *MethodTweener) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MethodTweener) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MethodTweener) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MethodTweener) SetDelay(delay float32, )  {
  panic("TODO: implement")
}

func  (me *MethodTweener) SetTrans(trans TweenTransitionType, )  {
  panic("TODO: implement")
}

func  (me *MethodTweener) SetEase(ease TweenEaseType, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
