package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nikhilkuyya/invoice-go-app/internal/store"
)

type TeamHandler struct {
	teamStore store.TeamStore
}

func NewTeamHandler(teamStore store.TeamStore) *TeamHandler {
	return &TeamHandler{
		teamStore: teamStore,
	}
}


func (handler *TeamHandler) HandleGetTeams(w http.ResponseWriter,r *http.Request) {
	teams, err := handler.teamStore.GetTeams()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to get teams", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(teams)
}

func (handler *TeamHandler) HandleCreateTeam(w http.ResponseWriter,r *http.Request) {
	var teamPayload store.Team
	err := json.NewDecoder(r.Body).Decode(&teamPayload)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create team",http.StatusBadRequest)
		return
	}

	createdTeam, err := handler.teamStore.CreateTeam(&teamPayload)

	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create team",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(createdTeam)
}

func (handler *TeamHandler) HandleGetTeamByID(w http.ResponseWriter, r *http.Request) {
	paramId := chi.URLParam(r,"id")

	if paramId == ""{
		http.NotFound(w,r)
		return;
	}

	teamId, err := strconv.ParseInt(paramId,10,64)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	team, err := handler.teamStore.GetTeamByID(teamId)

	if err != nil {
		fmt.Println(err)
		http.Error(w,"couldn't get the team", http.StatusInternalServerError)
		return ;
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}
