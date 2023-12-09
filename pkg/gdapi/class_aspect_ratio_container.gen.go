// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AspectRatioContainer struct {
  obj gdc.ObjectPtr
}

func (me *AspectRatioContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AspectRatioContainer) BaseClass() string {
  return "AspectRatioContainer"
}



// Enums

type AspectRatioContainerStretchMode int
const (
  AspectRatioContainerStretchModeStretchWidthControlsHeight AspectRatioContainerStretchMode = 0
  AspectRatioContainerStretchModeStretchHeightControlsWidth AspectRatioContainerStretchMode = 1
  AspectRatioContainerStretchModeStretchFit AspectRatioContainerStretchMode = 2
  AspectRatioContainerStretchModeStretchCover AspectRatioContainerStretchMode = 3
)

type AspectRatioContainerAlignmentMode int
const (
  AspectRatioContainerAlignmentModeAlignmentBegin AspectRatioContainerAlignmentMode = 0
  AspectRatioContainerAlignmentModeAlignmentCenter AspectRatioContainerAlignmentMode = 1
  AspectRatioContainerAlignmentModeAlignmentEnd AspectRatioContainerAlignmentMode = 2
)

func (me *AspectRatioContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AspectRatioContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AspectRatioContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AspectRatioContainer) SetRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *AspectRatioContainer) GetRatio()  {
  panic("TODO: implement")
}

func  (me *AspectRatioContainer) SetStretchMode(stretch_mode AspectRatioContainerStretchMode, )  {
  panic("TODO: implement")
}

func  (me *AspectRatioContainer) GetStretchMode()  {
  panic("TODO: implement")
}

func  (me *AspectRatioContainer) SetAlignmentHorizontal(alignment_horizontal AspectRatioContainerAlignmentMode, )  {
  panic("TODO: implement")
}

func  (me *AspectRatioContainer) GetAlignmentHorizontal()  {
  panic("TODO: implement")
}

func  (me *AspectRatioContainer) SetAlignmentVertical(alignment_vertical AspectRatioContainerAlignmentMode, )  {
  panic("TODO: implement")
}

func  (me *AspectRatioContainer) GetAlignmentVertical()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
