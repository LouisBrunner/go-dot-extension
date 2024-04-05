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

type ptrsForGLTFAnimationList struct {
  fnGetLoop gdc.MethodBindPtr
  fnSetLoop gdc.MethodBindPtr
}

var ptrsForGLTFAnimation ptrsForGLTFAnimationList

func initGLTFAnimationPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GLTFAnimation")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_loop")
    defer methodName.Destroy()
    ptrsForGLTFAnimation.fnGetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_loop")
    defer methodName.Destroy()
    ptrsForGLTFAnimation.fnSetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
}

type GLTFAnimation struct {
  Resource
}

func (me *GLTFAnimation) BaseClass() string {
  return "GLTFAnimation"
}

func NewGLTFAnimation() *GLTFAnimation {
  str := StringNameFromStr("GLTFAnimation") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFAnimation{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GLTFAnimation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFAnimation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFAnimation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFAnimation) GetLoop() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAnimation.fnGetLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAnimation) SetLoop(loop bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAnimation.fnSetLoop), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
