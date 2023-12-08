// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type X509Certificate struct {
  obj gdc.ObjectPtr
}

func (me *X509Certificate) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *X509Certificate) BaseClass() string {
  return "X509Certificate"
}
