package api

type Role struct {
	guildID string
	roleID  string
}

type Permission int

func (c *Connection) GetRole(guildID, roleID string) {}

func (g *Guild) GetName() {}

func (g *Guild) SetColor() {}

func (g *Guild) SetName() {}

func (g *Guild) SetPermission() {}

func (g *Guild) CheckPermission() {}

func (g *Guild) SetMentionable() {}

func getRole(guildID, roleID string) {
	// Get role by id
}
