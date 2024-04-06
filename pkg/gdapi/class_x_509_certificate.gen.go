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

type ptrsForX509CertificateList struct {
	fnSave           gdc.MethodBindPtr
	fnLoad           gdc.MethodBindPtr
	fnSaveToString   gdc.MethodBindPtr
	fnLoadFromString gdc.MethodBindPtr
}

var ptrsForX509Certificate ptrsForX509CertificateList

func initX509CertificatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("X509Certificate")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("save")
		defer methodName.Destroy()
		ptrsForX509Certificate.fnSave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("load")
		defer methodName.Destroy()
		ptrsForX509Certificate.fnLoad = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("save_to_string")
		defer methodName.Destroy()
		ptrsForX509Certificate.fnSaveToString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("load_from_string")
		defer methodName.Destroy()
		ptrsForX509Certificate.fnLoadFromString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
}

type X509Certificate struct {
	Resource
}

func (me *X509Certificate) BaseClass() string {
	return "X509Certificate"
}

func NewX509Certificate() *X509Certificate {
	str := StringNameFromStr("X509Certificate") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &X509Certificate{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *X509Certificate) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *X509Certificate) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *X509Certificate) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *X509Certificate) Save(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForX509Certificate.fnSave), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *X509Certificate) Load(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForX509Certificate.fnLoad), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *X509Certificate) SaveToString() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForX509Certificate.fnSaveToString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *X509Certificate) LoadFromString(string_ String) Error {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForX509Certificate.fnLoadFromString), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals
