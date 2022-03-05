package Models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"mongo-sandbox/database"
	"mongo-sandbox/utilities"

	"go.mongodb.org/mongo-driver/bson"
)

func FindOne[T Model](filter bson.M) (err error, record *T) {
	// Cancel the UpdateOne query if it takes longer than 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Cancel the timeout when this function (Update) returns

	// 1. Determine the collection name based on the provided Model type,
	// 2. Find a record matching the provided filter,
	// 3. Convert the record into a struct
	err = database.Collection(collName[T]()).
		FindOne(ctx, filter).
		Decode(&record)

	return
}

/* func Save[T Model](record T) (error, bool) {
	if record.GetID().IsZero() {
		err, _ := Create(record)

		return err, err == nil
	}

	return Update(record)
} */

func Update[T Model](record T) (error, *T) {
	fmt.Println("Updating...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call some basic bitch hook shits
	callHooks(UpdateHook, record)

	// Finally do the actual query... after figuring out the collection name via a mirror (reflection)
	_, err := database.Collection(collName[T]()).UpdateOne(ctx, bson.M{"_id": record.GetID()}, bson.M{"$set": record})

	// fmt.Printf("Updated: %v\n%v\n\n", result, record)

	return err, &record
}

func Create[T Model](record T) (error, *T) {
	fmt.Println("Creating...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	callHooks(CreateHook, record)

	result, err := database.Collection(collName[T]()).InsertOne(ctx, record)

	// If the query didn't throw a hissy fit, update the ID on the Model with what the DB gave us
	if err == nil {
		setID(record, result.InsertedID)
	}

	// fmt.Printf("Created: %v\n%v\n\n", result, record)

	return err, &record
}

/*
Private Functions
*/

func callHooks(event HookEvent, m Model) {
	// Check if the Model implements the ICreating interface
	// and if the event "contains" the CreateHook enum
	if hook, ok := m.(ICreating); (event&CreateHook == CreateHook) && ok {
		hook.Creating()
	}

	// Check if the Model implements the IUpdating interface
	// and if the event "contains" the UpdateHook enum
	if hook, ok := m.(IUpdating); (event&UpdateHook == UpdateHook) && ok {
		hook.Updating()
	}
}

func collName[T Model]() (name string) {
	name = strings.ToLower(utilities.TypeOf[T]().Name() + "s")

	return
}
