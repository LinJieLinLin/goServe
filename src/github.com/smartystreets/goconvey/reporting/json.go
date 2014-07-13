// TODO: under unit test

package reporting

import (
	"bytes"
	"encoding/json"
	"github.com/smartystreets/goconvey/printing"
	"strings"
)

func (self *JsonReporter) BeginStory(story *StoryReport) {}

func (self *JsonReporter) Enter(scope *ScopeReport) {
	if _, found := self.index[scope.ID]; !found {
		self.registerScope(scope)
	}
	self.current = self.index[scope.ID]
	self.depth++
}
func (self *JsonReporter) registerScope(scope *ScopeReport) {
	next := newScopeResult(scope.Title, self.depth, scope.File, scope.Line)
	self.scopes = append(self.scopes, next)
	self.index[scope.ID] = next
}

func (self *JsonReporter) Report(report *AssertionResult) {
	self.current.Assertions = append(self.current.Assertions, report)
}

func (self *JsonReporter) Exit() {
	self.depth--
}

func (self *JsonReporter) EndStory() {
	self.report()
	self.reset()
}
func (self *JsonReporter) report() {
	self.out.Print(OpenJson + "\n")
	scopes := []string{}
	for _, scope := range self.scopes {
		serialized, err := json.Marshal(scope)
		if err != nil {
			self.out.Println(jsonMarshalFailure)
			panic(err)
		}
		var buffer bytes.Buffer
		json.Indent(&buffer, serialized, "", "  ")
		scopes = append(scopes, buffer.String())
	}
	self.out.Print(strings.Join(scopes, ",") + ",\n")
	self.out.Print(CloseJson + "\n")
}
func (self *JsonReporter) reset() {
	self.scopes = []*ScopeResult{}
	self.index = map[string]*ScopeResult{}
	self.depth = 0
}

func NewJsonReporter(out *printing.Printer) *JsonReporter {
	self := &JsonReporter{}
	self.out = out
	self.reset()
	return self
}

type JsonReporter struct {
	out     *printing.Printer
	current *ScopeResult
	index   map[string]*ScopeResult
	scopes  []*ScopeResult
	depth   int
}

const OpenJson = ">>>>>"  // "⌦"
const CloseJson = "<<<<<" // "⌫"
const jsonMarshalFailure = `

GOCONVEY_JSON_MARSHALL_FAILURE: There was an error when attempting to convert test results to JSON.
Please file a bug report and reference the code that caused this failure if possible.

Here's the panic:

`
