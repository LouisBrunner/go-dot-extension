// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CodeEdit struct {
  obj gdc.ObjectPtr
}

func (me *CodeEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CodeEdit) BaseClass() string {
  return "CodeEdit"
}

type CodeEditCodeCompletionKind int
const (
  CodeEditKindClass CodeEditCodeCompletionKind = 0
  CodeEditKindFunction CodeEditCodeCompletionKind = 1
  CodeEditKindSignal CodeEditCodeCompletionKind = 2
  CodeEditKindVariable CodeEditCodeCompletionKind = 3
  CodeEditKindMember CodeEditCodeCompletionKind = 4
  CodeEditKindEnum CodeEditCodeCompletionKind = 5
  CodeEditKindConstant CodeEditCodeCompletionKind = 6
  CodeEditKindNodePath CodeEditCodeCompletionKind = 7
  CodeEditKindFilePath CodeEditCodeCompletionKind = 8
  CodeEditKindPlainText CodeEditCodeCompletionKind = 9
)

type CodeEditCodeCompletionLocation int
const (
  CodeEditLocationLocal CodeEditCodeCompletionLocation = 0
  CodeEditLocationParentMask CodeEditCodeCompletionLocation = 256
  CodeEditLocationOtherUserCode CodeEditCodeCompletionLocation = 512
  CodeEditLocationOther CodeEditCodeCompletionLocation = 1024
)
