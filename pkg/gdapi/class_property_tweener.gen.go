// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PropertyTweener struct {
  obj gdc.ObjectPtr
}

func (me *PropertyTweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PropertyTweener) BaseClass() string {
  return "PropertyTweener"
}
