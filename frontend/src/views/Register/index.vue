<script setup>
import { ref } from 'vue'
import 'element-plus/es/components/message/style/css'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user';
import { useRouter } from 'vue-router';

//注册表单对象
const registerForm = ref({
  account: '',
  password: '',
  confirmedPassword: '',
  nickname: '',
  gender: ''
})

//注册规则对象
const registerRules = ref({
  account:[
    {required: true, message: '账号不能为空！', trigger: 'blur'},
    {type: "string", len:11, pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: '请输入正确的手机号！', trigger: 'blur'}
  ],
  password:[
    {required: true, message: '密码不能为空！', trigger: 'blur'},
    {min: 6, max: 14, message: '密码长度应为6~14个字符！', trigger: 'blur'}
  ],
  confirmedPassword:[
    {required: true, message: '请再次确认密码！', trigger: 'blur'},
    {
      validator: (rule, value, callback) => {
        if(value === registerForm.value.password){
          callback()
        } else{
          callback(new Error('与密码不匹配！'))
        }
      }
    }
  ],
  nickname:[
    {required: true, message: '用户名不能为空！', trigger: 'blur'},
    {type: "string", max: 20, message: '用户名应为长度至多为20的字符！', trigger: 'blur'}
  ],
  gender:[
    {
      validator: (rule, value, callback) => {
        if(value === ''){
          callback(new Error('请选择性别！'))
        } else{
          callback()
        }
      }
    }
  ]
})

//获取用户实例
const userStore = useUserStore()

//获取注册表单实例进行统一检验
const registerFormRef = ref(null)

const router = useRouter()
//注册操作
const doRegister = () => {
    const {account, password, nickname, gender} = registerForm.value
    registerFormRef.value.validate( async (valid) => {
        if(valid){
            //通过用户实例调用注册接口
            await userStore.registerUser({account, password, nickname, gender})
            //消息提示
            ElMessage({ type:'success', message: '注册成功'})
            //跳转至主页
            router.replace('/')
        }
    })
}


</script>


<template>
    <div>
      <header class="login-header">
        <div class="container m-top-20">
          <h1 class="logo">
            <RouterLink to="/">小兔鲜</RouterLink>
          </h1>
          <RouterLink class="entry" to="/">
            进入网站首页
            <i class="iconfont icon-angle-right"></i>
            <i class="iconfont icon-angle-right"></i>
          </RouterLink>
        </div>
      </header>
      
      <section class="login-section">
        <div class="wrapper">
          <nav>
            <a href="javascript:;">用户注册</a>
          </nav>
          <div class="account-box">
            <div class="form">
              <el-form ref="registerFormRef" :model="registerForm" :rules="registerRules" label-position="right" label-width="70px"
                status-icon>
                <el-form-item prop="account" label="账户">
                  <el-input v-model="registerForm.account" />
                </el-form-item>
                <el-form-item prop="password" label="密码">
                  <el-input v-model="registerForm.password" />
                </el-form-item>
                <el-form-item prop="confirmedPassword" label="确认密码">
                  <el-input v-model="registerForm.confirmedPassword" />
                </el-form-item>
                <el-form-item prop="nickname" label="用户名">
                  <el-input v-model="registerForm.nickname" />
                </el-form-item>
                <el-form-item prop="gender" label="性别" >
                  <el-radio v-model="registerForm.gender" label="男">男</el-radio>
                  <el-radio v-model="registerForm.gender" label="女">女</el-radio>
                </el-form-item>
                <div style="text-align: right;font-size: smaller;" >
                  <el-link @click="$router.push('/login')">已有账号，立即登录</el-link>
                </div>
                <el-button size="large" class="subBtn" @click="doRegister()">点击注册</el-button>
              </el-form>
            </div>
          </div>
        </div>
      </section>
  
      <footer class="login-footer">
        <div class="container">
          <p>
            <a href="javascript:;">关于我们</a>
            <a href="javascript:;">帮助中心</a>
            <a href="javascript:;">售后服务</a>
            <a href="javascript:;">配送与验收</a>
            <a href="javascript:;">商务合作</a>
            <a href="javascript:;">搜索推荐</a>
            <a href="javascript:;">友情链接</a>
          </p>
          <p>CopyRight &copy; 小兔鲜儿</p>
        </div>
      </footer>
    </div>
  </template>
  
  <style scoped lang='scss'>
  .login-header {
    background: #fff;
    border-bottom: 1px solid #e4e4e4;
  
    .container {
      display: flex;
      align-items: flex-end;
      justify-content: space-between;
    }
  
    .logo {
      width: 200px;
  
      a {
        display: block;
        height: 132px;
        width: 100%;
        text-indent: -9999px;
        background: url("@/assets/images/logo.png") no-repeat center 18px / contain;
      }
    }
  
    .sub {
      flex: 1;
      font-size: 24px;
      font-weight: normal;
      margin-bottom: 38px;
      margin-left: 20px;
      color: #666;
    }
  
    .entry {
      width: 120px;
      margin-bottom: 38px;
      font-size: 16px;
  
      i {
        font-size: 14px;
        color: $xtxColor;
        letter-spacing: -5px;
      }
    }
  }
  
  .login-section {
    background: url('@/assets/images/login-bg.png') no-repeat center / cover;
    height: 488px;
    position: relative;
  
    .wrapper {
      width: 380px;
      background: #fff;
      position: absolute;
      left: 50%;
      top: 54px;
      transform: translate3d(100px, 0, 0);
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.15);
  
      nav {
        font-size: 14px;
        height: 55px;
        margin-bottom: 20px;
        border-bottom: 1px solid #f5f5f5;
        display: flex;
        padding: 0 40px;
        text-align: right;
        align-items: center;
  
        a {
          flex: 1;
          line-height: 1;
          display: inline-block;
          font-size: 18px;
          position: relative;
          text-align: center;
        }
      }
    }
  }
  
  .login-footer {
    padding: 30px 0 50px;
    background: #fff;
  
    p {
      text-align: center;
      color: #999;
      padding-top: 20px;
  
      a {
        line-height: 1;
        padding: 0 10px;
        color: #999;
        display: inline-block;
  
        ~a {
          border-left: 1px solid #ccc;
        }
      }
    }
  }
  
  .account-box {
    .toggle {
      padding: 15px 40px;
      text-align: right;
  
      a {
        color: $xtxColor;
  
        i {
          font-size: 14px;
        }
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
  
    .action {
      padding: 20px 40px;
      display: flex;
      justify-content: space-between;
      align-items: center;
  
      .url {
        a {
          color: #999;
          margin-left: 10px;
        }
      }
    }
  }
  
  .subBtn {
    background: $xtxColor;
    width: 100%;
    color: #fff;
  }
  </style>