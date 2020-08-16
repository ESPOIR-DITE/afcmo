package workout_io

import (
	"afcmo/api"
	"afcmo/domain/workout"
	"errors"
)

const workoutURL = api.BASE_URL + "workout/"

func CreateWorkOut(workoutObject workout.WorkOut) (workout.WorkOut, error) {
	entity := workout.WorkOut{}
	resp, _ := api.Rest().SetBody(workoutObject).Post(workoutURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadWorkOut(id string) (workout.WorkOut, error) {
	entity := workout.WorkOut{}
	resp, _ := api.Rest().Get(workoutURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteWorkOut(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(workoutURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadWorkOutAll() ([]workout.WorkOut, error) {
	entity := []workout.WorkOut{}
	resp, _ := api.Rest().Get(workoutURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWorkOutAllOf() ([]workout.WorkOut, error) {
	entity := []workout.WorkOut{}
	resp, _ := api.Rest().Get(workoutURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
