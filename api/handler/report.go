package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/api/models"
)

// ProductReportList godoc
// @Router       /products/report [GET]
// @Summary      Get  product report
// @Description  get product report
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        from query string false "from"
// @Param        to query string false "to"
// @Param        branch_id query string false "branch_id"
// @Success      201  {object}  models.ProductReportList
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) ProductReportList(c *gin.Context) {
	var (
		page, limit int
		branchID    string
		from, to    string
		err         error
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

	from = c.Query("from")
	to = c.Query("to")
	branchID = c.Query("branch_id")

	productReport, err := h.services.Report().ProductReportList(context.Background(), models.ProductRepoRequest{
		Page:     page,
		Limit:    limit,
		BranchID: branchID,
		From:     from,
		To:       to,
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting product report", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, productReport)
}
