// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDAttachmentFormat struct {
  obj gdc.ObjectPtr
}

func createRDAttachmentFormat(obj gdc.ObjectPtr) *RDAttachmentFormat {
  return &RDAttachmentFormat{
    obj: obj,
  }
}

func (me *RDAttachmentFormat) BaseClass() string {
  return "RDAttachmentFormat"
}
