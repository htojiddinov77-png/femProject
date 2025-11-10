package routes

import (

    "github.com/go-chi/chi/v5"
   
    "github.com/htojiddinov77-png/femProject/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
    r := chi.NewRouter()


    r.Get("/health", app.HealthCheck)
    r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)
    r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

    r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutById)
	r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkoutById)

	
    return r
}
