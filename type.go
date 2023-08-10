package residence

type WhatsappRequest struct {
	Message     string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Phonenumber string `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Delay       uint32 `json:"delay,omitempty" bson:"delay,omitempty"`
}
