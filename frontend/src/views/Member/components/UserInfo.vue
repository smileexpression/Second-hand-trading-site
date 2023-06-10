<script setup>
import { ref, onMounted } from "vue";
import { useUserStore } from "@/stores/user";
import { getLikeListAPI } from "@/apis/likelist";
import { ElMessage } from "element-plus";
import GoodsItem from "@/views/Home/components/GoodsItem.vue";
import { getImageUrl } from "@/apis/image";
import { useRouter } from "vue-router";


const userStore = useUserStore()
const token = userStore.userInfo.token
// const imageUrl = ref(userStore.userInfo.avatar)
// const imageUrl = ref('http://localhost:5173/92274fb1-9901-41b6-852b-1f70da06a6f0')
const uploadUrl = 'http://localhost:8080/image/upload'
const image = ref()

const likeList = ref([])

const getLikeList = async () => {
  const res = await getLikeListAPI({ limit: 4 })
  likeList.value = res.result
}

onMounted(() => getLikeList())

//上传头像标头
const headersObj = { 'Authorization': `Bearer ${token}` }

//上传文件类型限制
const beforeUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  if (!isImage) {
    ElMessage.error('只能上传图片文件')
  }
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    ElMessage.error('上传图片文件大小不能超过 2MB')
  }
  return isImage && isLt2M
}

//上传头像成功
const updateSuccess = (res, upload) => {
  //获取图片ID
  image.value = res.imageIds
  console.log(image);

  //用图片ID更改头像
  userStore.updateAvatar(image.value)
}

//上传头像失败
const updateError = (res, upload) => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
}

//安全管理修改密码
const passwordDialogVisible = ref(false)
//表单
const passwordForm = ref({
  oldpassword: '',
  newpassword: '',
  confirmed: ''
})
//密码规则
const passwordRules = ({
  oldpassword: [
    { required: true, message: '密码不能为空！', trigger: 'blur' },
    { min: 6, max: 14, message: '密码长度应为6~14个字符！', trigger: 'blur' }
  ],
  newpassword: [
    { required: true, message: '密码不能为空！', trigger: 'blur' },
    { min: 6, max: 14, message: '密码长度应为6~14个字符！', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if(value === passwordForm.value.oldpassword){
          callback(new Error('新的密码不能与输入的旧密码相同！'))
        } else{
          callback()
        }
      }
    }
  ],
  confirmed: [
    {required: true, message: '请再次确认密码！', trigger: 'blur'},
    {
      validator: (rule, value, callback) => {
        if(value === passwordForm.value.newpassword){
          callback()
        } else{
          callback(new Error('与密码不匹配！'))
        }
      }
    }
  ]
})

//获取密码表单实例
const passwordFormRef = ref(null)
const router = useRouter()
//确认修改密码
const doChange = () => {
  const { oldpassword, newpassword } = passwordForm.value
  passwordFormRef.value.validate(async (valid) => {
    if(valid){
      await userStore.changePassword({ oldpassword, newpassword })
      ElMessage({ type: 'success', message: '修改成功' })
      //退出登录
      userStore.clearUserInfo()
      router.push('/login')
    }
  })
}

</script>

<template>
  <div class="home-overview">
    <!-- 用户信息 -->
    <div class="user-meta">
      <div class="avatar">
        <el-upload :action="uploadUrl" :headers="headersObj"
        :show-file-list="false" :before-upload="beforeUpload" :on-success="updateSuccess" :on-error="updateError">
          <img :src="getImageUrl(userStore.userInfo.avatar)" />
        </el-upload>
      </div>
      <h4>{{ userStore.userInfo.nickname }}</h4>
    </div>
    <div class="item">
      <a href="javascript:;" @click="passwordDialogVisible = true">
        <span class="iconfont icon-aq"></span>
        <p>密码管理</p>
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

  <!-- 更改密码对话框 -->
  <el-dialog v-model="passwordDialogVisible" title="更改密码" width="450">
    <div class="form">
      <el-form ref="passwordFormRef" :model="passwordForm" :rules="passwordRules" label-position="right" label-width="80px" status-icon>
        <el-form-item prop="oldpassword" label="原密码">
          <el-input show-password v-model="passwordForm.oldpassword" />
        </el-form-item>
        <el-form-item prop="newpassword" label="新密码">
          <el-input show-password v-model="passwordForm.newpassword" />
        </el-form-item>
        <el-form-item prop="confirmed" label="确认密码">
          <el-input show-password v-model="passwordForm.confirmed" />
        </el-form-item>
        <div style="margin:0 auto; text-align: center;">
            <el-button size="large" class="subBtn" @click="doChange">确认更改</el-button>
            <el-button size="large" class="subBtn" @click="passwordDialogVisible = false">取消</el-button>
        </div>
      </el-form>
    </div>
  </el-dialog>
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

.form {
    padding: 0 20px 20px 20px;

    &-item {
      margin-bottom: 28px;

      .input {
        position: relative;
        height: 36px;

        >i {
          width: 34px;
          height: 34px;
          background: #cfcdcd;
          color: #fff;
          position: absolute;
          left: 1px;
          top: 1px;
          text-align: center;
          line-height: 34px;
          font-size: 18px;
        }

        input {
          padding-left: 44px;
          border: 1px solid #cfcdcd;
          height: 36px;
          line-height: 36px;
          width: 100%;

          &.error {
            border-color: $priceColor;
          }

          &.active,
          &:focus {
            border-color: $xtxColor;
          }
        }

        .code {
          position: absolute;
          right: 1px;
          top: 1px;
          text-align: center;
          line-height: 34px;
          font-size: 14px;
          background: #f5f5f5;
          color: #666;
          width: 90px;
          height: 34px;
          cursor: pointer;
        }
      }

      >.error {
        position: absolute;
        font-size: 12px;
        line-height: 28px;
        color: $priceColor;

        i {
          font-size: 14px;
          margin-right: 2px;
        }
      }
    }

    .agree {
      a {
        color: #069;
      }
    }

    .btn {
      display: block;
      width: 100%;
      height: 40px;
      color: #fff;
      text-align: center;
      line-height: 40px;
      background: $xtxColor;

      &.disabled {
        background: #cfcdcd;
      }
    }
  }

  .subBtn {
  background: $xtxColor;
  width: 40%;
  color: #fff;
}
</style>