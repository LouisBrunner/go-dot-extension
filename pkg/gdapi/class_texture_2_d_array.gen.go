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

type ptrsForTexture2DArrayList struct {
	fnCreatePlaceholder gdc.MethodBindPtr
}

var ptrsForTexture2DArray ptrsForTexture2DArrayList

func initTexture2DArrayPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Texture2DArray")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create_placeholder")
		defer methodName.Destroy()
		ptrsForTexture2DArray.fnCreatePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121922552))
	}
}

type Texture2DArray struct {
	ImageTextureLayered
}

func (me *Texture2DArray) BaseClass() string {
	return "Texture2DArray"
}

func NewTexture2DArray() *Texture2DArray {
	str := StringNameFromStr("Texture2DArray") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Texture2DArray{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Texture2DArray) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Texture2DArray) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Texture2DArray) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Texture2DArray) CreatePlaceholder() Resource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewResource()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2DArray.fnCreatePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
