package pkg

import (
	"github.com/AshokShau/TgMusicBot/pkg/config"
	"github.com/AshokShau/TgMusicBot/pkg/handlers"
	"github.com/AshokShau/TgMusicBot/pkg/vc"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func Init(client *tg.Client) error {
	for _, session := range config.Conf.SessionStrings {
		_, err := vc.Calls.StartClient(config.Conf.ApiId, config.Conf.ApiHash, session)
		if err != nil {
			return err
		}
	}

	vc.Calls.RegisterHandlers(client)
	handlers.LoadModules(client)
	return nil
}
