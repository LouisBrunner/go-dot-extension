// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ZIPReader struct {
  obj gdc.ObjectPtr
}

func createZIPReader(obj gdc.ObjectPtr) *ZIPReader {
  return &ZIPReader{
    obj: obj,
  }
}

func (me *ZIPReader) BaseClass() string {
  return "ZIPReader"
}
