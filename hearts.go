package main

import (
	"fmt"
	"time"
)

type Hearts struct {
	row     int
	column  int
	graphic rune
	color   string
}

var tickrate int = 10000
var heart1 *Hearts
var heart2 *Hearts
var heart3 *Hearts
var heart4 *Hearts
var heart5 *Hearts
var heart6 *Hearts
var heart7 *Hearts
var heart8 *Hearts
var heart9 *Hearts
var heart10 *Hearts
var heart11 *Hearts
var heart12 *Hearts
var heart13 *Hearts
var heart14 *Hearts
var heart15 *Hearts
var heart16 *Hearts
var heart17 *Hearts
var heart18 *Hearts
var heart19 *Hearts
var heart20 *Hearts
var heart21 *Hearts
var heart22 *Hearts
var heart23 *Hearts
var heart24 *Hearts
var heart25 *Hearts
var heart26 *Hearts
var heart27 *Hearts
var heart28 *Hearts
var heart29 *Hearts
var heart30 *Hearts
var heart31 *Hearts
var heart32 *Hearts
var heart33 *Hearts
var heart34 *Hearts
var heart35 *Hearts
var heart36 *Hearts
var heart37 *Hearts
var heart38 *Hearts
var heart39 *Hearts
var heart40 *Hearts
var heart41 *Hearts
var heart42 *Hearts
var heart43 *Hearts
var heart44 *Hearts
var heart45 *Hearts
var heart46 *Hearts
var heart47 *Hearts
var heart48 *Hearts
var heart49 *Hearts
var heart50 *Hearts
var heart51 *Hearts
var heart52 *Hearts
var heart53 *Hearts
var heart54 *Hearts
var heart55 *Hearts
var heart56 *Hearts
var heart57 *Hearts
var heart58 *Hearts
var heart59 *Hearts
var heart60 *Hearts
var heart61 *Hearts
var heart62 *Hearts
var heart63 *Hearts
var heart64 *Hearts
var heart65 *Hearts
var heart66 *Hearts
var heart67 *Hearts
var heart68 *Hearts
var heart69 *Hearts
var heart70 *Hearts
var heart71 *Hearts
var heart72 *Hearts
var heart73 *Hearts
var heart74 *Hearts
var heart75 *Hearts
var heart76 *Hearts
var heart77 *Hearts
var heart78 *Hearts
var heart79 *Hearts
var heart80 *Hearts

