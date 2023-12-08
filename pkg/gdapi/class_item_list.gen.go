// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ItemList struct {
  obj gdc.ObjectPtr
}

func (me *ItemList) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ItemList) BaseClass() string {
  return "ItemList"
}

type ItemListIconMode int
const (
  ItemListIconModeTop ItemListIconMode = 0
  ItemListIconModeLeft ItemListIconMode = 1
)

type ItemListSelectMode int
const (
  ItemListSelectSingle ItemListSelectMode = 0
  ItemListSelectMulti ItemListSelectMode = 1
)
