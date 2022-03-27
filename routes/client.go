package routers

import (
	d "textgopher/database"
	a "textgopher/pkg"
)

var mongoClient = d.Connection()
var awsClient = a.GetSession()
