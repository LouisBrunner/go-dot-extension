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

type ptrsForUniformSetCacheRDList struct {
	fnGetCache gdc.MethodBindPtr
}

var ptrsForUniformSetCacheRD ptrsForUniformSetCacheRDList

func initUniformSetCacheRDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("UniformSetCacheRD")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_cache")
		defer methodName.Destroy()
		ptrsForUniformSetCacheRD.fnGetCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 658571723))
	}

}

type UniformSetCacheRD struct {
	Object
}

func (me *UniformSetCacheRD) BaseClass() string {
	return "UniformSetCacheRD"
}

func NewUniformSetCacheRD() *UniformSetCacheRD {
	str := StringNameFromStr("UniformSetCacheRD") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &UniformSetCacheRD{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *UniformSetCacheRD) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *UniformSetCacheRD) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *UniformSetCacheRD) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func UniformSetCacheRDGetCache(shader RID, set int64, uniforms []RDUniform) RID {
	cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), gdc.ConstTypePtr(&set), gdc.ConstTypePtr(&uniforms)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&set)
	pinner.Pin(&uniforms)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUniformSetCacheRD.fnGetCache), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
