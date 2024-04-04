// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type ScriptLanguageExtension struct {
  ScriptLanguage
}

func (me *ScriptLanguageExtension) BaseClass() string {
  return "ScriptLanguageExtension"
}

func NewScriptLanguageExtension() *ScriptLanguageExtension {
  str := StringNameFromStr("ScriptLanguageExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ScriptLanguageExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ScriptLanguageExtensionLookupResultType int
const (
  ScriptLanguageExtensionLookupResultTypeLookupResultScriptLocation ScriptLanguageExtensionLookupResultType = 0
  ScriptLanguageExtensionLookupResultTypeLookupResultClass ScriptLanguageExtensionLookupResultType = 1
  ScriptLanguageExtensionLookupResultTypeLookupResultClassConstant ScriptLanguageExtensionLookupResultType = 2
  ScriptLanguageExtensionLookupResultTypeLookupResultClassProperty ScriptLanguageExtensionLookupResultType = 3
  ScriptLanguageExtensionLookupResultTypeLookupResultClassMethod ScriptLanguageExtensionLookupResultType = 4
  ScriptLanguageExtensionLookupResultTypeLookupResultClassSignal ScriptLanguageExtensionLookupResultType = 5
  ScriptLanguageExtensionLookupResultTypeLookupResultClassEnum ScriptLanguageExtensionLookupResultType = 6
  ScriptLanguageExtensionLookupResultTypeLookupResultClassTbdGlobalscope ScriptLanguageExtensionLookupResultType = 7
  ScriptLanguageExtensionLookupResultTypeLookupResultClassAnnotation ScriptLanguageExtensionLookupResultType = 8
  ScriptLanguageExtensionLookupResultTypeLookupResultMax ScriptLanguageExtensionLookupResultType = 9
)

type ScriptLanguageExtensionCodeCompletionLocation int
const (
  ScriptLanguageExtensionCodeCompletionLocationLocationLocal ScriptLanguageExtensionCodeCompletionLocation = 0
  ScriptLanguageExtensionCodeCompletionLocationLocationParentMask ScriptLanguageExtensionCodeCompletionLocation = 256
  ScriptLanguageExtensionCodeCompletionLocationLocationOtherUserCode ScriptLanguageExtensionCodeCompletionLocation = 512
  ScriptLanguageExtensionCodeCompletionLocationLocationOther ScriptLanguageExtensionCodeCompletionLocation = 1024
)

type ScriptLanguageExtensionCodeCompletionKind int
const (
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindClass ScriptLanguageExtensionCodeCompletionKind = 0
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindFunction ScriptLanguageExtensionCodeCompletionKind = 1
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindSignal ScriptLanguageExtensionCodeCompletionKind = 2
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindVariable ScriptLanguageExtensionCodeCompletionKind = 3
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindMember ScriptLanguageExtensionCodeCompletionKind = 4
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindEnum ScriptLanguageExtensionCodeCompletionKind = 5
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindConstant ScriptLanguageExtensionCodeCompletionKind = 6
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindNodePath ScriptLanguageExtensionCodeCompletionKind = 7
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindFilePath ScriptLanguageExtensionCodeCompletionKind = 8
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindPlainText ScriptLanguageExtensionCodeCompletionKind = 9
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindMax ScriptLanguageExtensionCodeCompletionKind = 10
)

func (me *ScriptLanguageExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScriptLanguageExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScriptLanguageExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
