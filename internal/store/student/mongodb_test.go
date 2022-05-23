package student_test

import (
	"context"
	"testing"
	"time"

	"githuh.com/cng-by-example/students/internal/db"
	"githuh.com/cng-by-example/students/internal/model"
	"githuh.com/cng-by-example/students/internal/store/student"
)

func TestMongoDD(t *testing.T) {
	d, err := db.New(db.Config{
		URL:               "mongodb://127.0.0.1:27017",
		ConnectionTimeout: time.Second,
	})
	if err != nil {
		t.Fatalf("mongodb connection failed %s", err)
	}

	store := student.NewMongodb(d)

	if err := store.Set(model.Student{
		ID:        9231058,
		FirstName: "Parham",
		LastName:  "Alvani",
	}); err != nil {
		t.Fatalf("mongodb student creation failed %s", err)
	}

	std, err := store.Get(9231058)
	if err != nil {
		t.Fatalf("mongodb student retrieve failed %s", err)
	}

	if !(std.FirstName == "Parham" && std.LastName == "Alvani" && std.ID == 9231058) {
		t.Fatalf("mongodb invalid student")
	}

	stds, err := store.GetAll()
	if err != nil {
		t.Fatalf("mongodb students retrieve failed %s", err)
	}

	if !(len(stds) == 1 && stds[0].ID == 9231058) {
		t.Fatalf("mongodb invalid students")
	}

	if err := d.Collection(student.MongoDBCollection).Drop(context.Background()); err != nil {
		t.Fatalf("mongodb collection drop failed %s", err)
	}
}
