package main

import (
	"api_test/config"
	"api_test/database"
	"api_test/logger"
	"api_test/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// setup fiber
	app := fiber.New(fiber.Config{
		BodyLimit: 500 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authrorization",
	}))

	// setup log
	logger.Initialize("develop", "cbs.log")
	logger.Log.Info("Initialized logger")

	// setup config and db
	config.Load_env()
	database.Initialize()

	// setup router
	router.Setup_router(app)

	// setup http
	err := app.Listen(config.Host)
	if err != nil {
		fmt.Println("fail on http")
	} else {
		fmt.Println("succeed on http")
	}

}

// ==========================================

// type event struct {
// 	ID          string `json:"ID"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// }

// type all_event []event

// var events = all_event{
// 	{
// 		ID:          "1",
// 		Title:       "Hello",
// 		Description: "hello hello",
// 	},
// }

// func create_event(w http.ResponseWriter, r *http.Request) {
// 	var new_event event
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Error of request body")
// 	}

// 	json.Unmarshal(reqBody, &new_event)
// 	events = append(events, new_event)
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(new_event)
// }

// func ServerHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world !!")
// }

// func main() {
// 	fmt.Println(events)
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", ServerHTTP)
// 	router.HandleFunc("/create_event", create_event).Methods("POST")
// 	log.Fatal(http.ListenAndServe(":8080", router))

// }
