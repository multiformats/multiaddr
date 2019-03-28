package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	godog "github.com/DATA-DOG/godog"
	gherkin "github.com/DATA-DOG/godog/gherkin"
	assistdog "github.com/hellomd/assistdog"

	_ "github.com/multiformats/go-multiaddr"
)

var multiaddrCmd = []string{"go", "run", "github.com/multiformats/go-multiaddr/multiaddr"}
var assist *assistdog.Assist

func init() {
	bin := os.Getenv("MULTIADDR_BIN")
	if len(bin) > 0 {
		multiaddrCmd = strings.Split(bin, " ")
	}

	assist = assistdog.NewDefault()
}

func TestGodog(t *testing.T) {
	fmt.Printf(`MULTIADDR_BIN="%s"`+"\n", strings.Join(multiaddrCmd, " "))

	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:        "pretty",
		Paths:         []string{"."},
		StopOnFailure: true,
		Strict:        true,
	})

	if status != 0 {
		t.Errorf("exit status: %d", status)
	}
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the multiaddr (.+)$`, theMultiaddr)
	s.Step(`^the packed form is (.+)$`, thePackedFormIs)
	s.Step(`^the packed size is (.+) bytes$`, thePackedSizeIs)
	s.Step(`^the string form is (.+)$`, theStringFormIs)
	s.Step(`^the string size is (.+) bytes$`, theStringSizeIs)
	s.Step(`^the components are:$`, theComponentsAre)
}

var theMaddr multiaddr

type multiaddr struct {
	String     string
	StringSize string
	Packed     string
	PackedSize string
	Components []component
}

type component struct {
	String       string
	StringSize   string
	Packed       string
	PackedSize   string
	Value        string
	RawValue     string
	ValueSize    string
	Protocol     string
	Codec        string
	Uvarint      string
	LengthPrefix string
}

func theMultiaddr(addr string) error {
	theMaddr = multiaddr{}

	cmd := exec.Command(multiaddrCmd[0], append(multiaddrCmd[1:], addr)...)
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command error: %s - stderr: %s", err, cmd.Stderr.(*bytes.Buffer).String())
	}

	stdout := cmd.Stdout.(*bytes.Buffer)
	if err := json.Unmarshal(stdout.Bytes(), &theMaddr); err != nil {
		return err
	}
	return nil
}

func thePackedFormIs(packed string) error {
	if theMaddr.Packed != packed {
		return fmt.Errorf("expected %s, got %s", packed, theMaddr.Packed)
	}
	return nil
}

func thePackedSizeIs(size string) error {
	if theMaddr.PackedSize != size {
		return fmt.Errorf("expected %s, got %s", size, theMaddr.PackedSize)
	}
	return nil
}

func theStringFormIs(str string) error {
	if theMaddr.String != str {
		return fmt.Errorf("expected %s, got %s", str, theMaddr.String)
	}
	return nil
}

func theStringSizeIs(size string) error {
	if theMaddr.StringSize != size {
		return fmt.Errorf("expected %s, got %s", size, theMaddr.StringSize)
	}
	return nil
}

func theComponentsAre(tbl *gherkin.DataTable) error {
	table, err := assist.ParseSlice(tbl)
	if err != nil {
		return err
	}

	actual := len(theMaddr.Components)
	expected := len(table)
	if actual != expected {
		return fmt.Errorf("expected %d row, got %d", expected, actual)
	}

	for i, _ := range table {
		actual, ok := table[i]["string"]
		if ok {
			expected := theMaddr.Components[i].String
			if actual != expected {
				return fmt.Errorf("expected string '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["stringSize"]
		if ok {
			expected := theMaddr.Components[i].StringSize
			if actual != expected {
				return fmt.Errorf("expected stringSize '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["packed"]
		if ok {
			expected := theMaddr.Components[i].Packed
			if actual != expected {
				return fmt.Errorf("expected packed '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["packedSize"]
		if ok {
			expected := theMaddr.Components[i].PackedSize
			if actual != expected {
				return fmt.Errorf("expected packedSize '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["value"]
		if ok {
			expected := theMaddr.Components[i].Value
			if actual != expected {
				return fmt.Errorf("expected value '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["rawValue"]
		if ok {
			expected := theMaddr.Components[i].RawValue
			if actual != expected {
				return fmt.Errorf("expected rawValue '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["valueSize"]
		if ok {
			expected := theMaddr.Components[i].ValueSize
			if actual != expected {
				return fmt.Errorf("expected valueSize '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["protocol"]
		if ok {
			expected := theMaddr.Components[i].Protocol
			if actual != expected {
				return fmt.Errorf("expected protocol '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["codec"]
		if ok {
			expected := theMaddr.Components[i].Codec
			if actual != expected {
				return fmt.Errorf("expected codec '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["uvarint"]
		if ok {
			expected := theMaddr.Components[i].Uvarint
			if actual != expected {
				return fmt.Errorf("expected uvarint '%s', got '%s'", expected, actual)
			}
		}
		actual, ok = table[i]["lengthPrefix"]
		if ok {
			expected := theMaddr.Components[i].LengthPrefix
			if actual != expected {
				return fmt.Errorf("expected lengthPrefix '%s', got '%s'", expected, actual)
			}
		}
	}

	return nil
}
