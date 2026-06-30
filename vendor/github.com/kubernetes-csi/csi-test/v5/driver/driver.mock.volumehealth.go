// TODO DELETE this file once csi-test is regenerated against spec and revendored.

package driver

import (
	context "context"
	reflect "reflect"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	gomock "github.com/golang/mock/gomock"
)

// ControllerListVolumeHealth mocks base method.
func (m *MockControllerServer) ControllerListVolumeHealth(arg0 context.Context, arg1 *csi.ControllerListVolumeHealthRequest) (*csi.ControllerListVolumeHealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerListVolumeHealth", arg0, arg1)
	ret0, _ := ret[0].(*csi.ControllerListVolumeHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerListVolumeHealth indicates an expected call of ControllerListVolumeHealth.
func (mr *MockControllerServerMockRecorder) ControllerListVolumeHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerListVolumeHealth", reflect.TypeOf((*MockControllerServer)(nil).ControllerListVolumeHealth), arg0, arg1)
}

// ControllerGetVolumeHealth mocks base method.
func (m *MockControllerServer) ControllerGetVolumeHealth(arg0 context.Context, arg1 *csi.ControllerGetVolumeHealthRequest) (*csi.ControllerGetVolumeHealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerGetVolumeHealth", arg0, arg1)
	ret0, _ := ret[0].(*csi.ControllerGetVolumeHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerGetVolumeHealth indicates an expected call of ControllerGetVolumeHealth.
func (mr *MockControllerServerMockRecorder) ControllerGetVolumeHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerGetVolumeHealth", reflect.TypeOf((*MockControllerServer)(nil).ControllerGetVolumeHealth), arg0, arg1)
}

// NodeGetVolumeHealth mocks base method.
func (m *MockNodeServer) NodeGetVolumeHealth(arg0 context.Context, arg1 *csi.NodeGetVolumeHealthRequest) (*csi.NodeGetVolumeHealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeGetVolumeHealth", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeGetVolumeHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeGetVolumeHealth indicates an expected call of NodeGetVolumeHealth.
func (mr *MockNodeServerMockRecorder) NodeGetVolumeHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeGetVolumeHealth", reflect.TypeOf((*MockNodeServer)(nil).NodeGetVolumeHealth), arg0, arg1)
}

// NodeGetStorageHealth mocks base method.
func (m *MockNodeServer) NodeGetStorageHealth(arg0 context.Context, arg1 *csi.NodeGetStorageHealthRequest) (*csi.NodeGetStorageHealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeGetStorageHealth", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeGetStorageHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeGetStorageHealth indicates an expected call of NodeGetStorageHealth.
func (mr *MockNodeServerMockRecorder) NodeGetStorageHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeGetStorageHealth", reflect.TypeOf((*MockNodeServer)(nil).NodeGetStorageHealth), arg0, arg1)
}
