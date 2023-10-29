package indexserver

import (
	_ "embed"
)

//go:embed rpc.swagger.json
var SwaggerJSONData string
