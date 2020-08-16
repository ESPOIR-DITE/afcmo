package workout_io

import (
	"afcmo/api"
	"afcmo/domain/workout"
	"errors"

)

const workoutVideoURL = api.BASE_URL + "workout/"

func CreateWorkOutVideo(workoutObject workout.WorkVideo) (workout.WorkVideo, error) {
	entity := workout.WorkVideo{}
	resp, _ := api.Rest().SetBody(workoutObject).Post(workoutVideoURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateWorkOutVideo(workoutObject workout.WorkVideo) (workout.WorkVideo, error) {
	entity := workout.WorkVideo{}
	resp, _ := api.Rest().SetBody(workoutObject).Post(workoutVideoURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadWorkOutVideo(id string) (workout.WorkVideo, error) {
	entity := workout.WorkVideo{}
	resp, _ := api.Rest().Get(workoutVideoURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteWorkOutVideo(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(workoutVideoURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadWorkOutVideoAll() ([]workout.WorkVideo, error) {
	entity := []workout.WorkVideo{}
	resp, _ := api.Rest().Get(workoutVideoURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWorkOutVideoAllOf() ([]workout.WorkVideo, error) {
	entity := []workout.WorkVideo{}
	resp, _ := api.Rest().Get(workoutVideoURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
