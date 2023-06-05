<script setup>
import PersonCard from '@/views/Chat/components/PersonCard.vue'
import ChatWindow from '@/views/Chat/components/chatwindow.vue'
import { ref, onMounted, onUnmounted } from "vue";
import { getChatListAPI } from "@/apis/chat";

const showChatWindow = ref(false)
const chatList = ref([])
const msgList = ref({})
const timer = ref()
const index = ref(-1)

const getChatList = async () => {
  const res = await getChatListAPI()
  chatList.value = res.result
}

const curPerson = (key) => {
  index.value = key
  // if (msgList.value?.Chat.length !== chatList.value?.[index.value]?.Chat.length) {
  //   console.log("change");
  //   msgList.value = chatList.value[index.value]
  // }
  // else {
  //   console.log("no change");
  // }
  msgList.value = chatList.value[index.value]
}

const clickPerson = (key) => {
  showChatWindow.value = true
  console.log("click")
  curPerson(key)
}

onMounted(() => {
  msgList.value = { Chat: ["test"] }
  timer.value = setInterval(() => {
    getChatList()
    curPerson(index.value)
  }, 500)
})

onUnmounted(() => {
  clearInterval(timer.value)
  timer.value = 0
})
</script>

<template>
  <div class="chatHome">
    <div class="chatLeft">
      <div class="title">
        <h1>聊天</h1>
      </div>
      <div class="online-person">
        <span class="onlin-text">聊天列表</span>
        <div class="person-cards-wrapper">
          <div class="personList" v-for="item, key in chatList" @click="clickPerson(key)">
            <PersonCard :imgUrl="item.Avatar" :nickname="item.Nickname"
              :lastMsg="item.Chat?.[item.Chat?.length - 1].Content">
            </PersonCard>
          </div>
        </div>
      </div>
    </div>
    <div class="chatRight">
      <div v-if="showChatWindow">
        <ChatWindow :chatList="msgList"></ChatWindow>
      </div>
      <div class="showIcon" v-else>
        <span class="iconfont icon-snapchat"></span>
        <!-- <img src="@/assets/img/snapchat.png" alt="" /> -->
      </div>
    </div>
    <!-- <el-col :span="4"><div class="grid-content bg-purple"></div></el-col> -->
  </div>
</template>



<style lang="scss" scoped>
.chatHome {
  // margin-top: 20px;
  display: flex;

  .chatLeft {
    width: 280px;

    .title {
      color: #fff;
      padding-left: 10px;
    }

    .online-person {
      margin-top: 100px;

      .onlin-text {
        padding-left: 10px;
        color: rgb(176, 178, 189);
      }

      .person-cards-wrapper {
        padding-left: 10px;
        height: 65vh;
        margin-top: 20px;
        overflow: hidden;
        overflow-y: scroll;
        box-sizing: border-box;

        &::-webkit-scrollbar {
          width: 0;
          /* Safari,Chrome 隐藏滚动条 */
          height: 0;
          /* Safari,Chrome 隐藏滚动条 */
          display: none;
          /* 移动端、pad 上Safari，Chrome，隐藏滚动条 */
        }
      }
    }
  }

  .chatRight {
    flex: 1;
    padding-right: 30px;

    .showIcon {
      position: absolute;
      top: calc(50% - 150px);
      /*垂直居中 */
      left: calc(50% - 50px);

      /*水平居中 */
      .icon-snapchat {
        width: 300px;
        height: 300px;
        font-size: 300px;
        // color: rgb(28, 30, 44);
      }
    }
  }
}
</style>