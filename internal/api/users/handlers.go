package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kashyaprahul94/go-boilerplate/internal/httperrors"
	"github.com/kashyaprahul94/go-boilerplate/internal/logger"
	"github.com/kashyaprahul94/go-boilerplate/pkg/users"
)

// GetUsers handles the http request to fetch all the users in the response
func GetUsers(c *gin.Context) {
	allUsers, err := users.GetUsers()

	if err != nil {
		httpError := err.ConvertToHttpError()
		logger.Error(httpError)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, allUsers)
}

// GetUserById handles the http request to fetch the user by provided id
func GetUserById(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, vErr := strconv.Atoi(userIdParam)

	if vErr != nil || userId <= 0 {
		errDetails := gin.H{
			"providedValue": userIdParam,
			"expectedValue": "A valid user id i.e. positive integer",
		}

		logger.Error("Invalid user id supplied in the param", errDetails)
		c.JSON(http.StatusBadRequest, httperrors.BadRequestError("Invalid user id supplied in the param", errDetails))
		return
	}

	user, err := users.GetUserById(userId)

	if err != nil {
		httpError := err.ConvertToHttpError()
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser handles the http request to create a new user in the system
func CreateUser(c *gin.Context) {
	var userPayload CreateUserRequestPayload

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		logger.Error("Invalid payload body", err.Error())
		c.JSON(http.StatusBadRequest, httperrors.BadRequestError("Invalid payload body", err.Error()))
		return
	}

	newUser, err := userPayload.ToUser().Insert()

	if err != nil {
		httpError := err.ConvertToHttpError()
		logger.Error(httpError)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, newUser)
}

// UpsertUser handles the http request to create a new user in the system if it doesn't exist already, otherwise updates the user
func UpsertUser(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, vErr := strconv.Atoi(userIdParam)

	if vErr != nil || userId <= 0 {
		errDetails := gin.H{
			"providedValue": userIdParam,
			"expectedValue": "A valid user id i.e. positive integer",
		}

		logger.Error("Invalid user id supplied in the param", errDetails)
		c.JSON(http.StatusBadRequest, httperrors.BadRequestError("Invalid user id supplied in the param", errDetails))
		return
	}

	var userPayload UpdateUserRequestPayload

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		logger.Error("Invalid payload body", err.Error())
		c.JSON(http.StatusBadRequest, httperrors.BadRequestError("Invalid payload body", err.Error()))
		return
	}

	userPayload.ID = uint(userId)

	newUser, err := userPayload.ToUser().Upsert()

	if err != nil {
		httpError := err.ConvertToHttpError()
		logger.Error(httpError)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, newUser)
}

// UpsertUser handles the http request to partially update a user record, if exists
func PartialUpdateUser(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, vErr := strconv.Atoi(userIdParam)

	if vErr != nil || userId <= 0 {
		errDetails := gin.H{
			"providedValue": userIdParam,
			"expectedValue": "A valid user id i.e. positive integer",
		}

		logger.Error("Invalid user id supplied in the param", errDetails)
		c.JSON(http.StatusBadRequest, httperrors.BadRequestError("Invalid user id supplied in the param", errDetails))
		return
	}

	var userPayload UpdateUserRequestPayload

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		logger.Error("Invalid payload body", err.Error())
		c.JSON(http.StatusBadRequest, httperrors.BadRequestError("Invalid payload body", err.Error()))
		return
	}

	userPayload.ID = uint(userId)

	newUser, err := userPayload.ToUser().PartialUpdate()

	if err != nil {
		httpError := err.ConvertToHttpError()
		logger.Error(httpError)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, newUser)
}

// DeleteUser handles the http request to delete the specified user from the system if it exists
func DeleteUser(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, vErr := strconv.Atoi(userIdParam)

	if vErr != nil || userId <= 0 {
		errDetails := gin.H{
			"providedValue": userIdParam,
			"expectedValue": "A valid user id i.e. positive integer",
		}

		logger.Error("Invalid user id supplied in the param", errDetails)
		c.JSON(http.StatusBadRequest, httperrors.BadRequestError("Invalid user id supplied in the param", errDetails))
		return
	}

	deletedUser, err := users.DeleteUserById(userId)

	if err != nil {
		httpError := err.ConvertToHttpError()
		logger.Error(httpError)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, deletedUser)
}