func setup() {
	heart1 = new(Hearts)
	heart1.row = 1
	heart1.column = 1
	heart1.graphic = '♥'
	heart1.color = "\u001b[31m"
	heart2 = new(Hearts)
	heart2.row = 1
	heart2.column = 2
	heart2.graphic = '♥'
	heart2.color = "\u001b[35m"
	heart3 = new(Hearts)
	heart3.row = 1
	heart3.column = 3
	heart3.graphic = '♥'
	heart3.color = "\u001b[34m"
	heart4 = new(Hearts)
	heart4.row = 1
	heart4.column = 4
	heart4.graphic = '♥'
	heart4.color = "\u001b[32m"
	heart5 = new(Hearts)
	heart5.row = 1
	heart5.column = 5
	heart5.graphic = '♥'
	heart5.color = "\u001b[36m"
	heart6 = new(Hearts)
	heart6.row = 1
	heart6.column = 6
	heart6.graphic = '♥'
	heart6.color = "\u001b[33m"
	heart7 = new(Hearts)
	heart7.row = 1
	heart7.column = 7
	heart7.graphic = '♥'
	heart7.color = "\u001b[31m"
	heart8 = new(Hearts)
	heart8.row = 1
	heart8.column = 8
	heart8.graphic = '♥'
	heart8.color = "\u001b[35m"
	heart9 = new(Hearts)
	heart9.row = 1
	heart9.column = 9
	heart9.graphic = '♥'
	heart9.color = "\u001b[34m"
	heart10 = new(Hearts)
	heart10.row = 1
	heart10.column = 10
	heart10.graphic = '♥'
	heart10.color = "\u001b[32m"
	heart11 = new(Hearts)
	heart11.row = 1
	heart11.column = 11
	heart11.graphic = '♥'
	heart11.color = "\u001b[36m"
	heart12 = new(Hearts)
	heart12.row = 1
	heart12.column = 12
	heart12.graphic = '♥'
	heart12.color = "\u001b[33m"
	heart13 = new(Hearts)
	heart13.row = 1
	heart13.column = 13
	heart13.graphic = '♥'
	heart13.color = "\u001b[31m"
	heart14 = new(Hearts)
	heart14.row = 1
	heart14.column = 14
	heart14.graphic = '♥'
	heart14.color = "\u001b[35m"
	heart15 = new(Hearts)
	heart15.row = 1
	heart15.column = 15
	heart15.graphic = '♥'
	heart15.color = "\u001b[34m"
	heart16 = new(Hearts)
	heart16.row = 1
	heart16.column = 16
	heart16.graphic = '♥'
	heart16.color = "\u001b[32m"
	heart17 = new(Hearts)
	heart17.row = 1
	heart17.column = 17
	heart17.graphic = '♥'
	heart17.color = "\u001b[36m"
	heart18 = new(Hearts)
	heart18.row = 1
	heart18.column = 18
	heart18.graphic = '♥'
	heart18.color = "\u001b[33m"
	heart19 = new(Hearts)
	heart19.row = 1
	heart19.column = 19
	heart19.graphic = '♥'
	heart19.color = "\u001b[31m"
	heart20 = new(Hearts)
	heart20.row = 1
	heart20.column = 20
	heart20.graphic = '♥'
	heart20.color = "\u001b[35m"
	heart21 = new(Hearts)
	heart21.row = 1
	heart21.column = 21
	heart21.graphic = '♥'
	heart21.color = "\u001b[34m"
	heart22 = new(Hearts)
	heart22.row = 1
	heart22.column = 22
	heart22.graphic = '♥'
	heart22.color = "\u001b[32m"
	heart23 = new(Hearts)
	heart23.row = 1
	heart23.column = 23
	heart23.graphic = '♥'
	heart23.color = "\u001b[36m"
	heart24 = new(Hearts)
	heart24.row = 1
	heart24.column = 24
	heart24.graphic = '♥'
	heart24.color = "\u001b[33m"
	heart25 = new(Hearts)
	heart25.row = 1
	heart25.column = 25
	heart25.graphic = '♥'
	heart25.color = "\u001b[31m"
	heart26 = new(Hearts)
	heart26.row = 1
	heart26.column = 26
	heart26.graphic = '♥'
	heart26.color = "\u001b[35m"
	heart27 = new(Hearts)
	heart27.row = 1
	heart27.column = 27
	heart27.graphic = '♥'
	heart27.color = "\u001b[34m"
	heart28 = new(Hearts)
	heart28.row = 1
	heart28.column = 28
	heart28.graphic = '♥'
	heart28.color = "\u001b[32m"
	heart29 = new(Hearts)
	heart29.row = 1
	heart29.column = 29
	heart29.graphic = '♥'
	heart29.color = "\u001b[36m"
	heart30 = new(Hearts)
	heart30.row = 1
	heart30.column = 30
	heart30.graphic = '♥'
	heart30.color = "\u001b[33m"
	heart31 = new(Hearts)
	heart31.row = 1
	heart31.column = 31
	heart31.graphic = '♥'
	heart31.color = "\u001b[31m"
	heart32 = new(Hearts)
	heart32.row = 1
	heart32.column = 32
	heart32.graphic = '♥'
	heart32.color = "\u001b[35m"
	heart33 = new(Hearts)
	heart33.row = 1
	heart33.column = 33
	heart33.graphic = '♥'
	heart33.color = "\u001b[34m"
	heart34 = new(Hearts)
	heart34.row = 1
	heart34.column = 34
	heart34.graphic = '♥'
	heart34.color = "\u001b[32m"
	heart35 = new(Hearts)
	heart35.row = 1
	heart35.column = 35
	heart35.graphic = '♥'
	heart35.color = "\u001b[36m"
	heart36 = new(Hearts)
	heart36.row = 1
	heart36.column = 36
	heart36.graphic = '♥'
	heart36.color = "\u001b[33m"
	heart37 = new(Hearts)
	heart37.row = 1
	heart37.column = 37
	heart37.graphic = '♥'
	heart37.color = "\u001b[31m"
	heart38 = new(Hearts)
	heart38.row = 1
	heart38.column = 38
	heart38.graphic = '♥'
	heart38.color = "\u001b[35m"
	heart39 = new(Hearts)
	heart39.row = 1
	heart39.column = 39
	heart39.graphic = '♥'
	heart39.color = "\u001b[34m"
	heart40 = new(Hearts)
	heart40.row = 1
	heart40.column = 40
	heart40.graphic = '♥'
	heart40.color = "\u001b[32m"
	heart41 = new(Hearts)
	heart41.row = 1
	heart41.column = 41
	heart41.graphic = '♥'
	heart41.color = "\u001b[36m"
	heart42 = new(Hearts)
	heart42.row = 1
	heart42.column = 42
	heart42.graphic = '♥'
	heart42.color = "\u001b[33m"
	heart43 = new(Hearts)
	heart43.row = 1
	heart43.column = 43
	heart43.graphic = '♥'
	heart43.color = "\u001b[31m"
	heart44 = new(Hearts)
	heart44.row = 1
	heart44.column = 44
	heart44.graphic = '♥'
	heart44.color = "\u001b[35m"
	heart45 = new(Hearts)
	heart45.row = 1
	heart45.column = 45
	heart45.graphic = '♥'
	heart45.color = "\u001b[34m"
	heart46 = new(Hearts)
	heart46.row = 1
	heart46.column = 46
	heart46.graphic = '♥'
	heart46.color = "\u001b[32m"
	heart47 = new(Hearts)
	heart47.row = 1
	heart47.column = 47
	heart47.graphic = '♥'
	heart47.color = "\u001b[36m"
	heart48 = new(Hearts)
	heart48.row = 1
	heart48.column = 48
	heart48.graphic = '♥'
	heart48.color = "\u001b[33m"
	heart49 = new(Hearts)
	heart49.row = 1
	heart49.column = 49
	heart49.graphic = '♥'
	heart49.color = "\u001b[31m"
	heart50 = new(Hearts)
	heart50.row = 1
	heart50.column = 50
	heart50.graphic = '♥'
	heart50.color = "\u001b[35m"
	heart51 = new(Hearts)
	heart51.row = 1
	heart51.column = 51
	heart51.graphic = '♥'
	heart51.color = "\u001b[34m"
	heart52 = new(Hearts)
	heart52.row = 1
	heart52.column = 52
	heart52.graphic = '♥'
	heart52.color = "\u001b[32m"
	heart53 = new(Hearts)
	heart53.row = 1
	heart53.column = 53
	heart53.graphic = '♥'
	heart53.color = "\u001b[36m"
	heart54 = new(Hearts)
	heart54.row = 1
	heart54.column = 54
	heart54.graphic = '♥'
	heart54.color = "\u001b[33m"
	heart55 = new(Hearts)
	heart55.row = 1
	heart55.column = 55
	heart55.graphic = '♥'
	heart55.color = "\u001b[31m"
	heart56 = new(Hearts)
	heart56.row = 1
	heart56.column = 56
	heart56.graphic = '♥'
	heart56.color = "\u001b[35m"
	heart57 = new(Hearts)
	heart57.row = 1
	heart57.column = 57
	heart57.graphic = '♥'
	heart57.color = "\u001b[34m"
	heart58 = new(Hearts)
	heart58.row = 1
	heart58.column = 58
	heart58.graphic = '♥'
	heart58.color = "\u001b[32m"
	heart59 = new(Hearts)
	heart59.row = 1
	heart59.column = 59
	heart59.graphic = '♥'
	heart59.color = "\u001b[36m"
	heart60 = new(Hearts)
	heart60.row = 1
	heart60.column = 60
	heart60.graphic = '♥'
	heart60.color = "\u001b[33m"
	heart61 = new(Hearts)
	heart61.row = 1
	heart61.column = 61
	heart61.graphic = '♥'
	heart61.color = "\u001b[31m"
	heart62 = new(Hearts)
	heart62.row = 1
	heart62.column = 62
	heart62.graphic = '♥'
	heart62.color = "\u001b[35m"
	heart63 = new(Hearts)
	heart63.row = 1
	heart63.column = 63
	heart63.graphic = '♥'
	heart63.color = "\u001b[34m"
	heart64 = new(Hearts)
	heart64.row = 1
	heart64.column = 64
	heart64.graphic = '♥'
	heart64.color = "\u001b[32m"
	heart65 = new(Hearts)
	heart65.row = 1
	heart65.column = 65
	heart65.graphic = '♥'
	heart65.color = "\u001b[36m"
	heart66 = new(Hearts)
	heart66.row = 1
	heart66.column = 66
	heart66.graphic = '♥'
	heart66.color = "\u001b[33m"
	heart67 = new(Hearts)
	heart67.row = 1
	heart67.column = 67
	heart67.graphic = '♥'
	heart67.color = "\u001b[31m"
	heart68 = new(Hearts)
	heart68.row = 1
	heart68.column = 68
	heart68.graphic = '♥'
	heart68.color = "\u001b[35m"
	heart69 = new(Hearts)
	heart69.row = 1
	heart69.column = 69
	heart69.graphic = '♥'
	heart69.color = "\u001b[34m"
	heart70 = new(Hearts)
	heart70.row = 1
	heart70.column = 70
	heart70.graphic = '♥'
	heart70.color = "\u001b[32m"
	heart71 = new(Hearts)
	heart71.row = 1
	heart71.column = 71
	heart71.graphic = '♥'
	heart71.color = "\u001b[36m"
	heart72 = new(Hearts)
	heart72.row = 1
	heart72.column = 72
	heart72.graphic = '♥'
	heart72.color = "\u001b[33m"
	heart73 = new(Hearts)
	heart73.row = 1
	heart73.column = 73
	heart73.graphic = '♥'
	heart73.color = "\u001b[31m"
	heart74 = new(Hearts)
	heart74.row = 1
	heart74.column = 74
	heart74.graphic = '♥'
	heart74.color = "\u001b[35m"
	heart75 = new(Hearts)
	heart75.row = 1
	heart75.column = 75
	heart75.graphic = '♥'
	heart75.color = "\u001b[34m"
	heart76 = new(Hearts)
	heart76.row = 1
	heart76.column = 76
	heart76.graphic = '♥'
	heart76.color = "\u001b[32m"
	heart77 = new(Hearts)
	heart77.row = 1
	heart77.column = 77
	heart77.graphic = '♥'
	heart77.color = "\u001b[36m"
	heart78 = new(Hearts)
	heart78.row = 1
	heart78.column = 78
	heart78.graphic = '♥'
	heart78.color = "\u001b[33m"
	heart79 = new(Hearts)
	heart79.row = 1
	heart79.column = 79
	heart79.graphic = '♥'
	heart79.color = "\u001b[31m"
	heart80 = new(Hearts)
	heart80.row = 1
	heart80.column = 80
	heart80.graphic = '♥'
	heart80.color = "\u001b[35m"
}

