package ubot

import "https://github.com/iamnolimit/tggomusicbot/pkg/vc/ntgcalls"

func (ctx *Context) Time(chatId any, streamMode ntgcalls.StreamMode) (uint64, error) {
	parsedChatId, err := ctx.parseChatId(chatId)
	if err != nil {
		return 0, err
	}
	return ctx.binding.Time(parsedChatId, streamMode)
}
