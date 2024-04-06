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

type ptrsForTLSOptionsList struct {
	fnClient       gdc.MethodBindPtr
	fnClientUnsafe gdc.MethodBindPtr
	fnServer       gdc.MethodBindPtr
}

var ptrsForTLSOptions ptrsForTLSOptionsList

func initTLSOptionsPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TLSOptions")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("client")
		defer methodName.Destroy()
		ptrsForTLSOptions.fnClient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3565000357))
	}
	{
		methodName := StringNameFromStr("client_unsafe")
		defer methodName.Destroy()
		ptrsForTLSOptions.fnClientUnsafe = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2090251749))
	}
	{
		methodName := StringNameFromStr("server")
		defer methodName.Destroy()
		ptrsForTLSOptions.fnServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36969539))
	}
}

type TLSOptions struct {
	RefCounted
}

func (me *TLSOptions) BaseClass() string {
	return "TLSOptions"
}

func NewTLSOptions() *TLSOptions {
	str := StringNameFromStr("TLSOptions") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TLSOptions{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TLSOptions) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TLSOptions) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TLSOptions) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func TLSOptionsClient(trusted_chain X509Certificate, common_name_override String) TLSOptions {
	cargs := []gdc.ConstTypePtr{trusted_chain.AsCTypePtr(), common_name_override.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTLSOptions()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTLSOptions.fnClient), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func TLSOptionsClientUnsafe(trusted_chain X509Certificate) TLSOptions {
	cargs := []gdc.ConstTypePtr{trusted_chain.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTLSOptions()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTLSOptions.fnClientUnsafe), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func TLSOptionsServer(key CryptoKey, certificate X509Certificate) TLSOptions {
	cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), certificate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTLSOptions()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTLSOptions.fnServer), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
