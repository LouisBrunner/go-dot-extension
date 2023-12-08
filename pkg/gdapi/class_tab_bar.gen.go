// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TabBar struct {
  obj gdc.ObjectPtr
}

func (me *TabBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TabBar) BaseClass() string {
  return "TabBar"
}

type TabBarAlignmentMode int
const (
  TabBarAlignmentLeft TabBarAlignmentMode = 0
  TabBarAlignmentCenter TabBarAlignmentMode = 1
  TabBarAlignmentRight TabBarAlignmentMode = 2
  TabBarAlignmentMax TabBarAlignmentMode = 3
)

type TabBarCloseButtonDisplayPolicy int
const (
  TabBarCloseButtonShowNever TabBarCloseButtonDisplayPolicy = 0
  TabBarCloseButtonShowActiveOnly TabBarCloseButtonDisplayPolicy = 1
  TabBarCloseButtonShowAlways TabBarCloseButtonDisplayPolicy = 2
  TabBarCloseButtonMax TabBarCloseButtonDisplayPolicy = 3
)
