// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Tweener struct {
  obj gdc.ObjectPtr
}

func createTweener(obj gdc.ObjectPtr) *Tweener {
  return &Tweener{
    obj: obj,
  }
}

func (me *Tweener) BaseClass() string {
  return "Tweener"
}
