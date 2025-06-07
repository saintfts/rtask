package structs

type Task struct {
    Id   uint   `json:"id"`   // ID implementation is a bit messy for now
    Text string `json:"text"`
}
