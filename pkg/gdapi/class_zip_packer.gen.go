// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  ZIPPackerZipAppendAppendCreate ZIPPackerZipAppend = 0
  ZIPPackerZipAppendAppendCreateafter ZIPPackerZipAppend = 1
  ZIPPackerZipAppendAppendAddinzip ZIPPackerZipAppend = 2
)

func  (me *ZIPPacker) Open(path String, append ZIPPackerZipAppend, ) { // TODO: return value
  // TODO: implement
}

func  (me *ZIPPacker) StartFile(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ZIPPacker) WriteFile(data PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *ZIPPacker) CloseFile() { // TODO: return value
  // TODO: implement
}

func  (me *ZIPPacker) Close() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
