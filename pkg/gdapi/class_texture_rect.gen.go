// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureRect struct {
  Control
}

func (me *TextureRect) BaseClass() string {
  return "TextureRect"
}

func NewTextureRect() *TextureRect {
  str := StringNameFromStr("TextureRect") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextureRect{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type TextureRectExpandMode int
const (
  TextureRectExpandModeExpandKeepSize TextureRectExpandMode = 0
  TextureRectExpandModeExpandIgnoreSize TextureRectExpandMode = 1
  TextureRectExpandModeExpandFitWidth TextureRectExpandMode = 2
  TextureRectExpandModeExpandFitWidthProportional TextureRectExpandMode = 3
  TextureRectExpandModeExpandFitHeight TextureRectExpandMode = 4
  TextureRectExpandModeExpandFitHeightProportional TextureRectExpandMode = 5
)

type TextureRectStretchMode int
const (
  TextureRectStretchModeStretchScale TextureRectStretchMode = 0
  TextureRectStretchModeStretchTile TextureRectStretchMode = 1
  TextureRectStretchModeStretchKeep TextureRectStretchMode = 2
  TextureRectStretchModeStretchKeepCentered TextureRectStretchMode = 3
  TextureRectStretchModeStretchKeepAspect TextureRectStretchMode = 4
  TextureRectStretchModeStretchKeepAspectCentered TextureRectStretchMode = 5
  TextureRectStretchModeStretchKeepAspectCovered TextureRectStretchMode = 6
)

func (me *TextureRect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureRect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureRect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextureRect) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureRect) GetTexture() Texture2D {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureRect) SetExpandMode(expand_mode TextureRectExpandMode, )  {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1870766882) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&expand_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureRect) GetExpandMode() TextureRectExpandMode {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_expand_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3863824733) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TextureRectExpandMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TextureRect) SetFlipH(enable bool, )  {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureRect) IsFlippedH() bool {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_h")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureRect) SetFlipV(enable bool, )  {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureRect) IsFlippedV() bool {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flipped_v")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureRect) SetStretchMode(stretch_mode TextureRectStretchMode, )  {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 58788729) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureRect) GetStretchMode() TextureRectStretchMode {
  classNameV := StringNameFromStr("TextureRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 346396079) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TextureRectStretchMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
