package models

import (
	"time"
)

type Problem struct {
	Name       string    `json:"name"`
	Tag        string    `json:"tag"`
	URL        string    `json:"url"`
	EaseFactor float64   `json:"ease_rating"`
	Interval   int       `json:"interval"`
	Repetition int       `json:"repetition"`
	ReviewAt   time.Time `json:"review_at"`
}

func NewProblem(name string, tag string, url string) *Problem {
	return &Problem{
		Name:       name,
		Tag:        tag,
		URL:        url,
		EaseFactor: 2.5,
		Interval:   0,
		Repetition: 0,
		ReviewAt:   time.Now(),
	}
}

func GetDueProblems(problems []*Problem) []*Problem {
	var dueProblems []*Problem
	now := time.Now()
	for _, p := range problems {
		if p.ReviewAt.Before(now) || p.ReviewAt.Equal(now) {
			dueProblems = append(dueProblems, p)
		}
	}
	return dueProblems
}

func nextReviewTime(interval int) time.Time {
	now := time.Now()

	// round down to midnight
	midnight := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	return midnight.Add(time.Duration(interval) * 24 * time.Hour)
}

func Review(p *Problem, quality int) {
	if quality < 3 {
		p.Repetition = 0
		p.Interval = 1
		p.ReviewAt = nextReviewTime(p.Interval)
		return
	}

	switch p.Repetition {
	case 0:
		p.Interval = 1
	case 1:
		p.Interval = 6
	default:
		p.Interval = int(float64(p.Interval) * p.EaseFactor)
	}

	p.Repetition++

	ef := p.EaseFactor + (0.1 - float64(5-quality)*(0.08+float64(5-quality)*0.02))

	if ef < 1.3 {
		ef = 1.3
	}

	p.EaseFactor = ef
	p.ReviewAt = nextReviewTime(p.Interval)
}
