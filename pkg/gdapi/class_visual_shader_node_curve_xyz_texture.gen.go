// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeCurveXYZTexture struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeCurveXYZTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeCurveXYZTexture) BaseClass() string {
  return "VisualShaderNodeCurveXYZTexture"
}



// Enums

func (me *VisualShaderNodeCurveXYZTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCurveXYZTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCurveXYZTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeCurveXYZTexture) SetTexture(texture CurveXYZTexture, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCurveXYZTexture) GetTexture()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
