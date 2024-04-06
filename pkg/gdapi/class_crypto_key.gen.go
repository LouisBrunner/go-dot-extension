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

type ptrsForCryptoKeyList struct {
	fnSave           gdc.MethodBindPtr
	fnLoad           gdc.MethodBindPtr
	fnIsPublicOnly   gdc.MethodBindPtr
	fnSaveToString   gdc.MethodBindPtr
	fnLoadFromString gdc.MethodBindPtr
}

var ptrsForCryptoKey ptrsForCryptoKeyList

func initCryptoKeyPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CryptoKey")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("save")
		defer methodName.Destroy()
		ptrsForCryptoKey.fnSave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 885841341))
	}
	{
		methodName := StringNameFromStr("load")
		defer methodName.Destroy()
		ptrsForCryptoKey.fnLoad = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 885841341))
	}
	{
		methodName := StringNameFromStr("is_public_only")
		defer methodName.Destroy()
		ptrsForCryptoKey.fnIsPublicOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("save_to_string")
		defer methodName.Destroy()
		ptrsForCryptoKey.fnSaveToString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 32795936))
	}
	{
		methodName := StringNameFromStr("load_from_string")
		defer methodName.Destroy()
		ptrsForCryptoKey.fnLoadFromString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 885841341))
	}

}

type CryptoKey struct {
	Resource
}

func (me *CryptoKey) BaseClass() string {
	return "CryptoKey"
}

func NewCryptoKey() *CryptoKey {
	str := StringNameFromStr("CryptoKey") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CryptoKey{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CryptoKey) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CryptoKey) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CryptoKey) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CryptoKey) Save(path String, public_only bool) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&public_only)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&public_only)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCryptoKey.fnSave), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CryptoKey) Load(path String, public_only bool) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&public_only)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&public_only)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCryptoKey.fnLoad), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CryptoKey) IsPublicOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCryptoKey.fnIsPublicOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CryptoKey) SaveToString(public_only bool) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&public_only)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&public_only)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCryptoKey.fnSaveToString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CryptoKey) LoadFromString(string_key String, public_only bool) Error {
	cargs := []gdc.ConstTypePtr{string_key.AsCTypePtr(), gdc.ConstTypePtr(&public_only)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&public_only)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCryptoKey.fnLoadFromString), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals
