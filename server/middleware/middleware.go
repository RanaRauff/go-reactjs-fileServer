package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "github.com/shijuvar/todo_api/views"
)

var Collection *mongo.Collection

func init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connectingString, err := ioutil.ReadFile("config.bin")
	if err != nil {
		fmt.Println("Error:", err)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(string(connectingString)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	Collection = client.Database("db1").Collection("coll1")
	fmt.Println(Collection.Name())
	fmt.Println(Collection.Database())

	fmt.Println("DataBase Created")
	// CreateData()
	// ReadData()

}

type Dd struct {
	Name string
	Size int
}

func CreateData(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)
	r.ParseMultipartForm(10 << 20)
	// fmt.Println(r.FormFile("files"))
	file, handler, err := r.FormFile("files")
	fmt.Println(handler.Filename)

	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	temp, err := ioutil.TempFile("static", "*"+handler.Filename)
	filename := temp.Name()
	if err != nil {
		panic(fmt.Sprintf("Error:PANIC CALLED", err))
	}
	defer temp.Close()
	fp, err := ioutil.ReadAll(file)
	temp.Write(fp)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if err = createDataFunc(strings.Split(filename, "/")[1], int(handler.Size)); err != nil {
		fmt.Println("Error in createDataFunc Call", err)
	} else {
		fmt.Println("File Uploaded..")
	}

}

func ReadData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reading Data...")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := readDataFunc()
	if err != nil {
		fmt.Println("ERROR in ReadData", err)
	}
	json.NewEncoder(w).Encode(data)

}

func Download(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We go here")
	type name struct {
		Data string `json:"data"`
	}
	var na name
	err := json.NewDecoder(r.Body).Decode(&na)
	if err != nil {
		fmt.Println("Error in Download", err)
	}
	fmt.Println(na)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+na.Data)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	// resp, err := os.Open("./static/" + na.Data)
	resp, err := ioutil.ReadFile("./static/" + na.Data)
	if err != nil {
		fmt.Println("Error in loading file...")
	}
	// w.Write(resp.)
	// defer resp.Close()
	// _, err = io.Copy(w, resp)
	fmt.Println(resp)
	http.ServeContent(w, r, "./static/"+na.Data, time.Now(), bytes.NewReader(resp))
	fmt.Println("Error...", err)
	fmt.Println("Uploaded")
}

func createDataFunc(name string, size int) error {
	var data = Dd{
		name,
		size,
	}

	fmt.Println("The Struct", data)

	crr, err := Collection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println("Error in createData", err)
		return err
	}
	fmt.Println("Successfully Created One Entry", crr)
	return nil

}

func readDataFunc() ([]primitive.M, error) {
	curr, err := Collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("Error in ReadData", err)
	}
	fmt.Println("The Data is :", curr)

	if err != nil {
		fmt.Println("The Error in Decoding...", err)
		return nil, err
	}
	// fmt.Println("The M Data is :", data)
	var data_list []primitive.M

	for curr.Next(context.Background()) {
		var data bson.M
		err = curr.Decode(&data)
		if err != nil {
			fmt.Println("ERror ", err)
		}
		fmt.Println("Data:", data)
		data_list = append(data_list, data)

	}

	return data_list, nil

}
