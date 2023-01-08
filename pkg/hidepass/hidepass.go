package hidepass

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	Regex []string `json:"regex"`
}

var conf = new(Config)

// SetConfig will set the global conf variable with content c
func SetConfig(c []byte) error {
	err := json.Unmarshal(c, conf)
	if err != nil {
		return err
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
	for _, expr := range conf.Regex {
		log.Printf("compile %s ...\n", expr)
		r := regexp.MustCompile(expr)
		if r.MatchString(rStr) {
			log.Printf("regex expression '%s' is ignored.", expr)
			continue
		}
		sms := r.FindAllStringSubmatch(str, -1)
		fmt.Printf("%#v\n", sms)
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
