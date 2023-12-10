// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorResourceTooltipPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorResourceTooltipPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorResourceTooltipPlugin) BaseClass() string {
  return "EditorResourceTooltipPlugin"
}



// Enums

func (me *EditorResourceTooltipPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorResourceTooltipPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorResourceTooltipPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorResourceTooltipPlugin) RequestThumbnail(path String, control TextureRect, )  {
  classNameV := StringNameFromStr("EditorResourceTooltipPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request_thumbnail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3245519720) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(control.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties