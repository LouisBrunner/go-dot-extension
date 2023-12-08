// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  AspectRatioContainerStretchWidthControlsHeight AspectRatioContainerStretchMode = 0
  AspectRatioContainerStretchHeightControlsWidth AspectRatioContainerStretchMode = 1
  AspectRatioContainerStretchFit AspectRatioContainerStretchMode = 2
  AspectRatioContainerStretchCover AspectRatioContainerStretchMode = 3
)

type AspectRatioContainerAlignmentMode int
const (
  AspectRatioContainerAlignmentBegin AspectRatioContainerAlignmentMode = 0
  AspectRatioContainerAlignmentCenter AspectRatioContainerAlignmentMode = 1
  AspectRatioContainerAlignmentEnd AspectRatioContainerAlignmentMode = 2
)
