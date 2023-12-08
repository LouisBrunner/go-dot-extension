// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JNISingleton struct {
  obj gdc.ObjectPtr
}

func (me *JNISingleton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JNISingleton) BaseClass() string {
  return "JNISingleton"
}
