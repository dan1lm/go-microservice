package main

import (
	"context"
	"fmt"

	"github.com/dan1lm/go-microservice.git/application"
)

func main(){
	app:= application.New()
	err:= app.Start(context.TODO())
	if err!=nil{
		fmt.Println("failed to start app: ", err)
	}
}