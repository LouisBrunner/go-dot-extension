// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptLanguageExtension struct {
  obj gdc.ObjectPtr
}

func (me *ScriptLanguageExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScriptLanguageExtension) BaseClass() string {
  return "ScriptLanguageExtension"
}

type ScriptLanguageExtensionLookupResultType int
const (
  ScriptLanguageExtensionLookupResultScriptLocation ScriptLanguageExtensionLookupResultType = 0
  ScriptLanguageExtensionLookupResultClass ScriptLanguageExtensionLookupResultType = 1
  ScriptLanguageExtensionLookupResultClassConstant ScriptLanguageExtensionLookupResultType = 2
  ScriptLanguageExtensionLookupResultClassProperty ScriptLanguageExtensionLookupResultType = 3
  ScriptLanguageExtensionLookupResultClassMethod ScriptLanguageExtensionLookupResultType = 4
  ScriptLanguageExtensionLookupResultClassSignal ScriptLanguageExtensionLookupResultType = 5
  ScriptLanguageExtensionLookupResultClassEnum ScriptLanguageExtensionLookupResultType = 6
  ScriptLanguageExtensionLookupResultClassTbdGlobalscope ScriptLanguageExtensionLookupResultType = 7
  ScriptLanguageExtensionLookupResultClassAnnotation ScriptLanguageExtensionLookupResultType = 8
  ScriptLanguageExtensionLookupResultMax ScriptLanguageExtensionLookupResultType = 9
)

type ScriptLanguageExtensionCodeCompletionLocation int
const (
  ScriptLanguageExtensionLocationLocal ScriptLanguageExtensionCodeCompletionLocation = 0
  ScriptLanguageExtensionLocationParentMask ScriptLanguageExtensionCodeCompletionLocation = 256
  ScriptLanguageExtensionLocationOtherUserCode ScriptLanguageExtensionCodeCompletionLocation = 512
  ScriptLanguageExtensionLocationOther ScriptLanguageExtensionCodeCompletionLocation = 1024
)

type ScriptLanguageExtensionCodeCompletionKind int
const (
  ScriptLanguageExtensionCodeCompletionKindClass ScriptLanguageExtensionCodeCompletionKind = 0
  ScriptLanguageExtensionCodeCompletionKindFunction ScriptLanguageExtensionCodeCompletionKind = 1
  ScriptLanguageExtensionCodeCompletionKindSignal ScriptLanguageExtensionCodeCompletionKind = 2
  ScriptLanguageExtensionCodeCompletionKindVariable ScriptLanguageExtensionCodeCompletionKind = 3
  ScriptLanguageExtensionCodeCompletionKindMember ScriptLanguageExtensionCodeCompletionKind = 4
  ScriptLanguageExtensionCodeCompletionKindEnum ScriptLanguageExtensionCodeCompletionKind = 5
  ScriptLanguageExtensionCodeCompletionKindConstant ScriptLanguageExtensionCodeCompletionKind = 6
  ScriptLanguageExtensionCodeCompletionKindNodePath ScriptLanguageExtensionCodeCompletionKind = 7
  ScriptLanguageExtensionCodeCompletionKindFilePath ScriptLanguageExtensionCodeCompletionKind = 8
  ScriptLanguageExtensionCodeCompletionKindPlainText ScriptLanguageExtensionCodeCompletionKind = 9
  ScriptLanguageExtensionCodeCompletionKindMax ScriptLanguageExtensionCodeCompletionKind = 10
)
