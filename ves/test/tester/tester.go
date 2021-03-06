package tester

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-ves/ves/server"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io"
	"testing"
)

type Tester struct {
	*server.Mocker

	ContextVars map[string]interface{}
}

type TesterContext struct {
	*server.MockerContext
	t *testing.T
	sugar.HandlerErrorLogger
}

func (tester Tester) Set(k string, v interface{}) interface{} {
	res, _ := tester.ContextVars[k]
	tester.ContextVars[k] = v
	return res
}

func (tester Tester) Get(k string) interface{} {
	return tester.ContextVars[k]
}

func (tester Tester) ShouldGet(k string) (v interface{}, ok bool) {
	v, ok = tester.ContextVars[k]
	return
}

func (tester Tester) MustGet(k string) interface{} {
	v, ok := tester.ContextVars[k]
	if !ok {
		panic(fmt.Errorf("could not get %v from context", k))
	}
	return v
}

func StartTester(serverOptions []server.Option) (tester *Tester) {
	tester = new(Tester)
	tester.ContextVars = make(map[string]interface{})
	tester.Mocker = server.Mock(serverOptions...)
	if tester.Mocker == nil {
		panic(errors.New("req mocker error"))
	}
	return tester
}

func (tester *Tester) Context(tt *testing.T) (s *TesterContext) {
	return &TesterContext{
		MockerContext:      tester.Mocker.Context(tt),
		t:                  tt,
		HandlerErrorLogger: sugar.NewHandlerErrorLogger(tt),
	}
}

func (t *TesterContext) AssertNoError(noErr bool) *TesterContext {
	t.MockerContext = t.MockerContext.AssertNoError(noErr)
	return t
}

type ErrorObject struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func (t *TesterContext) DecodeJSON(body io.Reader, req interface{}) interface{} {
	if err := json.NewDecoder(body).Decode(req); err != nil {
		t.t.Fatal(err)
	}
	return req
}

func (tester *Tester) Release() {
	tester.Mocker.ReleaseMock()
}

func (tester *Tester) MakeAdminContext() bool {
	return true
}

func (tester *Tester) MainM(m *testing.M) {
	tester.Main(func() {
		m.Run()
	})
}

func (tester *Tester) Main(doSomething func()) {
	defer func() {
		if err := recover(); err != nil {
			sugar.PrintStack()
			tester.Logger.Error("panic", "error", err)
		}
		tester.Release()
	}()
	if !tester.MakeAdminContext() {
		return
	}
	doSomething()
}

type GoStyleTestFunc func(*testing.T)
type MinimumStyleTestFunc func(ctx *TesterContext)

func (tester *Tester) HandleTest(testFunc MinimumStyleTestFunc) GoStyleTestFunc {
	return func(t *testing.T) {
		testFunc(tester.Context(t))
	}
}

func (tester *Tester) HandleTestWithOutError(testFunc MinimumStyleTestFunc) GoStyleTestFunc {
	return func(t *testing.T) {
		testFunc(tester.Context(t).AssertNoError(true))
	}
}
