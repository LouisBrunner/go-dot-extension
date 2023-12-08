// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorkerThreadPool struct {
  obj gdc.ObjectPtr
}

func (me *WorkerThreadPool) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WorkerThreadPool) BaseClass() string {
  return "WorkerThreadPool"
}
