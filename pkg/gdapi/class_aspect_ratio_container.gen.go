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

func  (me *AspectRatioContainer) SetRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AspectRatioContainer) GetRatio() { // TODO: return value
  // TODO: implement
}

func  (me *AspectRatioContainer) SetStretchMode(stretch_mode AspectRatioContainerStretchMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AspectRatioContainer) GetStretchMode() { // TODO: return value
  // TODO: implement
}

func  (me *AspectRatioContainer) SetAlignmentHorizontal(alignment_horizontal AspectRatioContainerAlignmentMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AspectRatioContainer) GetAlignmentHorizontal() { // TODO: return value
  // TODO: implement
}

func  (me *AspectRatioContainer) SetAlignmentVertical(alignment_vertical AspectRatioContainerAlignmentMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AspectRatioContainer) GetAlignmentVertical() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
