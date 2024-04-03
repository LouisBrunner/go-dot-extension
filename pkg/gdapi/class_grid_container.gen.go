// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GridContainer struct {
  Container
}

func (me *GridContainer) BaseClass() string {
  return "GridContainer"
}

func NewGridContainer() *GridContainer {
  str := StringNameFromStr("GridContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GridContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GridContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GridContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GridContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GridContainer) SetColumns(columns int64, )  {
  classNameV := StringNameFromStr("GridContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&columns), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GridContainer) GetColumns() int64 {
  classNameV := StringNameFromStr("GridContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
