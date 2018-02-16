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
	_ "github.com/heroku/x/hmetrics/onload"
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

	router.GET("/blog/ontechnology", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ontechnology.tmpl.html", nil)
	})

	router.GET("/blog/graphicshistory", func(c *gin.Context) {
		c.HTML(http.StatusOK, "graphicshistory.tmpl.html", nil)
	})

	router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.tmpl.html", nil)
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

	router.GET("/projects/graphtv", func(c *gin.Context) {
		c.HTML(http.StatusOK, "graphtv.tmpl.html", nil)
	})

	router.GET("/graphtvcode", func(c *gin.Context) {
		c.HTML(http.StatusOK, "graphtvcode.tmpl.html", nil)
	})

	router.GET("/rectangles", func(c *gin.Context) {
		c.HTML(http.StatusOK, "musicvis.tmpl.html", nil)
	})

	router.GET("/rectanglecode", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rectanglecode.tmpl.html", nil)
	})

	router.GET("/2018review", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"goals_and_expectations": gin.H{
				"goals": []string{
					"Get Minivan to a working state",
					"Deploy and secure the Daily Reports app",
					"Deploy, secure, and enhance the T&C Setup Tool",
					"Build a new XREF API that will serve as the source of truth for XREFs",
					"Retire the legacy DBLoader app",
					"Begin work to transfer FI4 to use the new XREF API",
					"Begin mapping companies to XREFs",
					"Deploy and secure IDMS",
				},
				"strategic_themes": "My goals align well with the strategic themes for the company. They all work towards replacing DC4UI by breaking it up into logical services, each of which represents a different piece of functionality. I also hooked all of these apps up to deploy using the company's BDP, and contributed effort towards improving the BDP. This all allows for better scaling, a clearer path to growth, and faster deployments and therefore faster customer success. My efforts have helped to transform the way that core applications at SPS interact with one another.",
				"summary":          "I'm very happy with the work that I did in 2017. Particularly with the work that went into the XREF API, which only took a few months to be completely finished functionally. I proved that I can work in any environment using any set of language, tools, and frameworks by working with Minivan (a Node app), daily reports (Spring Java), T&C Setup Tool (Python with Flask), and XREFs (Kotlin with Dropwizard). My enthusiasm and willingness to dive into new things has helped propel my work towards my goals. In addition, I proved that I'm light on my feet and capable of fast, high quality development by deploying POC Resets and the new Location app each in under 3 days. These small, one-off apps address a specific need and help to get rid of DC4UI as directly as possible. My work with mapping companies and updating FI4 got off to a good start, but definitely needs to continue into 2018. The only thing I'm disappointed about is that the FI4 work hasn't been able to get done faster, and part of that blame rests with me.",
			},
			"accomplishments_and_strengths": gin.H{
				"accomplishments": []string{
					"Finished the T&C Setup Tool",
					"Automated front-end Commerce Platform deployments",
					"Built XREF API",
					"Created a full test suite for XREF API",
					"Replaced DBLoader",
					"Replaced XTS tab in DC4UI",
					"Contributed to FI4",
					"Reviewed IDMS",
					"Deployed IDMS",
					"Moved all our services to ECS to fix performance problems",
					"Fixed lots of MiniVan performance problems",
					"Helped build out a stage environment for FI4",
					"Created Location Service for address lookups",
					"Created POC Reset Service",
					"Created a deploy tester/jira ticket creater/hipchat notifier for Internal Apps",
					"Spearheaded FI4 XREF changes",
				},
				"strengths": []string{
					"Persistence",
					"Empathy",
					"Comprehension and retention",
					"Getting results",
					"Working hard",
					"Vision",
					"Leadership",
					"Creativity",
				},
				"values":  "I've embodied multiple core SPS values over the course of the last year. I've shown a thirst for growth both from a high-level company perspective by creating more scalable and maintainable systems, as well as from a low-level personal perspective by pestering John about promoting me. Results clearly matter for me as well, otherwise I wouldn't have produced so many of them. I'm always learning more to be more; gaining knowledge both about SPS-related domain problems as well as everything in life is critical to me. I like to win, and my work will help SPS win in this industry. I 'get after it' in the sense that I'm always looking for the fastest and easiest thing we can do to bring value, and then trying to increment on that instead of taking on huge grandiose projects all at once. I've shown that I'm ready to succeed together with the rest of the SPS team by collaborating with all the other development and operations teams to produce software. I lead the way by often being the first to volunteer to do work that no one else wants to do like fixing/replacing old systems. Finally, I 'give back' by donating over $600 a year through United Way, which still isn't nearly enough.",
				"summary": "I've gotten a lot done this year. I'm especially proud of the XREF API and what a success that's been. I've been exited to do FI4 work and hope to continue to do high-value cross-team projects in the future.",
			},
			"development_opportunities": gin.H{
				"areas_for_improvement": []string{
					"Nudging teammates/people in the right direction",
					"Trusting my own decision-making and sticking to my guns when appropriate",
					"Organizational and logistical skills surrounding our apps",
					"Onboarding teammates quickly",
				},
				"summary": "While I feel that 2017 was a very successful year for me, there are still a few areas where I can improve. I'd like to take on more of a leadership role within my team by making sure that everyone is working on things that provide value. There have been a couple of instances this year where I felt that certain tasks teammates of mine were doing were taking too long or weren't valuable enough to work on. I'd like to get better at voicing my opinion (without being rude) and pushing work in a better direction. I can see what provides the most value extremely well; I embody that in my own work. But I need to get better at helping others see it too. A part of that is trusting myself more. I need to know when I've identified a good strategic plan for the team and be able to really push for it with confidence. When Nick came onto the team, it exposed a few areas for improvement that I have. Namely, organizational skills and onboarding. I need to get better at documenting software, and showing why I made certain decisions, not just what decisions were made. This comes primarily in the form of confluence pages, AD groups, hipchat integrations, etc... Also, with Nate, Petya, and Nick all coming on to the team at different times, I got some good experience onboarding new people. That being said, none of that onboarding happened as quickly as I'd like.",
			},
		})
	})

	router.Run(":" + port)
}
