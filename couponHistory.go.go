package CouponHistory

import "time"

type CouponNotify struct {
    custID   string
    servType  string
    Notification time.Time
    Expiration   time.Time
}
