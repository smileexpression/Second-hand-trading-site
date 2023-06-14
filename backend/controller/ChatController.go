package controller

import (
	"gin/common"
	"gin/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type chatDto struct {
	Type    string
	Content string
}

func ToChatDto(chat model.Chat) chatDto {
	return chatDto{
		Type:    chat.Type,
		Content: chat.Content,
	}
}

type single struct {
	Id       string
	Nickname string
	Avatar   string
	Chat     []chatDto
}

func GetMsg(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取当前用户的id
	user, _ := ctx.Get("user")
	userinfo := user.(model.User)
	id := userinfo.ID

	var chatList []model.ChatList
	DB.Table("chat_lists").Where("me = ?", id).Find(&chatList)

	var list []single
	for i := 0; i < len(chatList); i++ {
		var chat []model.Chat
		DB.Table("chats").Where("me = ? and you = ?", id, chatList[i].You).Find(&chat)
		// fmt.Println(chat)
		var chatDto []chatDto
		for j := 0; j < len(chat); j++ {
			chatDto = append(chatDto, ToChatDto(chat[j]))
		}
		var temp_user model.User
		DB.Table("users").Where("id = ?", chatList[i].You).First(&temp_user)
		newSingle := single{Id: chatList[i].You, Nickname: temp_user.Name, Avatar: temp_user.Avatar, Chat: chatDto}
		list = append(list, newSingle)
	}

	ctx.JSON(200, gin.H{
		"result": list,
	})
}

func SendMsg(ctx *gin.Context) {
	DB := common.GetDB()
	user, _ := ctx.Get("user")
	userinfo := user.(model.User)

	var chat model.Chat
	ctx.BindJSON(&chat)

	chat.Me = strconv.Itoa(int(userinfo.ID))
	chat.Type = "1"
	newChat := model.Chat{Me: chat.You, You: chat.Me, Type: "0", Content: chat.Content}
	DB.Create(&chat)
	DB.Create(&newChat)

	ctx.JSON(200, gin.H{
		"result": "suc",
	})
}

func AddChat(ctx *gin.Context) {
	DB := common.GetDB()
	user, _ := ctx.Get("user")
	userinfo := user.(model.User)

	var chatList, check model.ChatList
	ctx.BindJSON(&chatList)
	// 将string转为uint
	if chatList.You == strconv.Itoa(int(userinfo.ID)) {
		ctx.JSON(200, gin.H{
			"result": "bug",
		})
		return
	}
	DB.Table("chat_lists").Where("me = ? and you = ?", userinfo.ID, chatList.You).First(&check)
	if check.ID != 0 {
		ctx.JSON(200, gin.H{
			"result": "no need to add chat",
		})
		return
	}

	chatList.Me = strconv.Itoa(int(userinfo.ID))
	newChatList := model.ChatList{Me: chatList.You, You: chatList.Me}
	DB.Create(&chatList)
	DB.Create(&newChatList)
	ctx.JSON(200, gin.H{
		"result": "succeed in adding chat",
	})
}
