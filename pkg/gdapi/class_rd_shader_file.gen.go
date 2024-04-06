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

type ptrsForRDShaderFileList struct {
	fnSetBytecode    gdc.MethodBindPtr
	fnGetSpirv       gdc.MethodBindPtr
	fnGetVersionList gdc.MethodBindPtr
	fnSetBaseError   gdc.MethodBindPtr
	fnGetBaseError   gdc.MethodBindPtr
}

var ptrsForRDShaderFile ptrsForRDShaderFileList

func initRDShaderFilePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDShaderFile")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_bytecode")
		defer methodName.Destroy()
		ptrsForRDShaderFile.fnSetBytecode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1558064255))
	}
	{
		methodName := StringNameFromStr("get_spirv")
		defer methodName.Destroy()
		ptrsForRDShaderFile.fnGetSpirv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3340165340))
	}
	{
		methodName := StringNameFromStr("get_version_list")
		defer methodName.Destroy()
		ptrsForRDShaderFile.fnGetVersionList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_base_error")
		defer methodName.Destroy()
		ptrsForRDShaderFile.fnSetBaseError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_base_error")
		defer methodName.Destroy()
		ptrsForRDShaderFile.fnGetBaseError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
}

type RDShaderFile struct {
	Resource
}

func (me *RDShaderFile) BaseClass() string {
	return "RDShaderFile"
}

func NewRDShaderFile() *RDShaderFile {
	str := StringNameFromStr("RDShaderFile") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDShaderFile{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDShaderFile) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDShaderFile) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDShaderFile) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDShaderFile) SetBytecode(bytecode RDShaderSPIRV, version StringName) {
	cargs := []gdc.ConstTypePtr{bytecode.AsCTypePtr(), version.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderFile.fnSetBytecode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDShaderFile) GetSpirv(version StringName) RDShaderSPIRV {
	cargs := []gdc.ConstTypePtr{version.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRDShaderSPIRV()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderFile.fnGetSpirv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RDShaderFile) GetVersionList() []StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderFile.fnGetVersionList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[StringName](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *RDShaderFile) SetBaseError(error_ String) {
	cargs := []gdc.ConstTypePtr{error_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderFile.fnSetBaseError), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDShaderFile) GetBaseError() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDShaderFile.fnGetBaseError), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
