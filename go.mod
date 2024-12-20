module learn

go 1.23.4

require (
	github.com/jonjicjan/cryptit v0.0.0
	github.com/pborman/uuid v1.2.1
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/jonjicjan/cryptit => ../cryptit
