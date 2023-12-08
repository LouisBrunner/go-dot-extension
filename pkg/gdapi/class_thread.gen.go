// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Thread struct {
  obj gdc.ObjectPtr
}

func createThread(obj gdc.ObjectPtr) *Thread {
  return &Thread{
    obj: obj,
  }
}

func (me *Thread) BaseClass() string {
  return "Thread"
}
