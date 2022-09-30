package structs

import (
	"errors"
	"fmt"
)

type Booking interface {
	CreateTicket(t Ticket)
	ReadTickets() []Ticket
	ReadTicketById(idTicket int64) (Ticket, error)
	UpdateTicketById(idTicket int64, t Ticket) (Ticket, error)
	DeleteTicketById(idTicket int64) error
}

type bookings struct {
	CurrentTickets []Ticket
}

type Ticket struct {
	Id                             int64
	Name, Email, Destination, Hour string
	Price                          float64
}

func NewBooking(tickets []Ticket) Booking {
	return &bookings{CurrentTickets: tickets}
}

func (b *bookings) CreateTicket(t Ticket) {
	fmt.Println("Inicia creacion de ticket...")
	b.CurrentTickets = append(b.CurrentTickets, t)
	fmt.Printf("Se creo el ticket %v...\n", t.Id)
}

func (b *bookings) ReadTicketById(idTicket int64) (Ticket, error) {

	t := Ticket{}

	if idTicket <= 0 {
		return t, errors.New("el id del ticket proporcionado no es valido")
	}

	for _, ticket := range b.CurrentTickets {
		if ticket.Id == idTicket {
			t = ticket
		}
	}

	if t.Id == 0 {
		return t, errors.New("el id proporcionado no existe en el archivo de tickets")
	}

	return t, nil
}

func (b *bookings) ReadTickets() []Ticket {
	return b.CurrentTickets
}

func (b *bookings) UpdateTicketById(idTicket int64, t Ticket) (Ticket, error) {
	return t, nil
}

func (b *bookings) DeleteTicketById(idTicket int64) (err error) {

	fmt.Println("Inicia eliminacion de ticket...")

	if idTicket <= 0 {
		err = errors.New("el id del ticket proporcionado no es valido")
	}

	var index int
	for i, ticket := range b.CurrentTickets {
		if ticket.Id == idTicket {
			index = i
		}
	}

	b.CurrentTickets = append(b.CurrentTickets[:index], b.CurrentTickets[index+1:]...)

	fmt.Printf("Se elimino el ticket %v...\n", idTicket)

	return
}
