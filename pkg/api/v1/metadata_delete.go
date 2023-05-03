package api

import (
	"database/sql"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"go.infratographer.com/load-balancer-api/internal/models"
)

func (r *Router) metadataDeleteFactory(m *metadataCreateModels) func(c echo.Context) error {
	return func(c echo.Context) error {
		mods, err := r.metadataParamBidningFactory(m.pathID)(c)
		if err != nil {
			r.logger.Error("error parsing query params", zap.Error(err))
			return v1BadRequestResponse(c, err)
		}

		return m.objDel(c, mods)
	}
}

func (r *Router) oMetadataDelete(c echo.Context, mods []qm.QueryMod) error {
	ctx := c.Request().Context()

	mds, err := models.LoadBalancerMetadata(mods...).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return v1InternalServerErrorResponse(c, err)
		}
	}

	if len(mds) == 0 {
		return v1NotFoundResponse(c)
	} else if len(mds) > 1 {
		return v1BadRequestResponse(c, ErrAmbiguous)
	}

	metadata := mds[0]

	_, err = metadata.Delete(ctx, r.db, false)
	if err != nil {
		r.logger.Error("error deleting metadata", zap.Error(err))
	}

	return v1DeleteMetadataResponse(c)
}

func (r *Router) deleteOriginMetadata(c echo.Context) error {
	return r.metadataDeleteFactory(&metadataCreateModels{
		pathID: "origin_id",
		objDel: r.oMetadataDelete,
	})(c)
}

func (r *Router) deleteLBMetadata(c echo.Context) error {
	return r.metadataDeleteFactory(&metadataCreateModels{
		pathID: "load_balancer_id",
		objDel: r.lbMetadataDelete,
	})(c)
}

// lbMetadataDelete is the handler for the DELETE /loadbalancers/:load_balancer_id route
func (r *Router) lbMetadataDelete(c echo.Context, mods []qm.QueryMod) error {
	ctx := c.Request().Context()

	mds, err := models.LoadBalancerMetadata(mods...).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return v1InternalServerErrorResponse(c, err)
		}
	}

	if len(mds) == 0 {
		return v1NotFoundResponse(c)
	} else if len(mds) > 1 {
		return v1BadRequestResponse(c, ErrAmbiguous)
	}

	metadata := mds[0]

	_, err = metadata.Delete(ctx, r.db, false)
	if err != nil {
		r.logger.Error("error deleting metadata", zap.Error(err))
		return v1InternalServerErrorResponse(c, err)
	}

	return v1DeleteMetadataResponse(c)
}
