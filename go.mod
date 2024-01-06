module concertie/main

go 1.21.5

require concertie/rmq v0.0.0-00010101000000-000000000000

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/rabbitmq/amqp091-go v1.9.0 // indirect
)

replace concertie/rmq => ./rmq/
