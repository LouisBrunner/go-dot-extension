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

type ptrsForConfigFileList struct {
	fnSetValue          gdc.MethodBindPtr
	fnGetValue          gdc.MethodBindPtr
	fnHasSection        gdc.MethodBindPtr
	fnHasSectionKey     gdc.MethodBindPtr
	fnGetSections       gdc.MethodBindPtr
	fnGetSectionKeys    gdc.MethodBindPtr
	fnEraseSection      gdc.MethodBindPtr
	fnEraseSectionKey   gdc.MethodBindPtr
	fnLoad              gdc.MethodBindPtr
	fnParse             gdc.MethodBindPtr
	fnSave              gdc.MethodBindPtr
	fnEncodeToText      gdc.MethodBindPtr
	fnLoadEncrypted     gdc.MethodBindPtr
	fnLoadEncryptedPass gdc.MethodBindPtr
	fnSaveEncrypted     gdc.MethodBindPtr
	fnSaveEncryptedPass gdc.MethodBindPtr
	fnClear             gdc.MethodBindPtr
}

var ptrsForConfigFile ptrsForConfigFileList

func initConfigFilePtrs(iface gdc.Interface) {

	className := StringNameFromStr("ConfigFile")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_value")
		defer methodName.Destroy()
		ptrsForConfigFile.fnSetValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2504492430))
	}
	{
		methodName := StringNameFromStr("get_value")
		defer methodName.Destroy()
		ptrsForConfigFile.fnGetValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 89809366))
	}
	{
		methodName := StringNameFromStr("has_section")
		defer methodName.Destroy()
		ptrsForConfigFile.fnHasSection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("has_section_key")
		defer methodName.Destroy()
		ptrsForConfigFile.fnHasSectionKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 820780508))
	}
	{
		methodName := StringNameFromStr("get_sections")
		defer methodName.Destroy()
		ptrsForConfigFile.fnGetSections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("get_section_keys")
		defer methodName.Destroy()
		ptrsForConfigFile.fnGetSectionKeys = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291131558))
	}
	{
		methodName := StringNameFromStr("erase_section")
		defer methodName.Destroy()
		ptrsForConfigFile.fnEraseSection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("erase_section_key")
		defer methodName.Destroy()
		ptrsForConfigFile.fnEraseSectionKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3186203200))
	}
	{
		methodName := StringNameFromStr("load")
		defer methodName.Destroy()
		ptrsForConfigFile.fnLoad = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("parse")
		defer methodName.Destroy()
		ptrsForConfigFile.fnParse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("save")
		defer methodName.Destroy()
		ptrsForConfigFile.fnSave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("encode_to_text")
		defer methodName.Destroy()
		ptrsForConfigFile.fnEncodeToText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("load_encrypted")
		defer methodName.Destroy()
		ptrsForConfigFile.fnLoadEncrypted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 887037711))
	}
	{
		methodName := StringNameFromStr("load_encrypted_pass")
		defer methodName.Destroy()
		ptrsForConfigFile.fnLoadEncryptedPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
	{
		methodName := StringNameFromStr("save_encrypted")
		defer methodName.Destroy()
		ptrsForConfigFile.fnSaveEncrypted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 887037711))
	}
	{
		methodName := StringNameFromStr("save_encrypted_pass")
		defer methodName.Destroy()
		ptrsForConfigFile.fnSaveEncryptedPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForConfigFile.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type ConfigFile struct {
	RefCounted
}

func (me *ConfigFile) BaseClass() string {
	return "ConfigFile"
}

func NewConfigFile() *ConfigFile {
	str := StringNameFromStr("ConfigFile") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ConfigFile{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ConfigFile) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ConfigFile) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ConfigFile) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ConfigFile) SetValue(section String, key String, value Variant) {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr(), key.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnSetValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ConfigFile) GetValue(section String, key String, default_ Variant) Variant {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr(), key.AsCTypePtr(), default_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnGetValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ConfigFile) HasSection(section String) bool {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnHasSection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ConfigFile) HasSectionKey(section String, key String) bool {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr(), key.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnHasSectionKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ConfigFile) GetSections() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnGetSections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ConfigFile) GetSectionKeys(section String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnGetSectionKeys), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ConfigFile) EraseSection(section String) {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnEraseSection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ConfigFile) EraseSectionKey(section String, key String) {
	cargs := []gdc.ConstTypePtr{section.AsCTypePtr(), key.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnEraseSectionKey), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ConfigFile) Load(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnLoad), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ConfigFile) Parse(data String) Error {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnParse), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ConfigFile) Save(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnSave), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ConfigFile) EncodeToText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnEncodeToText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ConfigFile) LoadEncrypted(path String, key PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), key.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnLoadEncrypted), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ConfigFile) LoadEncryptedPass(path String, password String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), password.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnLoadEncryptedPass), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ConfigFile) SaveEncrypted(path String, key PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), key.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnSaveEncrypted), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ConfigFile) SaveEncryptedPass(path String, password String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), password.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnSaveEncryptedPass), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ConfigFile) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConfigFile.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
