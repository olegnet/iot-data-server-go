module postgres

go 1.18

replace config => ../config

replace network => ../network

require (
	config v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.6
)
