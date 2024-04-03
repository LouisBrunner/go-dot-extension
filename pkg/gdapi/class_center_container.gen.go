// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CenterContainer struct {
  Container
}

func (me *CenterContainer) BaseClass() string {
  return "CenterContainer"
}

func NewCenterContainer() *CenterContainer {
  str := StringNameFromStr("CenterContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CenterContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CenterContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CenterContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CenterContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CenterContainer) SetUseTopLeft(enable bool, )  {
  classNameV := StringNameFromStr("CenterContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_top_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CenterContainer) IsUsingTopLeft() bool {
  classNameV := StringNameFromStr("CenterContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_top_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
