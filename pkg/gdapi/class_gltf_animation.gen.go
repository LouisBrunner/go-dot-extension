// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  classNameV := StringNameFromStr("GLTFAnimation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAnimation) SetLoop(loop bool, )  {
  classNameV := StringNameFromStr("GLTFAnimation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
