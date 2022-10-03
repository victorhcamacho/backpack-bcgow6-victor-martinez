package main

import (
	"fmt"

	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/HackatonGoBases/structs"
)

const FILE_PATH string = "./tickets.csv"

func main() {

	csv := structs.GetCsvFile(FILE_PATH)

	tickets, err1 := csv.ReadFile()

	if err1 != nil {
		panic(err1)
	}

	ticketsRows := len(tickets)

	fmt.Printf("Total de tickets encontrado en el archivo: %v\n", ticketsRows)

	booking := structs.NewBooking(tickets)

	// createTicket(booking, int64(ticketsRows+1))

	// select by id
	/*ticket, errS := selectTicket(booking, int64(50))

	if errS != nil {
		panic(errS)
	}

	fmt.Printf("Ticket seleccionado: %+v\n", ticket)*/

	// delete by id
	/*errD := deleteTicket(booking, int64(ticketsRows))

	if errD != nil {
		panic(errD)
	}*/

	fmt.Println("Inicia escritura de archivo...")

	err2 := csv.WriteFile(booking.ReadTickets())

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Termina escritura de archivo...")

}

func createTicket(b structs.Booking, id int64) {

	ticket := structs.Ticket{
		Id:          id,
		Name:        "Víctor Hugo",
		Email:       "victor.martinez@domain.com",
		Destination: "Canada",
		Hour:        "20:45 hrs",
		Price:       1449.9,
	}

	b.CreateTicket(ticket)

	fmt.Printf("Nuevo ticket: %+v\n", ticket)
}

func deleteTicket(b structs.Booking, id int64) (err error) {
	err = b.DeleteTicketById(id)
	return
}

func selectTicket(b structs.Booking, id int64) (result structs.Ticket, err error) {
	result, err = b.ReadTicketById(id)
	return
}
