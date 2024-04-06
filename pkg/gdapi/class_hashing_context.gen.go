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

type ptrsForHashingContextList struct {
	fnStart  gdc.MethodBindPtr
	fnUpdate gdc.MethodBindPtr
	fnFinish gdc.MethodBindPtr
}

var ptrsForHashingContext ptrsForHashingContextList

func initHashingContextPtrs(iface gdc.Interface) {

	className := StringNameFromStr("HashingContext")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("start")
		defer methodName.Destroy()
		ptrsForHashingContext.fnStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3940338335))
	}
	{
		methodName := StringNameFromStr("update")
		defer methodName.Destroy()
		ptrsForHashingContext.fnUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
	}
	{
		methodName := StringNameFromStr("finish")
		defer methodName.Destroy()
		ptrsForHashingContext.fnFinish = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2115431945))
	}
}

type HashingContext struct {
	RefCounted
}

func (me *HashingContext) BaseClass() string {
	return "HashingContext"
}

func NewHashingContext() *HashingContext {
	str := StringNameFromStr("HashingContext") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &HashingContext{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type HashingContextHashType int

const (
	HashingContextHashTypeHashMd5    HashingContextHashType = 0
	HashingContextHashTypeHashSha1   HashingContextHashType = 1
	HashingContextHashTypeHashSha256 HashingContextHashType = 2
)

func (me *HashingContext) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *HashingContext) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *HashingContext) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *HashingContext) Start(type_ HashingContextHashType) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&type_)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHashingContext.fnStart), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *HashingContext) Update(chunk PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{chunk.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHashingContext.fnUpdate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *HashingContext) Finish() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHashingContext.fnFinish), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
