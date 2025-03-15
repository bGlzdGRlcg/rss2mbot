package bots

import (
	"encoding/json"
	"os"
)

func saveUserMap() error {
	userMapMutex.RLock()
	defer userMapMutex.RUnlock()

	data, err := json.Marshal(User_map)
	if err != nil {
		return err
	}

	return os.WriteFile("user.json", data, 0644)
}

func loadUserMap() error {
	userMapMutex.Lock()
	defer userMapMutex.Unlock()

	data, err := os.ReadFile("user.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &User_map)
}
