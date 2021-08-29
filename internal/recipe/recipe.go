package recipe

import "fmt"

type Recipe struct {
	id          uint64
	userId      uint64
	name        string
	description string
	actions     []string
}

func New(id uint64, userId uint64, name string, description string, actions []string) Recipe {
	return Recipe{id: id, userId: userId, name: name, description: description, actions: actions}
}

func (r *Recipe) Id() uint64 {
	return r.id
}

func (r *Recipe) UserId() uint64 {
	return r.userId
}

func (r *Recipe) Name() string {
	return r.name
}

func (r *Recipe) Description() string {
	return r.description
}

func (r *Recipe) Actions() []string {
	return r.actions
}

func (r *Recipe) String() string {
	return fmt.Sprintf("Recipe('%s')", r.name)
}
