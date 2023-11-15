package provider

import (
	"github.com/getopends/providerd"
	dstvmoz "github.com/getopends/providerd/provider/tv/dstv-moz"
)

var List = map[string]providerd.NewProviderFunc{
	// DSTv Mozambique
	"dstvmoz": dstvmoz.New,
}
