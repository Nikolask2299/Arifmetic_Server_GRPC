package sqlite

import (
	"context"
	"testing"
)


func TestDatabase(t *testing.T) {
	db, err := NewStorage("testbase/storage.db")
	if err != nil {
		t.Fatal(err)
	}
	
	ctx := context.Background()

	id, err := db.SaveUser(ctx, "Nikolas", []byte("eeesfsefsef"))
	if err != nil {
		t.Fatal(err)
	}

	if id == 0 {
		t.Fatal("wrong id")
	}

	user, err := db.GetUser(ctx, id, "")
	if err != nil {
		t.Fatal(err)
	}

	if user.Name != "Nikolas" {
		t.Fatal("wrong name")
	} 

	user, err = db.GetUser(ctx, 0, "Nikolas")
	if err != nil {
		t.Fatal(err)
	}

	if string(user.PassHash) != "eeesfsefsef" {
		t.Fatal("wrong name")
	}

	idtsk, err := db.SaveUserTask(ctx, 1, "2 + 2", "ANSWER")
	if err != nil {
		t.Fatal(err)
	}

	if idtsk == 0 {
		t.Fatal("wrong id")
	}

	task, err := db.GetTask(ctx, idtsk)
	if err != nil {
		t.Fatal(err)
	}

	if task.Task !="2 + 2" {
		t.Fatal("wrong name")
	}

	idans, err := db.SaveUserAnswer(ctx, idtsk, 4)
	if err != nil {
		t.Fatal(err)
	}

	if idans == 0 {
		t.Fatal("wrong id task")
	}

	answer, err := db.GetUserAnswer(ctx, idans, 0)
	if err != nil {
		t.Fatal(err)
	}

	if answer.Answer != 4 {
		t.Fatal("wrong answer")
	}

	answer, err = db.GetUserAnswer(ctx, 0, idtsk)
	if err != nil {
		t.Fatal(err)
	}

	if answer.Answer != 4 {
		t.Fatal("wrong answer")
	}
}
