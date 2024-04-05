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

type ptrsForTranslationServerList struct {
  fnSetLocale gdc.MethodBindPtr
  fnGetLocale gdc.MethodBindPtr
  fnGetToolLocale gdc.MethodBindPtr
  fnCompareLocales gdc.MethodBindPtr
  fnStandardizeLocale gdc.MethodBindPtr
  fnGetAllLanguages gdc.MethodBindPtr
  fnGetLanguageName gdc.MethodBindPtr
  fnGetAllScripts gdc.MethodBindPtr
  fnGetScriptName gdc.MethodBindPtr
  fnGetAllCountries gdc.MethodBindPtr
  fnGetCountryName gdc.MethodBindPtr
  fnGetLocaleName gdc.MethodBindPtr
  fnTranslate gdc.MethodBindPtr
  fnTranslatePlural gdc.MethodBindPtr
  fnAddTranslation gdc.MethodBindPtr
  fnRemoveTranslation gdc.MethodBindPtr
  fnGetTranslationObject gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
  fnGetLoadedLocales gdc.MethodBindPtr
  fnIsPseudolocalizationEnabled gdc.MethodBindPtr
  fnSetPseudolocalizationEnabled gdc.MethodBindPtr
  fnReloadPseudolocalization gdc.MethodBindPtr
  fnPseudolocalize gdc.MethodBindPtr
}

var ptrsForTranslationServer ptrsForTranslationServerList

func initTranslationServerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("TranslationServer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_locale")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnSetLocale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_locale")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetLocale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_tool_locale")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetToolLocale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
  }
  {
    methodName := StringNameFromStr("compare_locales")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnCompareLocales = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878152881))
  }
  {
    methodName := StringNameFromStr("standardize_locale")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnStandardizeLocale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
  }
  {
    methodName := StringNameFromStr("get_all_languages")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetAllLanguages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("get_language_name")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetLanguageName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
  }
  {
    methodName := StringNameFromStr("get_all_scripts")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetAllScripts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("get_script_name")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetScriptName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
  }
  {
    methodName := StringNameFromStr("get_all_countries")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetAllCountries = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("get_country_name")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetCountryName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
  }
  {
    methodName := StringNameFromStr("get_locale_name")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetLocaleName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
  }
  {
    methodName := StringNameFromStr("translate")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnTranslate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1829228469))
  }
  {
    methodName := StringNameFromStr("translate_plural")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnTranslatePlural = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 229954002))
  }
  {
    methodName := StringNameFromStr("add_translation")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnAddTranslation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1466479800))
  }
  {
    methodName := StringNameFromStr("remove_translation")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnRemoveTranslation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1466479800))
  }
  {
    methodName := StringNameFromStr("get_translation_object")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetTranslationObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2065240175))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_loaded_locales")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnGetLoadedLocales = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("is_pseudolocalization_enabled")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnIsPseudolocalizationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_pseudolocalization_enabled")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnSetPseudolocalizationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("reload_pseudolocalization")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnReloadPseudolocalization = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("pseudolocalize")
    defer methodName.Destroy()
    ptrsForTranslationServer.fnPseudolocalize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965194235))
  }
}

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
  cargs := []gdc.ConstTypePtr{locale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnSetLocale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) GetLocale() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetLocale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetToolLocale() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetToolLocale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) CompareLocales(locale_a String, locale_b String, ) int64 {
  cargs := []gdc.ConstTypePtr{locale_a.AsCTypePtr(), locale_b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnCompareLocales), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TranslationServer) StandardizeLocale(locale String, ) String {
  cargs := []gdc.ConstTypePtr{locale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnStandardizeLocale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetAllLanguages() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetAllLanguages), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetLanguageName(language String, ) String {
  cargs := []gdc.ConstTypePtr{language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetLanguageName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetAllScripts() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetAllScripts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetScriptName(script String, ) String {
  cargs := []gdc.ConstTypePtr{script.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetScriptName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetAllCountries() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetAllCountries), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetCountryName(country String, ) String {
  cargs := []gdc.ConstTypePtr{country.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetCountryName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) GetLocaleName(locale String, ) String {
  cargs := []gdc.ConstTypePtr{locale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetLocaleName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) Translate(message StringName, context StringName, ) StringName {
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), context.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnTranslate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) TranslatePlural(message StringName, plural_message StringName, n int64, context StringName, ) StringName {
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), plural_message.AsCTypePtr(), gdc.ConstTypePtr(&n) , context.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&n)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnTranslatePlural), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) AddTranslation(translation Translation, )  {
  cargs := []gdc.ConstTypePtr{translation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnAddTranslation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) RemoveTranslation(translation Translation, )  {
  cargs := []gdc.ConstTypePtr{translation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnRemoveTranslation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) GetTranslationObject(locale String, ) Translation {
  cargs := []gdc.ConstTypePtr{locale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTranslation()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetTranslationObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) GetLoadedLocales() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnGetLoadedLocales), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TranslationServer) IsPseudolocalizationEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnIsPseudolocalizationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TranslationServer) SetPseudolocalizationEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnSetPseudolocalizationEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) ReloadPseudolocalization()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnReloadPseudolocalization), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TranslationServer) Pseudolocalize(message StringName, ) StringName {
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslationServer.fnPseudolocalize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
