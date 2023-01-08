package hidepass

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"
)

type config struct {
	Regex []string `json:"regex"`
}

var re []*regexp.Regexp

// SetConfig will set the global conf variable with content c
func SetConfig(c []byte) error {
	var conf = new(config)
	err := json.Unmarshal(c, conf)
	if err != nil {
		return err
	}
	for _, expr := range conf.Regex {
		re = append(re, regexp.MustCompile(expr))
	}

	return nil
}

// ReadConfig will read all contents that encoded with json from path and
// set the global conf variable which used to Hide string
func ReadConfig(path string) error {
	c, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return SetConfig(c)
}

// Hide will replace all matched string with '**P.A.S.S.P.H.R.A.S.E**',
// if there are more than one matching groups, skip the first matching to
// make non-capture group working
func Hide(str string) string {
	rStr := "**P.A.S.S.P.H.R.A.S.E**"
	for _, r := range re {
		if r.MatchString(rStr) {
			log.Printf("regex expression '%s' is ignored.", r.String())
			continue
		}
		sms := r.FindAllStringSubmatch(str, -1)
		for _, sm := range sms {
			i := 1
			if len(sm) < 2 {
				i = 0
			}
			for ; i < len(sm); i++ {
				str = strings.ReplaceAll(str, sm[i], rStr)
			}
		}
	}
	return str
}
