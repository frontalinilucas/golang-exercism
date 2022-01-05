package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

const (
	win         = "win"
	draw        = "draw"
	loss        = "loss"
	writeFormat = "%-30s | %2v | %2v | %2v | %2v | %2v\n"
)

var ErrMalformedLine = errors.New("malformed line")

type Team struct {
	name                            string
	won, drawn, lost, total, points int
}

func Tally(reader io.Reader, writer io.Writer) error {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, reader)
	if err != nil {
		return err
	}
	lines := strings.Split(buf.String(), "\n")
	table := map[string]*Team{}
	for i := range lines {
		s, valid, err := parseLine(lines[i])
		if err != nil {
			return err
		}
		if !valid {
			continue
		}
		process(table, s)
	}

	sortedTable := getSortedTable(table)
	var sb strings.Builder
	sb.WriteString(getLine("Team", "MP", "W", "D", "L", "P"))
	for _, team := range sortedTable {
		sb.WriteString(
			getLine(
				team.name,
				strconv.Itoa(team.total),
				strconv.Itoa(team.won),
				strconv.Itoa(team.drawn),
				strconv.Itoa(team.lost),
				strconv.Itoa(team.points),
			),
		)
	}

	_, err = writer.Write([]byte(sb.String()))
	if err != nil {
		return err
	}
	return nil
}

func getLine(team, mp, w, d, l, p string) string {
	return fmt.Sprintf(writeFormat, team, mp, w, d, l, p)
}

func parseLine(line string) ([]string, bool, error) {
	s := strings.Split(line, ";")
	if len(s) == 1 {
		return nil, false, nil
	}
	if len(s) != 3 {
		return nil, false, ErrMalformedLine
	}
	result := s[2]
	switch result {
	case win, draw, loss:
		return s, true, nil
	default:
		return nil, false, ErrMalformedLine
	}
}

func process(table map[string]*Team, s []string) {
	nameTeam1 := s[0]
	nameTeam2 := s[1]
	result := s[2]
	team1, ok := table[nameTeam1]
	if !ok {
		team1 = &Team{name: nameTeam1}
		table[nameTeam1] = team1
	}
	team2, ok := table[nameTeam2]
	if !ok {
		team2 = &Team{name: nameTeam2}
		table[nameTeam2] = team2
	}

	switch result {
	case draw:
		team1.matchDrawn()
		team2.matchDrawn()
	case win:
		team1.matchWon()
		team2.matchLost()
	default:
		team1.matchLost()
		team2.matchWon()
	}
}

func getSortedTable(table map[string]*Team) []*Team {
	result := make([]*Team, len(table))
	i := 0
	for _, t := range table {
		result[i] = t
		i++
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].points == result[j].points {
			if result[i].total == result[j].total {
				if result[i].won == result[j].won {
					return result[i].name < result[j].name
				}
				return result[i].won > result[j].won
			}
			return result[i].total > result[j].total
		}
		return result[i].points > result[j].points
	})

	return result
}

func (t *Team) matchWon() {
	t.won++
	t.total++
	t.points += 3
}

func (t *Team) matchLost() {
	t.lost++
	t.total++
}

func (t *Team) matchDrawn() {
	t.drawn++
	t.total++
	t.points += 1
}
