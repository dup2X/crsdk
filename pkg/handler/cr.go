package handler

import (
	"context"

	"github.com/dup2X/crsdk/pkg/errors"
	"github.com/dup2X/crsdk/pkg/idl"
	"github.com/dup2X/gopkg/errcode"
	"github.com/dup2X/gopkg/logger"
)

func GetCR(ctx context.Context, req *idl.CRQueryReq) (*idl.CRRecord, *errors.CrErr) {
	logger.Debugf(ctx, logger.DLTagUndefined, "debug")
	return nil, errors.GenError(errcode.Success, nil)
}
