package mention

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrParseFailed = errors.New("failed to to parse @ mentions from the provided input")
)

type User struct {
	ID       string
	NameID   string
	HandleID string
}

type Parser struct {
	regex *regexp.Regexp
}

func NewParser() (*Parser, error) {
	regex, err := regexp.Compile(`<@([A-Z0-9]+)\|(.+)>`)
	if err != nil {
		return nil, fmt.Errorf("failed to set up @ mention parser: %w", err)
	}
	return &Parser{regex}, nil
}

func (p *Parser) Parse(input string) (User, error) {
	matches := p.regex.FindStringSubmatch(input)
	if matches == nil {
		return User{}, ErrParseFailed
	}

	if len(matches) != 3 {
		return User{}, ErrParseFailed
	}

	return User{
		ID:       matches[1],
		NameID:   matches[2],
		HandleID: input,
	}, nil
}

func (p *Parser) ParseList(input string) ([]User, error) {
	rawmentions := strings.Split(input, " ")

	// ensure we have at least one, but not too many
	if len(rawmentions) < 1 || len(rawmentions) >= 10 {
		return nil, ErrParseFailed
	}

	var users []User
	for _, mention := range rawmentions {
		user, err := p.Parse(mention)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
