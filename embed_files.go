package embedfiles

import "embed"

//go:embed .env
var ENV embed.FS

//go:embed internal/view/payment.html
var Tmpl embed.FS
