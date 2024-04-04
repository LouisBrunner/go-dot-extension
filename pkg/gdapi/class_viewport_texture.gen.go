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

type ViewportTexture struct {
  Texture2D
}

func (me *ViewportTexture) BaseClass() string {
  return "ViewportTexture"
}

func NewViewportTexture() *ViewportTexture {
  str := StringNameFromStr("ViewportTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ViewportTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ViewportTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ViewportTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ViewportTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ViewportTexture) SetViewportPathInScene(path NodePath, )  {
  classNameV := StringNameFromStr("ViewportTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_viewport_path_in_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ViewportTexture) GetViewportPathInScene() NodePath {
  classNameV := StringNameFromStr("ViewportTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_viewport_path_in_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
