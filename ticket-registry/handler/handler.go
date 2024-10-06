package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"ticket-registry/models"
	"ticket-registry/service"
	"time"

	"github.com/opentracing/opentracing-go"
)

type HTTPHeaderReader http.Header

// ForeachKey iterates over all key-value pairs in the HTTP headers
func (h HTTPHeaderReader) ForeachKey(handler func(key, val string) error) error {
	for key, values := range h {
		for _, value := range values {
			if err := handler(key, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func GetTicketsForGivenTypeAndQuantity(c *gin.Context) {

	parentSpanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, HTTPHeaderReader(c.Request.Header))

	span := opentracing.GlobalTracer().StartSpan("ServiceB-Request", opentracing.ChildOf(parentSpanCtx))
	defer span.Finish()

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketType := c.Query("ticketType")
	quantityStr := c.Query("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity value"})
		return
	}

	// Log the request for tracking
	log.Printf("Received GetTicketsForGivenTypeAndQuantity request: Type=%s, Quantity=%d", ticketType, quantity)

	db, _ := c.Get("db")

	// Call the service layer to handle the order placement
	tickets, err := service.GetTicketsForGivenTypeAndQuantity(db.(*gorm.DB), &models.GetTicketsRequest{
		TicketType: ticketType,
		Quantity:   quantity,
	})
	if err != nil {
		log.Printf("Error fetching tickets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func ShowTickets(c *gin.Context) {

	// parentSpanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.TextMapReader(r.Header))

	// span := opentracing.GlobalTracer().StartSpan("ShowTickets")
	// defer span.Finish()

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Log the request for tracking
	log.Printf("Received ShowTickets request")

	db, _ := c.Get("db")

	// Call the service layer to handle the order placement
	tickets, err := service.ShowTickets(db.(*gorm.DB))
	if err != nil {
		log.Printf("Error fetching tickets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}
