// Code generated by mockery v2.9.4. DO NOT EDIT.

package parser

import (
	antlr "github.com/antlr/antlr4/runtime/Go/antlr"
	mock "github.com/stretchr/testify/mock"
)

// ITypeParameterContext is an autogenerated mock type for the ITypeParameterContext type
type ITypeParameterContext struct {
	mock.Mock
}

// Accept provides a mock function with given fields: Visitor
func (_m *ITypeParameterContext) Accept(Visitor antlr.ParseTreeVisitor) interface{} {
	ret := _m.Called(Visitor)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(antlr.ParseTreeVisitor) interface{}); ok {
		r0 = rf(Visitor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// AddChild provides a mock function with given fields: child
func (_m *ITypeParameterContext) AddChild(child antlr.RuleContext) antlr.RuleContext {
	ret := _m.Called(child)

	var r0 antlr.RuleContext
	if rf, ok := ret.Get(0).(func(antlr.RuleContext) antlr.RuleContext); ok {
		r0 = rf(child)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(antlr.RuleContext)
		}
	}

	return r0
}

// AddErrorNode provides a mock function with given fields: badToken
func (_m *ITypeParameterContext) AddErrorNode(badToken antlr.Token) *antlr.ErrorNodeImpl {
	ret := _m.Called(badToken)

	var r0 *antlr.ErrorNodeImpl
	if rf, ok := ret.Get(0).(func(antlr.Token) *antlr.ErrorNodeImpl); ok {
		r0 = rf(badToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*antlr.ErrorNodeImpl)
		}
	}

	return r0
}

// AddTokenNode provides a mock function with given fields: token
func (_m *ITypeParameterContext) AddTokenNode(token antlr.Token) *antlr.TerminalNodeImpl {
	ret := _m.Called(token)

	var r0 *antlr.TerminalNodeImpl
	if rf, ok := ret.Get(0).(func(antlr.Token) *antlr.TerminalNodeImpl); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*antlr.TerminalNodeImpl)
		}
	}

	return r0
}

// EnterRule provides a mock function with given fields: listener
func (_m *ITypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	_m.Called(listener)
}

// ExitRule provides a mock function with given fields: listener
func (_m *ITypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	_m.Called(listener)
}

// GetAltNumber provides a mock function with given fields:
func (_m *ITypeParameterContext) GetAltNumber() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetBaseRuleContext provides a mock function with given fields:
func (_m *ITypeParameterContext) GetBaseRuleContext() *antlr.BaseRuleContext {
	ret := _m.Called()

	var r0 *antlr.BaseRuleContext
	if rf, ok := ret.Get(0).(func() *antlr.BaseRuleContext); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*antlr.BaseRuleContext)
		}
	}

	return r0
}

// GetChild provides a mock function with given fields: i
func (_m *ITypeParameterContext) GetChild(i int) antlr.Tree {
	ret := _m.Called(i)

	var r0 antlr.Tree
	if rf, ok := ret.Get(0).(func(int) antlr.Tree); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(antlr.Tree)
		}
	}

	return r0
}

// GetChildCount provides a mock function with given fields:
func (_m *ITypeParameterContext) GetChildCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetChildren provides a mock function with given fields:
func (_m *ITypeParameterContext) GetChildren() []antlr.Tree {
	ret := _m.Called()

	var r0 []antlr.Tree
	if rf, ok := ret.Get(0).(func() []antlr.Tree); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]antlr.Tree)
		}
	}

	return r0
}

// GetInvokingState provides a mock function with given fields:
func (_m *ITypeParameterContext) GetInvokingState() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetParent provides a mock function with given fields:
func (_m *ITypeParameterContext) GetParent() antlr.Tree {
	ret := _m.Called()

	var r0 antlr.Tree
	if rf, ok := ret.Get(0).(func() antlr.Tree); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(antlr.Tree)
		}
	}

	return r0
}

// GetParser provides a mock function with given fields:
func (_m *ITypeParameterContext) GetParser() antlr.Parser {
	ret := _m.Called()

	var r0 antlr.Parser
	if rf, ok := ret.Get(0).(func() antlr.Parser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(antlr.Parser)
		}
	}

	return r0
}

// GetPayload provides a mock function with given fields:
func (_m *ITypeParameterContext) GetPayload() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetRuleContext provides a mock function with given fields:
func (_m *ITypeParameterContext) GetRuleContext() antlr.RuleContext {
	ret := _m.Called()

	var r0 antlr.RuleContext
	if rf, ok := ret.Get(0).(func() antlr.RuleContext); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(antlr.RuleContext)
		}
	}

	return r0
}

// GetRuleIndex provides a mock function with given fields:
func (_m *ITypeParameterContext) GetRuleIndex() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetSourceInterval provides a mock function with given fields:
func (_m *ITypeParameterContext) GetSourceInterval() *antlr.Interval {
	ret := _m.Called()

	var r0 *antlr.Interval
	if rf, ok := ret.Get(0).(func() *antlr.Interval); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*antlr.Interval)
		}
	}

	return r0
}

// GetStart provides a mock function with given fields:
func (_m *ITypeParameterContext) GetStart() antlr.Token {
	ret := _m.Called()

	var r0 antlr.Token
	if rf, ok := ret.Get(0).(func() antlr.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(antlr.Token)
		}
	}

	return r0
}

// GetStop provides a mock function with given fields:
func (_m *ITypeParameterContext) GetStop() antlr.Token {
	ret := _m.Called()

	var r0 antlr.Token
	if rf, ok := ret.Get(0).(func() antlr.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(antlr.Token)
		}
	}

	return r0
}

// GetText provides a mock function with given fields:
func (_m *ITypeParameterContext) GetText() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsEmpty provides a mock function with given fields:
func (_m *ITypeParameterContext) IsEmpty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsTypeParameterContext provides a mock function with given fields:
func (_m *ITypeParameterContext) IsTypeParameterContext() {
	_m.Called()
}

// RemoveLastChild provides a mock function with given fields:
func (_m *ITypeParameterContext) RemoveLastChild() {
	_m.Called()
}

// SetAltNumber provides a mock function with given fields: altNumber
func (_m *ITypeParameterContext) SetAltNumber(altNumber int) {
	_m.Called(altNumber)
}

// SetException provides a mock function with given fields: _a0
func (_m *ITypeParameterContext) SetException(_a0 antlr.RecognitionException) {
	_m.Called(_a0)
}

// SetInvokingState provides a mock function with given fields: _a0
func (_m *ITypeParameterContext) SetInvokingState(_a0 int) {
	_m.Called(_a0)
}

// SetParent provides a mock function with given fields: _a0
func (_m *ITypeParameterContext) SetParent(_a0 antlr.Tree) {
	_m.Called(_a0)
}

// SetStart provides a mock function with given fields: _a0
func (_m *ITypeParameterContext) SetStart(_a0 antlr.Token) {
	_m.Called(_a0)
}

// SetStop provides a mock function with given fields: _a0
func (_m *ITypeParameterContext) SetStop(_a0 antlr.Token) {
	_m.Called(_a0)
}

// String provides a mock function with given fields: _a0, _a1
func (_m *ITypeParameterContext) String(_a0 []string, _a1 antlr.RuleContext) string {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string, antlr.RuleContext) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ToStringTree provides a mock function with given fields: _a0, _a1
func (_m *ITypeParameterContext) ToStringTree(_a0 []string, _a1 antlr.Recognizer) string {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string, antlr.Recognizer) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
