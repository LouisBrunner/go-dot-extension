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

type ptrsForRegExMatchList struct {
	fnGetSubject    gdc.MethodBindPtr
	fnGetGroupCount gdc.MethodBindPtr
	fnGetNames      gdc.MethodBindPtr
	fnGetStrings    gdc.MethodBindPtr
	fnGetString     gdc.MethodBindPtr
	fnGetStart      gdc.MethodBindPtr
	fnGetEnd        gdc.MethodBindPtr
}

var ptrsForRegExMatch ptrsForRegExMatchList

func initRegExMatchPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RegExMatch")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_subject")
		defer methodName.Destroy()
		ptrsForRegExMatch.fnGetSubject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_group_count")
		defer methodName.Destroy()
		ptrsForRegExMatch.fnGetGroupCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_names")
		defer methodName.Destroy()
		ptrsForRegExMatch.fnGetNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_strings")
		defer methodName.Destroy()
		ptrsForRegExMatch.fnGetStrings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("get_string")
		defer methodName.Destroy()
		ptrsForRegExMatch.fnGetString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 687115856))
	}
	{
		methodName := StringNameFromStr("get_start")
		defer methodName.Destroy()
		ptrsForRegExMatch.fnGetStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 490464691))
	}
	{
		methodName := StringNameFromStr("get_end")
		defer methodName.Destroy()
		ptrsForRegExMatch.fnGetEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 490464691))
	}

}

type RegExMatch struct {
	RefCounted
}

func (me *RegExMatch) BaseClass() string {
	return "RegExMatch"
}

func NewRegExMatch() *RegExMatch {
	str := StringNameFromStr("RegExMatch") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RegExMatch{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RegExMatch) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RegExMatch) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RegExMatch) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RegExMatch) GetSubject() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegExMatch.fnGetSubject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RegExMatch) GetGroupCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegExMatch.fnGetGroupCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RegExMatch) GetNames() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegExMatch.fnGetNames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RegExMatch) GetStrings() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegExMatch.fnGetStrings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RegExMatch) GetString(name Variant) String {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegExMatch.fnGetString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RegExMatch) GetStart(name Variant) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegExMatch.fnGetStart), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RegExMatch) GetEnd(name Variant) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRegExMatch.fnGetEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
