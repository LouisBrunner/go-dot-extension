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

type ptrsForObjectList struct {
	fnGetClass                  gdc.MethodBindPtr
	fnIsClass                   gdc.MethodBindPtr
	fnSet                       gdc.MethodBindPtr
	fnGet                       gdc.MethodBindPtr
	fnSetIndexed                gdc.MethodBindPtr
	fnGetIndexed                gdc.MethodBindPtr
	fnGetPropertyList           gdc.MethodBindPtr
	fnGetMethodList             gdc.MethodBindPtr
	fnPropertyCanRevert         gdc.MethodBindPtr
	fnPropertyGetRevert         gdc.MethodBindPtr
	fnNotification              gdc.MethodBindPtr
	fnToString                  gdc.MethodBindPtr
	fnGetInstanceId             gdc.MethodBindPtr
	fnSetScript                 gdc.MethodBindPtr
	fnGetScript                 gdc.MethodBindPtr
	fnSetMeta                   gdc.MethodBindPtr
	fnRemoveMeta                gdc.MethodBindPtr
	fnGetMeta                   gdc.MethodBindPtr
	fnHasMeta                   gdc.MethodBindPtr
	fnGetMetaList               gdc.MethodBindPtr
	fnAddUserSignal             gdc.MethodBindPtr
	fnHasUserSignal             gdc.MethodBindPtr
	fnEmitSignal                gdc.MethodBindPtr
	fnCall                      gdc.MethodBindPtr
	fnCallDeferred              gdc.MethodBindPtr
	fnSetDeferred               gdc.MethodBindPtr
	fnCallv                     gdc.MethodBindPtr
	fnHasMethod                 gdc.MethodBindPtr
	fnHasSignal                 gdc.MethodBindPtr
	fnGetSignalList             gdc.MethodBindPtr
	fnGetSignalConnectionList   gdc.MethodBindPtr
	fnGetIncomingConnections    gdc.MethodBindPtr
	fnConnect                   gdc.MethodBindPtr
	fnDisconnect                gdc.MethodBindPtr
	fnIsConnected               gdc.MethodBindPtr
	fnSetBlockSignals           gdc.MethodBindPtr
	fnIsBlockingSignals         gdc.MethodBindPtr
	fnNotifyPropertyListChanged gdc.MethodBindPtr
	fnSetMessageTranslation     gdc.MethodBindPtr
	fnCanTranslateMessages      gdc.MethodBindPtr
	fnTr                        gdc.MethodBindPtr
	fnTrN                       gdc.MethodBindPtr
	fnIsQueuedForDeletion       gdc.MethodBindPtr
	fnCancelFree                gdc.MethodBindPtr
}

var ptrsForObject ptrsForObjectList

func initObjectPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Object")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_class")
		defer methodName.Destroy()
		ptrsForObject.fnGetClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_class")
		defer methodName.Destroy()
		ptrsForObject.fnIsClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("set")
		defer methodName.Destroy()
		ptrsForObject.fnSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("get")
		defer methodName.Destroy()
		ptrsForObject.fnGet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
	{
		methodName := StringNameFromStr("set_indexed")
		defer methodName.Destroy()
		ptrsForObject.fnSetIndexed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3500910842))
	}
	{
		methodName := StringNameFromStr("get_indexed")
		defer methodName.Destroy()
		ptrsForObject.fnGetIndexed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4006125091))
	}
	{
		methodName := StringNameFromStr("get_property_list")
		defer methodName.Destroy()
		ptrsForObject.fnGetPropertyList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_method_list")
		defer methodName.Destroy()
		ptrsForObject.fnGetMethodList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("property_can_revert")
		defer methodName.Destroy()
		ptrsForObject.fnPropertyCanRevert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("property_get_revert")
		defer methodName.Destroy()
		ptrsForObject.fnPropertyGetRevert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
	{
		methodName := StringNameFromStr("notification")
		defer methodName.Destroy()
		ptrsForObject.fnNotification = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4023243586))
	}
	{
		methodName := StringNameFromStr("to_string")
		defer methodName.Destroy()
		ptrsForObject.fnToString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("get_instance_id")
		defer methodName.Destroy()
		ptrsForObject.fnGetInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_script")
		defer methodName.Destroy()
		ptrsForObject.fnSetScript = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1114965689))
	}
	{
		methodName := StringNameFromStr("get_script")
		defer methodName.Destroy()
		ptrsForObject.fnGetScript = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1214101251))
	}
	{
		methodName := StringNameFromStr("set_meta")
		defer methodName.Destroy()
		ptrsForObject.fnSetMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("remove_meta")
		defer methodName.Destroy()
		ptrsForObject.fnRemoveMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_meta")
		defer methodName.Destroy()
		ptrsForObject.fnGetMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3990617847))
	}
	{
		methodName := StringNameFromStr("has_meta")
		defer methodName.Destroy()
		ptrsForObject.fnHasMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_meta_list")
		defer methodName.Destroy()
		ptrsForObject.fnGetMetaList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("add_user_signal")
		defer methodName.Destroy()
		ptrsForObject.fnAddUserSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 85656714))
	}
	{
		methodName := StringNameFromStr("has_user_signal")
		defer methodName.Destroy()
		ptrsForObject.fnHasUserSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("emit_signal")
		defer methodName.Destroy()
		ptrsForObject.fnEmitSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4047867050))
	}
	{
		methodName := StringNameFromStr("call")
		defer methodName.Destroy()
		ptrsForObject.fnCall = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3400424181))
	}
	{
		methodName := StringNameFromStr("call_deferred")
		defer methodName.Destroy()
		ptrsForObject.fnCallDeferred = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3400424181))
	}
	{
		methodName := StringNameFromStr("set_deferred")
		defer methodName.Destroy()
		ptrsForObject.fnSetDeferred = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("callv")
		defer methodName.Destroy()
		ptrsForObject.fnCallv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1260104456))
	}
	{
		methodName := StringNameFromStr("has_method")
		defer methodName.Destroy()
		ptrsForObject.fnHasMethod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_signal")
		defer methodName.Destroy()
		ptrsForObject.fnHasSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_signal_list")
		defer methodName.Destroy()
		ptrsForObject.fnGetSignalList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_signal_connection_list")
		defer methodName.Destroy()
		ptrsForObject.fnGetSignalConnectionList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3147814860))
	}
	{
		methodName := StringNameFromStr("get_incoming_connections")
		defer methodName.Destroy()
		ptrsForObject.fnGetIncomingConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("connect")
		defer methodName.Destroy()
		ptrsForObject.fnConnect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1518946055))
	}
	{
		methodName := StringNameFromStr("disconnect")
		defer methodName.Destroy()
		ptrsForObject.fnDisconnect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1874754934))
	}
	{
		methodName := StringNameFromStr("is_connected")
		defer methodName.Destroy()
		ptrsForObject.fnIsConnected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 768136979))
	}
	{
		methodName := StringNameFromStr("set_block_signals")
		defer methodName.Destroy()
		ptrsForObject.fnSetBlockSignals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_blocking_signals")
		defer methodName.Destroy()
		ptrsForObject.fnIsBlockingSignals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("notify_property_list_changed")
		defer methodName.Destroy()
		ptrsForObject.fnNotifyPropertyListChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_message_translation")
		defer methodName.Destroy()
		ptrsForObject.fnSetMessageTranslation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("can_translate_messages")
		defer methodName.Destroy()
		ptrsForObject.fnCanTranslateMessages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("tr")
		defer methodName.Destroy()
		ptrsForObject.fnTr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1195764410))
	}
	{
		methodName := StringNameFromStr("tr_n")
		defer methodName.Destroy()
		ptrsForObject.fnTrN = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 162698058))
	}
	{
		methodName := StringNameFromStr("is_queued_for_deletion")
		defer methodName.Destroy()
		ptrsForObject.fnIsQueuedForDeletion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("cancel_free")
		defer methodName.Destroy()
		ptrsForObject.fnCancelFree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type Object struct {
	obj gdc.ObjectPtr
}

func (me *Object) SetBaseObject(obj gdc.ObjectPtr) {
	me.obj = obj
}

func (me *Object) BaseClass() string {
	return "Object"
}

func NewObject() *Object {
	str := StringNameFromStr("Object") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Object{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	ObjectNotificationPostinitialize = 0
	ObjectNotificationPredelete      = 1
)

// Enums

type ObjectConnectFlags int

const (
	ObjectConnectFlagsConnectDeferred         ObjectConnectFlags = 1
	ObjectConnectFlagsConnectPersist          ObjectConnectFlags = 2
	ObjectConnectFlagsConnectOneShot          ObjectConnectFlags = 4
	ObjectConnectFlagsConnectReferenceCounted ObjectConnectFlags = 8
)

func (me *Object) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Object) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Object) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Object) GetClass() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetClass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) IsClass(class String) bool {
	cargs := []gdc.ConstTypePtr{class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnIsClass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) Set(property StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) Get(property StringName) Variant {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) SetIndexed(property_path NodePath, value Variant) {
	cargs := []gdc.ConstTypePtr{property_path.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnSetIndexed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) GetIndexed(property_path NodePath) Variant {
	cargs := []gdc.ConstTypePtr{property_path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetIndexed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) GetPropertyList() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetPropertyList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Object) GetMethodList() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetMethodList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Object) PropertyCanRevert(property StringName) bool {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnPropertyCanRevert), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) PropertyGetRevert(property StringName) Variant {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnPropertyGetRevert), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) Notification(what int64, reversed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&what), gdc.ConstTypePtr(&reversed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnNotification), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) ToString() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnToString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) GetInstanceId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) SetScript(script Variant) {
	cargs := []gdc.ConstTypePtr{script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnSetScript), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) GetScript() Variant {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetScript), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) SetMeta(name StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnSetMeta), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) RemoveMeta(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnRemoveMeta), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) GetMeta(name StringName, default_ Variant) Variant {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), default_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetMeta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) HasMeta(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnHasMeta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) GetMetaList() []StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetMetaList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[StringName](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Object) AddUserSignal(signal String, arguments Array) {
	cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), arguments.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnAddUserSignal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) HasUserSignal(signal StringName) bool {
	cargs := []gdc.ConstTypePtr{signal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnHasUserSignal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) EmitSignal(signal StringName, varargs ...Variant) Error {
	cargs := make([]gdc.ConstVariantPtr, 0, 1+len(varargs))
	var0 := signal.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	defer ret.Destroy()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForObject.fnEmitSignal), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return Error(-1)
	}
	retInt, err := ret.AsInt()
	if err != nil {
		log.Printf("Error converting return value to int enum: %v", err) // FIXME: bad logging
		return Error(-1)
	}
	return Error(retInt.Get())
}

