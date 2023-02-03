package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/lcr2000/mchat/internal/config"
	"github.com/lcr2000/mchat/internal/model"
	"github.com/lcr2000/mchat/internal/utils"
	"github.com/spf13/cobra"
)

// ListOnlineCommand list cmd struct
type ListOnlineCommand struct {
	BaseCommand
}

// Init ListCommand
func (c *ListOnlineCommand) Init() {
	c.command = &cobra.Command{
		Use:     "list online",
		Short:   "List online users",
		Long:    "List online users",
		Aliases: []string{"lo"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.run(cmd, args)
		},
		Example: c.example(),
	}
	c.command.DisableFlagsInUseLine = true
}

func (c *ListOnlineCommand) run(command *cobra.Command, args []string) error {
	users, err := c.getOnlineUsers()
	if err != nil {
		return err
	}
	var count int
	var table [][]string
	for _, user := range users {
		conTmp := []string{user.UID, user.City, utils.TimeFormat(user.LastActiveTs)}
		count++
		table = append(table, conTmp)
	}
	err = utils.PrintTable(table, []string{"Name", "City", "LatActiveTime"})
	if err != nil {
		return err
	}
	utils.PrintKV(os.Stdout, "[Summary] ", map[string]interface{}{
		"online_count": count,
	})
	return nil
}

func (c *ListOnlineCommand) getOnlineUsers() ([]*model.User, error) {
	bytes, err := utils.HTTPGet(fmt.Sprintf("http://%s:%s/get_online_users", config.Cfg.Address, config.Cfg.HTTPPort))
	if err != nil {
		return nil, err
	}
	resp := &model.HTTPResponse{}
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != model.ErrCodeSuccess {
		return nil, errors.New(resp.Msg)
	}
	var user []*model.User
	err = json.Unmarshal([]byte(resp.Data.(string)), &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *ListOnlineCommand) example() string {
	return `
# List online users
mchat list online
# Aliases
mchat lo
`
}
