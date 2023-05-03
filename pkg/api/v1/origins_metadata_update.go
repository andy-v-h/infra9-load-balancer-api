package api

import (
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
	"go.uber.org/zap"

	"go.infratographer.com/load-balancer-api/internal/models"
)

// lbMetadataUpdate updates an metadata about a load balancer
func (r *Router) oMetadataUpdate(c echo.Context) error {
	ctx := c.Request().Context()

	mods, err := r.oMetadataParamsBinding(c)
	if err != nil {
		r.logger.Error("failed to bind metadata params", zap.Error(err))
		return v1BadRequestResponse(c, err)
	}

	mods = append(mods, qm.Load("Origin"))

	mds, err := models.OriginMetadata(mods...).All(ctx, r.db)
	if err != nil {
		r.logger.Error("failed to get metadata", zap.Error(err))
		return v1InternalServerErrorResponse(c, err)
	}

	if len(mds) == 0 {
		return v1NotFoundResponse(c)
	} else if len(mds) != 1 {
		r.logger.Warn("ambiguous query ", zap.Any("metadatas", mds))
		return v1BadRequestResponse(c, ErrAmbiguous)
	}

	metadata := mds[0]

	payload := struct {
		Data types.JSON `json:"data"`
	}{}

	if err := c.Bind(&payload); err != nil {
		return v1BadRequestResponse(c, err)
	}

	// update origin
	metadata.Data = payload.Data

	return r.updateOMetadata(c, metadata)
}

// lbMetadataPatch patches an origin
func (r *Router) oMetadataPatch(c echo.Context) error {
	ctx := c.Request().Context()

	mods, err := r.oMetadataParamsBinding(c)
	if err != nil {
		r.logger.Error("failed to bind metadata params", zap.Error(err))
		return v1BadRequestResponse(c, err)
	}

	mods = append(mods, qm.Load("Origin"))

	mds, err := models.OriginMetadata(mods...).All(ctx, r.db)
	if err != nil {
		r.logger.Error("failed to get metadata", zap.Error(err))
		return v1InternalServerErrorResponse(c, err)
	}

	if len(mds) == 0 {
		return v1NotFoundResponse(c)
	} else if len(mds) != 1 {
		r.logger.Warn("ambiguous query ", zap.Any("metadatas", mds))
		return v1BadRequestResponse(c, ErrAmbiguous)
	}

	metadata := mds[0]

	payload := struct {
		Data *types.JSON `json:"data"`
	}{}

	if err := c.Bind(&payload); err != nil {
		r.logger.Error("failed to bind origin patch input", zap.Error(err))
		return v1BadRequestResponse(c, err)
	}

	if payload.Data != nil {
		metadata.Data = *payload.Data
	}

	return r.updateOMetadata(c, metadata)
}

func (r *Router) updateOMetadata(c echo.Context, metadata *models.OriginMetadatum) error {
	ctx := c.Request().Context()

	if _, err := metadata.Update(ctx, r.db, boil.Infer()); err != nil {
		r.logger.Error("failed to update metadata", zap.Error(err))
		return v1InternalServerErrorResponse(c, err)
	}

	return v1UpdateLBMetadataResponse(c, metadata.MetadataID)
}
