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

type ptrsForGridContainerList struct {
	fnSetColumns gdc.MethodBindPtr
	fnGetColumns gdc.MethodBindPtr
}

var ptrsForGridContainer ptrsForGridContainerList

func initGridContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GridContainer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_columns")
		defer methodName.Destroy()
		ptrsForGridContainer.fnSetColumns = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_columns")
		defer methodName.Destroy()
		ptrsForGridContainer.fnGetColumns = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type GridContainer struct {
	Container
}

func (me *GridContainer) BaseClass() string {
	return "GridContainer"
}

func NewGridContainer() *GridContainer {
	str := StringNameFromStr("GridContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GridContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GridContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GridContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GridContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GridContainer) SetColumns(columns int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&columns)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridContainer.fnSetColumns), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GridContainer) GetColumns() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGridContainer.fnGetColumns), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
