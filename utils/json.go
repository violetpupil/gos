package utils

import (
	"encoding/json"
	"os"
)

func JSONLoad(name string, v any) error {
	bs, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, v)
	if err != nil {
		return err
	}
	return nil
}

func JSONDump(name string, v any) error {
	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(name, bs, 0666)
	if err != nil {
		return err
	}
	return nil
}

func JSONPrint(v any) {
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	e.Encode(v)
}
