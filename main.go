package main

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"
	handleCart "food_delivery/module/cart/transport"
	handleCategory "food_delivery/module/category/transport"
	handleFood "food_delivery/module/food/transport"
	handleLikeFood "food_delivery/module/foodlike/transport"
	handleOrder "food_delivery/module/order/transport"
	handleRestaurant "food_delivery/module/restaurant/transport"
	handleLikeRestaurant "food_delivery/module/restaurantlike/transport"
	handleUser "food_delivery/module/user/transport"
	"food_delivery/pubsub/pblocal"
	"food_delivery/subscriber"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	dsn := os.Getenv("CONNECT_DB_MYSQL")
	secretKey := os.Getenv("SECRET_KEY")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	if err != nil {
		log.Fatal(err)
	}

	// run pubsub
	appCtx := appctx.NewAppCtx(db, secretKey, pblocal.NewLocalPubSub())

	//subscriber.Setup(appCtx)
	if err := subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	v1 := router.Group("/v1")

	v1.POST("/register", handleUser.HandleRegister(appCtx))
	v1.POST("/login", handleUser.HandleLogin(appCtx))
	user := v1.Group("/user", middleware.RequiredAuth(appCtx))
	{
		user.PUT("/edit", handleUser.HandleUpdateInfoUser(appCtx))
		user.GET("/profile", handleUser.HandleGetProfile(appCtx))
	}

	restaurant := v1.Group("/restaurant", middleware.RequiredAuth(appCtx))
	{
		restaurant.POST("", handleRestaurant.HandleCreateRestaurant(appCtx))
		restaurant.GET("/:id", handleRestaurant.HandleGetRestaurant(appCtx))
		restaurant.GET("", handleRestaurant.HandleGetListRestaurants(appCtx))
		restaurant.PUT("/:id", handleRestaurant.HandleUpdateRestaurant(appCtx))
		restaurant.DELETE("/:id", handleRestaurant.HandleDeleteRestaurant(appCtx))

		restaurant.POST("/:id/like", handleLikeRestaurant.HandleUserLikeRestaurant(appCtx))
		restaurant.DELETE("/:id/dislike", handleLikeRestaurant.HandleUserDislikeRestaurant(appCtx))
		restaurant.GET("/:id/liked-users", handleLikeRestaurant.HandleListUserLikeRestaurant(appCtx))
	}

	food := v1.Group("/food", middleware.RequiredAuth(appCtx))
	{
		food.POST("", handleFood.HandleCreateFood(appCtx))
		food.GET("/:id", handleFood.HandleGetFood(appCtx))
		food.GET("", handleFood.HandleGetListFoods(appCtx))
		food.PUT("/:id", handleFood.HandleUpdateFood(appCtx))
		food.DELETE("/:id", handleFood.HandleDeleteFood(appCtx))

		food.POST("/:id/like", handleLikeFood.HandleUserLikeFood(appCtx))
		food.DELETE("/:id/dislike", handleLikeFood.HandleUserDislikeFood(appCtx))
		food.GET("/:id/liked-users", handleLikeFood.HandleListUserLikeFood(appCtx))
	}

	category := v1.Group("/category", middleware.RequiredAuth(appCtx))
	{
		category.POST("", handleCategory.HandleCreateCategory(appCtx))
		category.GET("/:id", handleCategory.HandleGetCategory(appCtx))
		category.GET("", handleCategory.HandleGetListCategories(appCtx))
		category.PUT("/:id", handleCategory.HandleUpdateCategory(appCtx))
		category.DELETE("/:id", handleCategory.HandleDeleteCategory(appCtx))

	}

	cart := v1.Group("/cart", middleware.RequiredAuth(appCtx))
	{
		cart.POST("", handleCart.HandleAddFoodToCart(appCtx))
		cart.DELETE("/:id", handleCart.HandleRemoveFoodFromCart(appCtx))
		cart.GET("", handleCart.HandleGetListFoodsFromCart(appCtx))
	}

	order := v1.Group("/order", middleware.RequiredAuth(appCtx))
	{
		order.POST("", handleOrder.HandleCreateOrder(appCtx))
	}

	router.Run(":5050")
}
