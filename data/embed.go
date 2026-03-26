package data

import "embed"

// DRCDatabase holds the embedded seed copy of the Douay-Rheims Catholic Bible
// SQLite database. This is extracted to the user's app config directory on
// first launch.
//
//go:embed DRC.db
var DRCDatabase embed.FS
