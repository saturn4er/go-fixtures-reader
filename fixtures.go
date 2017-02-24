package fixtures

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	Equal int = iota
	NotEqual
	StartsWith
	NotStartsWith
	EndsWith
	NotEndsWith
	Contains
	NotContains
)

type filter struct {
	key   string
	value string
	fType int // filter type
}

func (c *filter) CheckValue(s string) bool {
	switch c.fType {
	case Equal:
		return s == c.value
	case NotEqual:
		return s != c.value
	case StartsWith:
		return strings.HasPrefix(s, c.value)
	case NotStartsWith:
		return !strings.HasPrefix(s, c.value)
	case EndsWith:
		return strings.HasSuffix(s, c.value)
	case NotEndsWith:
		return !strings.HasSuffix(s, c.value)
	case Contains:
		return strings.Contains(s, c.value)
	case NotContains:
		return !strings.Contains(s, c.value)
	default:
		panic(fmt.Errorf("Unknown filter type: %v", c.fType))
	}
}

type Fixtures struct {
	fixtureName string
	data        []map[string]string
	filters     []filter
}

func (f *Fixtures) Filter(key string, t int, value string) *Fixtures {
	f.filters = append(f.filters, filter{
		key:   key,
		fType: t,
		value: value,
	})
	return f
}
func (f *Fixtures) checkFixture(fx map[string]string) bool {
	for _, f := range f.filters {
		val, ok := fx[f.key]
		if !ok {
			return false
		}
		if !f.CheckValue(val) {
			return false
		}
	}
	return true
}
func (f *Fixtures) All() []map[string]string {
	result := []map[string]string{}
	for _, d := range f.data {
		if f.checkFixture(d) {
			result = append(result, d)
		}
	}
	return result
}
func (f *Fixtures) First() (map[string]string, error) {
	data := f.All()
	if len(data) == 0 {
		return nil, errors.New("no such fixture")
	}
	return data[0], nil
}
func (f *Fixtures) Count() int {
	data := f.All()
	return len(data)
}

// GetFixture return Fixtures object to search rows in Fixtures
func GetFixture(path string) (*Fixtures, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't read fixtures file")
	}
	fixtures := []map[string]string{}
	err = yaml.Unmarshal(file, &fixtures)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse fixtures file")
	}
	result := new(Fixtures)
	result.data = fixtures
	result.fixtureName = path
	return result, nil
}
