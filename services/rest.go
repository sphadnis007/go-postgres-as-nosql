package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-postgres-as-nosql/configs"
	"github.com/google/uuid"
)

func StartRESTServer(pOps ProductOpsI) {
	router := gin.Default()

	router.GET("/cart/:cartID", func(c *gin.Context) {
		cartid := c.Param("cartID")
		obj, notPresent := pOps.Get(cartid)
		if obj != nil {
			c.JSON(http.StatusOK, obj)
		} else if notPresent {
			c.JSON(http.StatusNotFound, "{ Oops!!! No such cart available }")
		} else {
			c.JSON(http.StatusInternalServerError, "{ Error while getting cart }")
		}
	})

	router.GET("/cart", func(c *gin.Context) {
		carts, err := pOps.GetAllCarts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("{ Error: %s }", err))
		} else {
			if len(carts) == 0 {
				carts = make([]string, 0)
			}
			c.JSON(http.StatusOK, carts)
		}
	})

	router.POST("/cart/add/:numOfObjects", func(c *gin.Context) {
		numOfObjects := c.Param("numOfObjects")
		num, err := strconv.Atoi(numOfObjects)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Error: Invalid Number of objects")
		} else {
			cartID, err := pOps.AddFakeData(num)
			if err != nil {
				c.JSON(http.StatusInternalServerError, fmt.Sprintf("{ Error: %s }", err))
			} else {
				c.JSON(http.StatusAccepted, fmt.Sprintf("{ cart_id: %s }", cartID))
			}
		}
	})

	router.DELETE("/cart/:cartID/:productID", func(c *gin.Context) {
		cartID := c.Param("cartID")
		productID := c.Param("productID")

		if _, err := uuid.Parse(cartID); err != nil {
			c.JSON(http.StatusBadRequest, "Invalid Cart ID")
		} else if _, err := uuid.Parse(productID); err != nil {
			c.JSON(http.StatusBadRequest, "Invalid Product ID")
		} else {
			err := pOps.Delete(cartID, productID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err))
			} else {
				c.JSON(http.StatusOK, "Product Removed from the Cart!!!")
			}
		}
	})

	// Run() Starts REST router on localhost:8080
	// Assigning custom port value
	router.Run(configs.HTTPRouterPort)
}
