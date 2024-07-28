package passwords

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Password struct {
	m         sync.Mutex
	passwords map[string]string
	filename  string
}

func NewPassword(filename string) *Password {
	pm := &Password{
		passwords: make(map[string]string),
		filename:  filename,
	}
	pm.loadFromFile()
	return pm
}

func (pm *Password) SavePassword(name, password string) {
	pm.m.Lock()
	defer pm.m.Unlock()

	pm.passwords[name] = password
	pm.saveToFile()
	fmt.Println("Password saved successfully.")
}

func (pm *Password) GetPassword(name string) (string, bool) {
	pm.m.Lock()
	defer pm.m.Unlock()

	password, exists := pm.passwords[name]
	return password, exists
}

func (pm *Password) ListPasswords() []string {
	pm.m.Lock()
	defer pm.m.Unlock()

	names := make([]string, 0, len(pm.passwords))
	for name := range pm.passwords {
		names = append(names, name)
	}

	return names
}

func (pm *Password) saveToFile() {
	file, err := os.Create(pm.filename)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(pm.passwords)
	if err != nil {
		fmt.Printf("Failed to encode passwords to file: %v\n", err)
	}
}

func (pm *Password) loadFromFile() {
	file, err := os.Open(pm.filename)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("Failed to open file: %v\n", err)
		}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pm.passwords)
	if err != nil {
		fmt.Printf("Failed to decode passwords from file: %v\n", err)
	}
}
