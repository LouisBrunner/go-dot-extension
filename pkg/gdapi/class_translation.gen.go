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

type ptrsForTranslationList struct {
	fnXGetPluralMessage        gdc.MethodBindPtr
	fnXGetMessage              gdc.MethodBindPtr
	fnSetLocale                gdc.MethodBindPtr
	fnGetLocale                gdc.MethodBindPtr
	fnAddMessage               gdc.MethodBindPtr
	fnAddPluralMessage         gdc.MethodBindPtr
	fnGetMessage               gdc.MethodBindPtr
	fnGetPluralMessage         gdc.MethodBindPtr
	fnEraseMessage             gdc.MethodBindPtr
	fnGetMessageList           gdc.MethodBindPtr
	fnGetTranslatedMessageList gdc.MethodBindPtr
	fnGetMessageCount          gdc.MethodBindPtr
}

var ptrsForTranslation ptrsForTranslationList

func initTranslationPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Translation")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_locale")
		defer methodName.Destroy()
		ptrsForTranslation.fnSetLocale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_locale")
		defer methodName.Destroy()
		ptrsForTranslation.fnGetLocale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("add_message")
		defer methodName.Destroy()
		ptrsForTranslation.fnAddMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3898530326))
	}
	{
		methodName := StringNameFromStr("add_plural_message")
		defer methodName.Destroy()
		ptrsForTranslation.fnAddPluralMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2356982266))
	}
	{
		methodName := StringNameFromStr("get_message")
		defer methodName.Destroy()
		ptrsForTranslation.fnGetMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1829228469))
	}
	{
		methodName := StringNameFromStr("get_plural_message")
		defer methodName.Destroy()
		ptrsForTranslation.fnGetPluralMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 229954002))
	}
	{
		methodName := StringNameFromStr("erase_message")
		defer methodName.Destroy()
		ptrsForTranslation.fnEraseMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3959009644))
	}
	{
		methodName := StringNameFromStr("get_message_list")
		defer methodName.Destroy()
		ptrsForTranslation.fnGetMessageList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("get_translated_message_list")
		defer methodName.Destroy()
		ptrsForTranslation.fnGetTranslatedMessageList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("get_message_count")
		defer methodName.Destroy()
		ptrsForTranslation.fnGetMessageCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type Translation struct {
	Resource
}

func (me *Translation) BaseClass() string {
	return "Translation"
}

func NewTranslation() *Translation {
	str := StringNameFromStr("Translation") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Translation{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Translation) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Translation) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Translation) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Translation) SetLocale(locale String) {
	cargs := []gdc.ConstTypePtr{locale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnSetLocale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Translation) GetLocale() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnGetLocale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Translation) AddMessage(src_message StringName, xlated_message StringName, context StringName) {
	cargs := []gdc.ConstTypePtr{src_message.AsCTypePtr(), xlated_message.AsCTypePtr(), context.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnAddMessage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Translation) AddPluralMessage(src_message StringName, xlated_messages PackedStringArray, context StringName) {
	cargs := []gdc.ConstTypePtr{src_message.AsCTypePtr(), xlated_messages.AsCTypePtr(), context.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnAddPluralMessage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Translation) GetMessage(src_message StringName, context StringName) StringName {
	cargs := []gdc.ConstTypePtr{src_message.AsCTypePtr(), context.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnGetMessage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Translation) GetPluralMessage(src_message StringName, src_plural_message StringName, n int64, context StringName) StringName {
	cargs := []gdc.ConstTypePtr{src_message.AsCTypePtr(), src_plural_message.AsCTypePtr(), gdc.ConstTypePtr(&n), context.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&n)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnGetPluralMessage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Translation) EraseMessage(src_message StringName, context StringName) {
	cargs := []gdc.ConstTypePtr{src_message.AsCTypePtr(), context.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnEraseMessage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Translation) GetMessageList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnGetMessageList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Translation) GetTranslatedMessageList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnGetTranslatedMessageList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Translation) GetMessageCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTranslation.fnGetMessageCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
