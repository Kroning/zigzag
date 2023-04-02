package lib

import (
	"time"
	"fmt"
)

const (
	name = "ZigZag"
)

var (
	ErrorNotEnoughCandles	error = fmt.Errorf("%s: Not enough candles", name)
	ErrorMinCandles			error = fmt.Errorf("%s: minCandles must be more than 1", name)
)

type Candle struct {
	Date	time.Time
	High	float64
	Low		float64
}

type Swing struct {
	Date time.Time
	Price	float64
	Peak	bool // Peak (true) or valley (false)
}

func GetSwings(candles []*Candle, percent float64, minCandles int) ([]*Swing, error) {
	if len(candles) < minCandles {
		return nil, ErrorNotEnoughCandles
	}
	if minCandles<2 {
		return nil, ErrorMinCandles
	}

	// Initial peak/valley
	swings := []*Swing{}
	var lowPast, highPast float64
	pos := minCandles-1
	peak := false
	curSwing := &Swing{}
	for pos<len(candles) {
		low, high := candles[pos].Low, candles[pos].High
		for i:=0; i<pos; i++ {
			lowPast, highPast = candles[i].Low, candles[i].High

			// Peak to valley
			if (highPast - low)/highPast*100.0 >= percent {
				swing1 := &Swing{ candles[i].Date, highPast, true}	// Peak
				curSwing = &Swing{ candles[pos].Date, low, false} // Valley
				swings = append(swings, swing1)
				peak = false
				lowPast = low
				break
			}
			// Valley to peak
            if (high - lowPast)/lowPast*100.0 >= percent {
                swing1 := &Swing{ candles[i].Date, lowPast, false}  // Valley
                curSwing = &Swing{ candles[pos].Date, high, true} // Peak
				swings = append(swings, swing1)
                peak = true
				highPast = high
                break
            }
		}
		if len(swings) > 0 {
			break
		}
		pos++
	}
	if len(swings) == 0 {
		return swings, nil
	}

	lastPos := pos
	for pos<len(candles) {
		low, high := candles[pos].Low, candles[pos].High
		if peak {
			if	high >= highPast { // Peak higher
				curSwing = &Swing{ candles[pos].Date, high, true}
				highPast = high
				lastPos = pos
			} else if (highPast - low)/highPast*100.0 >= percent && (pos-lastPos+1) >= minCandles { // Changing from peak to valley
				swings = append(swings, curSwing)
				curSwing = &Swing{ candles[pos].Date, low, false}
				peak = false
				lastPos = pos
				lowPast = low
			}
		} else {
            if	low <= lowPast { // Valley lower
                curSwing = &Swing{ candles[pos].Date, low, false}
                lastPos = pos
                lowPast = low
            } else if (high - lowPast)/lowPast*100.0 >= percent	&& (pos-lastPos+1) >= minCandles { // Valley -> peak
                swings = append(swings, curSwing)				
				curSwing = &Swing{ candles[pos].Date, high, true}
				peak = true
                highPast = high
                lastPos = pos
			}
		}
		pos++
	}
	swings = append(swings, curSwing)
	return swings, nil 
}






