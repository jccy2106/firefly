// Code generated by mockery v1.0.0. DO NOT EDIT.

package operationmocks

import (
	context "context"

	core "github.com/hyperledger/firefly/pkg/core"
	dataexchange "github.com/hyperledger/firefly/pkg/dataexchange"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"

	operations "github.com/hyperledger/firefly/internal/operations"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// AddOrReuseOperation provides a mock function with given fields: ctx, op
func (_m *Manager) AddOrReuseOperation(ctx context.Context, op *core.Operation) error {
	ret := _m.Called(ctx, op)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) error); ok {
		r0 = rf(ctx, op)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrepareOperation provides a mock function with given fields: ctx, op
func (_m *Manager) PrepareOperation(ctx context.Context, op *core.Operation) (*core.PreparedOperation, error) {
	ret := _m.Called(ctx, op)

	var r0 *core.PreparedOperation
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) *core.PreparedOperation); ok {
		r0 = rf(ctx, op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.PreparedOperation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.Operation) error); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterHandler provides a mock function with given fields: ctx, handler, ops
func (_m *Manager) RegisterHandler(ctx context.Context, handler operations.OperationHandler, ops []fftypes.FFEnum) {
	_m.Called(ctx, handler, ops)
}

// ResolveOperationByID provides a mock function with given fields: ctx, opID, op
func (_m *Manager) ResolveOperationByID(ctx context.Context, opID *fftypes.UUID, op *core.OperationUpdateDTO) error {
	ret := _m.Called(ctx, opID, op)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, *core.OperationUpdateDTO) error); ok {
		r0 = rf(ctx, opID, op)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetryOperation provides a mock function with given fields: ctx, opID
func (_m *Manager) RetryOperation(ctx context.Context, opID *fftypes.UUID) (*core.Operation, error) {
	ret := _m.Called(ctx, opID)

	var r0 *core.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *core.Operation); ok {
		r0 = rf(ctx, opID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, opID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunOperation provides a mock function with given fields: ctx, op, options
func (_m *Manager) RunOperation(ctx context.Context, op *core.PreparedOperation, options ...operations.RunOperationOption) (fftypes.JSONObject, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, op)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 fftypes.JSONObject
	if rf, ok := ret.Get(0).(func(context.Context, *core.PreparedOperation, ...operations.RunOperationOption) fftypes.JSONObject); ok {
		r0 = rf(ctx, op, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fftypes.JSONObject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.PreparedOperation, ...operations.RunOperationOption) error); ok {
		r1 = rf(ctx, op, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Manager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubmitOperationUpdate provides a mock function with given fields: plugin, update
func (_m *Manager) SubmitOperationUpdate(plugin core.Named, update *operations.OperationUpdate) {
	_m.Called(plugin, update)
}

// TransferResult provides a mock function with given fields: dx, event
func (_m *Manager) TransferResult(dx dataexchange.Plugin, event dataexchange.DXEvent) {
	_m.Called(dx, event)
}

// WaitStop provides a mock function with given fields:
func (_m *Manager) WaitStop() {
	_m.Called()
}
