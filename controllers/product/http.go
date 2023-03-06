package product

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"THT-dagangan/businesses/productEntity"
	"THT-dagangan/controllers/product/request"
	"THT-dagangan/controllers/product/response"
	"THT-dagangan/helpers"
)

type ProductController struct {
	productService productEntity.Service
}

func NewProductController(s productEntity.Service) *ProductController {
	return &ProductController{
		productService: s,
	}
}

func (ctrl *ProductController) CreateNewProduct(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Product{}
	res := response.Product{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("The Data You Entered is Wrong", http.StatusBadRequest,
				err, helpers.EmptyObj{}))
	}

	domain := productEntity.Domain{}
	copier.Copy(&domain, &req)

	data, err := ctrl.productService.CreateNewProduct(ctx, &domain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator", http.StatusInternalServerError,
				err, helpers.EmptyObj{}))
	}
	copier.Copy(&res, &data)
	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created Product", http.StatusCreated,
			res))
}

func (ctrl *ProductController) GetAllProducts(c echo.Context) error {
	query := c.QueryParams()
	var page, filterPrice int
	var filterType, typePrice, qSort string
	offset := 0
	limit := 10
	params := productEntity.ParamGetProducts{}

	if pageStr := query.Get("page"); pageStr != "" {
		page, _ = strconv.Atoi(pageStr)

		if page <= 0 {
			page = 1
		}
	}

	if offsetStr := query.Get("offset"); offsetStr != "" {
		offset, _ = strconv.Atoi(offsetStr)
	}

	if limitStr := query.Get("limit"); limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	filter := query.Get("filter")
	switch filter {
	case "min_price":
		if minPriceStr := query.Get("min_price"); minPriceStr != "" {
			minPrice, _ := strconv.Atoi(minPriceStr)
			filterPrice = minPrice
			typePrice = "min"
		}
	case "max_price":
		if maxPriceStr := query.Get("max_price"); maxPriceStr != "" {
			maxPrice, _ := strconv.Atoi(maxPriceStr)
			filterPrice = maxPrice
			typePrice = "max"
		}
	}

	filterType = query.Get("type")

	sort := query.Get("sort")
	switch sort {
	case "price_asc":
		qSort = "price ASC"
	case "price_desc":
		qSort = "price DESC"
	}

	params = productEntity.ParamGetProducts{
		Page:      page,
		Offset:    offset,
		Limit:     limit,
		Price:     filterPrice,
		TypePrice: typePrice,
		Type:      filterType,
		Sort:      qSort,
	}
	data, offsetAfterGet, totalData, err := ctrl.productService.GetProducts(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Product Doesn't Exist", http.StatusNotFound,
				err, helpers.EmptyObj{}))
	}

	res := []response.Product{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offsetAfterGet,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(*data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("No Product have been made", http.StatusOK,
				data))
	}

	return c.JSON(http.StatusOK, helpers.BuildSuccessPageResponse(http.StatusOK, "success get all Product", res, resPage))
}

func (ctrl *ProductController) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, errGet := ctrl.productService.GetProductById(c.Request().Context(), uint(id))
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Product doesn't exist", http.StatusNotFound,
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.productService.DeleteProductById(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator", http.StatusInternalServerError,
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Deleted a Product", http.StatusOK,
			nil))
}

func (ctrl *ProductController) UpdateProduct(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	req := request.Product{}
	res := response.Product{}

	_, errGet := ctrl.productService.GetProductById(c.Request().Context(), uint(id))
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Product doesn't exist", http.StatusNotFound,
				errGet, helpers.EmptyObj{}))
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("The Data You Entered is Wrong", http.StatusBadRequest,
				err, helpers.EmptyObj{}))
	}

	domain := productEntity.Domain{}
	copier.Copy(&domain, &req)

	data, err := ctrl.productService.UpdateProductById(ctx, &domain, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator", http.StatusInternalServerError,
				err, helpers.EmptyObj{}))
	}
	copier.Copy(&res, &data)
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully updated Product", http.StatusOK,
			res))
}
