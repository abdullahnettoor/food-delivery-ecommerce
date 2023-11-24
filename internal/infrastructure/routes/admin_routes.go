package routes

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(
	f *fiber.App, 
	admin *handlers.AdminHandler,
	coupon *handlers.CouponHandler,
	) {

	f.Post("/admin/login", admin.Login)

	r := f.Group("/admin", middlewares.AuthenticateAdmin)

	r.Get("/sellers", admin.GetAllSellers)
	r.Patch("/sellers/:id/verify", admin.VerifySeller)
	r.Patch("/sellers/:id/block", admin.BlockSeller)
	r.Patch("/sellers/:id/unblock", admin.UnblockSeller)

	r.Get("/users", admin.GetAllUsers)
	r.Patch("/users/:id/block", admin.BlockUser)
	r.Patch("/users/:id/unblock", admin.UnblockUser)

	r.Post("/categories/addCategory", admin.AddCategory)
	r.Patch("/categories/:id/edit", admin.EditCategory)

	r.Get("/coupons", coupon.GetAllCoupons)
	r.Post("/coupons/add", coupon.CreateCoupon)
	r.Patch("/coupons/:id", coupon.UpdateCouponStatus)

}
