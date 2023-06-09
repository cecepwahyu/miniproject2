// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	entity "crud/entity"

	mock "github.com/stretchr/testify/mock"
)

// CustomerInterfaceRepo is an autogenerated mock type for the CustomerInterfaceRepo type
type CustomerInterfaceRepo struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: customer
func (_m *CustomerInterfaceRepo) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	ret := _m.Called(customer)

	var r0 *entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Customer) (*entity.Customer, error)); ok {
		return rf(customer)
	}
	if rf, ok := ret.Get(0).(func(*entity.Customer) *entity.Customer); ok {
		r0 = rf(customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Customer) error); ok {
		r1 = rf(customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomer provides a mock function with given fields: email
func (_m *CustomerInterfaceRepo) DeleteCustomer(email string) (interface{}, error) {
	ret := _m.Called(email)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomersById provides a mock function with given fields: id
func (_m *CustomerInterfaceRepo) GetCustomersById(id uint) (entity.Customer, error) {
	ret := _m.Called(id)

	var r0 entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.Customer, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.Customer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCustomer provides a mock function with given fields: customer
func (_m *CustomerInterfaceRepo) UpdateCustomer(customer *entity.Customer) (interface{}, error) {
	ret := _m.Called(customer)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Customer) (interface{}, error)); ok {
		return rf(customer)
	}
	if rf, ok := ret.Get(0).(func(*entity.Customer) interface{}); ok {
		r0 = rf(customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Customer) error); ok {
		r1 = rf(customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCustomerInterfaceRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewCustomerInterfaceRepo creates a new instance of CustomerInterfaceRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCustomerInterfaceRepo(t mockConstructorTestingTNewCustomerInterfaceRepo) *CustomerInterfaceRepo {
	mock := &CustomerInterfaceRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}