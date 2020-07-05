package parameter

import (
	"database/sql"
	"strconv"
	"sync"
	"time"
)

var (
	once    sync.Once
	ms      *model
	errData error
)

type Instance interface {
	GetString(name string) (string, bool)
	GetInt64(name string) (int64, bool)
	GetFloat64(name string) (float64, bool)
	GetTime(name, format string) (time.Time, bool)
}

type model struct {
	mutex sync.Mutex
	data  map[string]string
}

func New(engine string, db *sql.DB) error {
	once.Do(func() {
		if ms == nil {
			ms = &model{}
			errData = LoadData(engine, db)
		}
	})
	if errData != nil {
		return errData
	}

	return nil
}

func GetInstance() *model {
	return ms
}

func LoadData(engine string, db *sql.DB) error {
	params, err := NewRepository(engine, db).getAll()
	if err != nil {
		return err
	}

	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ms.data = make(map[string]string)
	for _, param := range params {
		ms.data[param.Name] = param.Value
	}

	return nil
}

// GetByName obtains the value by name
// in the parameter cache
func (m *model) GetString(name string) (string, bool) {
	v, ok := m.data[name]

	return v, ok
}

func (m *model) GetInt64(name string) (int64, bool) {
	value, ok := m.GetString(name)
	if !ok {
		return 0, ok
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, false
	}

	return intValue, true
}

func (m *model) GetInt(name string) (int, bool) {
	value, ok := m.GetString(name)
	if !ok {
		return 0, ok
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}

	return intValue, true
}

func (m *model) GetFloat64(name string) (float64, bool) {
	value, ok := m.GetString(name)
	if !ok {
		return 0, ok
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}

	return floatValue, true
}

func (m *model) GetTime(name, format string) (time.Time, bool) {
	value, ok := m.GetString(name)
	if !ok {
		return time.Time{}, ok
	}

	timeValue, err := time.Parse(format, value)
	if err != nil {
		return time.Time{}, false
	}

	return timeValue, true
}

func (m *model) GetBool(name string) (bool, bool) {
	value, ok := m.GetString(name)
	if !ok {
		return false, ok
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, false
	}

	return boolValue, true
}
