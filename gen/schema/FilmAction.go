package schema

type FilmActionTrait struct {

}

type FilmAction struct {
	MetaTrait
	FilmActionTrait
	CreateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewFilmAction() (x FilmAction) {
	x.Type = "http://schema.org/FilmAction"

	return
}
