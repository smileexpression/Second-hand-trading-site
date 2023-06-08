<script setup>
import { ref, onMounted } from "vue";
import { useUserStore } from "@/stores/user";
import { getLikeListAPI } from "@/apis/likelist";
import { ElMessage } from "element-plus";
import GoodsItem from "@/views/Home/components/GoodsItem.vue";

const userStore = useUserStore()
const token = userStore.userInfo.token
// const imageUrl = ref(userStore.userInfo.avatar)
const imageUrl = ref('http://localhost:5173/92274fb1-9901-41b6-852b-1f70da06a6f0')

const likeList = ref([])

const getLikeList = async () => {
  const res = await getLikeListAPI({ limit: 4 })
  likeList.value = res.result
}

onMounted(() => getLikeList())

//上传头像标头
const headersObj = { 'Authorization': `Bearer ${token}` }

//上传文件类型限制
const beforeUpdate = (rawFile) => {
  if (rawFile.type !== 'image/png' && rawFile.type !== 'image/jpeg') {
    ElMessage({
      type: 'error',
      message: '请选择png或jpeg格式文件'
    })
    return false
  }
  return true
}

//上传头像成功
const updateSuccess = (res, upload) => {
  ElMessage({
    type: 'success',
    message: '上传成功'
  })
  imageUrl.value = URL.createObjectURL(upload.raw)
  console.log(imageUrl.value, "image");
}

//上传头像失败
const updateError = (res, upload) => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
}

//安全管理修改密码
const dialogVisible = ref(false)
//表单
const userForm = ref({
  password: '',
  confirmed: ''

})

</script>

<template>
  <div class="home-overview">
    <!-- 用户信息 -->
    <div class="user-meta">
      <div class="avatar">
        <el-upload action="https://mock.apifox.cn/m1/2726765-0-default/member/updateavatar" :headers="headersObj"
          :show-file-list="false" :before-upload="beforeUpdate" :on-success="updateSuccess" :on-error="updateError"
          v-model="userStore.userInfo.avatar">
          <img :src="imageUrl" />
        </el-upload>
      </div>
      <h4>{{ userStore.userInfo.nickname }}</h4>
    </div>
    <div class="item">
      <a href="javascript:;">
        <span class="iconfont icon-aq"></span>
        <p>安全设置</p>
      </a>
      <a href="javascript:;">
        <span class="iconfont icon-dw"></span>
        <p>地址管理</p>
      </a>
    </div>
  </div>
  <div class="like-container">
    <div class="home-panel">
      <div class="header">
        <h4 data-v-bcb266e0="">猜你喜欢</h4>
      </div>
      <div class="goods-list">
        <GoodsItem v-for="good in likeList" :key="good.id" :goods="good" />
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.home-overview {
  height: 132px;
  background: url(@/assets/images/center-bg.png) no-repeat center / cover;
  display: flex;

  .user-meta {
    flex: 1;
    display: flex;
    align-items: center;

    .avatar {
      width: 85px;
      height: 85px;
      border-radius: 50%;
      overflow: hidden;
      margin-left: 60px;

      img {
        width: 100%;
        height: 100%;
      }
    }

    h4 {
      padding-left: 26px;
      font-size: 18px;
      font-weight: normal;
      color: white;
    }
  }

  .item {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-around;

    &:first-child {
      border-right: 1px solid #f4f4f4;
    }

    a {
      color: white;
      font-size: 16px;
      text-align: center;

      .iconfont {
        font-size: 32px;
      }

      p {
        line-height: 32px;
      }
    }
  }
}

.like-container {
  margin-top: 20px;
  border-radius: 4px;
  background-color: #fff;
}

.home-panel {
  background-color: #fff;
  padding: 0 20px;
  margin-top: 20px;
  height: 400px;

  .header {
    height: 66px;
    border-bottom: 1px solid #f5f5f5;
    padding: 18px 0;
    display: flex;
    justify-content: space-between;
    align-items: baseline;

    h4 {
      font-size: 22px;
      font-weight: 400;
    }

  }

  .goods-list {
    display: flex;
    justify-content: space-around;
  }
}
</style>