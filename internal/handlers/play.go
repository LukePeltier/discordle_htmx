package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lukepeltier/discordle/ent"
	"github.com/lukepeltier/discordle/ent/message"
	"github.com/lukepeltier/discordle/internal/templates/play"
)

func PlayHandler(c echo.Context) error {
	// Choose a random message

	client := c.Get("entClient").(*ent.Client)
	ctx := context.TODO()

	query := "SELECT messages.ID FROM messages WHERE (length(trim(text)) - length(REPLACE(trim(text), ' ', ''))) > 5 AND length(REPLACE(trim(text), ' ', '')) > 9 ORDER BY random() LIMIT 1;"

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

	return render(c, play.Play(randomMessage, discordMember))
}
