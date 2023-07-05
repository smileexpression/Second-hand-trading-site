<script setup>
import { getCheckInfoAPI, createOrderAPI } from '@/apis/checkout'
import { useRouter, useRoute } from 'vue-router'
import { ref, onMounted } from 'vue'
import { useCartStore } from '@/stores/cartStore';
import { getImageUrl } from '@/apis/image';
import { useUserStore } from "@/stores/user";
import { ElMessage } from "element-plus";
const cartStore = useCartStore()
const userStore = useUserStore()
const router = useRouter()
const route = useRoute()
const checkInfo = ref({})  // 订单对象
const curAddress = ref({})
const getCheckInfo = async () => {
  const i = route.query.goodID
  console.log("router id", i, typeof (i))
  const res = await getCheckInfoAPI(i)
  checkInfo.value = res.result
  console.log("checkInfo", checkInfo.value)
  // const item = checkInfo.value.userAddresses.find(item => item.isDefault === 0)
  // curAddress.value = item
}
onMounted(() => getCheckInfo())

// 控制弹框打开
const showDialog = ref(false)

// 切换地址
const activeAddress = ref({})
const switchAddress = (item) => {
  activeAddress.value = item
}
const confirm = () => {
  curAddress.value = activeAddress.value
  showDialog.value = false
  activeAddress.value = {}
}

// 添加地址
const addressDialogVisible = ref(false)
const addressForm = ref({
  receiver: '',
  contact: '',
  address: ''
})
const addressRules = ref({
  receiver: [
    { required: true, message: '收件人不能为空！', trigger: 'blur' }
  ],
  contact: [
    { required: true, message: '手机号不能为空！', trigger: 'blur' },
    { type: "string", len: 11, pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: '请输入正确的手机号！', trigger: 'blur' }
  ],
  address: [
    { required: true, message: '收货地址不能为空！', trigger: 'blur' }
  ]
})
//表单实例
const addressFormRef = ref(null)
const addAddress = () => {
  const { receiver, contact, address } = addressForm.value
  addressFormRef.value.validate(async (valid) => {
    if (valid) {
      await userStore.addAddress({ receiver, contact, address })
      ElMessage({ type: 'success', message: '操作成功' })
    }
    // console.log(address_last)
    checkInfo.value.userAddresses = userStore.userInfo.userAddresses
    // checkInfo.value.userAddresses.push({ address_last[address_last.length - 1], receiver, contact, address })
    addressFormRef.value.resetFields()
    addressDialogVisible.value = false
  })
}
//取消重置表单
const cancelAddress = () => {
  addressFormRef.value.resetFields()
  addressDialogVisible.value = false
}

// 创建订单
const createOrder = async () => {
  if (curAddress.value.id === undefined) {
    alert("请先选择收货地址")
    return
  }
  // console.log("checkInfo id", checkInfo.value.goods.id)
  // console.log("address id", curAddress.value.id)
  const res = await createOrderAPI({
    goodId: checkInfo.value.goods.id,
    addressId: curAddress.value.id
  })
  // console.log("createOrder", res)
  const orderId = res.id
  router.push({
    path: 'pay',
    query: {
      id: orderId
    }
  })
  // 更新购物车
  cartStore.delCart([checkInfo.value.goods.id])
  cartStore.updateCart()
}

</script>

