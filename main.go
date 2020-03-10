package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shijuvar/ter/server/middleware"
	// "github.com/shijuvar/todo_api/views"
)

// var collection *mongo.Collection

// func init() {

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
// 		"mongodb+srv://admin:admin@fileshare-z4phj.mongodb.net/test?retryWrites=true&w=majority",
// 	))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connected to MongoDB")
// 	collection = client.Database("db1").Collection("coll1")
// 	fmt.Println(collection.Name())
// 	fmt.Println(collection.Database())

// 	fmt.Println("DataBase Creadted")
// 	// CreateData()
// 	ReadData()

// }

// type Dd struct {
// 	Name string
// 	Age  int
// }

// func CreateData() {
// 	var d = Dd{
// 		Name: "Rahul",
// 		Age:  21,
// 	}

// 	fmt.Println("The Struct", d)
// 	collection.InsertOne(context.Background(), d)
// }

// func ReadData() {
// 	curr, err := collection.Find(context.Background(), bson.D{})
// 	if err != nil {
// 		fmt.Println("Error in ReadData", err)
// 	}
// 	fmt.Println("The Data is :", curr)
// 	// json.NewDecoder(curr).Decode()
// 	// var list []interface
// 	// var data bson.M
// 	// err = curr.Decode(&data)
// 	if err != nil {
// 		fmt.Println("The Error in Decoding...", err)
// 	}
// 	// fmt.Println("The M Data is :", data)

// 	for curr.Next(context.Background()) {
// 		var data bson.M
// 		err = curr.Decode(&data)
// 		if err != nil {
// 			fmt.Println("ERror ", err)
// 		}
// 		fmt.Println("Data:", data)

// 	}
// }

func main() {
	fmt.Println("hello")
	mux := mux.NewRouter()
	// var data views.DataBase
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// mux.Handle("/", http.HandlerFunc(data.UplData)).Methods("GET")
	// mux.Handle("/", http.HandlerFunc(data.UplData)).Methods("POST")
	mux.Handle("/", http.HandlerFunc(middleware.CreateData)).Methods("POST")
	mux.Handle("/", http.HandlerFunc(middleware.ReadData)).Methods("GET")
	mux.Handle("/dwnld", http.HandlerFunc(middleware.Download)).Methods("POST")
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/statics/{id}", http.StripPrefix("/statics/", fs))
	log.Println("Logging........")

	http.ListenAndServe(":8081", handlers.CORS(header, methods, origins)(mux))

}
