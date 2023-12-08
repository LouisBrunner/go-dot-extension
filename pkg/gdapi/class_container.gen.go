// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Container struct {
  obj gdc.ObjectPtr
}

func (me *Container) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Container) BaseClass() string {
  return "Container"
}

const (
  ContainerNOTIFICATION_PRE_SORT_CHILDREN = 50
  ContainerNOTIFICATION_SORT_CHILDREN = 51
)
