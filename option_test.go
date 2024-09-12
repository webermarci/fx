package fx

import (
	"encoding/json"
	"testing"
)

type testData struct {
	Primitive Option[int]            `json:"primitive"`
	Map       Option[map[string]int] `json:"map"`
	Slice     Option[[]int]          `json:"slice"`
	Pointer   Option[*int]           `json:"pointer"`
	Struct    Option[testStruct]     `json:"struct"`
}

type testStruct struct {
	Number int `json:"number"`
}

func TestOptionSome(t *testing.T) {
	data := testData{
		Primitive: Some(1),
		Map:       Some(map[string]int{"test": 1}),
		Slice:     Some([]int{1, 2, 3}),
		Pointer:   Some(new(int)),
		Struct:    Some(testStruct{Number: 1}),
	}

	if !data.Primitive.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if !data.Map.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if !data.Slice.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if !data.Pointer.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if !data.Struct.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if data.Primitive.Get() != 1 {
		t.Errorf("Expected 1, got %d", data.Primitive.Get())
	}

	if data.Map.Get()["test"] != 1 {
		t.Errorf("Expected 1, got %d", data.Map.Get()["test"])
	}

	if len(data.Slice.Get()) != 3 {
		t.Errorf("Expected 3, got %d", len(data.Slice.Get()))
	}

	if *data.Pointer.Get() != 0 {
		t.Errorf("Expected 0, got %d", *data.Pointer.Get())
	}

	if data.Struct.Get().Number != 1 {
		t.Errorf("Expected 1, got %d", data.Struct.Get().Number)
	}

	if value, found := data.Primitive.Value(); !found || value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	if value, found := data.Map.Value(); !found || value["test"] != 1 {
		t.Errorf("Expected 1, got %d", value["test"])
	}

	if value, found := data.Slice.Value(); !found || len(value) != 3 {
		t.Errorf("Expected 3, got %d", len(value))
	}

	if value, found := data.Pointer.Value(); !found || *value != 0 {
		t.Errorf("Expected 0, got %d", *value)
	}

	if value, found := data.Struct.Value(); !found || value.Number != 1 {
		t.Errorf("Expected 1, got %d", value.Number)
	}
}

func TestOptionSomeOrElse(t *testing.T) {
	data := testData{
		Primitive: Some(1),
		Map:       Some(map[string]int{"test": 1}),
		Slice:     Some([]int{1, 2, 3}),
		Pointer:   Some(new(int)),
		Struct:    Some(testStruct{Number: 1}),
	}

	if data.Primitive.OrElse(0) != 1 {
		t.Errorf("Expected 1, got %d", data.Primitive.OrElse(0))
	}

	if data.Map.OrElse(map[string]int{"test": 0})["test"] != 1 {
		t.Errorf("Expected 1, got %d", data.Map.OrElse(map[string]int{"test": 0})["test"])
	}

	if len(data.Slice.OrElse([]int{0})) != 3 {
		t.Errorf("Expected 3, got %d", len(data.Slice.OrElse([]int{0})))
	}

	if *data.Pointer.OrElse(new(int)) != 0 {
		t.Errorf("Expected 0, got %d", *data.Pointer.OrElse(new(int)))
	}

	if data.Struct.OrElse(testStruct{Number: 0}).Number != 1 {
		t.Errorf("Expected 1, got %d", data.Struct.OrElse(testStruct{Number: 0}).Number)
	}
}

func TestOptionNone(t *testing.T) {
	data := testData{
		Primitive: None[int](),
		Map:       None[map[string]int](),
		Slice:     None[[]int](),
		Pointer:   None[*int](),
		Struct:    None[testStruct](),
	}

	if data.Primitive.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if data.Map.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if data.Slice.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if data.Pointer.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if data.Struct.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if _, found := data.Primitive.Value(); found {
		t.Errorf("Expected false, got true")
	}

	if _, found := data.Map.Value(); found {
		t.Errorf("Expected false, got true")
	}

	if _, found := data.Slice.Value(); found {
		t.Errorf("Expected false, got true")
	}

	if _, found := data.Pointer.Value(); found {
		t.Errorf("Expected false, got true")
	}

	if _, found := data.Struct.Value(); found {
		t.Errorf("Expected false, got true")
	}
}

