package handler

import (
	"context"
	"net/http"
	"strconv"
	"test/api/models"

	"github.com/gin-gonic/gin"
)

// ReportProduct godoc
// @Router       /report [POST]]
// @Summary      Get  product report
// @Description  get product report
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        body models.Request false "request"
// @Success      201  {object}  models.ProductReportList
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) ReportProduct(c *gin.Context) {
	var (
		page, limit int
		err         error
		request     = models.Request{}
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting pageStr", http.StatusBadRequest, err)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting limitStr", http.StatusBadRequest, err)
		return
	}

	repo, err := h.services.Report().ReportProduct(context.Background(), models.ProductReportListRequest{
		From:     request.From,
		To:       request.To,
		BranchId: request.BranchID,
		Page:     page,
		Limit:    limit,
	})
	if err != nil {
		handleResponse(c, h.log, "error", 500, err.Error())
		return
	}
	handleResponse(c, h.log, "success", 200, repo)

}

// ReportIncome godoc
// @Router       /products/report [POST]
// @Summary      Get  income report
// @Description  get income_product report
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        body models.Request false "request"
// @Success      201  {object}  models.IncomeProductReportList
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) ReportIncome(c *gin.Context) {
	var (
		page, limit int
		err         error
		request     = models.Request{}
	)
	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting pageStr", http.StatusBadRequest, err)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting limitStr", http.StatusBadRequest, err)
		return
	}
	list, err := h.services.Report().ReportIncome(context.Background(), models.IncomeProductReportListRequest{
		From:     request.From,
		To:       request.To,
		BranchID: request.BranchID,
		Page:     page,
		Limit:    limit,
	})
	if err != nil {
		handleResponse(c, h.log, "error", 500, err.Error())
		return
	}
	handleResponse(c, h.log, "success", 200, list)
}
