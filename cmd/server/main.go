package main

import (
	"fmt"

	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"github.com/TutorialEdge/go-rest-api-course/internal/db"
	transportHttp "github.com/TutorialEdge/go-rest-api-course/internal/transport/http"
)

func Run() error {
	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }

	cmtService := comment.NewService(db)

	// cmtService.PostComment(
	// 	context.Background(),
	// 	comment.Comment{
	// 		ID:     "550e8400-e29b-41d4-a716-446655440000",
	// 		Slug:   "manuel-test",
	// 		Author: "Jhon Doe",
	// 		Body:   "Test",
	// 	},
	// )

	// fmt.Println(cmtService.GetComment(
	// 	context.Background(), "550e8400-e29b-41d4-a716-446655440000",
	// ))

	httpHandler := transportHttp.NewHandler(cmtService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("test")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
