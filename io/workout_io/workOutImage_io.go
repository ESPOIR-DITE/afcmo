package workout_io

import (
	"afcmo/api"
	"afcmo/domain/workout"
	"errors"
)

const workoutImageURL = api.BASE_URL + "workout_video/"

func CreateWorkOutImage(workoutObject workout.WorkOutImage) (workout.WorkOutImage, error) {
	entity := workout.WorkOutImage{}
	resp, _ := api.Rest().SetBody(workoutObject).Post(workoutImageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWorkOutImage(id string) (workout.WorkOutImage, error) {
	entity := workout.WorkOutImage{}
	resp, _ := api.Rest().Get(workoutImageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteWorkOutImage(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(workoutImageURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadWorkOutImageAll() ([]workout.WorkOutImage, error) {
	entity := []workout.WorkOutImage{}
	resp, _ := api.Rest().Get(workoutImageURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWorkOutImageAllOf() ([]workout.WorkOutImage, error) {
	entity := []workout.WorkOutImage{}
	resp, _ := api.Rest().Get(workoutImageURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