<template>
  <div class="xtx-pay-checkout-page">
    <div class="container">
      <div class="wrapper">
        <!-- 收货地址 -->
        <h3 class="box-title">收货地址</h3>
        <div class="box-body">
          <div class="address">
            <div class="text">
              <div class="none" v-if="!curAddress">您需要先选择收货地址才可提交订单。</div>
              <ul v-else>
                <li><span>收<i />货<i />人：</span>{{ curAddress.receiver }}</li>
                <li><span>联系方式：</span>{{ curAddress.contact }}</li>
                <li><span>收货地址：</span>{{ curAddress.fullLocation }} {{ curAddress.address }}</li>
              </ul>
            </div>
            <div class="action">
              <el-button size="large" @click="showDialog = true">选择地址</el-button>
              <el-button size="large" @click="addressDialogVisible = true">添加地址</el-button>
            </div>
          </div>
        </div>
        <!-- 商品信息 -->
        <h3 class="box-title">商品信息</h3>
        <div class="box-body">
          <table class="goods">
            <thead>
              <tr>
                <th width="520">商品信息</th>
                <th width="170">实付</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>
                  <a href="javascript:;" class="info">
                    <img :src="getImageUrl(checkInfo.goods?.picture)" alt="">
                    <div class="right">
                      <p>{{ checkInfo.goods?.name }}</p>
                      <p>{{ checkInfo.goods?.desc }}</p>
                    </div>
                  </a>
                </td>
                <td>&yen;{{ checkInfo.goods?.price }}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <!-- 支付方式 -->
        <h3 class="box-title">支付方式</h3>
        <div class="box-body">
          <a class="my-btn active" href="javascript:;">在线支付</a>
          <!-- <a class="my-btn" href="javascript:;">货到付款</a>
            <span style="color:#999">货到付款需付5元手续费</span> -->
        </div>
        <!-- 金额明细 -->
        <h3 class="box-title">金额明细</h3>
        <div class="box-body">
          <div class="total">
            <dl>
              <dt>商品件数：</dt>
              <dd>1件</dd>
            </dl>
            <dl>
              <dt>应付总额：</dt>
              <dd class="price">&yen;{{ checkInfo.goods?.price }}</dd>
            </dl>
          </div>
        </div>
        <!-- 提交订单 -->
        <div class="submit">
          <el-button @click="createOrder" type="primary" size="large">提交订单</el-button>
        </div>
      </div>
    </div>
  </div>
  <!-- 切换地址 -->
  <el-dialog v-model="showDialog" title="切换收货地址" width="30%" center>
    <div class="addressWrapper">
      <div class="text item" :class="{ active: activeAddress.id === item.id }" @click="switchAddress(item)"
        v-for="item in checkInfo.userAddresses" :key="item.id">
        <ul>
          <li><span>收<i />货<i />人：</span>{{ item.receiver }} </li>
          <li><span>联系方式：</span>{{ item.contact }}</li>
          <li><span>收货地址：</span>{{ item.fullLocation }} {{ item.address }}</li>
        </ul>
      </div>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="showDialog = false">取消</el-button>
        <el-button type="primary" @click="confirm">确定</el-button>
      </span>
    </template>
  </el-dialog>
  <!-- 添加地址 -->
  <!-- 添加收货地址对话框 -->
  <el-dialog v-model="addressDialogVisible" title="添加收货地址" width="500" @close="cancelAddress">
    <div class="form">
      <el-form ref="addressFormRef" :model="addressForm" :rules="addressRules" label-position="right" label-width="80px"
        status-icon>
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
.xtx-pay-checkout-page {
  margin-top: 20px;

  .wrapper {
    background: #fff;
    padding: 0 20px;

    .box-title {
      font-size: 16px;
      font-weight: normal;
      padding-left: 10px;
      line-height: 70px;
      border-bottom: 1px solid #f5f5f5;
    }

    .box-body {
      padding: 20px 0;
    }
  }
}

.address {
  border: 1px solid #f5f5f5;
  display: flex;
  align-items: center;

  .text {
    flex: 1;
    min-height: 90px;
    display: flex;
    align-items: center;

    .none {
      line-height: 90px;
      color: #999;
      text-align: center;
      width: 100%;
    }

    >ul {
      flex: 1;
      padding: 20px;

      li {
        line-height: 30px;

        span {
          color: #999;
          margin-right: 5px;

          >i {
            width: 0.5em;
            display: inline-block;
          }
        }
      }
    }

    >a {
      color: $xtxColor;
      width: 160px;
      text-align: center;
      height: 90px;
      line-height: 90px;
      border-right: 1px solid #f5f5f5;
    }
  }

  .action {
    width: 420px;
    text-align: center;

    .btn {
      width: 140px;
      height: 46px;
      line-height: 44px;
      font-size: 14px;

      &:first-child {
        margin-right: 10px;
      }
    }
  }
}

.goods {
  width: 100%;
  border-collapse: collapse;
  border-spacing: 0;

  .info {
    display: flex;
    text-align: left;

    img {
      width: 70px;
      height: 70px;
      margin-right: 20px;
    }

    .right {
      line-height: 24px;

      p {
        &:last-child {
          color: #999;
        }
      }
    }
  }

  tr {
    th {
      background: #f5f5f5;
      font-weight: normal;
    }

    td,
    th {
      text-align: center;
      padding: 20px;
      border-bottom: 1px solid #f5f5f5;

      &:first-child {
        border-left: 1px solid #f5f5f5;
      }

      &:last-child {
        border-right: 1px solid #f5f5f5;
      }
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

.subBtn {
  background: $xtxColor;
  width: 40%;
  color: #fff;
}

.my-btn {
  width: 228px;
  height: 50px;
  border: 1px solid #e4e4e4;
  text-align: center;
  line-height: 48px;
  margin-right: 25px;
  color: #666666;
  display: inline-block;

  &.active,
  &:hover {
    border-color: $xtxColor;
  }
}

.total {
  dl {
    display: flex;
    justify-content: flex-end;
    line-height: 50px;

    dt {
      i {
        display: inline-block;
        width: 2em;
      }
    }

    dd {
      width: 240px;
      text-align: right;
      padding-right: 70px;

      &.price {
        font-size: 20px;
        color: $priceColor;
      }
    }
  }
}

.submit {
  text-align: right;
  padding: 60px;
  border-top: 1px solid #f5f5f5;
}

.addressWrapper {
  max-height: 500px;
  overflow-y: auto;
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
</style>