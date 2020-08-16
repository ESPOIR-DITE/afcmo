package workout_io

import (
	"afcmo/api"
	"afcmo/domain/workout"
	"errors"
)

const workoutPlanURL = api.BASE_URL + "workout_plan/"

func CreateWorkOutPlan(workoutObject workout.WorkOutPlan) (workout.WorkOutPlan, error) {
	entity := workout.WorkOutPlan{}
	resp, _ := api.Rest().SetBody(workoutObject).Post(workoutPlanURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadWorkOutPlan(id string) (workout.WorkOutPlan, error) {
	entity := workout.WorkOutPlan{}
	resp, _ := api.Rest().Get(workoutPlanURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteWorkOutPlan(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(workoutPlanURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadWorkPlanOutAll() ([]workout.WorkOutPlan, error) {
	entity := []workout.WorkOutPlan{}
	resp, _ := api.Rest().Get(workoutPlanURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWorkPlanOutAllOf() ([]workout.WorkOutPlan, error) {
	entity := []workout.WorkOutPlan{}
	resp, _ := api.Rest().Get(workoutPlanURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
