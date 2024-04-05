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

type ptrsForFontVariationList struct {
  fnSetBaseFont gdc.MethodBindPtr
  fnGetBaseFont gdc.MethodBindPtr
  fnSetVariationOpentype gdc.MethodBindPtr
  fnGetVariationOpentype gdc.MethodBindPtr
  fnSetVariationEmbolden gdc.MethodBindPtr
  fnGetVariationEmbolden gdc.MethodBindPtr
  fnSetVariationFaceIndex gdc.MethodBindPtr
  fnGetVariationFaceIndex gdc.MethodBindPtr
  fnSetVariationTransform gdc.MethodBindPtr
  fnGetVariationTransform gdc.MethodBindPtr
  fnSetOpentypeFeatures gdc.MethodBindPtr
  fnSetSpacing gdc.MethodBindPtr
}

var ptrsForFontVariation ptrsForFontVariationList

func initFontVariationPtrs(iface gdc.Interface) {

  className := StringNameFromStr("FontVariation")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_base_font")
    defer methodName.Destroy()
    ptrsForFontVariation.fnSetBaseFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1262170328))
  }
  {
    methodName := StringNameFromStr("get_base_font")
    defer methodName.Destroy()
    ptrsForFontVariation.fnGetBaseFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229501585))
  }
  {
    methodName := StringNameFromStr("set_variation_opentype")
    defer methodName.Destroy()
    ptrsForFontVariation.fnSetVariationOpentype = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
  }
  {
    methodName := StringNameFromStr("get_variation_opentype")
    defer methodName.Destroy()
    ptrsForFontVariation.fnGetVariationOpentype = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
  }
  {
    methodName := StringNameFromStr("set_variation_embolden")
    defer methodName.Destroy()
    ptrsForFontVariation.fnSetVariationEmbolden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_variation_embolden")
    defer methodName.Destroy()
    ptrsForFontVariation.fnGetVariationEmbolden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_variation_face_index")
    defer methodName.Destroy()
    ptrsForFontVariation.fnSetVariationFaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_variation_face_index")
    defer methodName.Destroy()
    ptrsForFontVariation.fnGetVariationFaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_variation_transform")
    defer methodName.Destroy()
    ptrsForFontVariation.fnSetVariationTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
  }
  {
    methodName := StringNameFromStr("get_variation_transform")
    defer methodName.Destroy()
    ptrsForFontVariation.fnGetVariationTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
  }
  {
    methodName := StringNameFromStr("set_opentype_features")
    defer methodName.Destroy()
    ptrsForFontVariation.fnSetOpentypeFeatures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
  }
  {
    methodName := StringNameFromStr("set_spacing")
    defer methodName.Destroy()
    ptrsForFontVariation.fnSetSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3122339690))
  }
}

type FontVariation struct {
  Font
}

func (me *FontVariation) BaseClass() string {
  return "FontVariation"
}

func NewFontVariation() *FontVariation {
  str := StringNameFromStr("FontVariation") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &FontVariation{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *FontVariation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FontVariation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FontVariation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *FontVariation) SetBaseFont(font Font, )  {
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnSetBaseFont), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontVariation) GetBaseFont() Font {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFont()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnGetBaseFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontVariation) SetVariationOpentype(coords Dictionary, )  {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnSetVariationOpentype), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontVariation) GetVariationOpentype() Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnGetVariationOpentype), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontVariation) SetVariationEmbolden(strength float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnSetVariationEmbolden), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontVariation) GetVariationEmbolden() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnGetVariationEmbolden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontVariation) SetVariationFaceIndex(face_index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&face_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnSetVariationFaceIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontVariation) GetVariationFaceIndex() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnGetVariationFaceIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FontVariation) SetVariationTransform(transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnSetVariationTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontVariation) GetVariationTransform() Transform2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnGetVariationTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FontVariation) SetOpentypeFeatures(features Dictionary, )  {
  cargs := []gdc.ConstTypePtr{features.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnSetOpentypeFeatures), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FontVariation) SetSpacing(spacing TextServerSpacingType, value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spacing) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontVariation.fnSetSpacing), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
