package utils

import "rest-api/internal/models"

func BuildMultipleTeacherResponse(status string, count int, data []models.Teacher) models.MultiTeacherResp {
	return models.MultiTeacherResp{
		Status: status,
		Count:  count,
		Data:   data,
	}
}
