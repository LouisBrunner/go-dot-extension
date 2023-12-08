// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CallbackTweener struct {
  obj gdc.ObjectPtr
}

func (me *CallbackTweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CallbackTweener) BaseClass() string {
  return "CallbackTweener"
}
