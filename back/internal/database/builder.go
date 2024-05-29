package database

import "strings"

type BuilderValue struct {
	value string
}

func Builder() BuilderValue {
	return BuilderValue{value: ""}
}

func (b *BuilderValue) Select(field string) *BuilderValue {
	b.value = strings.Join([]string{"SELECT", field}, " ") + "\n"

	return b
}
func (b *BuilderValue) From(name string) *BuilderValue {
	b.value += strings.Join([]string{"FROM", name}, " ") + "\n"

	return b
}
func (b *BuilderValue) Where(state []string, operation string) *BuilderValue {
	if len(state) < 1 {
		return b
	}

	b.value += strings.Join([]string{"WHERE", strings.Join(state, " "+operation+" ")}, " ") + "\n"

	return b
}

func (b *BuilderValue) Join(state string) *BuilderValue {
	b.value += strings.Join([]string{"INNER JOIN", state}, " ") + "\n"

	return b
}

func (b *BuilderValue) Limit(state string) *BuilderValue {
	b.value += strings.Join([]string{"LIMIT", state}, " ") + "\n"

	return b
}

func (b *BuilderValue) ToSql() string {
	return b.value
}
