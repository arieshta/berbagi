package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RequestGift(c echo.Context) error {
	var newRequest models.NewGiftRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "children" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "Your role can't request gift"})
	}

	c.Bind(&newRequest)
	newRequest.UserID = uint(userId)

	res, err := libdb.CreateGiftRequest(newRequest)
	if err == errors.New("package doesn't exist") {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "package doesn't exist"})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "Can't create new request"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "Request has been submitted!",
		Data:    res})
}

func RequestDonation(c echo.Context) error {
	var newRequest models.NewDonationRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "Your role can't request donation"})
	}

	c.Bind(&newRequest)
	newRequest.FoundationID = uint(userId)

	res, err := libdb.CreateDonationRequest(newRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "Can't create new request"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "Request has been submitted!",
		Data:    res})
}

func RequestService(c echo.Context) error {
	var newRequest models.NewServiceRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "Your role can't request gift"})
	}

	c.Bind(&newRequest)
	newRequest.FoundationID = uint(userId)

	res, err := libdb.CreateServiceRequest(newRequest)
	if err == errors.New("service doesn't exist") {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "service doesn't exist"})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "Can't create new request"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "Request has been submitted!",
		Data:    res})
}

// Used by children & foundation role
func GetAllRequestListController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")
	resolved := c.QueryParams().Get("resolved")

	if role != "children" && role != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't reach this"})
	}

	res, err := libdb.GetBulkRequests(uint(userId), resolved)

	roleStatus, _ := libdb.CheckUserRoleRightness(uint(userId), role)
	if !roleStatus {
		res, err = libdb.GetBulkRequests(uint(userId), "no")
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "Failed",
			Message: "can't get request list"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success getting request list",
		Data:    res})
}

func GetTypeRequestListController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")
	reqType := c.Param("type")
	resolved := c.QueryParams().Get("resolved")

	if reqType != "gift" && reqType != "donation" && reqType != "service" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid request type"})
	}

	if role != "children" && role != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't reach this"})
	}

	var res interface{}
	var err error
	if (reqType == "gift" && role == "children") || (
		reqType == "donation" && role == "foundation") || (
		reqType == "service" && role == "foundation") {
		// Get gift requests list
		res, err = libdb.GetTypeRequests(uint(userId), reqType, resolved)
	} else {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't reach this"})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't get request list"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success getting request list",
		Data:    res})
}

func DeleteRequestController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	requestId, _ := strconv.ParseUint(c.Param("request_id"), 0, 0)

	get, err := libdb.GetRequestByIdResolve(uint(requestId), "no")
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't find unresolved request"})
	}
	if uint(userId) != get.ID {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't delete other's request"})
	}

	if err := libdb.DeleteRequest(uint(requestId)); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status: "failed",
			Message: "failed to delete request",
		})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status: "success",
		Message: "success delete request",
	})
}

func GetRequestByRecipientIdController(c echo.Context) error {
	recipientId, _ := strconv.ParseUint(c.Param("recipient_id"), 0, 0)

	res, err := libdb.GetRequestByRecipientIdResolve(uint(recipientId), "no")
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't find request list"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status: "success",
		Message: "success getting request list",
		Data: res,
	})
}