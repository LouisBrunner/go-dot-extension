// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SplitContainer struct {
  obj gdc.ObjectPtr
}

func (me *SplitContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SplitContainer) BaseClass() string {
  return "SplitContainer"
}



// Enums

type SplitContainerDraggerVisibility int
const (
  SplitContainerDraggerVisibilityDraggerVisible SplitContainerDraggerVisibility = 0
  SplitContainerDraggerVisibilityDraggerHidden SplitContainerDraggerVisibility = 1
  SplitContainerDraggerVisibilityDraggerHiddenCollapsed SplitContainerDraggerVisibility = 2
)

func (me *SplitContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SplitContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SplitContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SplitContainer) SetSplitOffset(offset int, )  {
  panic("TODO: implement")
}

func  (me *SplitContainer) GetSplitOffset()  {
  panic("TODO: implement")
}

func  (me *SplitContainer) ClampSplitOffset()  {
  panic("TODO: implement")
}

func  (me *SplitContainer) SetCollapsed(collapsed bool, )  {
  panic("TODO: implement")
}

func  (me *SplitContainer) IsCollapsed()  {
  panic("TODO: implement")
}

func  (me *SplitContainer) SetDraggerVisibility(mode SplitContainerDraggerVisibility, )  {
  panic("TODO: implement")
}

func  (me *SplitContainer) GetDraggerVisibility()  {
  panic("TODO: implement")
}

func  (me *SplitContainer) SetVertical(vertical bool, )  {
  panic("TODO: implement")
}

func  (me *SplitContainer) IsVertical()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
