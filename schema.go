package textsplitter

type Document struct {
	PageContent string
	Metadata    map[string]any
	Score       float32
}
