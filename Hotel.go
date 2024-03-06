package main

import "fmt"

type Room struct {
	ID       int
	Type     string
	Status   bool
	BedCount int
	Price    int
}

// slice of structs
var Rooms []Room = GenerateRoom()

// for / switch / funcs
func main() {

	input := ""
	for input != "exit" {
		fmt.Println("Enter a Command")
		fmt.Println("1: Room List ")
		fmt.Println("2: Add Room")
		fmt.Println("3: Reserve Room")
		fmt.Scanln(&input)
		switch input {
		case "1":
			GetRoomList()
		case "2":
			AddRoom()
		case "3":
			ReserveRoom()
		case "exit":
			fmt.Println("Exiting...")
			break
		default:
			fmt.Println("Invalid Command")
		}
	}

}

// 1--Get Room List
// implement with loop (+v) : show fields
func GetRoomList() {
	for _, room := range Rooms {
		fmt.Printf("%+v\n", room)

	}
}

// 2--AddRoom (struct)
func GetRoomFromInput() Room {
	//struct Room
	var room Room = Room{Status: false}
	fmt.Println("Enter the room information line by line(ID , Type , BedCount, Price)")
	fmt.Scanln(&room.ID)
	fmt.Scanln(&room.Type)
	fmt.Scanln(&room.BedCount)
	fmt.Scanln(&room.Price)

	return room

}

// 2 --Add Room
// Slice & Struct
func AddRoom() {
	room := GetRoomFromInput()
	Rooms = append(Rooms, room)

}

// 3--Reserve Room(if , CalculateRoomPrice() , GetRoom())
func ReserveRoom() {
	id := 0
	nights := 0
	personCount := 0
	fmt.Println("Please Enter  room Id for reservation")
	fmt.Scanln(&id)

	room := GetRoom(id)

	if room == nil {
		fmt.Println("Room Not Found")
		return
	}
	if room.Status {
		fmt.Println("Room is already reserved")
		return
	}
	fmt.Println("Please enter reserve information line by line(nights , personCount)")
	fmt.Scanln(&nights)
	fmt.Scanln(&personCount)
	roomPrice, tax, discountAmount, finalPrice := room.CalculateRoomPrice(nights, personCount)
	room.Status = true

	fmt.Printf("Room Price : %f , Tax : %f , Discount: %f , FinalPrice: %f\n", roomPrice, tax, discountAmount, finalPrice)

}

// 3--Reserve Room(if , switch)
func (room *Room) CalculateRoomPrice(nights int, personCount int) (roomPrice float64, tax float64, discountAmount float64, finalPrice float64) {

	discountPercentage := 0.0
	if nights >= 7 && nights <= 15 {
		discountPercentage = 0.1
	} else if nights > 15 && nights <= 30 {
		discountPercentage = 0.15
	} else if nights > 30 {
		discountPercentage = 0.2
	}
	switch room.Type {
	case "Single":
		roomPrice = float64(nights*room.Price*personCount) * 1.0

	case "Double":
		roomPrice = float64(nights*room.Price*personCount) * 1.2

	case "Standard":
		roomPrice = float64(nights*room.Price*personCount) * 1.3

	case "Suite":
		roomPrice = float64(nights*room.Price*personCount) * 1.5

	}

	tax = roomPrice * 0.09
	discountAmount = roomPrice * discountPercentage
	finalPrice = roomPrice + tax - discountAmount

	return
}

// 3-Reserve Room(slice struct for)
func GetRoom(id int) *Room {
	for i := 0; i < len(Rooms); i++ {
		if Rooms[i].ID == id {
			return &Rooms[i]
		}

	}
	return nil
}

// Base Generate room
// slice & Append struct

func GenerateRoom() []Room {

	rooms := []Room{}

	rooms = append(rooms, Room{ID: 1, Type: "Single", Status: false, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 2, Type: "Single", Status: false, BedCount: 1, Price: 120})
	rooms = append(rooms, Room{ID: 3, Type: "Single", Status: false, BedCount: 1, Price: 150})
	rooms = append(rooms, Room{ID: 4, Type: "Double", Status: false, BedCount: 2, Price: 200})
	rooms = append(rooms, Room{ID: 5, Type: "Double", Status: false, BedCount: 2, Price: 220})
	rooms = append(rooms, Room{ID: 6, Type: "Double", Status: false, BedCount: 2, Price: 250})
	rooms = append(rooms, Room{ID: 7, Type: "Double", Status: false, BedCount: 3, Price: 230})
	rooms = append(rooms, Room{ID: 8, Type: "Double", Status: false, BedCount: 3, Price: 250})
	rooms = append(rooms, Room{ID: 9, Type: "Double", Status: false, BedCount: 3, Price: 280})
	rooms = append(rooms, Room{ID: 10, Type: "Standard", Status: false, BedCount: 4, Price: 300})
	rooms = append(rooms, Room{ID: 11, Type: "Standard", Status: false, BedCount: 4, Price: 320})
	rooms = append(rooms, Room{ID: 12, Type: "Standard", Status: false, BedCount: 4, Price: 360})

	return rooms

}
