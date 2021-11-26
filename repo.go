package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
)


type UserRegisterParams struct {
	Email string `json:"email"`
	Password string `json:"password"`
	FavoriteCake string `json:"favorite_cake"`
}

type InMemoryStorage struct {
	lock sync.RWMutex
	userStorage map[int]User
	toDoStorage map[int]ToDoList

}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		lock: sync.RWMutex{},
		userStorage: make(map[int]User),
		toDoStorage: make(map[int]ToDoList),
	}
}

func (rep *InMemoryStorage) NewUser(userEmail string, userPassword string) *User {
	return &User{
		id: len(rep.userStorage),
		email: userEmail,
		password: userPassword,
	}
}

type User struct {
	id int
	email string
	imageLocation string
	password string
	todos map[int]ToDoList
}

type ToDoList struct {
	id int // List unique identifier
	name string
	tasks map[int]Task // Map taskId -> Task
}

type Task struct {
	id int // Unique task id
	name string
	description string
	status string
}

func (repository *InMemoryStorage) AddNewUser(w http.ResponseWriter, usr User) {
	repository.lock.Lock()
	defer repository.lock.Unlock()
	for userFound := range repository.userStorage {
		if userFound.email == usr.email {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The user already exists"))
		}
	}

	repository.userStorage[len(repository.userStorage)] = usr
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Successfully created"))
}

func (repository *InMemoryStorage) AddNewToDoList(w http.ResponseWriter, usr User, list ToDoList) {
	repository.lock.Lock()
	defer repository.lock.Unlock()
	for toDoList := range usr.todos {
		if toDoList.name == list.name {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The list already exists"))
		}
	}

	repository.toDoStorage[list.id] = list
	repository.userStorage[usr.id].todos[list.id] = list
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Successfully created"))
}

func (repository *InMemoryStorage) AddNewTask(w http.ResponseWriter, usr User, listId int, task Task) {
	repository.lock.Lock()
	defer repository.lock.Unlock()
	if _, ok := usr.todos[listId]; ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The list does not exist"))
	}
	for taskFound := range usr.todos[listId].tasks {
		if taskFound.name == task.name {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Task already exists"))
		}
	}

	repository.toDoStorage[listId].tasks[task.id] = task
	repository.userStorage[usr.id].todos[listId].tasks[task.id] = task
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Successfully created"))
}

func (repository *InMemoryStorage) UpdateToDoList(w http.ResponseWriter, usr User, listId int, name string) {
	repository.lock.Lock()
	defer repository.lock.Unlock()

	repository.userStorage[usr.id].todos[listId].name = name
	repository.toDoStorage[listId].name = name
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully updated"))
}

func (repository *InMemoryStorage) UpdateTask(w http.ResponseWriter, usr User, listId int, taskId int, task Task) {
	repository.lock.Lock()
	defer repository.lock.Unlock()

	repository.userStorage[usr.id].todos[listId].tasks[taskId] = task
	repository.toDoStorage[listId].tasks[taskId] = task
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully updated"))
}

func (repository *InMemoryStorage) DeleteToDoList(w http.ResponseWriter, usr User, listId int) {
	repository.lock.Lock()
	defer repository.lock.Unlock()

	delete(repository.userStorage[usr.id].todos, listId)
	delete(repository.toDoStorage, listId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted"))
}

func (repository *InMemoryStorage) DeleteTask(w http.ResponseWriter, usr User, listId int, taskId int) {
	repository.lock.Lock()
	defer repository.lock.Unlock()

	delete(repository.userStorage[usr.id].todos[listId].tasks, taskId)
	delete(repository.toDoStorage[listId].tasks, taskId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted"))
}

func (repository *InMemoryStorage) GetToDoLists(w http.ResponseWriter, usr User) {

	jsonMap, _ := json.Marshal(usr.todos)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonMap)
}

func (repository *InMemoryStorage) GetTasks(w http.ResponseWriter, usr User, listId int) {

	jsonMap, _ := json.Marshal(repository.userStorage[usr.id].todos[listId].tasks)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonMap)
}

func (repository *InMemoryStorage) Get(email string) (User, error) {
	repository.lock.Lock()
	ind := 0
	defer repository.lock.Unlock()
	var returnValue User
	for userFound := range repository.userStorage {
		if userFound.email == email {
			ind = userFound.id
		}
	}

	if _, ok := repository.userStorage[ind]; !ok {
		return returnValue, errors.New("The user doesn't exist")
	}
	returnValue = repository.userStorage[ind]
	return returnValue, nil

}