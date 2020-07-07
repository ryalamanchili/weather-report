package weather_test

import (
	"context"
	"time"

	"github.com/google/uuid"

	weather "github.com/derekkenney/weather-report/business/data/weather"
	data "github.com/derekkenney/weather-report/foundation/data"
	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Weather", func() {
	var (
		w          weather.Weather
		db         *mongo.Client
		weatherId  string
		locationId string
	)

	BeforeEach(func() {
		weatherId = uuid.New().String()
		locationId = uuid.New().String()
		db = data.ConnectToDB()
		w = weather.Weather{
			ID:          weatherId,
			Temperature: 76.5,
			Description: "Sunny and mild",
			Location: weather.Location{
				ID:   locationId,
				Name: "Austin, TX",
			},
			DateCreated: time.Now(),
		}

	})

	Describe("Adding a new weather report", func() {
		Context("For Austin, TX", func() {
			It("should be able to create a new weather report", func() {
				err := w.CreateReport(context.TODO(), db)
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("should return an error if weather id is missing", func() {
				w.ID = ""
				err := w.CreateReport(context.TODO(), db)
				Ω(err).Should(HaveOccurred())
			})
			It("should be sunny and mild", func() {
				w.CreateReport(context.TODO(), db)
				Expect(w.Description).To(Equal("Sunny and mild"))
			})
			It("should have a temperature of 76.5", func() {
				w.CreateReport(context.TODO(), db)
				Expect(w.Temperature).To(Equal(76.5))
			})
			It("should return an error if temperature is missing", func() {
				w.Temperature = 0
				err := w.CreateReport(context.TODO(), db)
				Ω(err).Should(HaveOccurred())
			})
			It("should have a weather id", func() {
				w.CreateReport(context.TODO(), db)
				Expect(w.ID).To(Equal(weatherId))
			})
			It("should have a location", func() {
				w.CreateReport(context.TODO(), db)
				Expect(w.Location).ToNot(BeNil())
			})
			It("should have a location with a location id", func() {
				w.CreateReport(context.TODO(), db)
				Expect(w.Location.ID).To(Equal(locationId))
			})
			It("should return an error if location id is missing", func() {
				w.Location.ID = ""
				err := w.CreateReport(context.TODO(), db)
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
