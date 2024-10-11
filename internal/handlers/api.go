package handlers

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lukepeltier/discordle/ent"
	"github.com/lukepeltier/discordle/ent/message"
)

type RandomMessageResponse struct {
	Message *ent.Message
	Sender  *ent.DiscordMember
}

func RandomMessageAPIHandler(c echo.Context) error {
	client := c.Get("entClient").(*ent.Client)
	ctx := c.Request().Context()

	query := "SELECT messages.ID FROM messages join message_scores on messages.id=message_scores.message_score_message join discord_members on messages.discord_member_messages=discord_members.id WHERE (length(trim(text)) - length(REPLACE(trim(text), ' ', ''))) > 5 AND length(REPLACE(trim(text), ' ', '')) > 9 and message_scores.score > 1 and discord_members.blacklisted=0 and NOT EXISTS (select 1 from blacklists where messages.text like '%' || blacklists.bad || '%') ORDER BY random() LIMIT 1;"

	rows, err := client.QueryContext(ctx, query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error querying db: [%s]", err.Error()))
	}
	defer rows.Close()

	var messageID int

	for rows.Next() {
		if err := rows.Scan(&messageID); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error parsing row return: [%s]", err.Error()))

		}
		break
	}

	randomMessage, err := client.Message.Query().Where(message.ID(messageID)).Only(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting message: [%s]", err.Error()))

	}
	discordMember, err := randomMessage.QuerySender().Only(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting sender: [%s]", err.Error()))

	}

	response := RandomMessageResponse{
		Message: randomMessage,
		Sender:  discordMember,
	}

	return c.JSON(http.StatusOK, response)
}

type MessageGuessResponse struct {
	Correct bool `json:"correct"`
}

type MessageGuessRequest struct {
	Guess string `json:"guess"`
}

func MessageGuessAPIHandler(c echo.Context) error {
	client := c.Get("entClient").(*ent.Client)
	ctx := c.Request().Context()
	messageIDStr := c.Param("id")
	guessRequest := new(MessageGuessRequest)
	if err := c.Bind(guessRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Bad request: [%s]", err.Error()))
	}
	usernameGuess := guessRequest.Guess

	messageID, err := strconv.Atoi(messageIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Error parsing message id: [%s]", err.Error()))
	}

	randomMessage, err := client.Message.Query().Where(message.ID(messageID)).Only(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Message [%s] not found: [%s]", messageIDStr, err.Error()))

	}

	discordMember, err := randomMessage.QuerySender().Only(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting sender: [%s]", err.Error()))

	}

	log.Printf("Guess: [%s]\n", usernameGuess)

	responseObj := MessageGuessResponse{
		Correct: slices.Contains(discordMember.Nicknames, usernameGuess),
	}

	return c.JSON(http.StatusOK, responseObj)

}
