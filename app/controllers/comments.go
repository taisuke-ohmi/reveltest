package controllers

import (
	"myapp/app/models"

	"github.com/revel/revel"
	"gopkg.in/validator.v2"
)

type ApiV1Comments struct {
	ApiV1Controller
}

func (c ApiV1Comments) Index() revel.Result {
	comments := []models.Comment{}

	if err := DB.Order("id desc").Find(&comments).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}

	r := Response{comments}
	return c.RenderJSON(r)
}

func (c ApiV1Comments) Show() revel.Result {
	comment := &models.Comment{}

	if err := DB.First(&comment, c.Params.Get("id")).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	r := Response{comment}
	return c.RenderJSON(r)
}

func (c ApiV1Comments) Create() revel.Result {
	comment := &models.Comment{}

	if err := c.BindParams(comment); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := validator.Validate(comment); err != nil {
		return c.HandleBadRequestError(err.Error())
	}

	if err := DB.Create(comment).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}

	r := Response{comment}
	return c.RenderJSON(r)
}

func (c ApiV1Comments) Delete() revel.Result {
	comment := &models.Comment{}

	if err := DB.First(&comment, c.Params.Get("id")).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if err := DB.Delete(&comment).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}

	r := Response{"success"}
	return c.RenderJSON(r)
}
