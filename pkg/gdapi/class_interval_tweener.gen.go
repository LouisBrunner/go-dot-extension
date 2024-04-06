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

type ptrsForIntervalTweenerList struct {
}

var ptrsForIntervalTweener ptrsForIntervalTweenerList

func initIntervalTweenerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("IntervalTweener")
	defer className.Destroy()
}

type IntervalTweener struct {
	Tweener
}

func (me *IntervalTweener) BaseClass() string {
	return "IntervalTweener"
}

func NewIntervalTweener() *IntervalTweener {
	str := StringNameFromStr("IntervalTweener") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &IntervalTweener{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *IntervalTweener) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *IntervalTweener) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *IntervalTweener) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
