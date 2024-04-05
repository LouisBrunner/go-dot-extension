// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForCenterContainerList struct {
  fnSetUseTopLeft gdc.MethodBindPtr
  fnIsUsingTopLeft gdc.MethodBindPtr
}

var ptrsForCenterContainer ptrsForCenterContainerList

func initCenterContainerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CenterContainer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_use_top_left")
    defer methodName.Destroy()
    ptrsForCenterContainer.fnSetUseTopLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_top_left")
    defer methodName.Destroy()
    ptrsForCenterContainer.fnIsUsingTopLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCenterContainer.fnSetUseTopLeft), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CenterContainer) IsUsingTopLeft() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCenterContainer.fnIsUsingTopLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
