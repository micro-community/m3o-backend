package handler

import (
	"context"

	pb "github.com/m3o/services/admin/proto"
	"github.com/m3o/services/pkg/auth"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/errors"
)

type Admin struct{}

func New(svc *service.Service) *Admin {
	a := Admin{}
	a.consumeEvents()
	return &a
}

func (a *Admin) DeleteData(ctx context.Context, request *pb.DeleteDataRequest, response *pb.DeleteDataResponse) error {
	method := "admin.DeleteData"
	_, err := auth.VerifyMicroAdmin(ctx, method)
	if err != nil {
		return err
	}
	if len(request.TenantId) == 0 {
		return errors.BadRequest(method, "Missing tenant ID")
	}

	return a.deleteData(ctx, request.TenantId)
}
