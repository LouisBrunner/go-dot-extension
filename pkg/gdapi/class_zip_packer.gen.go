// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ZIPPacker struct {
  obj gdc.ObjectPtr
}

func (me *ZIPPacker) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ZIPPacker) BaseClass() string {
  return "ZIPPacker"
}

type ZIPPackerZipAppend int
const (
  ZIPPackerAppendCreate ZIPPackerZipAppend = 0
  ZIPPackerAppendCreateafter ZIPPackerZipAppend = 1
  ZIPPackerAppendAddinzip ZIPPackerZipAppend = 2
)
