package manticore

import (
	"fmt"
	"strings"
)

type procedureOption struct {
	name  string
	value string
}

func (o *procedureOption) string() string {
	return fmt.Sprintf("%s as %s", o.value, o.name)
}

func optsToString(opts []procedureOption) string {
	sb := strings.Builder{}
	for _, opt := range opts {
		sb.WriteString(",")
		sb.WriteString(opt.string())
	}
	return sb.String()
}

func Option(name string, value string) procedureOption {
	return procedureOption{
		name:  name,
		value: value,
	}
}