func TestOptionNoneOrElse(t *testing.T) {
	data := testData{
		Primitive: None[int](),
		Map:       None[map[string]int](),
		Slice:     None[[]int](),
		Pointer:   None[*int](),
		Struct:    None[testStruct](),
	}

	if data.Primitive.OrElse(1) != 1 {
		t.Errorf("Expected 1, got %d", data.Primitive.OrElse(1))
	}

	if data.Map.OrElse(map[string]int{"test": 1})["test"] != 1 {
		t.Errorf("Expected 1, got %d", data.Map.OrElse(map[string]int{"test": 1})["test"])
	}

	if len(data.Slice.OrElse([]int{1})) != 1 {
		t.Errorf("Expected 1, got %d", len(data.Slice.OrElse([]int{1})))
	}

	if *data.Pointer.OrElse(new(int)) != 0 {
		t.Errorf("Expected 0, got %d", *data.Pointer.OrElse(new(int)))
	}

	if data.Struct.OrElse(testStruct{Number: 1}).Number != 1 {
		t.Errorf("Expected 1, got %d", data.Struct.OrElse(testStruct{Number: 1}).Number)
	}
}

func TestOptionSomeJsonMarshaling(t *testing.T) {
	data := testData{
		Primitive: Some(1),
		Map:       Some(map[string]int{"test": 1}),
		Slice:     Some([]int{1, 2, 3}),
		Pointer:   Some(new(int)),
		Struct:    Some(testStruct{Number: 1}),
	}

	marshaledData, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	if string(marshaledData) != `{"primitive":1,"map":{"test":1},"slice":[1,2,3],"pointer":0,"struct":{"number":1}}` {
		t.Errorf("Expected %s, got %s", `{"primitive":1,"map":{"test":1},"slice":[1,2,3],"pointer":0,"struct":{"number":1}}`, string(marshaledData))
	}
}

func TestOptionNoneJsonMarshaling(t *testing.T) {
	data := testData{
		Primitive: None[int](),
		Map:       None[map[string]int](),
		Slice:     None[[]int](),
		Pointer:   None[*int](),
		Struct:    None[testStruct](),
	}

	marshaledData, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	if string(marshaledData) != `{"primitive":null,"map":null,"slice":null,"pointer":null,"struct":null}` {
		t.Errorf("Expected %s, got %s", `{"primitive":null,"map":null,"slice":null,"pointer":null,"struct":null}`, string(marshaledData))
	}
}

func TestOptionSomeJsonUnmarshaling(t *testing.T) {
	data := testData{
		Primitive: Some(1),
		Map:       Some(map[string]int{"test": 1}),
		Slice:     Some([]int{1, 2, 3}),
		Pointer:   Some(new(int)),
		Struct:    Some(testStruct{Number: 1}),
	}

	marshaledData, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	var unmarshaledData testData
	err = json.Unmarshal(marshaledData, &unmarshaledData)
	if err != nil {
		t.Error(err)
	}

	if !unmarshaledData.Primitive.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if !unmarshaledData.Map.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if !unmarshaledData.Slice.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if !unmarshaledData.Pointer.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if !unmarshaledData.Struct.IsSome() {
		t.Errorf("Expected true, got false")
	}

	if unmarshaledData.Primitive.Get() != 1 {
		t.Errorf("Expected 1, got %d", unmarshaledData.Primitive.Get())
	}

	if unmarshaledData.Map.Get()["test"] != 1 {
		t.Errorf("Expected 1, got %d", unmarshaledData.Map.Get()["test"])
	}

	if len(unmarshaledData.Slice.Get()) != 3 {
		t.Errorf("Expected 3, got %d", len(unmarshaledData.Slice.Get()))
	}

	if *unmarshaledData.Pointer.Get() != 0 {
		t.Errorf("Expected 0, got %d", *unmarshaledData.Pointer.Get())
	}

	if unmarshaledData.Struct.Get().Number != 1 {
		t.Errorf("Expected 1, got %d", unmarshaledData.Struct.Get().Number)
	}
}

func TestOptionNoneJsonUnmarshaling(t *testing.T) {
	data := testData{
		Primitive: None[int](),
		Map:       None[map[string]int](),
		Slice:     None[[]int](),
		Pointer:   None[*int](),
		Struct:    None[testStruct](),
	}

	marshaledData, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	var unmarshaledData testData
	err = json.Unmarshal(marshaledData, &unmarshaledData)
	if err != nil {
		t.Error(err)
	}

	if unmarshaledData.Primitive.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if unmarshaledData.Map.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if unmarshaledData.Slice.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if unmarshaledData.Pointer.IsSome() {
		t.Errorf("Expected false, got true")
	}

	if unmarshaledData.Struct.IsSome() {
		t.Errorf("Expected false, got true")
	}
}
