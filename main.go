package main

import (

	"concertie/rmq"

)

func main() {

	rmq.PublishMessage("purchase_order", "purchase!")
}

