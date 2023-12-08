// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ItemList struct {
  obj gdc.ObjectPtr
}

func createItemList(obj gdc.ObjectPtr) *ItemList {
  return &ItemList{
    obj: obj,
  }
}

func (me *ItemList) BaseClass() string {
  return "ItemList"
}
