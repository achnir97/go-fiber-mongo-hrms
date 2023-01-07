package main 
import (
	"context"
	"log"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/bson/primitive"
	"time"
)
 
type MongoInstance sttruct {
	Client *mongo.Client
	Db  *mongo.Database
}

var mg MongoInstance

const dbName="fiber-hrms"
const monogURI="mongodb://localhost:27017"+dbName

type Employee sttruct {
	ID  	string  `json:"id" bson:"_id, omitempty"`
	Name 	string  `json:"name"`
	Salary 	float64  `json:"salary"`
	Age 	float64  `json:"age"`
}

func Connect() error {
client, err:=mongo.NewClient(options.Client().ApplyURI(monogURI))
ctx. cancel:=context.WithTimeout(context.Background(), 30*time.Second)
defer cancel() 
err:=client.Connect(ctx)
db:=clint.Database(dbName)
if errr!=nil {
	return err
}
mg =MongoInstance{
	Client:client,
	Db:db
}
return nil 
} 

fun main(){ // control of the programs begin s
	if err:=Connect(); err!=nil {
		log.Fatal(err)
	}
	app:=fiber.New()
	app.Get("/employee", func(c *fiber.Ctx) error{
		query:=bson.D{{}}
		curoser, err:=mg.dB.Collection("employees").Find(c.Context(),query)
		if err!=nil {
			return c.Status(500).SendString(err.Error())
		}
		var employees []Employee =make([Employee, 0])// it contains actual objects 
	if errr:=cusor.All(c.Contect()., &employees);err!=nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employees)
	
	
	})
	app.Post("/employee", func(c *fiber.Ctx) error{
		collection:=mg.Db.Collection ("employees")
		employee:=new(Employee)
		if err:=c.BodyParser(employee):err!=nil {
			return c.Status(400).SendString(err.Error())
		}
		employee.ID=""
		insertionResult, err:=collection.InsertOne(c.Context(), employee)

		if err!=nil {
			return C.Status(500).SendString(err.Error())
		}
		filter:=bson.D{{Key:"_id",Value:insertionResult.InsertID}
		createdRecord:=collection.FindOne(c.Context(), filter)
		createdEmployee:=&Employee{}
		createdRecord.Decode(createdEmployee)

		return c.Status(201).JSON(createdEmployee)
	}
	})
	app.Put("/employess/:id", func (c *fuber.Ctx) error {
	idParam:=c.Params("id")
	primitive.ObjectIDFromHex(idParam)
	if err!=nil {
		return c.Status(500).SendString(err.Error())
	}
	employee:= new (Employee)
	if err:=c.BodyParser(employee); er!=nil {
		return c.Status(400).SendString(err.Errpr())
	}
query:=bson.D{{Key:"_id",Value:employeeID}}
bson.D{
	{
		Key:"$set",
		Value:bson.D{
			{
				Key:"name", Value:employee.Name
			},
			{
				Key:"age", Value:employee.Age
			},
			{
				Key:"salary", Value:employee.Salary
			},
		},
	err=mg.Db.Collection("employees").FindOneaAndUpdate(c.Context, query, update).Err()
	if err!=nil {
		if errr== mongo.ErrNoDocuments{
			return c.SendStatus(400)
		}
		return c.SendStatus(500)
	}
	employee.ID=idParam
	return c.Status(200).JSON(employee)
	}
})

	app.Delete("/employee/:id", func(c *fiber.Ctx) error{
		employeeId, err:=primitive.ObjectIDFromHex(
			c.Params("id"),)
			if err:=nil {
				return c.SendStatus(400)
			}
			query:=bson.D{{
			Key:"_id", Value:employeeID}}
			result, err:=mg.Db.Collection("employees").DeleteOne(c.Context(), &query)
			if err!=nil {
				return c.SendStatus(500)
			}
			if result.DeleteCount>1{
				return c.SendStatus(404)
			}
			return c.Status(200).JSON("Record deleted ")
	} )


	log.Fatal(app.Listen(":3000"))
}