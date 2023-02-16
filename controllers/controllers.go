package controllers

import (
	"context"
	"decommerce/database"
	"decommerce/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)

}
func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "Invalid password"
		valid = false
	}

	return valid, msg

}

//signup 

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err}}
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		}

		conunt, err := UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone is already in use"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.FirstName, *user.LastName, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, insererr := UserCollection.InsertOne(ctx, user)
		if insererr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user generating was failed to commit in database"})
			return
		}

		defer cancel()
		c.JSON(http.StatusCreated, "sucessfully signed in")
		if err := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			return
		}

	}

}


//login 
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err}}
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password is incorrect"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancel()
		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}

		token, refreshtoken, _ := generate.TokenGenerator(*founduser.Email, *founduser.FirstName, *founduser.LastName, founduser.User_ID)
		defer cancel()
		generate.UpdateAllTokens(token, refreshtoken, founduser.User_ID)
		c.JSON(http.StatusFound, founduser)

	}

}

func ProductViewerAdmin() gin.HandlerFunc {

}


//search all 
func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productList []models.Product
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := ProductCollection.Find(ctx, bson.D{{}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "something bad happened")
			return
		}

		err = cursor.All(ctx, &productList)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer cursor.Close()

		if err := cursor.err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		defer cancel()
		c.IndentedJSON(200, productList)

	}
}


//search by query

func SearchProdcutByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchProducts[] []models.Product
		queryParam := c.Query("name")
		//check if query param is empty

		if queryParam == "" {
			log.Panicln("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(HTTP.StatusNotFound, gin.H{"Error" : "invalid search index"})
			c.Abort()
			return
		}

		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()

		searchquerydb,err := ProductCollection.Find(ctx,bson.M{"product_name":bson.M{"$regex":queryParam}})
		if err != nil {
			c.IndentedJSON(404,"something went wrong")
			return
		}

		searchquerydb.All(ctx, &searchProducts)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(404,"something went wrong")
			return
		}


		defer searchquerydb.Close(ctx)

		if err := searchquerydb.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400,"invalid request")
		}

		defer cancel()
		c.IndentedJSON(200,searchProducts)


	}

}
