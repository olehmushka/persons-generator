package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	cont "persons_generator/core/context"
	hs "persons_generator/core/http_server"
	"persons_generator/core/http_server_tools"
	"persons_generator/core/wrapped_error"
	"strings"

	"github.com/google/uuid"
)

func (h *handlers) CreateWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := CreateWorldPersonsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	amount, maleAmount, femaleAmount, err := getAmounts(req.PersonsAmount, req.MalePersonsAmount, req.FemalePersonsAmount)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	religionCultureRels, err := parseReligionCultureRelations(req.CultureReligionIDsPairs)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	world, err := h.worldSrv.CreateWorld(ctx, amount, maleAmount, femaleAmount, religionCultureRels)
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}
	if world == nil {
		http_server_tools.SendErrorResp(ctx, w, wrapped_error.New(http.StatusInternalServerError, nil, "can not create world"))
		return
	}

	respJSON, err := json.Marshal(CreateWorldPersonsResponse{
		WorldID: world.ID,
	})
	if err != nil {
		http_server_tools.SendErrorResp(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(hs.TraceIDHeader, cont.GetTraceID(ctx))
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(respJSON)
}

func getAmounts(personsAmount *int, malePersonsAmount *int, femalePersonsAmount *int) (int, int, int, error) {
	if personsAmount == nil && malePersonsAmount == nil && femalePersonsAmount == nil {
		var (
			personsA       = 100
			malePersonsA   = personsA / 2
			femalePersonsA = personsA - malePersonsA
		)

		return personsA, malePersonsA, femalePersonsA, nil
	}
	if personsAmount != nil && malePersonsAmount == nil && femalePersonsAmount == nil {
		if *personsAmount < 1 {
			return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("persons_amount can not be less than 1 (persons_amount=%d)", *personsAmount))
		}
		if *personsAmount == 1 {
			return 1, 1, 0, nil
		}
		var (
			personsA       = *personsAmount
			malePersonsA   = personsA / 2
			femalePersonsA = personsA - malePersonsA
		)

		return personsA, malePersonsA, femalePersonsA, nil
	}
	if personsAmount != nil && malePersonsAmount != nil && femalePersonsAmount == nil {
		var (
			personsA     = *personsAmount
			malePersonsA = *malePersonsAmount
		)
		if personsA < malePersonsA {
			return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("persons_amount can not be less than male_persons_amount (persons_amount=%d, male_persons_amount=%d)", personsA, malePersonsA))
		}

		return personsA, malePersonsA, personsA - malePersonsA, nil
	}
	if personsAmount == nil && malePersonsAmount != nil && femalePersonsAmount == nil {
		if *malePersonsAmount < 1 {
			return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("male_persons_amount can not be less than 1 (male_persons_amount=%d)", *malePersonsAmount))
		}

		return *malePersonsAmount, *malePersonsAmount, 0, nil
	}
	if personsAmount == nil && malePersonsAmount != nil && femalePersonsAmount != nil {
		if *malePersonsAmount < 1 {
			return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("male_persons_amount can not be less than 1 (male_persons_amount=%d)", *malePersonsAmount))
		}
		if *femalePersonsAmount < 1 {
			return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("female_persons_amount can not be less than 1 (female_persons_amount=%d)", *femalePersonsAmount))
		}

		return *malePersonsAmount + *femalePersonsAmount, *malePersonsAmount, *femalePersonsAmount, nil
	}
	if personsAmount != nil && malePersonsAmount == nil && femalePersonsAmount != nil {
		var (
			personsA       = *personsAmount
			femalePersonsA = *femalePersonsAmount
		)
		if personsA < femalePersonsA {
			return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("persons_amount can not be less than female_persons_amount (persons_amount=%d, female_persons_amount=%d)", personsA, femalePersonsA))
		}

		return personsA, personsA - femalePersonsA, femalePersonsA, nil
	}
	if personsAmount == nil && malePersonsAmount == nil && femalePersonsAmount != nil {
		if *femalePersonsAmount < 1 {
			return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("female_persons_amount can not be less than 1 (female_persons_amount=%d)", *femalePersonsAmount))
		}

		return *femalePersonsAmount, 0, *femalePersonsAmount, nil
	}
	if personsAmount != nil && malePersonsAmount != nil && femalePersonsAmount != nil {
		var (
			personsA       = *personsAmount
			malePersonsA   = *malePersonsAmount
			femalePersonsA = *femalePersonsAmount
		)
		if malePersonsA+femalePersonsA != personsA {
			return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("sum of male_persons_amount and female_persons_amount must be equal persons_amount (persons_amount=%d, male=persons_amount=%d, female_persons_amount=%d)", personsA, malePersonsA, femalePersonsA))
		}

		return personsA, malePersonsA, femalePersonsA, nil
	}

	return 0, 0, 0, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("incorrect input (persons_amount=%+v, male_persons_amount=%+v, female_persons_amount=%+v)", personsAmount, malePersonsAmount, femalePersonsAmount))
}

func parseReligionCultureRelations(in []string) (map[uuid.UUID]uuid.UUID, error) {
	out := make(map[uuid.UUID]uuid.UUID, len(in))
	for _, rel := range in {
		relParts := strings.Split(rel, ":")
		if len(relParts) != 2 {
			return nil, wrapped_error.New(http.StatusBadRequest, nil, fmt.Sprintf("invalid religion_culture rel signature (rel=%s)", rel))
		}
		cultureID, err := uuid.Parse(relParts[0])
		if err != nil {
			return nil, wrapped_error.New(http.StatusBadRequest, err, fmt.Sprintf("can not parse culture_id (culture_id=%v)", rel[0]))
		}
		religionID, err := uuid.Parse(relParts[1])
		if err != nil {
			return nil, wrapped_error.New(http.StatusBadRequest, err, fmt.Sprintf("can not parse religion_id (religion_id=%v)", rel[1]))
		}
		out[religionID] = cultureID
	}

	return out, nil
}
