// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TreeItem struct {
  obj gdc.ObjectPtr
}

func createTreeItem(obj gdc.ObjectPtr) *TreeItem {
  return &TreeItem{
    obj: obj,
  }
}

func (me *TreeItem) BaseClass() string {
  return "TreeItem"
}
