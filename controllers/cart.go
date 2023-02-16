package controllers

import (
	"context"
	"decommerce/database"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("productQueryID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("productQueryID is empty"))
			return

		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("userQueryID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id  is empty"))
			return

		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)

		}
		c.IndentedJSON(200, "sucessfully added to the cart")

	}

}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("productQueryID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("productQueryID is empty"))
			return

		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("userQueryID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id  is empty"))
			return

		}
		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "successfully removed from cart")

	}

}
func GetItemFromCart() gin.HandlerFunc {

}
func (app *Application)BuyFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" {
			log.Panicln("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("user id  is empty"))
		}
		context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()
		err := database.BuyItemFromCart(ctx,app.userCollection,userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,err)
		}

		c.IndentedJSON("sucessfuly placed the order")


	}

}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("productQueryID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("productQueryID is empty"))
			return

		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("userQueryID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id  is empty"))
			return

		}
		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.InstantBuyItemFromCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "successfully placed the order")

	}

}
