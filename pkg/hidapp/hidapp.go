package hidapp

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Processor struct {
	ReplacedWith string   `json:"replaced_with"`
	Regex        []string `json:"regex"`

	re []*regexp.Regexp
}

const defaultReplacedWith = "**********"

// NewProcessor creates a default new parser.
func NewProcessor() *Processor {
	c := new(Processor)
	c.ReplacedWith = defaultReplacedWith
	return c
}

// NewProcessorFrom creates a new parser from a Reader, normally a json file.
func NewProcessorFrom(r io.Reader) (*Processor, error) {
	c := new(Processor)
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	if c.ReplacedWith == "" {
		c.ReplacedWith = defaultReplacedWith
	}
	for _, re := range c.Regex {
		err = c.AppendRegexp(re)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// AppendRegexp appends new regexps to this instance,
// the arguments can be string or *regexp.Regexp.
// here are two situations:
//  1. regexp without sub-match will replace the whole matched string with
//     Processor.ReplacedWith
//  2. regexp with sub-match will replace the every second sub-matched string
//     with Processor.ReplacedWith
func (p *Processor) AppendRegexp(reg ...any) error {
	for _, v := range reg {
		switch v.(type) {
		case string:
			vv, err := regexp.Compile(v.(string))
			if err != nil {
				return err
			}
			p.re = append(p.re, vv)
		case *regexp.Regexp:
			p.re = append(p.re, v.(*regexp.Regexp))
		case regexp.Regexp:
			vv := v.(regexp.Regexp)
			p.re = append(p.re, &vv)
		default:
			return fmt.Errorf("unsupported type: %#v", v)
		}
	}
	return nil
}

// Process will replace all matched string with Processor.ReplacedWith,
// if there are more than one matching groups, skip the first matching to
// make non-capture group working
func (p *Processor) Process(str string) string {
	// TODO: improve performance
	for _, re := range p.re {
		sms := re.FindAllStringSubmatch(str, -1)
		for _, sm := range sms {
			i := 1
			if len(sm) < 2 {
				i = 0
			}
			for ; i < len(sm) && len(sm[i]) > 0; i++ {
				str = strings.ReplaceAll(str, sm[i], p.ReplacedWith)
			}
		}
	}
	return str
}
