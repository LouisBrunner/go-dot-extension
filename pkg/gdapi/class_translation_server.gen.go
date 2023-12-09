// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TranslationServer struct {
  obj gdc.ObjectPtr
}

func (me *TranslationServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TranslationServer) BaseClass() string {
  return "TranslationServer"
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
  panic("TODO: implement")
}

func  (me *TranslationServer) GetLocale()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetToolLocale()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) CompareLocales(locale_a String, locale_b String, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) StandardizeLocale(locale String, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetAllLanguages()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetLanguageName(language String, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetAllScripts()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetScriptName(script String, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetAllCountries()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetCountryName(country String, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetLocaleName(locale String, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) Translate(message StringName, context StringName, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) TranslatePlural(message StringName, plural_message StringName, n int, context StringName, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) AddTranslation(translation Translation, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) RemoveTranslation(translation Translation, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetTranslationObject(locale String, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) Clear()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) GetLoadedLocales()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) IsPseudolocalizationEnabled()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) SetPseudolocalizationEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TranslationServer) ReloadPseudolocalization()  {
  panic("TODO: implement")
}

func  (me *TranslationServer) Pseudolocalize(message StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
