package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/akwrs/lunch-delivery/types"
	"github.com/gorilla/websocket"
)

const wsEndPoint = "ws://127.0.0.1:30000/ws"

var timeInterval = time.Second * 3

func main() {
	fmt.Println("Rider en route!...")
	conn, _, err := websocket.DefaultDialer.Dial(wsEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	riderIDS := generateRiderIDS(20)

	for {
		for i := 0; i < len(riderIDS); i++ {
			lat, long := genLatLong()
			name := genName()
			data := types.Rider{
				RiderID: riderIDS[i],
				Name:    name,
				Lat:     lat,
				Long:    long,
			}

			fmt.Printf("\n%+v\n", data)
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Created Rider data [%d], <latitude: %.3f :: longitude %.3f \n>", data.RiderID, data.Lat, data.Long)
			time.Sleep(timeInterval)
		}
	}
}

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func genName() string {

	names := []string{
		"Ram", "Shyam", "Ramesh", "Suresh", "Mahesh",
		"Babu", "Amit", "Rahul", "Raj", "Vijay",
		"Arjun", "Kumar", "Deepak", "Sunil", "Anil",
		"Sanjay", "Prakash", "Mohan", "Gopal", "Vishal",
		"Ajay", "Nitin", "Pankaj", "Dinesh", "Harish",
		"Manish", "Alok", "Vinod", "Satish", "Rohit",
	}
	randomName := names[rand.Intn(len(names))]
	return randomName

}

func generateRiderIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
