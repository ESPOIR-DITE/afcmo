package subscription

import (
	"afcmo/api"
	"afcmo/domain/user"
	"errors"
)

const workoutSubscriptionURL = api.BASE_URL + "WorkOutSubscription/ "

func CreateWorkOutSubscription(users user.WorkoutSubscription) (user.WorkoutSubscription, error) {
	entity := user.WorkoutSubscription{}
	resp, _ := api.Rest().SetBody(users).Post(workoutSubscriptionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadWorkOutSubscription(id string) (user.WorkoutSubscription, error) {
	entity := user.WorkoutSubscription{}
	resp, _ := api.Rest().Get(workoutSubscriptionURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteWorkOutSubscription(id string) (bool, error) {
	var entity bool
	resp, _ := api.Rest().Get(workoutSubscriptionURL + "delete?id=" + id)
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func ReadWorkOutSubscriptionAll() ([]user.WorkoutSubscription, error) {
	entity := []user.WorkoutSubscription{}
	resp, _ := api.Rest().Get(workoutSubscriptionURL + "readAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWorkOutSubscriptionAllOf() ([]user.WorkoutSubscription, error) {
	entity := []user.WorkoutSubscription{}
	resp, _ := api.Rest().Get(workoutSubscriptionURL + "readAllOf")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