func (me *Object) Call(method StringName, varargs ...Variant) Variant {
	cargs := make([]gdc.ConstVariantPtr, 0, 1+len(varargs))
	var0 := method.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForObject.fnCall), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return *ret
	}
	return *ret
}

func (me *Object) CallDeferred(method StringName, varargs ...Variant) Variant {
	cargs := make([]gdc.ConstVariantPtr, 0, 1+len(varargs))
	var0 := method.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForObject.fnCallDeferred), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return *ret
	}
	return *ret
}

func (me *Object) SetDeferred(property StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnSetDeferred), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) Callv(method StringName, arg_array Array) Variant {
	cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), arg_array.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnCallv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) HasMethod(method StringName) bool {
	cargs := []gdc.ConstTypePtr{method.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnHasMethod), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) HasSignal(signal StringName) bool {
	cargs := []gdc.ConstTypePtr{signal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnHasSignal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) GetSignalList() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetSignalList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Object) GetSignalConnectionList(signal StringName) []Dictionary {
	cargs := []gdc.ConstTypePtr{signal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetSignalConnectionList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Object) GetIncomingConnections() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnGetIncomingConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Object) Connect(signal StringName, callable Callable, flags int64) Error {
	cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), callable.AsCTypePtr(), gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnConnect), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Object) Disconnect(signal StringName, callable Callable) {
	cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), callable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnDisconnect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) IsConnected(signal StringName, callable Callable) bool {
	cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), callable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnIsConnected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) SetBlockSignals(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnSetBlockSignals), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) IsBlockingSignals() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnIsBlockingSignals), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) NotifyPropertyListChanged() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnNotifyPropertyListChanged), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) SetMessageTranslation(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnSetMessageTranslation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Object) CanTranslateMessages() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnCanTranslateMessages), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) Tr(message StringName, context StringName) String {
	cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), context.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnTr), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) TrN(message StringName, plural_message StringName, n int64, context StringName) String {
	cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), plural_message.AsCTypePtr(), gdc.ConstTypePtr(&n), context.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&n)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnTrN), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Object) IsQueuedForDeletion() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnIsQueuedForDeletion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Object) CancelFree() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForObject.fnCancelFree), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type ObjectScriptChangedSignalFn func()

func (me *Object) ConnectScriptChanged(subs SignalSubscribers, fn ObjectScriptChangedSignalFn) {
	sig := StringNameFromStr("script_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Object) DisconnectScriptChanged(subs SignalSubscribers, fn ObjectScriptChangedSignalFn) {
	sig := StringNameFromStr("script_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ObjectPropertyListChangedSignalFn func()

func (me *Object) ConnectPropertyListChanged(subs SignalSubscribers, fn ObjectPropertyListChangedSignalFn) {
	sig := StringNameFromStr("property_list_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Object) DisconnectPropertyListChanged(subs SignalSubscribers, fn ObjectPropertyListChangedSignalFn) {
	sig := StringNameFromStr("property_list_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
