package reporting

import "fmt"

import (
	"github.com/smartystreets/goconvey/printing"
)

func (self *problem) BeginStory(story *StoryReport) {}

func (self *problem) Enter(scope *ScopeReport) {}

func (self *problem) Report(report *AssertionResult) {
	if report.Error != nil {
		self.errors = append(self.errors, report)
	} else if report.Failure != "" {
		self.failures = append(self.failures, report)
	}
}

func (self *problem) Exit() {}

func (self *problem) EndStory() {
	self.out.Println("")
	self.show(self.showErrors, redColor)
	self.show(self.showFailures, yellowColor)
	self.prepareForNextStory()
}
func (self *problem) show(display func(), color string) {
	fmt.Print(color)
	display()
	fmt.Print(resetColor)
	self.out.Dedent()
}
func (self *problem) showErrors() {
	for i, e := range self.errors {
		if i == 0 {
			self.out.Println("\nErrors:\n")
			self.out.Indent()
		}
		self.out.Println(errorTemplate, e.File, e.Line, e.Error, e.StackTrace)
	}
}
func (self *problem) showFailures() {
	for i, f := range self.failures {
		if i == 0 {
			self.out.Println("\nFailures:\n")
			self.out.Indent()
		}
		self.out.Println(failureTemplate, f.File, f.Line, f.Failure)
	}
}

func NewProblemReporter(out *printing.Printer) *problem {
	self := problem{}
	self.out = out
	self.prepareForNextStory()
	return &self
}
func (self *problem) prepareForNextStory() {
	self.errors = []*AssertionResult{}
	self.failures = []*AssertionResult{}
}

type problem struct {
	out      *printing.Printer
	errors   []*AssertionResult
	failures []*AssertionResult
}
