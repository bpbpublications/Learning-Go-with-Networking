package main

import (
    "fmt"
    "log"

    "github.com/linkedin/goavro/v2"
)

const personSchema = `{
"type": "record",
"name": "Person",
"fields": [
{"name": "name", "type": "string"},
{"name": "age", "type": "int"},
{"name": "hobbies", "type": {"type": "array", "items": "string"}}
]}
`

func main() {
    // Create a new Avro codec with the schema
    codec, err := goavro.NewCodec(personSchema)
    if err != nil {
        log.Fatal("Error while creating Avro codec:", err)
    }

    // Create a new Avro record
    record := map[string]interface{}{
        "name":    "John",
        "age":     25,
        "hobbies": []interface{}{"Reading", "Gaming"},
    }

    // Serialize the record to Avro binary
    avroData, err := codec.BinaryFromNative(nil, record)
    if err != nil {
        log.Fatal("Error while serializing Avro record:", err)
    }

    fmt.Println("Serialized Avro Data:", avroData)

    // Deserialize the Avro binary to a new record
    newRecord, _, err := codec.NativeFromBinary(avroData)
    if err != nil {
        log.Fatal("Error while deserializing Avro data:", err)
    }

    fmt.Println("Deserialized Avro Record:", newRecord)
}