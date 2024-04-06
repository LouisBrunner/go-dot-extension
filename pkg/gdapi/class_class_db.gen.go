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

type ptrsForClassDBList struct {
	fnGetClassList                gdc.MethodBindPtr
	fnGetInheritersFromClass      gdc.MethodBindPtr
	fnGetParentClass              gdc.MethodBindPtr
	fnClassExists                 gdc.MethodBindPtr
	fnIsParentClass               gdc.MethodBindPtr
	fnCanInstantiate              gdc.MethodBindPtr
	fnInstantiate                 gdc.MethodBindPtr
	fnClassHasSignal              gdc.MethodBindPtr
	fnClassGetSignal              gdc.MethodBindPtr
	fnClassGetSignalList          gdc.MethodBindPtr
	fnClassGetPropertyList        gdc.MethodBindPtr
	fnClassGetProperty            gdc.MethodBindPtr
	fnClassSetProperty            gdc.MethodBindPtr
	fnClassHasMethod              gdc.MethodBindPtr
	fnClassGetMethodList          gdc.MethodBindPtr
	fnClassGetIntegerConstantList gdc.MethodBindPtr
	fnClassHasIntegerConstant     gdc.MethodBindPtr
	fnClassGetIntegerConstant     gdc.MethodBindPtr
	fnClassHasEnum                gdc.MethodBindPtr
	fnClassGetEnumList            gdc.MethodBindPtr
	fnClassGetEnumConstants       gdc.MethodBindPtr
	fnClassGetIntegerConstantEnum gdc.MethodBindPtr
	fnIsClassEnabled              gdc.MethodBindPtr
}

var ptrsForClassDB ptrsForClassDBList

func initClassDBPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ClassDB")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_class_list")
		defer methodName.Destroy()
		ptrsForClassDB.fnGetClassList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("get_inheriters_from_class")
		defer methodName.Destroy()
		ptrsForClassDB.fnGetInheritersFromClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1761182771))
	}
	{
		methodName := StringNameFromStr("get_parent_class")
		defer methodName.Destroy()
		ptrsForClassDB.fnGetParentClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965194235))
	}
	{
		methodName := StringNameFromStr("class_exists")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassExists = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("is_parent_class")
		defer methodName.Destroy()
		ptrsForClassDB.fnIsParentClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("can_instantiate")
		defer methodName.Destroy()
		ptrsForClassDB.fnCanInstantiate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("instantiate")
		defer methodName.Destroy()
		ptrsForClassDB.fnInstantiate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
	{
		methodName := StringNameFromStr("class_has_signal")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassHasSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("class_get_signal")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3061114238))
	}
	{
		methodName := StringNameFromStr("class_get_signal_list")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetSignalList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3504980660))
	}
	{
		methodName := StringNameFromStr("class_get_property_list")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetPropertyList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3504980660))
	}
	{
		methodName := StringNameFromStr("class_get_property")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2498641674))
	}
	{
		methodName := StringNameFromStr("class_set_property")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassSetProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1690314931))
	}
	{
		methodName := StringNameFromStr("class_has_method")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassHasMethod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3860701026))
	}
	{
		methodName := StringNameFromStr("class_get_method_list")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetMethodList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3504980660))
	}
	{
		methodName := StringNameFromStr("class_get_integer_constant_list")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetIntegerConstantList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3031669221))
	}
	{
		methodName := StringNameFromStr("class_has_integer_constant")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassHasIntegerConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 471820014))
	}
	{
		methodName := StringNameFromStr("class_get_integer_constant")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetIntegerConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2419549490))
	}
	{
		methodName := StringNameFromStr("class_has_enum")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassHasEnum = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3860701026))
	}
	{
		methodName := StringNameFromStr("class_get_enum_list")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetEnumList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3031669221))
	}
	{
		methodName := StringNameFromStr("class_get_enum_constants")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetEnumConstants = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 661528303))
	}
	{
		methodName := StringNameFromStr("class_get_integer_constant_enum")
		defer methodName.Destroy()
		ptrsForClassDB.fnClassGetIntegerConstantEnum = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2457504236))
	}
	{
		methodName := StringNameFromStr("is_class_enabled")
		defer methodName.Destroy()
		ptrsForClassDB.fnIsClassEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
}

type ClassDB struct {
	Object
}

func (me *ClassDB) BaseClass() string {
	return "ClassDB"
}

func NewClassDB() *ClassDB {
	str := StringNameFromStr("ClassDB") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ClassDB{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ClassDB) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ClassDB) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ClassDB) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ClassDB) GetClassList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnGetClassList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) GetInheritersFromClass(class StringName) PackedStringArray {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnGetInheritersFromClass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) GetParentClass(class StringName) StringName {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnGetParentClass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) ClassExists(class StringName) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassExists), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ClassDB) IsParentClass(class StringName, inherits StringName) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), inherits.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnIsParentClass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ClassDB) CanInstantiate(class StringName) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnCanInstantiate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ClassDB) Instantiate(class StringName) Variant {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnInstantiate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) ClassHasSignal(class StringName, signal StringName) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), signal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassHasSignal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ClassDB) ClassGetSignal(class StringName, signal StringName) Dictionary {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), signal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetSignal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) ClassGetSignalList(class StringName, no_inheritance bool) []Dictionary {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetSignalList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *ClassDB) ClassGetPropertyList(class StringName, no_inheritance bool) []Dictionary {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetPropertyList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *ClassDB) ClassGetProperty(object Object, property StringName) Variant {
	cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetProperty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) ClassSetProperty(object Object, property StringName, value Variant) Error {
	cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassSetProperty), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ClassDB) ClassHasMethod(class StringName, method StringName, no_inheritance bool) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), method.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassHasMethod), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ClassDB) ClassGetMethodList(class StringName, no_inheritance bool) []Dictionary {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetMethodList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *ClassDB) ClassGetIntegerConstantList(class StringName, no_inheritance bool) PackedStringArray {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetIntegerConstantList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) ClassHasIntegerConstant(class StringName, name StringName) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassHasIntegerConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ClassDB) ClassGetIntegerConstant(class StringName, name StringName) int64 {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetIntegerConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ClassDB) ClassHasEnum(class StringName, name StringName, no_inheritance bool) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassHasEnum), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ClassDB) ClassGetEnumList(class StringName, no_inheritance bool) PackedStringArray {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetEnumList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) ClassGetEnumConstants(class StringName, enum StringName, no_inheritance bool) PackedStringArray {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), enum.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetEnumConstants), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) ClassGetIntegerConstantEnum(class StringName, name StringName, no_inheritance bool) StringName {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&no_inheritance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&no_inheritance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnClassGetIntegerConstantEnum), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ClassDB) IsClassEnabled(class StringName) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForClassDB.fnIsClassEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
