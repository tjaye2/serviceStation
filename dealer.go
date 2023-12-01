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
    "github.com/tjaye2/serviceStation/service.go"
    "github.com/tjaye2/service-station/servicehistory.go"
    "github.com/tjaye2/service-station/tech.go"
)

type Dealer struct {
    Name      string
    Branches  []string
    Customers []customer.Customer
    Services  []services.Service
    Techs     []technician.Technician
    History   []serviceHistory.ServiceHistory
    Coupons   []couponHistory.CouponHistory
}

func NewDealer(name string, branches []string) *Dealer {
    return &Dealer{
        Name:     name,
        Branches: branches,
        Customers: []customer.Customer{
            {custId: "J1", custName: "George Johnson", custAddress: "5555 Elm St", custPhone: "940-555-5555"},
            {custId: "D50", custName: "Pam Downy, custAddress: "3285 Swisher St", custPhone: "940-555-9999"},
        },
        Services: []service.Service{
            {Name: "Oil Change", servType: "Synthetic oil change"},
            {Name: "Car Wash", servType: "Deluxe car wash"},
        },
        Techs: []technician.Technician{
            {techId: "T1001", techName: "Mike Smith", techPhone: "555-555-5555", Service: service.Service{custName: "Pam Downy", custAddress: "3285 Swisher", custPhone: "940-555-9999"},
        },
}
