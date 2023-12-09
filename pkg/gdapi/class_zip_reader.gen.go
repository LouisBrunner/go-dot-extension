// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ZIPReader struct {
  obj gdc.ObjectPtr
}

func (me *ZIPReader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ZIPReader) BaseClass() string {
  return "ZIPReader"
}



// Enums

func (me *ZIPReader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ZIPReader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ZIPReader) Open(path String, )  {
  panic("TODO: implement")
}

func  (me *ZIPReader) Close()  {
  panic("TODO: implement")
}

func  (me *ZIPReader) GetFiles()  {
  panic("TODO: implement")
}

func  (me *ZIPReader) ReadFile(path String, case_sensitive bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
