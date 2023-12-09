// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PCKPacker struct {
  obj gdc.ObjectPtr
}

func (me *PCKPacker) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PCKPacker) BaseClass() string {
  return "PCKPacker"
}



// Enums

func (me *PCKPacker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PCKPacker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PCKPacker) PckStart(pck_name String, alignment int, key String, encrypt_directory bool, )  {
  panic("TODO: implement")
}

func  (me *PCKPacker) AddFile(pck_path String, source_path String, encrypt bool, )  {
  panic("TODO: implement")
}

func  (me *PCKPacker) Flush(verbose bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
