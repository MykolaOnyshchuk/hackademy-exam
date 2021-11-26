CREATE DATABASE "ToDoApp";

CREATE TABLE "toDoTask" (
                            "id" SERIAL PRIMARY KEY NOT NULL,
                            "name" varchar(200),
                            "description" varchar(10000),
                            "status" varchar(100) NOT NULL,
                            "updatedAt" timestamp NOT NULL,
                            "createdAt" timestamp NOT NULL
);

CREATE TABLE "ToDoList" (
                            "id" SERIAL PRIMARY KEY NOT NULL,
                            "name" varchar(200),
                            "taskId" INT,
                            "updatedAt" timestamp NOT NULL,
                            "createdAt" timestamp NOT NULL
);

CREATE TABLE "User" (
                        "userID" SERIAL PRIMARY KEY NOT NULL,
                        "email" varchar(60) NOT NULL,
                        "imageLoc" varchar(1000),
                        "toDoListID" INT,
                        "updatedAt" timestamp NOT NULL,
                        "createdAt" timestamp NOT NULL,
                        CONSTRAINT "ToDoListID_fk0" FOREIGN KEY ("toDoListID") REFERENCES "ToDoList"("id") ON DELETE CASCADE
);
