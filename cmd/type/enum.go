package _type

type Finished string
type Type int16

const (
	Beat Type = iota
	Song
	Remix
	Voiceover
	Edit
)

func ProjectTypes() []Type {
	return []Type{Beat, Song, Remix, Voiceover, Edit}
}

func (t Type) String() string {
	switch t {
	case Beat:
		return "Beat"
	case Song:
		return "Song"
	case Remix:
		return "Remix"
	case Voiceover:
		return "Voiceover or Podcast"
	case Edit:
		return "Edit or Mashup"
	}
	return "unknown ProjectType"
}

func (t Type) FilterValue() string {
	return t.String()
}

func (t Type) Title() string {
	return t.String()
}
