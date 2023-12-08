// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Semaphore struct {
  obj gdc.ObjectPtr
}

func createSemaphore(obj gdc.ObjectPtr) *Semaphore {
  return &Semaphore{
    obj: obj,
  }
}

func (me *Semaphore) BaseClass() string {
  return "Semaphore"
}
