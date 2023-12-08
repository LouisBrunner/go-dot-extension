// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CallbackTweener struct {
  obj gdc.ObjectPtr
}

func createCallbackTweener(obj gdc.ObjectPtr) *CallbackTweener {
  return &CallbackTweener{
    obj: obj,
  }
}

func (me *CallbackTweener) BaseClass() string {
  return "CallbackTweener"
}
