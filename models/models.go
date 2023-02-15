package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName       *string            `bson:"first_name,omitempty" json:"first_name,omitempty validate:"required,min=4,max=30"`
	LastName        *string            `bson:"last_name,omitempty" json:"last_name,omitempty validate:"required,min=4,max=30"`
	Password        *string            `bson:"password,omitempty" json:"password,omitempty validate:"required,min=6,max=20"`
	Email           *string            `bson:"email,omitempty" json:"email,omitempty validate:"required,min=4,max=30"`
	Phone           *string            `bson:"phone,omitempty" json:"phone,omitempty" validate:"required,min=7,max=13`
	Token           *string            `bson:"token,omitempty" json:"token,omitempty"`
	Refresh_Token   *string            `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	Created_At      time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	Updated_At      time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	User_ID         string            `bson:"user_id,omitempty" json:"user_id,omitempty"`
	UserCart        []ProductUser      `bson:"user_cart,omitempty" json:"user_cart,omitempty"`
	Address_Details []Address          `bson:"address_details,omitempty" json:"address_details,omitempty"`
	Order_Status    []Order            `bson:"order_status,omitempty" json:"order_detail,omitempty"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"`
	Product_Name *string            `bson:"product_name,omitempty" json:"product_name,omitempty"`
	Price        *uint64            `bson:"price,omitempty" json:"price,omit`
	Rating       *uint8             `bson:"rating,omitempty" json:"rating,omit`
	Image        *string            `bson:"image,omitempty" json:"image,omitempty"`
}

type ProductUser struct { //list of Product
	Product_ID   primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"`
	Product_Name *string            `bson:"product_name,omitempty" json:"product_name,omitempty"`
	Price        *uint64            `bson:"price,omitempty" json:"price,omit`
	Rating       *uint8             `bson:"rating,omitempty" json:"rating,omit`
	Image        *string            `bson:"image,omitempty" json:"image,omitempty"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"address_id,omitempty" json:"address_id,omitempty"`
	House      *string            `bson:"house,omitempty" json:"house,omit`
	Street     *string            `bson:"street,omitempty" json:"street,omitempty"`
	City       *string            `bson:"city,omitempty" json:"city,omit`
	Code       *string            `bson:"code,omitempty" json:"code,omit`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"order_id,omitempty" json:"order_id,omitempty"`
	Order_Cart     []ProductUser      `bson:"order_cart,omitempty" json:"order_cart,omitempty"`
	Order_At       time.Time          `bson:"order_at,omitempty" json:"order_at,omitempty"`
	Price          *uint64            `bson:"price,omitempty" json:"price,omitempty"`
	Discont        *uint64            `bson:"discont,omitempty" json:"discont,omitempty"`
	Payment_Method Payment            `bson:"payment_method,omitempty" json:"payment_method,omitempty"`
}

type Payment struct {
	Digital bool `bson:"digital,omitempty" json:"digital,omitempty"`
	COD     bool `bson:"cod,omitempty" json:"cod,omitempty"`
}
