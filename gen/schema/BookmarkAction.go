package schema

type BookmarkActionTrait struct {

}

type BookmarkAction struct {
	MetaTrait
	BookmarkActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewBookmarkAction() (x BookmarkAction) {
	x.Type = "http://schema.org/BookmarkAction"

	return
}
