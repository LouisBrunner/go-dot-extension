// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TreeItem struct {
  obj gdc.ObjectPtr
}

func (me *TreeItem) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TreeItem) BaseClass() string {
  return "TreeItem"
}
