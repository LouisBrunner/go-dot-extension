package gdapi

import "github.com/LouisBrunner/go-dot-extension/pkg/gdc"

// String

func StringFromStr(str string) *String {
	ptr := (gdc.UninitializedStringPtr)(cmalloc(classSizeString))
	giface.StringNewWithUtf8CharsAndLen(ptr, str, gdc.Int(len(str)))
	return &String{
		iface: giface,
		ptr:   gdc.TypePtr(ptr),
	}
}

func (me *String) AsPtr() gdc.StringPtr {
	return gdc.StringPtr(me.ptr)
}

func (me *String) AsCPtr() gdc.ConstStringPtr {
	return gdc.ConstStringPtr(me.ptr)
}

// StringName

func StringNameFromStr(str string) StringName {
	strPtr := StringFromStr(str)
	defer strPtr.Destroy()
	return NewStringNameFromString(*strPtr)
}

func StringNameFromPtr(ptr gdc.StringNamePtr) StringName {
	return StringName{
		iface: giface,
		ptr:   gdc.TypePtr(ptr),
	}
}

func (me *StringName) AsPtr() gdc.StringNamePtr {
	return gdc.StringNamePtr(me.ptr)
}

func (me *StringName) AsCPtr() gdc.ConstStringNamePtr {
	return gdc.ConstStringNamePtr(me.ptr)
}
