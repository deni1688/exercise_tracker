package main

import (
	"github.com/deni1688/exercise_tracker/config"
	"github.com/deni1688/exercise_tracker/domain"
	"github.com/deni1688/exercise_tracker/mykafka"
	"github.com/deni1688/exercise_tracker/rest"
	"github.com/deni1688/exercise_tracker/storage"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	env := config.GetEnv()

	err := config.LoadConfigFile(env)
	if err != nil {
		log.Fatal("error loading the config file:", err)
	}

	er, err := storage.NewExerciseRepository(config.GetString("storage.driver"))
	if err != nil {
		log.Fatal("error creating new repository:", err)
	}

	br := mykafka.NewKafkaProducer(config.GetString("brokers.kafka.server"))
	defer br.Close()
	es := domain.NewExerciseService(er, br)

	router := gin.Default()
	rest.NewResource(router, rest.GetHandlerFor(es, domain.ExerciseCollection))

	log.Fatal("error starting server:", router.Run(config.GetString("server.port")))
}
