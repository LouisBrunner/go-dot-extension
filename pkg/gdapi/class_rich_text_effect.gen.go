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

type ptrsForRichTextEffectList struct {
	fnXProcessCustomFx gdc.MethodBindPtr
}

var ptrsForRichTextEffect ptrsForRichTextEffectList

func initRichTextEffectPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RichTextEffect")
	defer className.Destroy()

}

type RichTextEffect struct {
	Resource
}

func (me *RichTextEffect) BaseClass() string {
	return "RichTextEffect"
}

func NewRichTextEffect() *RichTextEffect {
	str := StringNameFromStr("RichTextEffect") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RichTextEffect{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RichTextEffect) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RichTextEffect) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RichTextEffect) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
