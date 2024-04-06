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

type ptrsForEncodedObjectAsIDList struct {
	fnSetObjectId gdc.MethodBindPtr
	fnGetObjectId gdc.MethodBindPtr
}

var ptrsForEncodedObjectAsID ptrsForEncodedObjectAsIDList

func initEncodedObjectAsIDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EncodedObjectAsID")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_object_id")
		defer methodName.Destroy()
		ptrsForEncodedObjectAsID.fnSetObjectId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_object_id")
		defer methodName.Destroy()
		ptrsForEncodedObjectAsID.fnGetObjectId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type EncodedObjectAsID struct {
	RefCounted
}

func (me *EncodedObjectAsID) BaseClass() string {
	return "EncodedObjectAsID"
}

func NewEncodedObjectAsID() *EncodedObjectAsID {
	str := StringNameFromStr("EncodedObjectAsID") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EncodedObjectAsID{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EncodedObjectAsID) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EncodedObjectAsID) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EncodedObjectAsID) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EncodedObjectAsID) SetObjectId(id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEncodedObjectAsID.fnSetObjectId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EncodedObjectAsID) GetObjectId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEncodedObjectAsID.fnGetObjectId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
