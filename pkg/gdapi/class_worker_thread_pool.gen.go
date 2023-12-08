// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorkerThreadPool struct {
  obj gdc.ObjectPtr
}

func createWorkerThreadPool(obj gdc.ObjectPtr) *WorkerThreadPool {
  return &WorkerThreadPool{
    obj: obj,
  }
}

func (me *WorkerThreadPool) BaseClass() string {
  return "WorkerThreadPool"
}
