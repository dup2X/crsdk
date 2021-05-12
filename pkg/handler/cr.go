package handler

import (
	"context"

	"github.com/dup2X/crsdk/pkg/errors"
	"github.com/dup2X/crsdk/pkg/idl"
	"github.com/dup2X/crsdk/pkg/model"
	"github.com/dup2X/gopkg/errcode"
	"github.com/dup2X/gopkg/logger"
)

func GetCR(ctx context.Context, req *idl.CRQueryReq) (*idl.CRRecord, *errors.CrErr) {
	logger.Debugf(ctx, logger.DLTagUndefined, "debug")
	return nil, errors.GenError(errcode.Success, nil)
}

func CreateCR(ctx context.Context, req *idl.CRCreateReq) (*idl.CRRecord, *errors.CrErr) {
	id, err := model.CreateRevision(ctx, req)
	if err != nil {
		return nil, errors.GenError(errcode.RCommonParamInvalid, err)
	}
	r := &idl.CRRecord{DemandID: id}
	return r, errors.GenError(errcode.Success, nil)
}
