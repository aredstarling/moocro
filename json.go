package moocro

// ToJSON returns a string representation of the value
func ToJSON(value interface{}) string {
	serializer := CreateJSONSerializer()
	v, err := serializer.Marshal(value)
	if err != nil {
		return err.Error()
	}

	return string(v)
}
