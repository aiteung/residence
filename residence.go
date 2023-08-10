package residence

import (
	"fmt"
	"strings"

	"github.com/aiteung/atmessage"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

const Keyword string = "r3s1d3nce"

func HasKeyword(Info *types.MessageInfo, Message *waProto.Message) (whmsg bool) {
	if Message.ExtendedTextMessage != nil && Info.Chat.Server == "s.whatsapp.net" {
		if strings.Contains(*Message.ExtendedTextMessage.Text, Keyword) {
			if Message.ExtendedTextMessage.ContextInfo != nil {
				if *Message.ExtendedTextMessage.ContextInfo.EntryPointConversionSource == "click_to_chat_link" {
					whmsg = true
				}
			}
		}
	}
	return
}

func HandlerResidence(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, urlwauthreq string, urlwauthrole string, prefixurlapiwa string) {
	RunModule(waclient, Info, Message, urlwauthreq, prefixurlapiwa)
}

func RunModule(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, urlwauthreq string, prefixurlapiwa string) {
	var watemp WhatsappRequest
	watemp.Phonenumber = Info.Sender.User
	watemp.Name = Info.PushName
	watemp.Message = Message.GetConversation()
	watemp.Message = strings.Replace(*Message.ExtendedTextMessage.Text, Keyword, "", 1)
	watemp.Delay = *Message.ExtendedTextMessage.ContextInfo.EntryPointConversionDelaySeconds
	fmt.Println(watemp)
	atmessage.SendMessage(watemp.Message, Info.Sender, waclient)
}
