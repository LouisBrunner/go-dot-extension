// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RichTextLabel struct {
  obj gdc.ObjectPtr
}

func (me *RichTextLabel) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RichTextLabel) BaseClass() string {
  return "RichTextLabel"
}

type RichTextLabelListType int
const (
  RichTextLabelListNumbers RichTextLabelListType = 0
  RichTextLabelListLetters RichTextLabelListType = 1
  RichTextLabelListRoman RichTextLabelListType = 2
  RichTextLabelListDots RichTextLabelListType = 3
)

type RichTextLabelMenuItems int
const (
  RichTextLabelMenuCopy RichTextLabelMenuItems = 0
  RichTextLabelMenuSelectAll RichTextLabelMenuItems = 1
  RichTextLabelMenuMax RichTextLabelMenuItems = 2
)
