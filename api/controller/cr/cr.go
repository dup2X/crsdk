package cr

import (
	"context"

	"github.com/dup2X/crsdk/pkg/handler"
	vidl "github.com/dup2X/crsdk/pkg/idl"
	"github.com/dup2X/gopkg/idl"
)

// QueryController cr controller
type QueryController struct {
}

// GetRequestIDL return requets idl obj
func (c *QueryController) GetRequestIDL() interface{} {
	return &vidl.CRQueryReq{}
}

// Do controller func
func (c *QueryController) Do(ctx context.Context, req interface{}) (interface{}, idl.APIErr) {
	// call handlers
	mreq := req.(*vidl.CRQueryReq)
	resp, err := handler.GetCR(ctx, mreq)
	return resp, err
}
