package ServHistory

inport (
	"time"
	"github.com/tjaye2/serviceStation/customer.go"
	"github.com/tjaye2/serviceStation/service.go"
    "github.com/tjaye2/serviceStation/tech.go"
)

type ServiceTrans struct {
    Customer   customer.Customer
    Service    service.Service
    Technician technician.Technician
    Date       time.Time
}