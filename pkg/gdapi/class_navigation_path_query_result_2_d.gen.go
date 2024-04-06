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

type ptrsForNavigationPathQueryResult2DList struct {
	fnSetPath         gdc.MethodBindPtr
	fnGetPath         gdc.MethodBindPtr
	fnSetPathTypes    gdc.MethodBindPtr
	fnGetPathTypes    gdc.MethodBindPtr
	fnSetPathRids     gdc.MethodBindPtr
	fnGetPathRids     gdc.MethodBindPtr
	fnSetPathOwnerIds gdc.MethodBindPtr
	fnGetPathOwnerIds gdc.MethodBindPtr
	fnReset           gdc.MethodBindPtr
}

var ptrsForNavigationPathQueryResult2D ptrsForNavigationPathQueryResult2DList

func initNavigationPathQueryResult2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationPathQueryResult2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_path")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnSetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_path")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("set_path_types")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnSetPathTypes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_path_types")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnGetPathTypes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}
	{
		methodName := StringNameFromStr("set_path_rids")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnSetPathRids = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_path_rids")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnGetPathRids = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_path_owner_ids")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnSetPathOwnerIds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3709968205))
	}
	{
		methodName := StringNameFromStr("get_path_owner_ids")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnGetPathOwnerIds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 235988956))
	}
	{
		methodName := StringNameFromStr("reset")
		defer methodName.Destroy()
		ptrsForNavigationPathQueryResult2D.fnReset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type NavigationPathQueryResult2D struct {
	RefCounted
}

func (me *NavigationPathQueryResult2D) BaseClass() string {
	return "NavigationPathQueryResult2D"
}

func NewNavigationPathQueryResult2D() *NavigationPathQueryResult2D {
	str := StringNameFromStr("NavigationPathQueryResult2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationPathQueryResult2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type NavigationPathQueryResult2DPathSegmentType int

const (
	NavigationPathQueryResult2DPathSegmentTypePathSegmentTypeRegion NavigationPathQueryResult2DPathSegmentType = 0
	NavigationPathQueryResult2DPathSegmentTypePathSegmentTypeLink   NavigationPathQueryResult2DPathSegmentType = 1
)

func (me *NavigationPathQueryResult2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationPathQueryResult2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryResult2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationPathQueryResult2D) SetPath(path PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnSetPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryResult2D) GetPath() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryResult2D) SetPathTypes(path_types PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{path_types.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnSetPathTypes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryResult2D) GetPathTypes() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnGetPathTypes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryResult2D) SetPathRids(path_rids []RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_rids)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnSetPathRids), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryResult2D) GetPathRids() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnGetPathRids), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *NavigationPathQueryResult2D) SetPathOwnerIds(path_owner_ids PackedInt64Array) {
	cargs := []gdc.ConstTypePtr{path_owner_ids.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnSetPathOwnerIds), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationPathQueryResult2D) GetPathOwnerIds() PackedInt64Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt64Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnGetPathOwnerIds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationPathQueryResult2D) Reset() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPathQueryResult2D.fnReset), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
