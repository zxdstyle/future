package db

import (
	"future-admin/pkg/log"
	"github.com/samber/do"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

type Manager struct {
	connections map[string]*gorm.DB
	mutex       sync.RWMutex
}

func NewManager(i *do.Injector) (*Manager, error) {
	return &Manager{
		mutex:       sync.RWMutex{},
		connections: make(map[string]*gorm.DB),
	}, nil
}

func (m *Manager) ConnectByConfig() {
	databases := viper.GetStringMap("databases")

	for name := range databases {
		cfg, ok := databases[name].(map[string]any)
		if !ok {
			continue
		}
		if cast.ToString(cfg["driver"]) != "mysql" {
			log.Errorf("not support db driver %s, name: %s", cfg["driver"], name)
		}

		db, err := Connect(cast.ToString(cfg["dsn"]))
		if err != nil {
			log.Error(err)
		}
		m.Set(name, db)
	}
}

func (m *Manager) Connect(name string) *gorm.DB {
	databases := viper.GetStringMap("databases")

	cfg, ok := databases[name].(map[string]any)
	if !ok {
		log.Errorf("not found database `%s` config", name)
		return nil
	}
	if cast.ToString(cfg["driver"]) != "mysql" {
		log.Errorf("not support db driver %s, name: %s", cfg["driver"], name)
		return nil
	}

	db, err := Connect(cast.ToString(cfg["dsn"]))
	if err != nil {
		log.Error(err)
	}
	m.Set(name, db)
	return db
}

func (m *Manager) Set(name string, conn *gorm.DB) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.connections[name] = conn
}

func (m *Manager) Get(name string) *gorm.DB {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.connections[name]
}

func (m *Manager) Healthcheck() error {
	m.mutex.RLock()
	defer m.mutex.RLock()

	for _, db := range m.connections {
		instance, err := db.DB()
		if err != nil {
			zap.Error(err)
			continue
		}
		if err := instance.Ping(); err != nil {
			zap.Error(err)
		}
	}
	return nil
}

func (m *Manager) DB(names ...string) *gorm.DB {
	name := "default"
	if len(names) > 0 {
		name = names[0]
	}
	db := m.Get(name)
	if db != nil {
		return db
	}
	return m.Connect(name)
}
