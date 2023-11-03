module AnneDB

go 1.20

require pkg/database v1.0.0

require github.com/google/uuid v1.3.0 // indirect

replace pkg/database => ./pkg/database
