package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/kitalabs/go-2gaijin/channels"
	"gitlab.com/kitalabs/go-2gaijin/middleware"
)

// Router is exported and used in main.go
func Router() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", middleware.GetHome)
	r.GET("/products/:id", middleware.GetProductDetail)
	r.GET("/wishlist", middleware.GetWishlistPage)
	r.POST("/add_product", middleware.PostNewProduct)
	r.POST("/mark_as_sold", middleware.MarkAsSold)
	r.POST("/delete_product", middleware.DeleteProduct)
	r.POST("/edit_product", middleware.EditProduct)
	r.POST("/like_product", middleware.LikeProduct)
	r.GET("/get_categories", middleware.GetAllCategories)

	r.POST("/sign_in", middleware.LoginHandler)
	r.POST("/sign_up", middleware.RegisterHandler)
	r.POST("/sign_out", middleware.LogoutHandler)
	r.POST("/refresh_token", middleware.RefreshToken)
	r.POST("/reset_password", middleware.ResetPasswordHandler)
	r.POST("/update_password", middleware.UpdatePasswordHandler)
	r.POST("/profile", middleware.ProfileHandler)
	r.POST("/update_profile", middleware.UpdateProfile)
	r.POST("/confirm_identity", middleware.GenerateConfirmToken)
	r.GET("/confirm_email", middleware.EmailConfirmation)
	r.GET("/confirm_phone", middleware.PhoneConfirmation)

	r.GET("/profile_visitor", middleware.GetProfileForVisitorPage)

	r.POST("/chat_lobby", middleware.GetChatLobby)
	r.GET("/chat_messages", middleware.GetChatRoomMsg)
	r.GET("/initiate_chat", middleware.ChatUser)
	r.POST("/insert_message", middleware.InsertMessage)
	r.GET("/ws", channels.ServeChat)

	r.GET("/search", middleware.GetSearch)

	r.POST("/insert_notification", middleware.InsertNotification)
	r.POST("/insert_appointment", middleware.InsertAppointment)
	r.POST("/insert_trust_coin", middleware.InsertTrustCoin)
	r.POST("/confirm_appointment", middleware.AppointmentConfirmation)
	r.POST("/reschedule_appointment", middleware.RescheduleAppointment)
	r.POST("/finish_appointment", middleware.FinishAppointment)

	r.GET("/get_seller_appointments", middleware.GetSellerAppointmentPage)
	r.GET("/get_buyer_appointments", middleware.GetBuyerAppointmentPage)
	r.GET("/get_notifications", middleware.GetNotificationPage)

	return r
}
