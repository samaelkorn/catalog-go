package database

type BuilderValue struct {
	value string
}

func Builder() BuilderValue {
	return BuilderValue{value: ""}
}

func (b *BuilderValue) Select() *BuilderValue {
	b.value = `SELECT * `

	return b
}
func (b *BuilderValue) From(name string) *BuilderValue {
	b.value += `FROM ` + name

	return b
}

func (b *BuilderValue) ToSql() string {
	return b.value
}
