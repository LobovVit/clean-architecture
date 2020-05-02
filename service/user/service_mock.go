package user

import (
	"context"

	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
)

// Make sure ServiceMock implement Service interface
var _ Service = &ServiceMock{}

// ServiceMock struct used to mock test
type ServiceMock struct {
	FindByIDFunc func(ctx context.Context, req userReq.FindByID) (*userRes.FindByID, error)
}

// FindByID mock function
func (sm *ServiceMock) FindByID(ctx context.Context, req userReq.FindByID) (*userRes.FindByID, error) {
	if sm.FindByIDFunc == nil {
		panic("ServiceMock not declare FindByID function")
	}
	return sm.FindByIDFunc(ctx, req)
}
