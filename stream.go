package gonginx

import (
	"errors"
)

//Stream represents stream block
type Stream struct {
	Servers    []*Server
	Directives []IDirective
}

//NewStream create an http block from a directive which has a block
func NewStream(directive IDirective) (*Stream, error) {
	if block := directive.GetBlock(); block != nil {
		stream := &Stream{
			Servers:    []*Server{},
			Directives: []IDirective{},
		}
		for _, directive := range block.GetDirectives() {
			if server, ok := directive.(*Server); ok {
				stream.Servers = append(stream.Servers, server)
				continue
			}
			stream.Directives = append(stream.Directives, directive)
		}

		return stream, nil
	}
	return nil, errors.New("http directive must have a block")
}

//GetName get directive name to construct the statment string
func (h *Stream) GetName() string { //the directive name.
	return "stream"
}

//GetParameters get directive parameters if any
func (h *Stream) GetParameters() []string {
	return []string{}
}

//GetDirectives get all directives in http
func (h *Stream) GetDirectives() []IDirective {
	directives := make([]IDirective, 0)
	for _, directive := range h.Directives {
		directives = append(directives, directive)
	}
	for _, directive := range h.Servers {
		directives = append(directives, directive)
	}
	return directives
}

//FindDirectives find directives
func (h *Stream) FindDirectives(directiveName string) []IDirective {
	directives := make([]IDirective, 0)
	for _, directive := range h.GetDirectives() {
		if directive.GetName() == directiveName {
			directives = append(directives, directive)
		}
		if directive.GetBlock() != nil {
			directives = append(directives, directive.GetBlock().FindDirectives(directiveName)...)
		}
	}

	return directives
}

//GetBlock get block if any
func (h *Stream) GetBlock() IBlock {
	return h
}