// func resetrow() {
// 	heart1 = new(Hearts)
// 	heart1.row = 1
// 	heart1.graphic = '♥'
// 	heart1.color = "\u001b[31m"
// 	heart2 = new(Hearts)
// 	heart2.row = 1
// 	heart2.graphic = '♥'
// 	heart2.color = "\u001b[35m"
// 	heart3 = new(Hearts)
// 	heart3.row = 1
// 	heart3.graphic = '♥'
// 	heart3.color = "\u001b[34m"
// 	heart4 = new(Hearts)
// 	heart4.row = 1
// 	heart4.graphic = '♥'
// 	heart4.color = "\u001b[32m"
// 	heart5 = new(Hearts)
// 	heart5.row = 1
// 	heart5.graphic = '♥'
// 	heart5.color = "\u001b[36m"
// 	heart6 = new(Hearts)
// 	heart6.row = 1
// 	heart6.graphic = '♥'
// 	heart6.color = "\u001b[33m"
// 	heart7 = new(Hearts)
// 	heart7.row = 1
// 	heart7.graphic = '♥'
// 	heart7.color = "\u001b[31m"
// 	heart8 = new(Hearts)
// 	heart8.row = 1
// 	heart8.graphic = '♥'
// 	heart8.color = "\u001b[35m"
// 	heart9 = new(Hearts)
// 	heart9.row = 1
// 	heart9.graphic = '♥'
// 	heart9.color = "\u001b[34m"
// 	heart10 = new(Hearts)
// 	heart10.row = 1
// 	heart10.graphic = '♥'
// 	heart10.color = "\u001b[32m"
// 	heart11 = new(Hearts)
// 	heart11.row = 1
// 	heart11.graphic = '♥'
// 	heart11.color = "\u001b[36m"
// 	heart12 = new(Hearts)
// 	heart12.row = 1
// 	heart12.graphic = '♥'
// 	heart12.color = "\u001b[33m"
// 	heart13 = new(Hearts)
// 	heart13.row = 1
// 	heart13.graphic = '♥'
// 	heart13.color = "\u001b[31m"
// 	heart14 = new(Hearts)
// 	heart14.row = 1
// 	heart14.graphic = '♥'
// 	heart14.color = "\u001b[35m"
// 	heart15 = new(Hearts)
// 	heart15.row = 1
// 	heart15.graphic = '♥'
// 	heart15.color = "\u001b[34m"
// 	heart16 = new(Hearts)
// 	heart16.row = 1
// 	heart16.graphic = '♥'
// 	heart16.color = "\u001b[32m"
// 	heart17 = new(Hearts)
// 	heart17.row = 1
// 	heart17.graphic = '♥'
// 	heart17.color = "\u001b[36m"
// 	heart18 = new(Hearts)
// 	heart18.row = 1
// 	heart18.graphic = '♥'
// 	heart18.color = "\u001b[33m"
// 	heart19 = new(Hearts)
// 	heart19.row = 1
// 	heart19.graphic = '♥'
// 	heart19.color = "\u001b[31m"
// 	heart20 = new(Hearts)
// 	heart20.row = 1
// 	heart20.graphic = '♥'
// 	heart20.color = "\u001b[35m"
// 	heart21 = new(Hearts)
// 	heart21.row = 1
// 	heart21.graphic = '♥'
// 	heart21.color = "\u001b[34m"
// 	heart22 = new(Hearts)
// 	heart22.row = 1
// 	heart22.graphic = '♥'
// 	heart22.color = "\u001b[32m"
// 	heart23 = new(Hearts)
// 	heart23.row = 1
// 	heart23.graphic = '♥'
// 	heart23.color = "\u001b[36m"
// 	heart24 = new(Hearts)
// 	heart24.row = 1
// 	heart24.graphic = '♥'
// 	heart24.color = "\u001b[33m"
// 	heart25 = new(Hearts)
// 	heart25.row = 1
// 	heart25.graphic = '♥'
// 	heart25.color = "\u001b[31m"
// 	heart26 = new(Hearts)
// 	heart26.row = 1
// 	heart26.graphic = '♥'
// 	heart26.color = "\u001b[35m"
// 	heart27 = new(Hearts)
// 	heart27.row = 1
// 	heart27.graphic = '♥'
// 	heart27.color = "\u001b[34m"
// 	heart28 = new(Hearts)
// 	heart28.row = 1
// 	heart28.graphic = '♥'
// 	heart28.color = "\u001b[32m"
// 	heart29 = new(Hearts)
// 	heart29.row = 1
// 	heart29.graphic = '♥'
// 	heart29.color = "\u001b[36m"
// 	heart30 = new(Hearts)
// 	heart30.row = 1
// 	heart30.graphic = '♥'
// 	heart30.color = "\u001b[33m"
// 	heart31 = new(Hearts)
// 	heart31.row = 1
// 	heart31.graphic = '♥'
// 	heart31.color = "\u001b[31m"
// 	heart32 = new(Hearts)
// 	heart32.row = 1
// 	heart32.graphic = '♥'
// 	heart32.color = "\u001b[35m"
// 	heart33 = new(Hearts)
// 	heart33.row = 1
// 	heart33.graphic = '♥'
// 	heart33.color = "\u001b[34m"
// 	heart34 = new(Hearts)
// 	heart34.row = 1
// 	heart34.graphic = '♥'
// 	heart34.color = "\u001b[32m"
// 	heart35 = new(Hearts)
// 	heart35.row = 1
// 	heart35.graphic = '♥'
// 	heart35.color = "\u001b[36m"
// 	heart36 = new(Hearts)
// 	heart36.row = 1
// 	heart36.graphic = '♥'
// 	heart36.color = "\u001b[33m"
// 	heart37 = new(Hearts)
// 	heart37.row = 1
// 	heart37.graphic = '♥'
// 	heart37.color = "\u001b[31m"
// 	heart38 = new(Hearts)
// 	heart38.row = 1
// 	heart38.graphic = '♥'
// 	heart38.color = "\u001b[35m"
// 	heart39 = new(Hearts)
// 	heart39.row = 1
// 	heart39.graphic = '♥'
// 	heart39.color = "\u001b[34m"
// 	heart40 = new(Hearts)
// 	heart40.row = 1
// 	heart40.graphic = '♥'
// 	heart40.color = "\u001b[32m"
// 	heart41 = new(Hearts)
// 	heart41.row = 1
// 	heart41.graphic = '♥'
// 	heart41.color = "\u001b[36m"
// 	heart42 = new(Hearts)
// 	heart42.row = 1
// 	heart42.graphic = '♥'
// 	heart42.color = "\u001b[33m"
// 	heart43 = new(Hearts)
// 	heart43.row = 1
// 	heart43.graphic = '♥'
// 	heart43.color = "\u001b[31m"
// 	heart44 = new(Hearts)
// 	heart44.row = 1
// 	heart44.graphic = '♥'
// 	heart44.color = "\u001b[35m"
// 	heart45 = new(Hearts)
// 	heart45.row = 1
// 	heart45.graphic = '♥'
// 	heart45.color = "\u001b[34m"
// 	heart46 = new(Hearts)
// 	heart46.row = 1
// 	heart46.graphic = '♥'
// 	heart46.color = "\u001b[32m"
// 	heart47 = new(Hearts)
// 	heart47.row = 1
// 	heart47.graphic = '♥'
// 	heart47.color = "\u001b[36m"
// 	heart48 = new(Hearts)
// 	heart48.row = 1
// 	heart48.graphic = '♥'
// 	heart48.color = "\u001b[33m"
// 	heart49 = new(Hearts)
// 	heart49.row = 1
// 	heart49.graphic = '♥'
// 	heart49.color = "\u001b[31m"
// 	heart50 = new(Hearts)
// 	heart50.row = 1
// 	heart50.graphic = '♥'
// 	heart50.color = "\u001b[35m"
// 	heart51 = new(Hearts)
// 	heart51.row = 1
// 	heart51.graphic = '♥'
// 	heart51.color = "\u001b[34m"
// 	heart52 = new(Hearts)
// 	heart52.row = 1
// 	heart52.graphic = '♥'
// 	heart52.color = "\u001b[32m"
// 	heart53 = new(Hearts)
// 	heart53.row = 1
// 	heart53.graphic = '♥'
// 	heart53.color = "\u001b[36m"
// 	heart54 = new(Hearts)
// 	heart54.row = 1
// 	heart54.graphic = '♥'
// 	heart54.color = "\u001b[33m"
// 	heart55 = new(Hearts)
// 	heart55.row = 1
// 	heart55.graphic = '♥'
// 	heart55.color = "\u001b[31m"
// 	heart56 = new(Hearts)
// 	heart56.row = 1
// 	heart56.graphic = '♥'
// 	heart56.color = "\u001b[35m"
// 	heart57 = new(Hearts)
// 	heart57.row = 1
// 	heart57.graphic = '♥'
// 	heart57.color = "\u001b[34m"
// 	heart58 = new(Hearts)
// 	heart58.row = 1
// 	heart58.graphic = '♥'
// 	heart58.color = "\u001b[32m"
// 	heart59 = new(Hearts)
// 	heart59.row = 1
// 	heart59.graphic = '♥'
// 	heart59.color = "\u001b[36m"
// 	heart60 = new(Hearts)
// 	heart60.row = 1
// 	heart60.graphic = '♥'
// 	heart60.color = "\u001b[33m"
// 	heart61 = new(Hearts)
// 	heart61.row = 1
// 	heart61.graphic = '♥'
// 	heart61.color = "\u001b[31m"
// 	heart62 = new(Hearts)
// 	heart62.row = 1
// 	heart62.graphic = '♥'
// 	heart62.color = "\u001b[35m"
// 	heart63 = new(Hearts)
// 	heart63.row = 1
// 	heart63.graphic = '♥'
// 	heart63.color = "\u001b[34m"
// 	heart64 = new(Hearts)
// 	heart64.row = 1
// 	heart64.graphic = '♥'
// 	heart64.color = "\u001b[32m"
// 	heart65 = new(Hearts)
// 	heart65.row = 1
// 	heart65.graphic = '♥'
// 	heart65.color = "\u001b[36m"
// 	heart66 = new(Hearts)
// 	heart66.row = 1
// 	heart66.graphic = '♥'
// 	heart66.color = "\u001b[33m"
// 	heart67 = new(Hearts)
// 	heart67.row = 1
// 	heart67.graphic = '♥'
// 	heart67.color = "\u001b[31m"
// 	heart68 = new(Hearts)
// 	heart68.row = 1
// 	heart68.graphic = '♥'
// 	heart68.color = "\u001b[35m"
// 	heart69 = new(Hearts)
// 	heart69.row = 1
// 	heart69.graphic = '♥'
// 	heart69.color = "\u001b[34m"
// 	heart70 = new(Hearts)
// 	heart70.row = 1
// 	heart70.graphic = '♥'
// 	heart70.color = "\u001b[32m"
// 	heart71 = new(Hearts)
// 	heart71.row = 1
// 	heart71.graphic = '♥'
// 	heart71.color = "\u001b[36m"
// 	heart72 = new(Hearts)
// 	heart72.row = 1
// 	heart72.graphic = '♥'
// 	heart72.color = "\u001b[33m"
// 	heart73 = new(Hearts)
// 	heart73.row = 1
// 	heart73.graphic = '♥'
// 	heart73.color = "\u001b[31m"
// 	heart74 = new(Hearts)
// 	heart74.row = 1
// 	heart74.graphic = '♥'
// 	heart74.color = "\u001b[35m"
// 	heart75 = new(Hearts)
// 	heart75.row = 1
// 	heart75.graphic = '♥'
// 	heart75.color = "\u001b[34m"
// 	heart76 = new(Hearts)
// 	heart76.row = 1
// 	heart76.graphic = '♥'
// 	heart76.color = "\u001b[32m"
// 	heart77 = new(Hearts)
// 	heart77.row = 1
// 	heart77.graphic = '♥'
// 	heart77.color = "\u001b[36m"
// 	heart78 = new(Hearts)
// 	heart78.row = 1
// 	heart78.graphic = '♥'
// 	heart78.color = "\u001b[33m"
// 	heart79 = new(Hearts)
// 	heart79.row = 1
// 	heart79.graphic = '♥'
// 	heart79.color = "\u001b[31m"
// 	heart80 = new(Hearts)
// 	heart80.row = 1
// 	heart80.graphic = '♥'
// 	heart80.color = "\u001b[35m"
// }
func firstrowdraw() {
	fmt.Printf("\033[%d;%dH%s%c", heart1.row, heart1.column, heart1.color, heart1.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart2.row, heart2.column, heart2.color, heart2.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart3.row, heart3.column, heart3.color, heart3.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart4.row, heart4.column, heart4.color, heart4.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart5.row, heart5.column, heart5.color, heart5.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart6.row, heart6.column, heart6.color, heart6.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart7.row, heart7.column, heart7.color, heart7.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart8.row, heart8.column, heart8.color, heart8.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart9.row, heart9.column, heart9.color, heart9.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart10.row, heart10.column, heart10.color, heart10.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart11.row, heart11.column, heart11.color, heart11.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart12.row, heart12.column, heart12.color, heart12.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart13.row, heart13.column, heart13.color, heart13.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart14.row, heart14.column, heart14.color, heart14.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart15.row, heart15.column, heart15.color, heart15.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart16.row, heart16.column, heart16.color, heart16.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart17.row, heart17.column, heart17.color, heart17.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart18.row, heart18.column, heart18.color, heart18.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart19.row, heart19.column, heart19.color, heart19.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart20.row, heart20.column, heart20.color, heart20.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart21.row, heart21.column, heart21.color, heart21.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart22.row, heart22.column, heart22.color, heart22.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart23.row, heart23.column, heart23.color, heart23.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart24.row, heart24.column, heart24.color, heart24.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart25.row, heart25.column, heart25.color, heart25.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart26.row, heart26.column, heart26.color, heart26.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart27.row, heart27.column, heart27.color, heart27.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart28.row, heart28.column, heart28.color, heart28.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart29.row, heart29.column, heart29.color, heart29.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart30.row, heart30.column, heart30.color, heart30.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart31.row, heart31.column, heart31.color, heart31.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart32.row, heart32.column, heart32.color, heart32.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart33.row, heart33.column, heart33.color, heart33.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart34.row, heart34.column, heart34.color, heart34.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart35.row, heart35.column, heart35.color, heart35.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart36.row, heart36.column, heart36.color, heart36.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart37.row, heart37.column, heart37.color, heart37.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart38.row, heart38.column, heart38.color, heart38.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart39.row, heart39.column, heart39.color, heart39.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart40.row, heart40.column, heart40.color, heart40.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart41.row, heart41.column, heart41.color, heart41.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart42.row, heart42.column, heart42.color, heart42.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart43.row, heart43.column, heart43.color, heart43.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart44.row, heart44.column, heart44.color, heart44.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart45.row, heart45.column, heart45.color, heart45.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart46.row, heart46.column, heart46.color, heart46.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart47.row, heart47.column, heart47.color, heart47.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart48.row, heart48.column, heart48.color, heart48.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart49.row, heart49.column, heart49.color, heart49.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart50.row, heart50.column, heart50.color, heart50.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart51.row, heart51.column, heart51.color, heart51.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart52.row, heart52.column, heart52.color, heart52.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart53.row, heart53.column, heart53.color, heart53.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart54.row, heart54.column, heart54.color, heart54.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart55.row, heart55.column, heart55.color, heart55.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart56.row, heart56.column, heart56.color, heart56.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart57.row, heart57.column, heart57.color, heart57.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart58.row, heart58.column, heart58.color, heart58.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart59.row, heart59.column, heart59.color, heart59.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart60.row, heart60.column, heart60.color, heart60.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart61.row, heart61.column, heart61.color, heart61.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart62.row, heart62.column, heart62.color, heart62.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart63.row, heart63.column, heart63.color, heart63.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart64.row, heart64.column, heart64.color, heart64.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart65.row, heart65.column, heart65.color, heart65.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart66.row, heart66.column, heart66.color, heart66.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart67.row, heart67.column, heart67.color, heart67.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart68.row, heart68.column, heart68.color, heart68.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart69.row, heart69.column, heart69.color, heart69.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart70.row, heart70.column, heart70.color, heart70.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart71.row, heart71.column, heart71.color, heart71.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart72.row, heart72.column, heart72.color, heart72.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart73.row, heart73.column, heart73.color, heart73.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart74.row, heart74.column, heart74.color, heart74.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart75.row, heart75.column, heart75.color, heart75.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart76.row, heart76.column, heart76.color, heart76.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart77.row, heart77.column, heart77.color, heart77.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart78.row, heart78.column, heart78.color, heart78.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart79.row, heart79.column, heart79.color, heart79.graphic)
	fmt.Printf("\033[%d;%dH%s%c", heart80.row, heart80.column, heart80.color, heart80.graphic)
}
func fall() {
	for heart80.row < 25 {
		heart80.row++
		fmt.Printf("\033[%d;%dH%s%c", heart80.row, heart80.column, heart80.color, heart80.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart80.row = 1
	for heart79.row < 25 {
		heart79.row++
		fmt.Printf("\033[%d;%dH%s%c", heart79.row, heart79.column, heart79.color, heart79.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart79.row = 1
	for heart78.row < 25 {
		heart78.row++
		fmt.Printf("\033[%d;%dH%s%c", heart78.row, heart78.column, heart78.color, heart78.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart78.row = 1
	for heart77.row < 25 {
		heart77.row++
		fmt.Printf("\033[%d;%dH%s%c", heart77.row, heart77.column, heart77.color, heart77.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart77.row = 1
	for heart76.row < 25 {
		heart76.row++
		fmt.Printf("\033[%d;%dH%s%c", heart76.row, heart76.column, heart76.color, heart76.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart76.row = 1
	for heart75.row < 25 {
		heart75.row++
		fmt.Printf("\033[%d;%dH%s%c", heart75.row, heart75.column, heart75.color, heart75.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart75.row = 1
	for heart74.row < 25 {
		heart74.row++
		fmt.Printf("\033[%d;%dH%s%c", heart74.row, heart74.column, heart74.color, heart74.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart74.row = 1
	for heart73.row < 25 {
		heart73.row++
		fmt.Printf("\033[%d;%dH%s%c", heart73.row, heart73.column, heart73.color, heart73.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart73.row = 1
	for heart72.row < 25 {
		heart72.row++
		fmt.Printf("\033[%d;%dH%s%c", heart72.row, heart72.column, heart72.color, heart72.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart72.row = 1
	for heart71.row < 25 {
		heart71.row++
		fmt.Printf("\033[%d;%dH%s%c", heart71.row, heart71.column, heart71.color, heart71.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart71.row = 1
	for heart70.row < 25 {
		heart70.row++
		fmt.Printf("\033[%d;%dH%s%c", heart70.row, heart70.column, heart70.color, heart70.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart70.row = 1
	for heart69.row < 25 {
		heart69.row++
		fmt.Printf("\033[%d;%dH%s%c", heart69.row, heart69.column, heart69.color, heart69.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart69.row = 1
	for heart68.row < 25 {
		heart68.row++
		fmt.Printf("\033[%d;%dH%s%c", heart68.row, heart68.column, heart68.color, heart68.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart68.row = 1
	for heart67.row < 25 {
		heart67.row++
		fmt.Printf("\033[%d;%dH%s%c", heart67.row, heart67.column, heart67.color, heart67.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart67.row = 1
	for heart66.row < 25 {
		heart66.row++
		fmt.Printf("\033[%d;%dH%s%c", heart66.row, heart66.column, heart66.color, heart66.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart66.row = 1
	for heart65.row < 25 {
		heart65.row++
		fmt.Printf("\033[%d;%dH%s%c", heart65.row, heart65.column, heart65.color, heart65.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart65.row = 1
	for heart64.row < 25 {
		heart64.row++
		fmt.Printf("\033[%d;%dH%s%c", heart64.row, heart64.column, heart64.color, heart64.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart64.row = 1
	for heart63.row < 25 {
		heart63.row++
		fmt.Printf("\033[%d;%dH%s%c", heart63.row, heart63.column, heart63.color, heart63.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart63.row = 1
	for heart62.row < 25 {
		heart62.row++
		fmt.Printf("\033[%d;%dH%s%c", heart62.row, heart62.column, heart62.color, heart62.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart62.row = 1
	for heart61.row < 25 {
		heart61.row++
		fmt.Printf("\033[%d;%dH%s%c", heart61.row, heart61.column, heart61.color, heart61.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart61.row = 1
	for heart60.row < 25 {
		heart60.row++
		fmt.Printf("\033[%d;%dH%s%c", heart60.row, heart60.column, heart60.color, heart60.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart60.row = 1
	for heart59.row < 25 {
		heart59.row++
		fmt.Printf("\033[%d;%dH%s%c", heart59.row, heart59.column, heart59.color, heart59.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart59.row = 1
	for heart58.row < 25 {
		heart58.row++
		fmt.Printf("\033[%d;%dH%s%c", heart58.row, heart58.column, heart58.color, heart58.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart58.row = 1
	for heart57.row < 25 {
		heart57.row++
		fmt.Printf("\033[%d;%dH%s%c", heart57.row, heart57.column, heart57.color, heart57.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart57.row = 1
	for heart56.row < 25 {
		heart56.row++
		fmt.Printf("\033[%d;%dH%s%c", heart56.row, heart56.column, heart56.color, heart56.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart56.row = 1
	for heart55.row < 25 {
		heart55.row++
		fmt.Printf("\033[%d;%dH%s%c", heart55.row, heart55.column, heart55.color, heart55.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart55.row = 1
	for heart54.row < 25 {
		heart54.row++
		fmt.Printf("\033[%d;%dH%s%c", heart54.row, heart54.column, heart54.color, heart54.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart54.row = 1
	for heart53.row < 25 {
		heart53.row++
		fmt.Printf("\033[%d;%dH%s%c", heart53.row, heart53.column, heart53.color, heart53.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart53.row = 1
	for heart52.row < 25 {
		heart52.row++
		fmt.Printf("\033[%d;%dH%s%c", heart52.row, heart52.column, heart52.color, heart52.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart52.row = 1
	for heart51.row < 25 {
		heart51.row++
		fmt.Printf("\033[%d;%dH%s%c", heart51.row, heart51.column, heart51.color, heart51.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart51.row = 1
	for heart50.row < 25 {
		heart50.row++
		fmt.Printf("\033[%d;%dH%s%c", heart50.row, heart50.column, heart50.color, heart50.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart50.row = 1
	for heart49.row < 25 {
		heart49.row++
		fmt.Printf("\033[%d;%dH%s%c", heart49.row, heart49.column, heart49.color, heart49.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart49.row = 1
	for heart48.row < 25 {
		heart48.row++
		fmt.Printf("\033[%d;%dH%s%c", heart48.row, heart48.column, heart48.color, heart48.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart48.row = 1
	for heart47.row < 25 {
		heart47.row++
		fmt.Printf("\033[%d;%dH%s%c", heart47.row, heart47.column, heart47.color, heart47.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart47.row = 1
	for heart46.row < 25 {
		heart46.row++
		fmt.Printf("\033[%d;%dH%s%c", heart46.row, heart46.column, heart46.color, heart46.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart46.row = 1
	for heart45.row < 25 {
		heart45.row++
		fmt.Printf("\033[%d;%dH%s%c", heart45.row, heart45.column, heart45.color, heart45.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart45.row = 1
	for heart44.row < 25 {
		heart44.row++
		fmt.Printf("\033[%d;%dH%s%c", heart44.row, heart44.column, heart44.color, heart44.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart44.row = 1
	for heart43.row < 25 {
		heart43.row++
		fmt.Printf("\033[%d;%dH%s%c", heart43.row, heart43.column, heart43.color, heart43.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart43.row = 1
	for heart42.row < 25 {
		heart42.row++
		fmt.Printf("\033[%d;%dH%s%c", heart42.row, heart42.column, heart42.color, heart42.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart42.row = 1
	for heart41.row < 25 {
		heart41.row++
		fmt.Printf("\033[%d;%dH%s%c", heart41.row, heart41.column, heart41.color, heart41.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart41.row = 1
	for heart40.row < 25 {
		heart40.row++
		fmt.Printf("\033[%d;%dH%s%c", heart40.row, heart40.column, heart40.color, heart40.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart40.row = 1
	for heart39.row < 25 {
		heart39.row++
		fmt.Printf("\033[%d;%dH%s%c", heart39.row, heart39.column, heart39.color, heart39.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart39.row = 1
	for heart38.row < 25 {
		heart38.row++
		fmt.Printf("\033[%d;%dH%s%c", heart38.row, heart38.column, heart38.color, heart38.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart38.row = 1
	for heart37.row < 25 {
		heart37.row++
		fmt.Printf("\033[%d;%dH%s%c", heart37.row, heart37.column, heart37.color, heart37.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart37.row = 1
	for heart36.row < 25 {
		heart36.row++
		fmt.Printf("\033[%d;%dH%s%c", heart36.row, heart36.column, heart36.color, heart36.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart36.row = 1
	for heart35.row < 25 {
		heart35.row++
		fmt.Printf("\033[%d;%dH%s%c", heart35.row, heart35.column, heart35.color, heart35.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart35.row = 1
	for heart34.row < 25 {
		heart34.row++
		fmt.Printf("\033[%d;%dH%s%c", heart34.row, heart34.column, heart34.color, heart34.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart34.row = 1
	for heart33.row < 25 {
		heart33.row++
		fmt.Printf("\033[%d;%dH%s%c", heart33.row, heart33.column, heart33.color, heart33.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart33.row = 1
	for heart32.row < 25 {
		heart32.row++
		fmt.Printf("\033[%d;%dH%s%c", heart32.row, heart32.column, heart32.color, heart32.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart32.row = 1
	for heart31.row < 25 {
		heart31.row++
		fmt.Printf("\033[%d;%dH%s%c", heart31.row, heart31.column, heart31.color, heart31.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart31.row = 1
	for heart30.row < 25 {
		heart30.row++
		fmt.Printf("\033[%d;%dH%s%c", heart30.row, heart30.column, heart30.color, heart30.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart30.row = 1
	for heart29.row < 25 {
		heart29.row++
		fmt.Printf("\033[%d;%dH%s%c", heart29.row, heart29.column, heart29.color, heart29.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart29.row = 1
	for heart28.row < 25 {
		heart28.row++
		fmt.Printf("\033[%d;%dH%s%c", heart28.row, heart28.column, heart28.color, heart28.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart28.row = 1
	for heart27.row < 25 {
		heart27.row++
		fmt.Printf("\033[%d;%dH%s%c", heart27.row, heart27.column, heart27.color, heart27.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart27.row = 1
	for heart26.row < 25 {
		heart26.row++
		fmt.Printf("\033[%d;%dH%s%c", heart26.row, heart26.column, heart26.color, heart26.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart26.row = 1
	for heart25.row < 25 {
		heart25.row++
		fmt.Printf("\033[%d;%dH%s%c", heart25.row, heart25.column, heart25.color, heart25.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart25.row = 1
	for heart24.row < 25 {
		heart24.row++
		fmt.Printf("\033[%d;%dH%s%c", heart24.row, heart24.column, heart24.color, heart24.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart24.row = 1
	for heart23.row < 25 {
		heart23.row++
		fmt.Printf("\033[%d;%dH%s%c", heart23.row, heart23.column, heart23.color, heart23.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart23.row = 1
	for heart22.row < 25 {
		heart22.row++
		fmt.Printf("\033[%d;%dH%s%c", heart22.row, heart22.column, heart22.color, heart22.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart22.row = 1
	for heart21.row < 25 {
		heart21.row++
		fmt.Printf("\033[%d;%dH%s%c", heart21.row, heart21.column, heart21.color, heart21.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart21.row = 1
	for heart20.row < 25 {
		heart20.row++
		fmt.Printf("\033[%d;%dH%s%c", heart20.row, heart20.column, heart20.color, heart20.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart20.row = 1
	for heart19.row < 25 {
		heart19.row++
		fmt.Printf("\033[%d;%dH%s%c", heart19.row, heart19.column, heart19.color, heart19.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart19.row = 1
	for heart18.row < 25 {
		heart18.row++
		fmt.Printf("\033[%d;%dH%s%c", heart18.row, heart18.column, heart18.color, heart18.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart18.row = 1
	for heart17.row < 25 {
		heart17.row++
		fmt.Printf("\033[%d;%dH%s%c", heart17.row, heart17.column, heart17.color, heart17.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart17.row = 1
	for heart16.row < 25 {
		heart16.row++
		fmt.Printf("\033[%d;%dH%s%c", heart16.row, heart16.column, heart16.color, heart16.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart16.row = 1
	for heart15.row < 25 {
		heart15.row++
		fmt.Printf("\033[%d;%dH%s%c", heart15.row, heart15.column, heart15.color, heart15.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart15.row = 1
	for heart14.row < 25 {
		heart14.row++
		fmt.Printf("\033[%d;%dH%s%c", heart14.row, heart14.column, heart14.color, heart14.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart14.row = 1
	for heart13.row < 25 {
		heart13.row++
		fmt.Printf("\033[%d;%dH%s%c", heart13.row, heart13.column, heart13.color, heart13.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart13.row = 1
	for heart12.row < 25 {
		heart12.row++
		fmt.Printf("\033[%d;%dH%s%c", heart12.row, heart12.column, heart12.color, heart12.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart12.row = 1
	for heart11.row < 25 {
		heart11.row++
		fmt.Printf("\033[%d;%dH%s%c", heart11.row, heart11.column, heart11.color, heart11.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart11.row = 1
	for heart10.row < 25 {
		heart10.row++
		fmt.Printf("\033[%d;%dH%s%c", heart10.row, heart10.column, heart10.color, heart10.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart10.row = 1
	for heart9.row < 25 {
		heart9.row++
		fmt.Printf("\033[%d;%dH%s%c", heart9.row, heart9.column, heart9.color, heart9.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart9.row = 1
	for heart8.row < 25 {
		heart8.row++
		fmt.Printf("\033[%d;%dH%s%c", heart8.row, heart8.column, heart8.color, heart8.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart8.row = 1
	for heart7.row < 25 {
		heart7.row++
		fmt.Printf("\033[%d;%dH%s%c", heart7.row, heart7.column, heart7.color, heart7.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart7.row = 1
	for heart6.row < 25 {
		heart6.row++
		fmt.Printf("\033[%d;%dH%s%c", heart6.row, heart6.column, heart6.color, heart6.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart6.row = 1
	for heart5.row < 25 {
		heart5.row++
		fmt.Printf("\033[%d;%dH%s%c", heart5.row, heart5.column, heart5.color, heart5.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart5.row = 1
	for heart4.row < 25 {
		heart4.row++
		fmt.Printf("\033[%d;%dH%s%c", heart4.row, heart4.column, heart4.color, heart4.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart4.row = 1
	for heart3.row < 25 {
		heart3.row++
		fmt.Printf("\033[%d;%dH%s%c", heart3.row, heart3.column, heart3.color, heart3.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart3.row = 1
	for heart2.row < 25 {
		heart2.row++
		fmt.Printf("\033[%d;%dH%s%c", heart2.row, heart2.column, heart2.color, heart2.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart2.row = 1
	for heart1.row < 25 {
		heart1.row++
		fmt.Printf("\033[%d;%dH%s%c", heart1.row, heart1.column, heart1.color, heart1.graphic)
		<-time.After(time.Duration(tickrate) * time.Nanosecond)
	}
	heart1.row = 1
}
func main() {
	fmt.Print("\033[2J")
	setup()
	for i := 0; i < 80; i++ {
		firstrowdraw()
		fall()
		heart80.column--
		heart79.column--
		heart78.column--
		heart77.column--
		heart76.column--
		heart75.column--
		heart74.column--
		heart73.column--
		heart72.column--
		heart71.column--
		heart70.column--
		heart69.column--
		heart68.column--
		heart67.column--
		heart66.column--
		heart65.column--
		heart64.column--
		heart63.column--
		heart62.column--
		heart61.column--
		heart60.column--
		heart59.column--
		heart58.column--
		heart57.column--
		heart56.column--
		heart55.column--
		heart54.column--
		heart53.column--
		heart52.column--
		heart51.column--
		heart50.column--
		heart49.column--
		heart48.column--
		heart47.column--
		heart46.column--
		heart45.column--
		heart44.column--
		heart43.column--
		heart42.column--
		heart41.column--
		heart40.column--
		heart39.column--
		heart38.column--
		heart37.column--
		heart36.column--
		heart35.column--
		heart34.column--
		heart33.column--
		heart32.column--
		heart31.column--
		heart30.column--
		heart29.column--
		heart28.column--
		heart27.column--
		heart26.column--
		heart25.column--
		heart24.column--
		heart23.column--
		heart22.column--
		heart21.column--
		heart20.column--
		heart19.column--
		heart18.column--
		heart17.column--
		heart16.column--
		heart15.column--
		heart14.column--
		heart13.column--
		heart12.column--
		heart11.column--
		heart10.column--
		heart9.column--
		heart8.column--
		heart7.column--
		heart6.column--
		heart5.column--
		heart4.column--
		heart3.column--
		heart2.column--
		heart1.column--
		//firstrowdraw()
		//fall()
	}
	for i := 0; i < 80; i++ {
		firstrowdraw()
		fall()
		heart80.column--
		heart79.column--
		heart78.column--
		heart77.column--
		heart76.column--
		heart75.column--
		heart74.column--
		heart73.column--
		heart72.column--
		heart71.column--
		heart70.column--
		heart69.column--
		heart68.column--
		heart67.column--
		heart66.column--
		heart65.column--
		heart64.column--
		heart63.column--
		heart62.column--
		heart61.column--
		heart60.column--
		heart59.column--
		heart58.column--
		heart57.column--
		heart56.column--
		heart55.column--
		heart54.column--
		heart53.column--
		heart52.column--
		heart51.column--
		heart50.column--
		heart49.column--
		heart48.column--
		heart47.column--
		heart46.column--
		heart45.column--
		heart44.column--
		heart43.column--
		heart42.column--
		heart41.column--
		heart40.column--
		heart39.column--
		heart38.column--
		heart37.column--
		heart36.column--
		heart35.column--
		heart34.column--
		heart33.column--
		heart32.column--
		heart31.column--
		heart30.column--
		heart29.column--
		heart28.column--
		heart27.column--
		heart26.column--
		heart25.column--
		heart24.column--
		heart23.column--
		heart22.column--
		heart21.column--
		heart20.column--
		heart19.column--
		heart18.column--
		heart17.column--
		heart16.column--
		heart15.column--
		heart14.column--
		heart13.column--
		heart12.column--
		heart11.column--
		heart10.column--
		heart9.column--
		heart8.column--
		heart7.column--
		heart6.column--
		heart5.column--
		heart4.column--
		heart3.column--
		heart2.column--
		heart1.column--
		//firstrowdraw()
		//fall()
	}
	for i := 0; i < 80; i++ {
		firstrowdraw()
		fall()
		heart80.column++
		heart79.column++
		heart78.column++
		heart77.column++
		heart76.column++
		heart75.column++
		heart74.column++
		heart73.column++
		heart72.column++
		heart71.column++
		heart70.column++
		heart69.column++
		heart68.column++
		heart67.column++
		heart66.column++
		heart65.column++
		heart64.column++
		heart63.column++
		heart62.column++
		heart61.column++
		heart60.column++
		heart59.column++
		heart58.column++
		heart57.column++
		heart56.column++
		heart55.column++
		heart54.column++
		heart53.column++
		heart52.column++
		heart51.column++
		heart50.column++
		heart49.column++
		heart48.column++
		heart47.column++
		heart46.column++
		heart45.column++
		heart44.column++
		heart43.column++
		heart42.column++
		heart41.column++
		heart40.column++
		heart39.column++
		heart38.column++
		heart37.column++
		heart36.column++
		heart35.column++
		heart34.column++
		heart33.column++
		heart32.column++
		heart31.column++
		heart30.column++
		heart29.column++
		heart28.column++
		heart27.column++
		heart26.column++
		heart25.column++
		heart24.column++
		heart23.column++
		heart22.column++
		heart21.column++
		heart20.column++
		heart19.column++
		heart18.column++
		heart17.column++
		heart16.column++
		heart15.column++
		heart14.column++
		heart13.column++
		heart12.column++
		heart11.column++
		heart10.column++
		heart9.column++
		heart8.column++
		heart7.column++
		heart6.column++
		heart5.column++
		heart4.column++
		heart3.column++
		heart2.column++
		heart1.column++
		//firstrowdraw()
		//fall()
	}
	for i := 0; i < 80; i++ {
		firstrowdraw()
		fall()
		heart80.column++
		heart79.column++
		heart78.column++
		heart77.column++
		heart76.column++
		heart75.column++
		heart74.column++
		heart73.column++
		heart72.column++
		heart71.column++
		heart70.column++
		heart69.column++
		heart68.column++
		heart67.column++
		heart66.column++
		heart65.column++
		heart64.column++
		heart63.column++
		heart62.column++
		heart61.column++
		heart60.column++
		heart59.column++
		heart58.column++
		heart57.column++
		heart56.column++
		heart55.column++
		heart54.column++
		heart53.column++
		heart52.column++
		heart51.column++
		heart50.column++
		heart49.column++
		heart48.column++
		heart47.column++
		heart46.column++
		heart45.column++
		heart44.column++
		heart43.column++
		heart42.column++
		heart41.column++
		heart40.column++
		heart39.column++
		heart38.column++
		heart37.column++
		heart36.column++
		heart35.column++
		heart34.column++
		heart33.column++
		heart32.column++
		heart31.column++
		heart30.column++
		heart29.column++
		heart28.column++
		heart27.column++
		heart26.column++
		heart25.column++
		heart24.column++
		heart23.column++
		heart22.column++
		heart21.column++
		heart20.column++
		heart19.column++
		heart18.column++
		heart17.column++
		heart16.column++
		heart15.column++
		heart14.column++
		heart13.column++
		heart12.column++
		heart11.column++
		heart10.column++
		heart9.column++
		heart8.column++
		heart7.column++
		heart6.column++
		heart5.column++
		heart4.column++
		heart3.column++
		heart2.column++
		heart1.column++
		//firstrowdraw()
		//fall()
	}
	for i := 0; i < 80; i++ {
		firstrowdraw()
		fall()
		heart80.column++
		heart79.column++
		heart78.column++
		heart77.column++
		heart76.column++
		heart75.column++
		heart74.column++
		heart73.column++
		heart72.column++
		heart71.column++
		heart70.column++
		heart69.column++
		heart68.column++
		heart67.column++
		heart66.column++
		heart65.column++
		heart64.column++
		heart63.column++
		heart62.column++
		heart61.column++
		heart60.column++
		heart59.column++
		heart58.column++
		heart57.column++
		heart56.column++
		heart55.column++
		heart54.column++
		heart53.column++
		heart52.column++
		heart51.column++
		heart50.column++
		heart49.column++
		heart48.column++
		heart47.column++
		heart46.column++
		heart45.column++
		heart44.column++
		heart43.column++
		heart42.column++
		heart41.column++
		heart40.column++
		heart39.column++
		heart38.column++
		heart37.column++
		heart36.column++
		heart35.column++
		heart34.column++
		heart33.column++
		heart32.column++
		heart31.column++
		heart30.column++
		heart29.column++
		heart28.column++
		heart27.column++
		heart26.column++
		heart25.column++
		heart24.column++
		heart23.column++
		heart22.column++
		heart21.column++
		heart20.column++
		heart19.column++
		heart18.column++
		heart17.column++
		heart16.column++
		heart15.column++
		heart14.column++
		heart13.column++
		heart12.column++
		heart11.column++
		heart10.column++
		heart9.column++
		heart8.column++
		heart7.column++
		heart6.column++
		heart5.column++
		heart4.column++
		heart3.column++
		heart2.column++
		heart1.column++
		//firstrowdraw()
		//fall()
	}
	for i := 0; i < 80; i++ {
		firstrowdraw()
		fall()
		heart80.column--
		heart79.column--
		heart78.column--
		heart77.column--
		heart76.column--
		heart75.column--
		heart74.column--
		heart73.column--
		heart72.column--
		heart71.column--
		heart70.column--
		heart69.column--
		heart68.column--
		heart67.column--
		heart66.column--
		heart65.column--
		heart64.column--
		heart63.column--
		heart62.column--
		heart61.column--
		heart60.column--
		heart59.column--
		heart58.column--
		heart57.column--
		heart56.column--
		heart55.column--
		heart54.column--
		heart53.column--
		heart52.column--
		heart51.column--
		heart50.column--
		heart49.column--
		heart48.column--
		heart47.column--
		heart46.column--
		heart45.column--
		heart44.column--
		heart43.column--
		heart42.column--
		heart41.column--
		heart40.column--
		heart39.column--
		heart38.column--
		heart37.column--
		heart36.column--
		heart35.column--
		heart34.column--
		heart33.column--
		heart32.column--
		heart31.column--
		heart30.column--
		heart29.column--
		heart28.column--
		heart27.column--
		heart26.column--
		heart25.column--
		heart24.column--
		heart23.column--
		heart22.column--
		heart21.column--
		heart20.column--
		heart19.column--
		heart18.column--
		heart17.column--
		heart16.column--
		heart15.column--
		heart14.column--
		heart13.column--
		heart12.column--
		heart11.column--
		heart10.column--
		heart9.column--
		heart8.column--
		heart7.column--
		heart6.column--
		heart5.column--
		heart4.column--
		heart3.column--
		heart2.column--
		heart1.column--
		//firstrowdraw()
		//fall()
	}
}
