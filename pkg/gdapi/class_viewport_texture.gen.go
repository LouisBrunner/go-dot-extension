// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ViewportTexture struct {
  obj gdc.ObjectPtr
}

func (me *ViewportTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ViewportTexture) BaseClass() string {
  return "ViewportTexture"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ViewportTexture) GetViewportPathInScene() NodePath {
  classNameV := StringNameFromStr("ViewportTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_viewport_path_in_scene")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ViewportTexture) GetPropViewportPath() NodePath {
  panic("TODO: implement")
}

func (me *ViewportTexture) SetPropViewportPath(value NodePath) {
  panic("TODO: implement")
}