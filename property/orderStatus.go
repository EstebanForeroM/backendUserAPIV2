package property

import "fmt"

type Status string

const (
    Pending Status = "Pending"
    InDelivery Status = "In delivery"
    Delivered Status = "Delivered"
)

func StringToStatus(s string) (Status, error) {
    switch s {
    case string(Pending):
        return Pending, nil
    case string(InDelivery):
        return InDelivery, nil
    case string(Delivered):
        return Delivered, nil
    default:
        return "", fmt.Errorf("Invalid status: %s", s)
    }
}

func (Status) Values() (kinds []string) {
    for _, s := range []Status{Pending, InDelivery, Delivered} {
        kinds = append(kinds, string(s))
    }
    return
}

