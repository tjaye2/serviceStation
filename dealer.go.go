package dealer

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/tjaye2/serviceStation/couponHistory.go"
    "github.com/tjaye2/serviceStation/customer.go"
    "github.com/your-username/serviceStation/service.go"
    "github.com/your-username/service-station/servicehistory.go"
    "github.com/your-username/service-station/tech.go"
)

type Dealer struct {
    Name      string
    Branches  []string
    Customers []customer.Customer
    Services  []services.Service
    Techs     []technician.Technician
    History   []servicehistory.ServiceTransaction
    Coupons   []couponhistory.CouponNotification
}

func NewDealer(name string, branches []string) *Dealer {
    return &Dealer{
        Name:     name,
        Branches: branches,
        Customers: []customer.Customer{
            {ID: "C1001", Name: "John Doe", Address: "123 Main St, Anytown, USA", Phone: "555-555-1234"},
            {ID: "C1002", Name: "Jane Doe", Address: "456 Elm St, Anytown, USA", Phone: "555-555-5678"},
        },
        Services: []service.Service{
            {Name: "Oil Change", Description: "Basic oil change service"},
            {Name: "Car Wash", Description: "Basic car wash service"},
        },
        Techs: []technician.Technician{
            {ID: "T1001", Name: "Mike Smith", Phone: "555-555-5555", Service: service.Service{Name: "Jane Doe", Address: "456 Elm St, Anytown, USA", Phone: "555-555-5678"},
        },
}
