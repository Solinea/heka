// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: PluginHelper)

package pipeline

import (
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of PluginHelper interface
type MockPluginHelper struct {
	ctrl     *gomock.Controller
	recorder *_MockPluginHelperRecorder
}

// Recorder for MockPluginHelper (not exported)
type _MockPluginHelperRecorder struct {
	mock *MockPluginHelper
}

func NewMockPluginHelper(ctrl *gomock.Controller) *MockPluginHelper {
	mock := &MockPluginHelper{ctrl: ctrl}
	mock.recorder = &_MockPluginHelperRecorder{mock}
	return mock
}

func (_m *MockPluginHelper) EXPECT() *_MockPluginHelperRecorder {
	return _m.recorder
}

func (_m *MockPluginHelper) ChainRouter() *ChainRouter {
	ret := _m.ctrl.Call(_m, "ChainRouter")
	ret0, _ := ret[0].(*ChainRouter)
	return ret0
}

func (_mr *_MockPluginHelperRecorder) ChainRouter() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ChainRouter")
}

func (_m *MockPluginHelper) NewDecoder(_param0 string) (DecoderRunner, bool) {
	ret := _m.ctrl.Call(_m, "NewDecoder", _param0)
	ret0, _ := ret[0].(DecoderRunner)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockPluginHelperRecorder) NewDecoder(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewDecoder", arg0)
}

func (_m *MockPluginHelper) NewDecoderSet() []DecoderRunner {
	ret := _m.ctrl.Call(_m, "NewDecoderSet")
	ret0, _ := ret[0].([]DecoderRunner)
	return ret0
}

func (_mr *_MockPluginHelperRecorder) NewDecoderSet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewDecoderSet")
}

func (_m *MockPluginHelper) PackSupply() chan *PipelinePack {
	ret := _m.ctrl.Call(_m, "PackSupply")
	ret0, _ := ret[0].(chan *PipelinePack)
	return ret0
}

func (_mr *_MockPluginHelperRecorder) PackSupply() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PackSupply")
}
