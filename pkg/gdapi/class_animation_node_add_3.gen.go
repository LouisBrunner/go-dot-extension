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

type ptrsForAnimationNodeAdd3List struct {
}

var ptrsForAnimationNodeAdd3 ptrsForAnimationNodeAdd3List

func initAnimationNodeAdd3Ptrs(iface gdc.Interface) {

  className := StringNameFromStr("AnimationNodeAdd3")
  defer className.Destroy()
}

type AnimationNodeAdd3 struct {
  AnimationNodeSync
}

func (me *AnimationNodeAdd3) BaseClass() string {
  return "AnimationNodeAdd3"
}

func NewAnimationNodeAdd3() *AnimationNodeAdd3 {
  str := StringNameFromStr("AnimationNodeAdd3") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeAdd3{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AnimationNodeAdd3) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeAdd3) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeAdd3) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
