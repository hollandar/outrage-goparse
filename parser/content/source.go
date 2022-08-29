package content

type Source struct {
	memory   []rune
	Line     int
	Column   int
	Position int
}

func NewSource(source string) *Source {
	var result Source
	result.Position = 0
	result.Line = 0
	result.Column = 0
	result.memory = []rune(source)

	return &result
}

func (s *Source) Constrain(constrainLength int) *Source {
	var result Source
	result.memory = s.memory[:constrainLength]
	result.Line = s.Line
	result.Column = s.Column
	result.Position = s.Position

	return &result
}

func (s *Source) Copy() *Source {
	var result Source
	result.memory = []rune(s.memory)
	result.Line = s.Line
	result.Column = s.Column
	result.Position = s.Position

	return &result
}

func (s *Source) Clone() *Source {
	var result Source
	result.memory = s.memory
	result.Line = s.Line
	result.Column = s.Column
	result.Position = s.Position

	return &result
}

func (s *Source) Length() int {
	return len(s.memory) - s.Position
}

func (s *Source) Advance(length int) {
	var slice = s.memory[s.Position : s.Position+length]
	for _, element := range slice {
		s.Column += 1
		if element == '\n' {
			s.Column = 0
			s.Line += 1
		}
	}

	s.Position += length
}

func (s *Source) Slice(start int, length int) []rune {
	return s.memory[start : start+length]
}

func (s *Source) Take(length int) []rune {
	return s.Slice(s.Position, length)
}

func (s Source) AtPosition(count int) string {
	endPosition := s.Length()
	if s.Position+count < s.Length() {
		endPosition = s.Position + count
	}
	return string(s.memory[s.Position:endPosition])
}
