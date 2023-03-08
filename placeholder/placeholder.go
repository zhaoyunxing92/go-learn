package placeholder

type Placeholder interface {
	Resolver(text string) string
}
