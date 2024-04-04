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

type TranslationServer struct {
  Object
}

func (me *TranslationServer) BaseClass() string {
  return "TranslationServer"
}

func NewTranslationServer() *TranslationServer {
  str := StringNameFromStr("TranslationServer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TranslationServer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *TranslationServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TranslationServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TranslationServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TranslationServer) SetLocale(locale String, )  {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_locale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{locale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) GetLocale() String {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_locale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetToolLocale() String {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tool_locale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) CompareLocales(locale_a String, locale_b String, ) int64 {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compare_locales")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878152881) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{locale_a.AsCTypePtr(), locale_b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TranslationServer) StandardizeLocale(locale String, ) String {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("standardize_locale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{locale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetAllLanguages() PackedStringArray {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_all_languages")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetLanguageName(language String, ) String {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetAllScripts() PackedStringArray {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_all_scripts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetScriptName(script String, ) String {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_script_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{script.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetAllCountries() PackedStringArray {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_all_countries")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetCountryName(country String, ) String {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_country_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{country.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetLocaleName(locale String, ) String {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_locale_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{locale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) Translate(message StringName, context StringName, ) StringName {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("translate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1829228469) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), context.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) TranslatePlural(message StringName, plural_message StringName, n int64, context StringName, ) StringName {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("translate_plural")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229954002) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), plural_message.AsCTypePtr(), gdc.ConstTypePtr(&n) , context.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&n)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) AddTranslation(translation Translation, )  {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_translation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1466479800) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{translation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) RemoveTranslation(translation Translation, )  {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_translation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1466479800) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{translation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) GetTranslationObject(locale String, ) Translation {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_translation_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2065240175) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{locale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTranslation()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) Clear()  {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) GetLoadedLocales() PackedStringArray {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loaded_locales")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) IsPseudolocalizationEnabled() bool {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_pseudolocalization_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TranslationServer) SetPseudolocalizationEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pseudolocalization_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) ReloadPseudolocalization()  {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reload_pseudolocalization")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) Pseudolocalize(message StringName, ) StringName {
  classNameV := StringNameFromStr("TranslationServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pseudolocalize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
