// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeOutput struct {
  AnimationNode
}

func (me *AnimationNodeOutput) BaseClass() string {
  return "AnimationNodeOutput"
}

func NewAnimationNodeOutput() *AnimationNodeOutput {
  str := StringNameFromStr("AnimationNodeOutput") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeOutput{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimationNodeOutput) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeOutput) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeOutput) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
