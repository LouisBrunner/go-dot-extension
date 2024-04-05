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

type ptrsForCharFXTransformList struct {
  fnGetTransform gdc.MethodBindPtr
  fnSetTransform gdc.MethodBindPtr
  fnGetRange gdc.MethodBindPtr
  fnSetRange gdc.MethodBindPtr
  fnGetElapsedTime gdc.MethodBindPtr
  fnSetElapsedTime gdc.MethodBindPtr
  fnIsVisible gdc.MethodBindPtr
  fnSetVisibility gdc.MethodBindPtr
  fnIsOutline gdc.MethodBindPtr
  fnSetOutline gdc.MethodBindPtr
  fnGetOffset gdc.MethodBindPtr
  fnSetOffset gdc.MethodBindPtr
  fnGetColor gdc.MethodBindPtr
  fnSetColor gdc.MethodBindPtr
  fnGetEnvironment gdc.MethodBindPtr
  fnSetEnvironment gdc.MethodBindPtr
  fnGetGlyphIndex gdc.MethodBindPtr
  fnSetGlyphIndex gdc.MethodBindPtr
  fnGetRelativeIndex gdc.MethodBindPtr
  fnSetRelativeIndex gdc.MethodBindPtr
  fnGetGlyphCount gdc.MethodBindPtr
  fnSetGlyphCount gdc.MethodBindPtr
  fnGetGlyphFlags gdc.MethodBindPtr
  fnSetGlyphFlags gdc.MethodBindPtr
  fnGetFont gdc.MethodBindPtr
  fnSetFont gdc.MethodBindPtr
}

var ptrsForCharFXTransform ptrsForCharFXTransformList

func initCharFXTransformPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CharFXTransform")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_transform")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3761352769))
  }
  {
    methodName := StringNameFromStr("set_transform")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
  }
  {
    methodName := StringNameFromStr("get_range")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2741790807))
  }
  {
    methodName := StringNameFromStr("set_range")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_elapsed_time")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetElapsedTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_elapsed_time")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetElapsedTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("is_visible")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_visibility")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetVisibility = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_outline")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnIsOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_outline")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_offset")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497962370))
  }
  {
    methodName := StringNameFromStr("set_offset")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_color")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200896285))
  }
  {
    methodName := StringNameFromStr("set_color")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_environment")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
  }
  {
    methodName := StringNameFromStr("set_environment")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
  }
  {
    methodName := StringNameFromStr("get_glyph_index")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetGlyphIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_glyph_index")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetGlyphIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_relative_index")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetRelativeIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_relative_index")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetRelativeIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_glyph_count")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetGlyphCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_glyph_count")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetGlyphCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_glyph_flags")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetGlyphFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_glyph_flags")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetGlyphFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_font")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnGetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("set_font")
    defer methodName.Destroy()
    ptrsForCharFXTransform.fnSetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
}

type CharFXTransform struct {
  RefCounted
}

func (me *CharFXTransform) BaseClass() string {
  return "CharFXTransform"
}

func NewCharFXTransform() *CharFXTransform {
  str := StringNameFromStr("CharFXTransform") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CharFXTransform{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CharFXTransform) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CharFXTransform) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CharFXTransform) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CharFXTransform) GetTransform() Transform2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharFXTransform) SetTransform(transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetRange() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharFXTransform) SetRange(range_ Vector2i, )  {
  cargs := []gdc.ConstTypePtr{range_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetRange), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetElapsedTime() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetElapsedTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharFXTransform) SetElapsedTime(time float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetElapsedTime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) IsVisible() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnIsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharFXTransform) SetVisibility(visibility bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visibility) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetVisibility), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) IsOutline() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnIsOutline), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharFXTransform) SetOutline(outline bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&outline) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharFXTransform) SetOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharFXTransform) SetColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetEnvironment() Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetEnvironment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharFXTransform) SetEnvironment(environment Dictionary, )  {
  cargs := []gdc.ConstTypePtr{environment.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetGlyphIndex() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetGlyphIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharFXTransform) SetGlyphIndex(glyph_index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&glyph_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetGlyphIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetRelativeIndex() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetRelativeIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharFXTransform) SetRelativeIndex(relative_index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&relative_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetRelativeIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetGlyphCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetGlyphCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharFXTransform) SetGlyphCount(glyph_count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&glyph_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetGlyphCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetGlyphFlags() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetGlyphFlags), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CharFXTransform) SetGlyphFlags(glyph_flags int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&glyph_flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetGlyphFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CharFXTransform) GetFont() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnGetFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CharFXTransform) SetFont(font RID, )  {
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharFXTransform.fnSetFont), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
