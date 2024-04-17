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

type ptrsForScriptList struct {
	fnCanInstantiate          gdc.MethodBindPtr
	fnInstanceHas             gdc.MethodBindPtr
	fnHasSourceCode           gdc.MethodBindPtr
	fnGetSourceCode           gdc.MethodBindPtr
	fnSetSourceCode           gdc.MethodBindPtr
	fnReload                  gdc.MethodBindPtr
	fnGetBaseScript           gdc.MethodBindPtr
	fnGetInstanceBaseType     gdc.MethodBindPtr
	fnGetGlobalName           gdc.MethodBindPtr
	fnHasScriptSignal         gdc.MethodBindPtr
	fnGetScriptPropertyList   gdc.MethodBindPtr
	fnGetScriptMethodList     gdc.MethodBindPtr
	fnGetScriptSignalList     gdc.MethodBindPtr
	fnGetScriptConstantMap    gdc.MethodBindPtr
	fnGetPropertyDefaultValue gdc.MethodBindPtr
	fnIsTool                  gdc.MethodBindPtr
	fnIsAbstract              gdc.MethodBindPtr
}

var ptrsForScript ptrsForScriptList

func initScriptPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Script")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("can_instantiate")
		defer methodName.Destroy()
		ptrsForScript.fnCanInstantiate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("instance_has")
		defer methodName.Destroy()
		ptrsForScript.fnInstanceHas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 397768994))
	}
	{
		methodName := StringNameFromStr("has_source_code")
		defer methodName.Destroy()
		ptrsForScript.fnHasSourceCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_source_code")
		defer methodName.Destroy()
		ptrsForScript.fnGetSourceCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_source_code")
		defer methodName.Destroy()
		ptrsForScript.fnSetSourceCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("reload")
		defer methodName.Destroy()
		ptrsForScript.fnReload = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1633102583))
	}
	{
		methodName := StringNameFromStr("get_base_script")
		defer methodName.Destroy()
		ptrsForScript.fnGetBaseScript = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 278624046))
	}
	{
		methodName := StringNameFromStr("get_instance_base_type")
		defer methodName.Destroy()
		ptrsForScript.fnGetInstanceBaseType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("get_global_name")
		defer methodName.Destroy()
		ptrsForScript.fnGetGlobalName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("has_script_signal")
		defer methodName.Destroy()
		ptrsForScript.fnHasScriptSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_script_property_list")
		defer methodName.Destroy()
		ptrsForScript.fnGetScriptPropertyList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("get_script_method_list")
		defer methodName.Destroy()
		ptrsForScript.fnGetScriptMethodList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("get_script_signal_list")
		defer methodName.Destroy()
		ptrsForScript.fnGetScriptSignalList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("get_script_constant_map")
		defer methodName.Destroy()
		ptrsForScript.fnGetScriptConstantMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
	}
	{
		methodName := StringNameFromStr("get_property_default_value")
		defer methodName.Destroy()
		ptrsForScript.fnGetPropertyDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2138907829))
	}
	{
		methodName := StringNameFromStr("is_tool")
		defer methodName.Destroy()
		ptrsForScript.fnIsTool = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_abstract")
		defer methodName.Destroy()
		ptrsForScript.fnIsAbstract = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type Script struct {
	Resource
}

func (me *Script) BaseClass() string {
	return "Script"
}

func NewScript() *Script {
	str := StringNameFromStr("Script") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Script{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Script) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Script) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Script) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Script) CanInstantiate() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnCanInstantiate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Script) InstanceHas(base_object Object) bool {
	cargs := []gdc.ConstTypePtr{base_object.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnInstanceHas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Script) HasSourceCode() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnHasSourceCode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Script) GetSourceCode() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetSourceCode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Script) SetSourceCode(source String) {
	cargs := []gdc.ConstTypePtr{source.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnSetSourceCode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Script) Reload(keep_state bool) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep_state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&keep_state)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnReload), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Script) GetBaseScript() Script {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewScript()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetBaseScript), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Script) GetInstanceBaseType() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetInstanceBaseType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Script) GetGlobalName() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetGlobalName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Script) HasScriptSignal(signal_name StringName) bool {
	cargs := []gdc.ConstTypePtr{signal_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnHasScriptSignal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Script) GetScriptPropertyList() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetScriptPropertyList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Script) GetScriptMethodList() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetScriptMethodList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Script) GetScriptSignalList() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetScriptSignalList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Script) GetScriptConstantMap() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetScriptConstantMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Script) GetPropertyDefaultValue(property StringName) Variant {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnGetPropertyDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Script) IsTool() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnIsTool), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Script) IsAbstract() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScript.fnIsAbstract), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
