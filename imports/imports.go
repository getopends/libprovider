package imports

import (
	"github.com/getopends/providerd"
	dstvmoz "github.com/getopends/providerd/provider/tv/dstv-moz"
)

func enrol() error {
	r := providerd.Registry{}

	m := map[string]providerd.NewProviderFunc{
		"dstv-moz": dstvmoz.New,
	}

	for k, v := range m {
		if err := r.Register(k, v); err != nil {
			return err
		}
	}

	return nil
}
