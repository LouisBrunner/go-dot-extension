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



// Enums

type ZIPPackerZipAppend int
const (
  ZIPPackerZipAppendAppendCreate ZIPPackerZipAppend = 0
  ZIPPackerZipAppendAppendCreateafter ZIPPackerZipAppend = 1
  ZIPPackerZipAppendAppendAddinzip ZIPPackerZipAppend = 2
)

func (me *ZIPPacker) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ZIPPacker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ZIPPacker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ZIPPacker) Open(path String, append ZIPPackerZipAppend, )  {
  panic("TODO: implement")
}

func  (me *ZIPPacker) StartFile(path String, )  {
  panic("TODO: implement")
}

func  (me *ZIPPacker) WriteFile(data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *ZIPPacker) CloseFile()  {
  panic("TODO: implement")
}

func  (me *ZIPPacker) Close()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
