package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/api/models"
)

// CreateIncomeProducts godoc
// @Router       /income_products [POST]
// @Summary      Creates a new income products
// @Description  create a new income products
// @Tags         income_products
// @Accept       json
// @Produce      json
// @Param		 income_products body models.CreateIncomeProducts false "income_products"
// @Success      201  {object}  string
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateIncomeProducts(c *gin.Context) {
	var incomeProducts = models.CreateIncomeProducts{}

	if err := c.ShouldBindJSON(&incomeProducts); err != nil {
		handleResponse(c, h.log, "error while binding json", http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.IncomeProduct().CreateMultiple(context.Background(), incomeProducts)
	if err != nil {
		handleResponse(c, h.log, "error while creating incomeProducts", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, "created")
}

// GetIncomeProductsList godoc
// @Router       /income_products [GET]
// @Summary      Get income products list
// @Description  get income products list
// @Tags         income_products
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.IncomesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetIncomeProductsList(c *gin.Context) {
	var (
		page, limit int
		search      string
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

	search = c.Query("search")
	resp, err := h.services.IncomeProduct().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, resp)
}

// UpdateIncomeProducts godoc
// @Router       /income_products [PUT]
// @Summary      Update income products
// @Description  update income products
// @Tags         income_products
// @Accept       json
// @Produce      json
// @Param        income_products body models.UpdateIncomeProducts false "income_products"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateIncomeProducts(c *gin.Context) {
	body := models.UpdateIncomeProducts{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handleResponse(c, h.log, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.IncomeProduct().UpdateMultiple(context.Background(), body); err != nil {
		handleResponse(c, h.log, "error is while updating multiple income products", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "success", http.StatusOK, "success")
}

// DeleteIncomeProducts godoc
// @Router       /income_products [Delete]
// @Summary      Delete income products
// @Description  delete income products
// @Tags         income_products
// @Accept       json
// @Produce      json
// @Param        ids body models.DeleteIncomeProducts false "ids"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteIncomeProducts(c *gin.Context) {
	body := models.DeleteIncomeProducts{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handleResponse(c, h.log, "error is while reading body", http.StatusBadRequest, err.Error())
	}
	if err := h.services.IncomeProduct().DeleteMultiple(context.Background(), body); err != nil {
		handleResponse(c, h.log, "error is deleting income product", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "success", http.StatusOK, "income products deleted!")
}

// IncomeProductReport godoc
// @Router       /income_product/report [GET]
// @Summary      Get  income product report
// @Description  get income product report
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        from query string false "from"
// @Param        to query string false "to"
// @Param        branch_id query string false "branch_id"
// @Success      201  {object}  models.IncomeProductReportList
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) IncomeProductReport(c *gin.Context) {
	var (
		page, limit int
		branchID    string
		from, to    string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error while converting pege", http.StatusBadRequest, err)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error while converting limit", http.StatusBadRequest, err)
		return
	}

	from = c.Query("from")
	to = c.Query("to")
	branchID = c.Query("branch_id")

	incomeProductReport, err := h.services.IncomeProduct().IncomeProductList(context.Background(), models.IncomeProductReportRequest{
		Page:     page,
		Limit:    limit,
		BranchID: branchID,
		From:     from,
		To:       to,
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting income product report", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, incomeProductReport)
}