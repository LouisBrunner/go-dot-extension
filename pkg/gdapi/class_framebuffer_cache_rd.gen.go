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

type ptrsForFramebufferCacheRDList struct {
	fnGetCacheMultipass gdc.MethodBindPtr
}

var ptrsForFramebufferCacheRD ptrsForFramebufferCacheRDList

func initFramebufferCacheRDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("FramebufferCacheRD")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_cache_multipass")
		defer methodName.Destroy()
		ptrsForFramebufferCacheRD.fnGetCacheMultipass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3437881813))
	}

}

type FramebufferCacheRD struct {
	Object
}

func (me *FramebufferCacheRD) BaseClass() string {
	return "FramebufferCacheRD"
}

func NewFramebufferCacheRD() *FramebufferCacheRD {
	str := StringNameFromStr("FramebufferCacheRD") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &FramebufferCacheRD{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *FramebufferCacheRD) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *FramebufferCacheRD) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *FramebufferCacheRD) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func FramebufferCacheRDGetCacheMultipass(textures []RID, passes []RDFramebufferPass, views int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&textures), gdc.ConstTypePtr(&passes), gdc.ConstTypePtr(&views)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&textures)
	pinner.Pin(&passes)
	pinner.Pin(&views)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFramebufferCacheRD.fnGetCacheMultipass), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
