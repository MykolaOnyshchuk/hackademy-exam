package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := mux.NewRouter()




	srv := http.Server{
		Addr: ":8080",
		Handler: router,
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(),
			5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()
	log.Println("Server started, hit Ctrl+C to stop")
	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Server exited with error:", err)
	}



	jwtService, err := NewJWTService("pubkey.rsa", "privkey.rsa")
	if err != nil {
		panic(err)
	}

	ims := NewInMemoryStorage()

	router.HandleFunc("user/signup", logRequest(ims.userStorage[0].Register)).
		Methods(http.MethodPost)
	router.HandleFunc("/user/signin", logRequest(jwtService.jwtAuth())).
		Methods(http.MethodPost)
	router.HandleFunc("/todo/lists", logRequest(jwtService.jwtAuth(ims.AddNewToDoList))).
		Methods(http.MethodPost)
	router.HandleFunc("/todo/lists/*list_id*", logRequest(jwtService.jwtAuth(ims.UpdateToDoList))).
		Methods(http.MethodPut)
	router.HandleFunc("/todo/lists/*list_id*", logRequest(jwtService.jwtAuth(DeleteToDoList))).
		Methods(http.MethodDelete)
	router.HandleFunc("/todo/lists", logRequest(jwtService.jwtAuth(GetToDoLists))).
		Methods(http.MethodGet)
	router.HandleFunc("/todo/lists/*list_id*/tasks", logRequest(jwtService.jwtAuth(AddNewTask))).
		Methods(http.MethodPost)
	router.HandleFunc("/todo/lists/*list_id*/tasks/*task_id*", logRequest(jwtService.jwtAuth(UpdateTask))).
		Methods(http.MethodPut)
	router.HandleFunc("/todo/lists/*list_id*/tasks/*task_id*", logRequest(jwtService.jwtAuth(DeleteTask))).
		Methods(http.MethodDelete)
	router.HandleFunc("/todo/lists/*list_id*/tasks", logRequest(jwtService.jwtAuth(GetTasks))).
		Methods(http.MethodGet)
}


func wrapJwt(
	jwt *JWTService,
	f func(http.ResponseWriter, *http.Request, *JWTService),
) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		f(rw, r, jwt)
	}
}
