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
//取消重置表单
const cancelChange = () => {
  passwordFormRef.value.resetFields()
  passwordDialogVisible.value = false
}

//个人信息管理
const infoDialogVisible = ref(false)
const infoForm = ref({
  account: `${userStore.userInfo.account}`,
  nickname: `${userStore.userInfo.nickname}`,
  gender: `${userStore.userInfo.gender}`
})
const infoRules = ref({
  nickname: [
    {required: true, message: '用户名不能为空！', trigger: 'blur'},
    {type: "string", max: 20, message: '用户名应为长度至多为20的字符！', trigger: 'blur'}
  ]
})
//表单实例
const infoFormRef = ref(null)
const doInfo = () => {
  const { nickname, gender } = infoForm.value
  infoFormRef.value.validate(async (valid) => {
    if(valid){
      await userStore.changeInfo({ nickname, gender })
      ElMessage({ type: 'success', message: '修改成功' })
    }
  })
}
//取消重置表单
const cancelInfo = () => {
  infoFormRef.value.resetFields()
  infoDialogVisible.value = false
}

//添加地址
const addressDialogVisible = ref(false)
const addressForm = ref({
  receiver: '',
  contact: '',
  address: ''
})
const addressRules = ref({
  receiver:[
    {required: true, message: '收件人不能为空！', trigger: 'blur'}
  ],
  contact:[
    { required: true, message: '手机号不能为空！', trigger: 'blur' },
    { type: "string", len: 11, pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: '请输入正确的手机号！', trigger: 'blur' }
  ],
  address:[
    {required: true, message: '收货地址不能为空！', trigger: 'blur'}
  ]
})
//表单实例
const addressFormRef = ref(null)
const addAddress = () => {
  const { receiver, contact, address } = addressForm.value
  addressFormRef.value.validate(async (valid) => {
    if(valid){
      await userStore.addAddress({ receiver, contact, address })
      ElMessage({ type: 'success', message: '操作成功' })
    }
  })
}
//取消重置表单
const cancelAddress = () => {
  addressFormRef.value.resetFields()
  addressDialogVisible.value = false
}

//删除地址
const deleteAddress = async (id) => {
  await userStore.deladdress(id)
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
      <h4 @click="infoDialogVisible = true">{{ userStore.userInfo.nickname }}</h4>
    </div>
    <div class="item">
      <a href="javascript:;" @click="passwordDialogVisible = true">
        <span class="iconfont icon-aq"></span>
        <p>密码管理</p>
      </a>
      <a href="javascript:;" @click="addressDialogVisible = true">
        <span class="iconfont icon-dw"></span>
        <p>添加地址</p>
      </a>
    </div>
  </div>
  <div class="holder-container" v-if="userStore.userInfo.userAddresses === null">
    <el-empty description="暂无地址数据" />
  </div>
  <div v-else class="like-container">
    <div class="header">
      <h4 style="font-size: 22px; font-weight: 400; padding: 18px ;">收货地址</h4>
    </div>
    <div class="addressWrapper">
      <div class="text item" v-for="item in userStore.userInfo.userAddresses" :key="item.id">
        <ul>
          <li><span>收<i />货<i />人：</span>{{ item.receiver }} </li>
          <li><span>联系方式：</span>{{ item.contact }}</li>
          <li><span>收货地址：</span>{{ item.address }}</li>
        </ul>
        <div style="margin:auto; position: absolute; right: 30px;">
          <el-button size="large" class="subBtn" style="width:100px;" @click="deleteAddress(item.id)">删除地址</el-button>
        </div>
      </div>
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
  <el-dialog v-model="passwordDialogVisible" title="更改密码" width="450" @close="cancelChange">
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
            <el-button size="large" class="subBtn" @click="cancelChange">取消</el-button>
        </div>
      </el-form>
    </div>
  </el-dialog>

  <!-- 修改个人信息对话框 -->
  <el-dialog v-model="infoDialogVisible" title="修改个人信息" width="450" @close="cancelInfo">
    <div class="form">
      <el-form ref="infoFormRef" :model="infoForm" :rules="infoRules" label-position="right" label-width="70px" status-icon>
        <el-form-item prop="account" label="账号">
          <el-input v-model="infoForm.account" disabled />
        </el-form-item>
        <el-form-item prop="nickname" label="用户名">
          <el-input v-model="infoForm.nickname" />
        </el-form-item>
        <el-form-item prop="gender" label="性别">
          <el-radio v-model="infoForm.gender" label="男">男</el-radio>
          <el-radio v-model="infoForm.gender" label="女">女</el-radio>
        </el-form-item>
        <div style="margin:0 auto; text-align: center;">
            <el-button size="large" class="subBtn" @click="doInfo">确认更改</el-button>
          <el-button size="large" class="subBtn" @click="cancelInfo">取消</el-button>
        </div>
      </el-form>
    </div>
  </el-dialog>

  <!-- 添加收货地址对话框 -->
  <el-dialog v-model="addressDialogVisible" title="添加收货地址" width="500" @close="cancelAddress">
    <div class="form">
      <el-form ref="addressFormRef" :model="addressForm" :rules="addressRules" label-position="right" label-width="80px" status-icon>
        <el-form-item prop="receiver" label="收货人">
          <el-input v-model="addressForm.receiver" />
        </el-form-item>
        <el-form-item prop="contact" label="手机号">
          <el-input v-model="addressForm.contact" />
        </el-form-item>
        <el-form-item prop="address" label="收货地址">
          <el-input v-model="addressForm.address" />
        </el-form-item>
        <div style="margin:0 auto; text-align: center;">
            <el-button size="large" class="subBtn" @click="addAddress">确认添加</el-button>
          <el-button size="large" class="subBtn" @click="cancelAddress">取消</el-button>
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

.addressWrapper {
  max-height: 1000px;
  overflow-y: auto;
  position: relative;
}

.text {
  flex: 1;
  min-height: 90px;
  display: flex;
  align-items: center;

  &.item {
    border: 1px solid #f5f5f5;
    margin-bottom: 10px;
    cursor: pointer;

    &.active,
    &:hover {
      border-color: $xtxColor;
      background: lighten($xtxColor, 50%);
    }

    >ul {
      padding: 10px;
      font-size: 14px;
      line-height: 30px;
    }
  }
}

.holder-container {
      min-height: 500px;
      display: flex;
      justify-content: center;
      align-items: center;
    }
</style>