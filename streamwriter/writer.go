package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

type writer struct {
	emitter *goka.Emitter
}

func NewWriter(topic string, broker []string) (*writer, error) {
	emitter, err := goka.NewEmitter(broker, goka.Stream(topic), &codec.Bytes{})
	if err != nil {
		return nil, err
	}

	return &writer{
		emitter: emitter,
	}, nil
}

func (w *writer) Write(msg Message) error {
	data, err := msg.MarshalKafka()
	if err != nil {
		return err
	}
	return w.emitter.EmitSync(msg.ID(), data)
}

func (w *writer) Shutdown() error {
	return w.emitter.Finish()
}

type Message interface {
	ID() string
	MarshalKafka() ([]byte, error)
}

type Writer interface {
	Write(msg Message) error
}

type Profile struct {
	Age        int                    `json:"age"`
	Email      string                 `json:"email"`
	Properties map[string]interface{} `json:"properties"`
}

type TriggerEvent struct {
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
	User       Profile                `json:"user"`
}

func (t TriggerEvent) ID() string {
	return t.Id
}

func (t TriggerEvent) MarshalKafka() ([]byte, error) {
	return json.Marshal(t)
}

func main() {
	argsWithProg := os.Args

	jsonFile, err := os.Open(argsWithProg[1])
	if err != nil {
		log.Fatalln(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var events []TriggerEvent

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &events)
	w, _ := NewWriter("events", []string{"localhost:9092"})

	for _, e := range events {
		w.Write(e)
	}

}
