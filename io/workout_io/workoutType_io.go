package workout_io

import (
	"afcmo/api"
	"afcmo/domain/workout"
	"errors"
)

const workoutTypeURL = api.BASE_URL + "workout_type/"

func CreateWorkOutType(workoutObject workout.WorkoutType) (workout.WorkoutType, error) {
	entity := workout.WorkoutType{}
	resp, _ := api.Rest().SetBody(workoutObject).Post(workoutTypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadWorkOutType(id string) (workout.WorkoutType, error) {
	entity := workout.WorkoutType{}
	resp, _ := api.Rest().Get(workoutTypeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteWorkOutType(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(workoutTypeURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadWorkOutTypeAll() ([]workout.WorkoutType, error) {
	entity := []workout.WorkoutType{}
	resp, _ := api.Rest().Get(workoutTypeURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWorkOutTypeAllOf() ([]workout.WorkoutType, error) {
	entity := []workout.WorkoutType{}
	resp, _ := api.Rest().Get(workoutTypeURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
