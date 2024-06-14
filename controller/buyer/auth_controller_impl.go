package buyer_controller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/ticketing_app/helper"
	buyer_web "github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	buyer_service "github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// BuyerControllerImpl is the implementation of BuyerController interface
type BuyerControllerImpl struct {
	Service buyer_service.BuyerServiceInterface
}

// NewBuyerController creates a new instance of BuyerControllerImpl
func NewBuyerController(service buyer_service.BuyerServiceInterface) *BuyerControllerImpl {
	return &BuyerControllerImpl{Service: service}
}

// Register godoc
// @Summary Register a new buyer
// @Description Register a new buyer with the given details
// @Tags buyer
// @Accept json
// @Produce json
// @Param buyer body buyer_web.BuyerRequest true "Buyer details"
// @Success 201 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/auth/register [post]
func (controller *BuyerControllerImpl) Register(c echo.Context) error {
	newUser := new(buyer_web.BuyerRequest)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(newUser); err != nil {
		return err
	}

	result, err := controller.Service.Register(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseClient(http.StatusCreated, "Success", result))
}

// Login godoc
// @Summary Login a buyer
// @Description Login a buyer with email and password
// @Tags buyer
// @Accept json
// @Produce json
// @Param buyer body buyer_web.BuyerLoginRequest true "Login details"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/auth/login [post]
func (controller *BuyerControllerImpl) Login(c echo.Context) error {
	loginUser := new(buyer_web.BuyerLoginRequest)

	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(loginUser); err != nil {
		return err
	}

	userLogin, errLogin := controller.Service.Login(loginUser.Email, loginUser.Password)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", userLogin))
}

// Update godoc
// @Summary Update a buyer's information
// @Description Update a buyer's information by ID
// @Tags buyer
// @Accept json
// @Produce json
// @Param id path int true "Buyer ID"
// @Param buyer body buyer_web.BuyerUpdateRequest true "Updated buyer details"
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/me/update [put]
func (controller *BuyerControllerImpl) Update(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	updateUser := new(buyer_web.BuyerUpdateRequest)

	if err := c.Bind(updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(updateUser); err != nil {
		return err
	}

	result, err := controller.Service.Update(userID, *updateUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}

// GetAll godoc
// @Summary Get all buyers
// @Description Get a list of all buyers
// @Tags buyer
// @Accept json
// @Produce json
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/me/buyers [get]
func (controller *BuyerControllerImpl) GetAll(c echo.Context) error {
	result, err := controller.Service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}

// GetHistory godoc
// @Summary Get buyer's purchase history
// @Description Get the purchase history of the logged-in buyer
// @Tags buyer
// @Accept json
// @Produce json
// @Success 200 {object} helper.ResponseClientModel
// @Failure 400 {object} helper.ResponseClientModel
// @Router /buyer/me/history [get]
func (controller *BuyerControllerImpl) GetHistory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helper.JwtCustomClaims)
	userID, _ := strconv.Atoi(claims.ID)

	result, err := controller.Service.GetHistory(userID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseClient(http.StatusOK, "Success", result))
}
