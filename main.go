package main

import (
	"bytes"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"thisis": "atest",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.tmpl.html", nil)
	})

	router.GET("/blog/ontravel", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ontravel.tmpl.html", nil)
	})

	router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.tmpl.html", nil)
	})

	router.GET("/projects/metaclicker", func(c *gin.Context) {
		c.HTML(http.StatusOK, "metaclicker.tmpl.html", nil)
	})

	router.GET("/projects/hamming", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hamming.tmpl.html", nil)
	})

	router.GET("/hamming", func(c *gin.Context) {
		inputCode := c.Query("code")
		if inputCode == "" {
			c.JSON(500, gin.H{
				"error": "Hamming code input is required",
			})
			return
		}
		match, _ := regexp.MatchString("[^01]+", inputCode)
		if match == true {
			c.JSON(500, gin.H{
				"error": "Only 1s and 0s allowed",
			})
			return
		}
		if len(inputCode) > 99 {
			c.JSON(500, gin.H{
				"error": "Max length of 100",
			})
			return
		}
		var inputBits []int
		for _, char := range inputCode {
			inputBits = append(inputBits, int(char-'0'))
		}

		// Calculate the length of the hamming code
		numberOfParityBits := 1
		for (float64(len(inputCode) + numberOfParityBits + 1)) >= math.Pow(2, float64(numberOfParityBits)) {
			numberOfParityBits++
		}
		hammingLength := len(inputCode) + numberOfParityBits
		hammingCode := make([]int, hammingLength)

		// Insert the parity bits (denoted as 2 because we don't know their value yet)
		inputBitPosition := 0
		parityBitPosition := 0
		for hammingPosition := 0; hammingPosition < hammingLength; hammingPosition++ {
			currentParity := math.Pow(2, float64(parityBitPosition))
			if float64(hammingPosition+1) == currentParity {
				hammingCode[hammingPosition] = 2
				parityBitPosition++
			} else {
				hammingCode[hammingPosition] = inputBits[inputBitPosition]
				inputBitPosition++
			}
		}

		// Calculate the correct values for the parity bits
		parityCalculator := 0
		parityBitPositionCalculator := 1
		parityBitPosition = 1
		bitsConsumed := 0
		for i := 0; i < hammingLength; i++ {
			if hammingCode[i] == 2 {
				j := i
				for j < hammingLength {
					if bitsConsumed >= parityBitPosition {
						bitsConsumed -= parityBitPosition
						j += parityBitPosition - 1
					} else {
						if hammingCode[j] == 1 {
							parityCalculator++
						}
						bitsConsumed++
					}
					j++
				}
				if parityCalculator%2 == 0 {
					hammingCode[i] = 0
				} else {
					hammingCode[i] = 1
				}
				parityBitPosition = int(math.Pow(2, float64(parityBitPositionCalculator)))
				parityBitPositionCalculator++
				bitsConsumed = 0
				parityCalculator = 0
			}
		}
		var finalHammingCodeStringBuffer bytes.Buffer
		for _, e := range hammingCode {
			finalHammingCodeStringBuffer.WriteString(strconv.Itoa(e))
		}

		c.JSON(200, gin.H{
			"hamming_code": finalHammingCodeStringBuffer.String(),
		})
	})

	router.GET("/hammingcodecode", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hammingcodecode.tmpl.html", nil)
	})

	router.GET("/rectangles", func(c *gin.Context) {
		c.HTML(http.StatusOK, "musicvis.tmpl.html", nil)
	})

	router.GET("/rectanglecode", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rectanglecode.tmpl.html", nil)
	})

	router.Run(":" + port)
}
