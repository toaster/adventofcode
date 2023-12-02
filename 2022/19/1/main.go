package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	bps := parseBlueprints(io.ReadLines())
	qualityLevelSum := 0
	for i, bp := range bps {
		geodeCount := computeMaxOpenGeodes(bp, 24)
		qualityLevelSum += geodeCount * (i + 1)
		break
	}
	fmt.Println(qualityLevelSum)
}

func computeMaxOpenGeodes(bp *blueprint, steps int) int {
	oreRobots := robotStock{name: "ore-collecting", amount: 1, price: bp.oreRobotPrice}
	clayRobots := robotStock{name: "clay-collecting", price: bp.clayRobotPrice}
	obsidianRobots := robotStock{name: "obsidian-collecting", price: bp.obsidianRobotPrice}
	geodeRobots := robotStock{name: "geode-cracking", price: bp.geodeRobotPrice}
	resourceStock := resources{}
	for i := 0; i < steps; i++ {
		fmt.Println("== Minute", i+1, "==")
		var buildGeodeRobot, buildObsidianRobot, buildClayRobot, buildOreRobot bool
		if clayRobots.amount == 0 {
			buildClayRobot = orderRobotIfAffordable(clayRobots, &resourceStock)
		} else if obsidianRobots.amount == 0 {
			buildObsidianRobot = orderRobotIfAffordable(obsidianRobots, &resourceStock)

			if !buildObsidianRobot {
				remainingOreRounds := math.Ceil((float64(obsidianRobots.price.ore) - float64(resourceStock.ore)) / float64(oreRobots.amount))
				remainingClayRounds := math.Ceil((float64(obsidianRobots.price.clay) - float64(resourceStock.clay)) / float64(clayRobots.amount))
				oreRate := float64(oreRobots.amount) / float64(obsidianRobots.price.ore)
				clayRate := float64(clayRobots.amount) / float64(obsidianRobots.price.clay)
				fmt.Println("ror", remainingOreRounds, "rcr", remainingClayRounds, "or", oreRate*float64(steps-i-1), "cr", clayRate*float64(steps-i-1))
				if remainingOreRounds < remainingClayRounds {
					buildClayRobot = orderRobotIfAffordable(clayRobots, &resourceStock)
				}
			}
		} else {
			buildGeodeRobot = orderRobotIfAffordable(geodeRobots, &resourceStock)

			if !buildGeodeRobot {
				remainingOreRounds := math.Ceil((float64(geodeRobots.price.ore) - float64(resourceStock.ore)) / float64(oreRobots.amount))
				remainingObsidianRounds := math.Ceil((float64(geodeRobots.price.obsidian) - float64(resourceStock.obsidian)) / float64(obsidianRobots.amount))
				if remainingOreRounds < remainingObsidianRounds {
					buildObsidianRobot = orderRobotIfAffordable(obsidianRobots, &resourceStock)

					if !buildObsidianRobot {
						remainingOreRounds := math.Ceil((float64(obsidianRobots.price.ore) - float64(resourceStock.ore)) / float64(oreRobots.amount))
						remainingClayRounds := math.Ceil((float64(obsidianRobots.price.clay) - float64(resourceStock.clay)) / float64(clayRobots.amount))
						if remainingOreRounds < remainingClayRounds {
							buildClayRobot = orderRobotIfAffordable(clayRobots, &resourceStock)
						}
					}
				}
				// oreRate := float64(oreRobots.amount) / float64(geodeRobots.price.ore)
				// obsidianRate := float64(obsidianRobots.amount) / float64(geodeRobots.price.obsidian)
			}
		}

		updateResource(&resourceStock.ore, oreRobots)
		updateResource(&resourceStock.clay, clayRobots)
		updateResource(&resourceStock.obsidian, obsidianRobots)
		updateResource(&resourceStock.geode, geodeRobots)

		updateRobotStock(&oreRobots, buildOreRobot)
		updateRobotStock(&clayRobots, buildClayRobot)
		updateRobotStock(&obsidianRobots, buildObsidianRobot)
		updateRobotStock(&geodeRobots, buildGeodeRobot)
		fmt.Println()
	}
	return resourceStock.geode
}

func updateResource(amount *int, robots robotStock) {
	if robots.amount == 0 {
		return
	}

	*amount = *amount + robots.amount
	action := "robots collect"
	if robots.amount == 1 {
		action = "robot collects"
	}
	resourceName := strings.Split(robots.name, "-")[0]
	fmt.Printf("%d %s %s %d %s; you now have %d %s.\n", robots.amount, robots.name, action, robots.amount, resourceName, *amount, resourceName)
}

func orderRobotIfAffordable(stock robotStock, res *resources) bool {
	p := stock.price
	if p.ore <= res.ore && p.obsidian <= res.obsidian && p.clay <= res.clay {
		fmt.Print("Spend ")
		appendAnd := false
		if p.ore > 0 {
			if appendAnd {
				fmt.Print(" and ")
			}
			fmt.Print(p.ore, " ore")
			appendAnd = true
		}
		if p.clay > 0 {
			if appendAnd {
				fmt.Print(" and ")
			}
			fmt.Print(p.clay, " clay")
			appendAnd = true
		}
		if p.obsidian > 0 {
			if appendAnd {
				fmt.Print(" and ")
			}
			fmt.Print(p.obsidian, " obsidian")
			appendAnd = true
		}
		fmt.Println(" to start building a", stock.name, "robot.")
		res.Subtract(p)
		return true
	}

	return false
}

func parseBlueprints(lines []string) []*blueprint {
	var bps []*blueprint
	for _, line := range lines {
		d := strings.Split(line, ":")[1]
		c := strings.Split(d, ".")
		bp := &blueprint{
			oreRobotPrice:      resources{ore: io.ParseInt(strings.Split(c[0], " ")[5])},
			clayRobotPrice:     resources{ore: io.ParseInt(strings.Split(c[1], " ")[5])},
			obsidianRobotPrice: resources{ore: io.ParseInt(strings.Split(c[2], " ")[5]), clay: io.ParseInt(strings.Split(c[2], " ")[8])},
			geodeRobotPrice:    resources{ore: io.ParseInt(strings.Split(c[3], " ")[5]), obsidian: io.ParseInt(strings.Split(c[3], " ")[8])},
		}
		bps = append(bps, bp)
	}
	return bps
}

func updateRobotStock(stock *robotStock, build bool) {
	if !build {
		return
	}

	stock.amount++
	fmt.Println("The new", stock.name, "robot is ready; you now have", stock.amount, "of them.")
}

type blueprint struct {
	oreRobotPrice      resources
	clayRobotPrice     resources
	obsidianRobotPrice resources
	geodeRobotPrice    resources
}

type resources struct {
	clay     int
	geode    int
	obsidian int
	ore      int
}

func (r *resources) Subtract(p resources) {
	r.ore -= p.ore
	r.clay -= p.clay
	r.obsidian -= p.obsidian
	r.geode -= p.geode
}

type robotStock struct {
	amount int
	name   string
	price  resources
}
