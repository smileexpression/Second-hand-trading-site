<script setup>
import { getOrderAPI } from '@/apis/pay'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
// 获取订单数据
const route = useRoute()
const orderInfo = ref({})
const getOrderInfo = async () => {
  const res = await getOrderAPI(route.query.orderId)
  console.log("orderInfo", res)
  orderInfo.value = res
}
onMounted(() => getOrderInfo())
</script>


<template>
  <div class="xtx-pay-page">
    <div class="container">
      <!-- 支付结果 -->
      <div class="pay-result">
        <!-- 路由参数获取的是字符串 -->
        <span class="iconfont icon-queren2 green" v-if="$route.query.payResult === 'true'"></span>
        <span class="iconfont icon-shanchu red" v-else></span>
        <p class="tit">支付{{ $route.query.payResult === 'true' ? '成功' : '失败' }}</p>
        <p class="tip">将尽快为您发货，收货期间请保持手机畅通</p>
        <!-- <p>支付方式：<span>支付宝</span></p> -->
        <p>支付金额：<span>¥{{ orderInfo.payMoney }}</span></p>
        <div class="btn">
          <el-button type="primary" style="margin-right:20px">
            <RouterLink to="/member/order">查看订单</RouterLink>
          </el-button>
          <el-button>
            <RouterLink to="/">回到首页</RouterLink>
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.pay-result {
  padding: 100px 0;
  background: #fff;
  text-align: center;
  margin-top: 20px;

  >.iconfont {
    font-size: 100px;
  }

  .green {
    color: #1dc779;
  }

  .red {
    color: $priceColor;
  }

  .tit {
    font-size: 24px;
  }

  .tip {
    color: #999;
  }

  p {
    line-height: 40px;
    font-size: 16px;
  }

  .btn {
    margin-top: 50px;
  }

  .alert {
    font-size: 12px;
    color: #999;
    margin-top: 50px;
  }
}
</style>