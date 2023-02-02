package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lcr2000/mchat/internal/client"
	"github.com/lcr2000/mchat/internal/model"
	"github.com/lcr2000/mchat/internal/utils"
	"github.com/spf13/cobra"
)

// ChatCommand chat cmd struct
type ChatCommand struct {
	BaseCommand
}

// Init ChatCommand
func (c *ChatCommand) Init() {
	c.command = &cobra.Command{
		Use:     "chat",
		Short:   "Chat with other",
		Long:    "Chat with other",
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.runChat(cmd, args)
		},
		Example: chatExample(),
	}
	c.command.DisableFlagsInUseLine = true
}

func (c *ChatCommand) runChat(command *cobra.Command, args []string) error {
	fmt.Println("Enter the server address.")
	address := utils.PromptUI("address", "127.0.0.1")

	var username string

	for {
		fmt.Println("Enter your username.")
		username = utils.PromptUI("username", "")
		err := c.login(address, username)
		if err == nil {
			break
		}
		utils.PrintWarning(os.Stdout, fmt.Sprintf("%s.\n", err.Error()))
	}

	client.Dial(address, username)
	fmt.Println("Connect to the server.")

	err := c.enter(username)
	if err != nil {
		return err
	}

	utils.PrintYellow(os.Stdout, username)
	utils.PrintString(os.Stdout, " Login succeeded!\n")

	return c.process()
}

func (c *ChatCommand) login(address, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req := struct {
		Username string `json:"username"`
	}{
		Username: username,
	}
	marshal, _ := json.Marshal(req)
	bytes, err := utils.HTTPPostWithContext(ctx, fmt.Sprintf("http://%s:8080/login", address), "application/json", string(marshal))
	if err != nil {
		return err
	}
	resp := &model.HTTPResponse{}
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return err
	}
	switch resp.Code {
	case model.ErrCodeSuccess:
		return nil
	case model.ErrCodeAccountExist:
		return errors.New("account exist")
	default:
		return errors.New("try again")
	}
}

func (c *ChatCommand) enter(username string) error {
	packet := model.BuildClientPacket(model.CmdChatEnter, username)
	marshal, _ := json.Marshal(packet)
	_, err := client.GetClientConn().GetConn().Write(marshal)
	return err
}

func (c *ChatCommand) process() error {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // Read user input.
		inputInfo := strings.Trim(input, "\r\n")
		if inputInfo == "" {
			continue
		}
		if strings.ToUpper(inputInfo) == "q" { // Exit if enter q.
			return nil
		}
		packet := model.BuildClientPacket(model.CmdChat, inputInfo)
		marshal, _ := json.Marshal(packet)
		_, err := client.GetClientConn().GetConn().Write(marshal)
		if err != nil {
			return err
		}
	}
}

func chatExample() string {
	return `
# Chat with other
mchat chat
# Aliases
mchat c
`
}
