package controllers

import (
	"github.com/astaxie/beego"
	"home24/recommendation.api/models"
	"home24/recommendation/product"
)

type ProductRecommendationController struct {
	beego.Controller
}

func (productRecommendationController *ProductRecommendationController) GetRecommendations() {
	productId := productRecommendationController.Ctx.Input.Param(":id")

	recommendedProducts, err := product.GetRecommendations(productId)

	if err != nil {
		productRecommendationController.Ctx.Output.SetStatus(500)
		productRecommendationController.Ctx.Output.JSON(models.Error{Message: err.Error()}, false, false)
	} else {
		productRecommendationController.Ctx.Output.JSON(recommendedProducts, false, false)
	}
}
