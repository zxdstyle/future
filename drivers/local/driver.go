package local

import "encoding/json"

type Local struct {
	Addition
}

func (d *Local) Init(addition string) error {
	if err := json.Unmarshal([]byte(addition), &d.Addition); err != nil {
		return err
	}

	return nil
}
