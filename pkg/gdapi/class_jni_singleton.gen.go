// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JNISingleton struct {
  obj gdc.ObjectPtr
}

func createJNISingleton(obj gdc.ObjectPtr) *JNISingleton {
  return &JNISingleton{
    obj: obj,
  }
}

func (me *JNISingleton) BaseClass() string {
  return "JNISingleton"
}
