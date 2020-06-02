package extension

import (
	"github.com/navinds25/EviveInDesignServer/pkg/pdf"
	"github.com/navinds25/EviveInDesignServer/pkg/splitter"
)

// Impl is the interface for all Extensions
type Impl interface {
	Name() string
	Run(map[string]string) (map[string]string, error)
}

// Directory is the variable to hold all extensions
var Directory map[string]Impl

// LoadExtensions loads extensions, must be called at startup
func LoadExtensions() error {
	directory := make(map[string]Impl)
	directory["PDF"] = &pdf.Extension{}
	directory["SPLITTER"] = &splitter.Extension{}
	Directory = directory
	return nil
}
