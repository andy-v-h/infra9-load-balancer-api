package api

import (
	"database/sql"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"go.infratographer.com/load-balancer-api/internal/models"
)

// listMetadata is the factory for /$thing/:thing_id/metadata routes
func (r *Router) listMetadata(m *metadataCreateModels) func(c echo.Context) error {
	return func(c echo.Context) error {
		mods, err := r.metadataParamBidningFactory(m.pathID)(c)
		if err != nil {
			r.logger.Error("error parsing query params", zap.Error(err))
			return v1BadRequestResponse(c, err)
		}

		return m.objList(c, mods)
	}
}

func (r *Router) oMetadataList(c echo.Context) error {
	return r.listMetadata(&metadataCreateModels{
		pathID:  "origin_id",
		objList: r.omList,
	})(c)
}

func (r *Router) omList(c echo.Context, mods []qm.QueryMod) error {
	ctx := c.Request().Context()

	mds, err := models.OriginMetadata(mods...).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return v1NotFoundResponse(c)
		}

		r.logger.Error("error loading metadata", zap.Error(err))

		return v1InternalServerErrorResponse(c, err)
	}

	return v1OMetadatasResponse(c, mds)
}

func (r *Router) lbMeatadataList(c echo.Context) error {
	return r.listMetadata(&metadataCreateModels{
		pathID:  "load_balancer_id",
		objList: r.lbmList,
	})(c)
}

// lbMetadataList is the handler for the GET /loadbalancers/:load_balancer_id/metadata route
func (r *Router) lbmList(c echo.Context, mods []qm.QueryMod) error {
	ctx := c.Request().Context()

	mds, err := models.LoadBalancerMetadata(mods...).All(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return v1NotFoundResponse(c)
		}

		r.logger.Error("error loading metadata", zap.Error(err))

		return v1InternalServerErrorResponse(c, err)
	}

	return v1LBMetadatasResponse(c, mds)
}
