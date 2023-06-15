<script setup>
import HeadPortrait from '@/views/Chat/components/HeadPortrait.vue'
import { ref, watch, onMounted } from "vue";
import { addChatListAPI } from '@/apis/chat'

// const imgUrl = require("@/assets/img/head_portrait1.jpg");
const inputMsg = ref("");
const msgList = ref({});

const props = defineProps({
  chatList: {
    type: Object,
  }
})
// vue3使用watch监听父组件的参数
watch(
  () => props.chatList,
  (newVal, oldVal) => {
    // console.log(newVal, oldVal);
    msgList.value = newVal
    turnToBottom()
  }
)

const addChatList = async () => {
  const res = await addChatListAPI({
    you: `${msgList.value.Id}`,
    content: inputMsg.value,
  })
  // console.log(res);
}

const turnToBottom = () => {
  const chatContent = document.querySelector(".chat-content");
  chatContent.scrollTop = chatContent.scrollHeight;
}

const sendText = () => {
  // console.log("sendText");
  if (inputMsg.value === '') {
    return
  }
  addChatList()
  inputMsg.value = ''
  // setTimeout(() => {
  //   turnToBottom()
  // }, 1000);
}

onMounted(() => {
  setTimeout(() => {
    turnToBottom()
  }, 750);
})
</script>


<template>
  <div class="chat-window">
    <div class="top">
      <div class="head-pic">
        <HeadPortrait :imgUrl="msgList.Avatar"></HeadPortrait>
      </div>
      <div class="info-detail">
        <div class="name">{{ msgList.Nickname }}</div>
        <div class="detail">A human</div>
      </div>
    </div>
    <div class="botoom">
      <div class="chat-content" ref="chatContent">
        <div class="chat-wrapper" v-for="item in msgList.Chat">
          <div class="chat-friend" v-if="item.Type === '0'">
            <div class="chat-text">{{ item.Content }}</div>
          </div>
          <div class="chat-me" v-else>
            <div class="chat-text">{{ item.Content }}</div>
          </div>
        </div>
      </div>
      <div class="chatInputs">
        <input class="inputs" v-model="inputMsg" @keyup.enter="sendText" />
        <div class="send boxinput" @click="sendText">
          <img src="@/assets/img/emoji/rocket.png" alt="" />
        </div>
      </div>
    </div>
  </div>
</template>


<style lang="scss" scoped>
.chat-window {
  width: 100%;
  height: 100%;
  margin-left: 20px;
  position: relative;

  .top {
    margin-bottom: 50px;

    &::after {
      content: "";
      display: block;
      clear: both;
    }

    .head-pic {
      float: left;
    }

    .info-detail {
      float: left;
      margin: 5px 20px 0;

      .name {
        font-size: 20px;
        font-weight: 600;
        color: #000000;
      }

      .detail {
        color: #9e9e9e;
        font-size: 12px;
        margin-top: 2px;
      }
    }

    .other-fun {
      float: right;
      margin-top: 20px;

      span {
        margin-left: 30px;
        cursor: pointer;
      }

      // .icon-tupian {

      // }
      input {
        display: none;
      }
    }
  }

  .botoom {
    width: 100%;
    height: 70vh;
    background-color: rgb(255, 255, 255);
    border-radius: 20px;
    padding: 20px;
    box-sizing: border-box;
    position: relative;

    .chat-content {
      width: 100%;
      height: 85%;
      overflow-y: scroll;
      padding: 20px;
      box-sizing: border-box;


      &::-webkit-scrollbar {
        width: 0;
        /* Safari,Chrome 隐藏滚动条 */
        height: 0;
        /* Safari,Chrome 隐藏滚动条 */
        display: none;
        /* 移动端、pad 上Safari，Chrome，隐藏滚动条 */
      }

      .chat-wrapper {
        position: relative;
        word-break: break-all;

        .chat-friend {
          width: 100%;
          float: left;
          margin-bottom: 20px;
          display: flex;
          flex-direction: column;
          justify-content: flex-start;
          align-items: flex-start;

          .chat-text {
            max-width: 90%;
            padding: 20px;
            border-radius: 20px 20px 20px 5px;
            background-color: #F5F5F5;
            color: #000000;
            font-weight: 550;
            font-size: 15px;

            &:hover {
              background-color: #d6d6d6;
            }
          }

          .chat-img {
            img {
              width: 100px;
              height: 100px;
            }
          }

          .info-time {
            margin: 10px 0;
            color: #fff;
            font-size: 14px;

            img {
              width: 30px;
              height: 30px;
              border-radius: 50%;
              vertical-align: middle;
              margin-right: 10px;
            }

            span:last-child {
              color: rgb(101, 104, 115);
              margin-left: 10px;
              vertical-align: middle;
            }
          }
        }

        .chat-me {
          width: 100%;
          float: right;
          margin-bottom: 20px;
          position: relative;
          display: flex;
          flex-direction: column;
          justify-content: flex-end;
          align-items: flex-end;

          .chat-text {
            float: right;
            max-width: 90%;
            padding: 20px;
            border-radius: 20px 20px 5px 20px;
            background-color: #53d6ba;
            color: #fff;
            font-weight: 550;
            font-size: 15px;

            &:hover {
              background-color: #32bc9e;
            }
          }

          .chat-img {
            img {
              max-width: 300px;
              max-height: 200px;
              border-radius: 10px;
            }
          }

          .info-time {
            margin: 10px 0;
            color: #fff;
            font-size: 14px;
            display: flex;
            justify-content: flex-end;

            img {
              width: 30px;
              height: 30px;
              border-radius: 50%;
              vertical-align: middle;
              margin-left: 10px;
            }

            span {
              line-height: 30px;
            }

            span:first-child {
              color: rgb(101, 104, 115);
              margin-right: 10px;
              vertical-align: middle;
            }
          }
        }
      }
    }

    .chatInputs {
      width: 90%;
      position: absolute;
      bottom: 0;
      margin: 3%;
      display: flex;

      .boxinput {
        width: 50px;
        height: 50px;
        background-color: rgb(66, 70, 86);
        border-radius: 15px;
        border: 1px solid rgb(80, 85, 103);
        position: relative;
        cursor: pointer;

        img {
          width: 30px;
          height: 30px;
          position: absolute;
          left: 50%;
          top: 50%;
          transform: translate(-50%, -50%);
          border-radius: 5px;
        }
      }

      .inputs {
        width: 90%;
        height: 50px;
        background-color: #F5F5F5;
        border-radius: 15px;
        border: 2px solid #53d6ba;
        padding: 10px;
        box-sizing: border-box;
        transition: 0.2s;
        font-size: 20px;
        color: #fff;
        font-weight: 100;
        margin: 0 20px;

        &:focus {
          outline: none;
        }
      }

      .send {
        background-color: #53d6ba;
        border: 0;
        transition: 0.3s;
        box-shadow: 0px 0px 5px 0px #32bc9e;

        &:hover {
          box-shadow: 0px 0px 10px 0px #32bc9e;
        }
      }
    }
  }
}
</style>