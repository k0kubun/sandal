package data

import (
	"strings"
)

type (
	Type interface {
		typetype()
		Equal(Type) bool
		String() string
	}

	NamedType struct {
		Name string
	}

	CallableType struct {
		Parameters []Type
	}

	ArrayType struct {
		ElemType Type
	}

	HandshakeChannelType struct {
		Elems []Type
	}

	BufferedChannelType struct {
		BufferSize Expr
		Elems      []Type
	}
)

func (x NamedType) typetype()            {}
func (x CallableType) typetype()         {}
func (x ArrayType) typetype()            {}
func (x HandshakeChannelType) typetype() {}
func (x BufferedChannelType) typetype()  {}

func (x NamedType) Equal(ty Type) bool {
	if ty, b := ty.(NamedType); b {
		return (ty.Name == x.Name)
	} else {
		return false
	}
}

func (x CallableType) Equal(ty Type) bool {
	if ty, b := ty.(CallableType); b {
		if len(ty.Parameters) != len(x.Parameters) {
			return false
		}
		for i := 0; i < len(x.Parameters); i++ {
			if !ty.Parameters[i].Equal(x.Parameters[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (x ArrayType) Equal(ty Type) bool {
	if ty, b := ty.(ArrayType); b {
		return ty.ElemType.Equal(x.ElemType)
	} else {
		return false
	}
}

func (x HandshakeChannelType) Equal(ty Type) bool {
	if ty, b := ty.(HandshakeChannelType); b {
		if len(ty.Elems) != len(x.Elems) {
			return false
		}
		for i := 0; i < len(x.Elems); i++ {
			if !ty.Elems[i].Equal(x.Elems[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (x BufferedChannelType) Equal(ty Type) bool {
	if ty, b := ty.(BufferedChannelType); b {
		if len(ty.Elems) != len(x.Elems) {
			return false
		}
		for i := 0; i < len(x.Elems); i++ {
			if !ty.Elems[i].Equal(x.Elems[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (x NamedType) String() string { return x.Name }
func (x ArrayType) String() string { return "[]" + x.ElemType.String() }

func (x CallableType) String() string {
	params := []string{}
	for _, param := range x.Parameters {
		params = append(params, param.String())
	}
	return "callable(" + strings.Join(params, ", ") + ")"
}

func (x HandshakeChannelType) String() string {
	elems := []string{}
	for _, elem := range x.Elems {
		elems = append(elems, elem.String())
	}
	return "channel {" + strings.Join(elems, ", ") + "}"
}

func (x BufferedChannelType) String() string {
	bufsize := ""
	if x.BufferSize != nil {
		bufsize = x.BufferSize.String()
	}

	elems := []string{}
	for _, elem := range x.Elems {
		elems = append(elems, elem.String())
	}

	return "channel [" + bufsize + "] {" + strings.Join(elems, ", ") + "}"
}
