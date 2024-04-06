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

type ptrsForSemaphoreList struct {
	fnWait    gdc.MethodBindPtr
	fnTryWait gdc.MethodBindPtr
	fnPost    gdc.MethodBindPtr
}

var ptrsForSemaphore ptrsForSemaphoreList

func initSemaphorePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Semaphore")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("wait")
		defer methodName.Destroy()
		ptrsForSemaphore.fnWait = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("try_wait")
		defer methodName.Destroy()
		ptrsForSemaphore.fnTryWait = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("post")
		defer methodName.Destroy()
		ptrsForSemaphore.fnPost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type Semaphore struct {
	RefCounted
}

func (me *Semaphore) BaseClass() string {
	return "Semaphore"
}

func NewSemaphore() *Semaphore {
	str := StringNameFromStr("Semaphore") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Semaphore{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Semaphore) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Semaphore) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Semaphore) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Semaphore) Wait() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSemaphore.fnWait), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Semaphore) TryWait() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSemaphore.fnTryWait), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Semaphore) Post() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSemaphore.fnPost), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
