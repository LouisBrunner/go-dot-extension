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

type ptrsForZIPPackerList struct {
	fnOpen      gdc.MethodBindPtr
	fnStartFile gdc.MethodBindPtr
	fnWriteFile gdc.MethodBindPtr
	fnCloseFile gdc.MethodBindPtr
	fnClose     gdc.MethodBindPtr
}

var ptrsForZIPPacker ptrsForZIPPackerList

func initZIPPackerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ZIPPacker")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("open")
		defer methodName.Destroy()
		ptrsForZIPPacker.fnOpen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1936816515))
	}
	{
		methodName := StringNameFromStr("start_file")
		defer methodName.Destroy()
		ptrsForZIPPacker.fnStartFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("write_file")
		defer methodName.Destroy()
		ptrsForZIPPacker.fnWriteFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
	}
	{
		methodName := StringNameFromStr("close_file")
		defer methodName.Destroy()
		ptrsForZIPPacker.fnCloseFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("close")
		defer methodName.Destroy()
		ptrsForZIPPacker.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}

}

type ZIPPacker struct {
	RefCounted
}

func (me *ZIPPacker) BaseClass() string {
	return "ZIPPacker"
}

func NewZIPPacker() *ZIPPacker {
	str := StringNameFromStr("ZIPPacker") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ZIPPacker{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type ZIPPackerZipAppend int

const (
	ZIPPackerZipAppendAppendCreate      ZIPPackerZipAppend = 0
	ZIPPackerZipAppendAppendCreateafter ZIPPackerZipAppend = 1
	ZIPPackerZipAppendAppendAddinzip    ZIPPackerZipAppend = 2
)

func (me *ZIPPacker) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ZIPPacker) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ZIPPacker) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ZIPPacker) Open(path String, append ZIPPackerZipAppend) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&append)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&append)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPPacker.fnOpen), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ZIPPacker) StartFile(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPPacker.fnStartFile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ZIPPacker) WriteFile(data PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPPacker.fnWriteFile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ZIPPacker) CloseFile() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPPacker.fnCloseFile), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ZIPPacker) Close() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForZIPPacker.fnClose), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals
