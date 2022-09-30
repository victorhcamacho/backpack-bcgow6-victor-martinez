package structs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CsvFile struct {
	path string
}

func GetCsvFile(path string) *CsvFile {
	return &CsvFile{path: path}
}

func (f *CsvFile) ReadFile() (tickets []Ticket, err error) {

	fmt.Println("Inicia lectura de archivo...")

	data, errReading := os.ReadFile(f.path)

	if errReading != nil {
		err = errReading
		return
	}

	fmt.Println("Termina lectura de archivo...")

	tickets = getTickets(data)

	return
}

func (f *CsvFile) WriteFile(tickets []Ticket) (err error) {
	var lines string
	for _, ticket := range tickets {
		lines += fmt.Sprintf("%d,%v,%v,%v,%v,%.2f\n", ticket.Id, ticket.Name, ticket.Email, ticket.Destination, ticket.Hour, ticket.Price)
	}
	err = os.WriteFile(f.path, []byte(lines), 0644)
	return
}

func getTickets(data []byte) (tickets []Ticket) {

	ticket := &Ticket{}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line != "" {
			// fmt.Printf("%+v\n", line)
			pieces := strings.Split(line, ",")

			ticket.Id, _ = strconv.ParseInt(pieces[0], 10, 64)
			ticket.Name = pieces[1]
			ticket.Email = pieces[2]
			ticket.Destination = pieces[3]
			ticket.Hour = pieces[4]
			ticket.Price, _ = strconv.ParseFloat(pieces[5], 64)

			tickets = append(tickets, *ticket)
		}
	}

	return
}

/*func (f *CsvFile) GetTicketById(id int) (ticket Ticket, err error) {
	bytes, errReading := f.ReadFile()

	if errReading != nil {
		err = errReading
	}

	lines := strings.Fields(string(bytes))

	for _, line := range lines {
		pieces := strings.Split(line, ",")
		if idTicket := pieces[0]; idTicket == string(id) {
			ticket = Ticket{}
			return
		}
	}

	return
}*/
