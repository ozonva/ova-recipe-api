package recipe

import "fmt"

type Recipe struct {
	id          uint64
	userId      uint64
	name        string
	description string
	actions     []Action
}

func New(id uint64, userId uint64, name string, description string, actions []Action) Recipe {
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

func (r *Recipe) Cook() error {
	for idx, _ := range r.actions {
		if err := r.actions[idx].DoAction(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Recipe) String() string {
	return fmt.Sprintf("Recipe('%s')", r.name)
}
