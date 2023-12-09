// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *EditorResourceTooltipPlugin) XHandles(type_ String, )  {
  panic("TODO: implement")
}

func  (me *EditorResourceTooltipPlugin) XMakeTooltipForPath(path String, metadata Dictionary, base Control, )  {
  panic("TODO: implement")
}

func  (me *EditorResourceTooltipPlugin) RequestThumbnail(path String, control TextureRect, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
