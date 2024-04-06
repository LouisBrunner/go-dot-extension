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

type ptrsForResourceUIDList struct {
	fnIdToText  gdc.MethodBindPtr
	fnTextToId  gdc.MethodBindPtr
	fnCreateId  gdc.MethodBindPtr
	fnHasId     gdc.MethodBindPtr
	fnAddId     gdc.MethodBindPtr
	fnSetId     gdc.MethodBindPtr
	fnGetIdPath gdc.MethodBindPtr
	fnRemoveId  gdc.MethodBindPtr
}

var ptrsForResourceUID ptrsForResourceUIDList

func initResourceUIDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceUID")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("id_to_text")
		defer methodName.Destroy()
		ptrsForResourceUID.fnIdToText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("text_to_id")
		defer methodName.Destroy()
		ptrsForResourceUID.fnTextToId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("create_id")
		defer methodName.Destroy()
		ptrsForResourceUID.fnCreateId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("has_id")
		defer methodName.Destroy()
		ptrsForResourceUID.fnHasId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("add_id")
		defer methodName.Destroy()
		ptrsForResourceUID.fnAddId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("set_id")
		defer methodName.Destroy()
		ptrsForResourceUID.fnSetId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_id_path")
		defer methodName.Destroy()
		ptrsForResourceUID.fnGetIdPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("remove_id")
		defer methodName.Destroy()
		ptrsForResourceUID.fnRemoveId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}

}

type ResourceUID struct {
	Object
}

func (me *ResourceUID) BaseClass() string {
	return "ResourceUID"
}

func NewResourceUID() *ResourceUID {
	str := StringNameFromStr("ResourceUID") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ResourceUID{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	ResourceUIDInvalidId = -1
)

// Enums

func (me *ResourceUID) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ResourceUID) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ResourceUID) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ResourceUID) IdToText(id int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceUID.fnIdToText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ResourceUID) TextToId(text_id String) int64 {
	cargs := []gdc.ConstTypePtr{text_id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceUID.fnTextToId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ResourceUID) CreateId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceUID.fnCreateId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ResourceUID) HasId(id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceUID.fnHasId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ResourceUID) AddId(id int64, path String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceUID.fnAddId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ResourceUID) SetId(id int64, path String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceUID.fnSetId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ResourceUID) GetIdPath(id int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceUID.fnGetIdPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ResourceUID) RemoveId(id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceUID.fnRemoveId), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
