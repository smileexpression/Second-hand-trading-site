<script setup>
import { ref, onMounted } from 'vue'
import { getUserOrder } from '@/apis/order';
import { getImageUrl } from '@/apis/image';


// 订单列表
const orderList = ref([])
const total =ref(0)

const params = ref({
    page: 0,
    pageSize: 2
})

const getOrderList = async () => {
  const res = await getUserOrder(params.value)
  orderList.value = res.result.item
  total.value = res.result.counts
}

onMounted(() => getOrderList())

//页数切换
const pageChange = (page) => {
  params.value.page = page
  getOrderList()
}

</script>

<template>
  <div class="order-container">
      <div class="main-container">
        <div class="holder-container" v-if="orderList.length === 0">
          <el-empty description="暂无订单数据" />
        </div>
        <div v-else>
          <!-- 订单列表 -->
          <div class="order-item" v-for="order in orderList" :key="order.id">
            <div class="head">
              <span>下单时间：{{ order.createTime }}</span>
              <span>订单编号：{{ order.id }}</span>
            </div>
            <div class="body">
              <div class="column goods">
                <ul>
                  <li :key="order.skus.id">
                    <a class="image" href="javascript:;">
                      <img :src="getImageUrl(order.skus.image)" alt="" />
                    </a>
                    <div class="info">
                      <p class="name ellipsis-2">
                        {{ order.skus.name }}
                      </p>
                      <p class="attr ellipsis">
                        <span>{{ order.skus.attrsText }}</span>
                      </p>
                    </div>
                    <div class="price">¥{{ order.skus.realPay?.toFixed(2) }}</div>
                    <div class="count">x1</div>
                  </li>
                </ul>
              </div>
              <div class="column amount">
                <p class="red">¥{{ order.payMoney?.toFixed(2) }}</p>
                <p>在线支付</p>
              </div>
              <div class="column action">
                <p><a href="javascript:;">查看详情</a></p>
              </div>
            </div>
          </div>
          <!-- 分页 -->
          <div class="pagination-container">
            <el-pagination :total="total" :page-size="params.pageSize" @current-change="pageChange" background layout="prev, pager, next" />
          </div>
        </div>
      </div>
  </div>

</template>

<style scoped lang="scss">
.order-container {
  padding: 10px 20px;

  .pagination-container {
    display: flex;
    justify-content: center;
  }

  .main-container {
    min-height: 500px;

    .holder-container {
      min-height: 500px;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }
}

.order-item {
  margin-bottom: 20px;
  border: 1px solid #f5f5f5;

  .head {
    height: 50px;
    line-height: 50px;
    background: #f5f5f5;
    padding: 0 20px;
    overflow: hidden;

    span {
      margin-right: 20px;

      &.down-time {
        margin-right: 0;
        float: right;

        i {
          vertical-align: middle;
          margin-right: 3px;
        }

        b {
          vertical-align: middle;
          font-weight: normal;
        }
      }
    }

    .del {
      margin-right: 0;
      float: right;
      color: #999;
    }
  }

  .body {
    display: flex;
    align-items: stretch;

    .column {
      border-left: 1px solid #f5f5f5;
      text-align: center;
      padding: 20px;

      >p {
        padding-top: 10px;
      }

      &:first-child {
        border-left: none;
      }

      &.goods {
        flex: 1;
        padding: 0;
        align-self: center;

        ul {
          li {
            border-bottom: 1px solid #f5f5f5;
            padding: 10px;
            display: flex;

            &:last-child {
              border-bottom: none;
            }

            .image {
              width: 70px;
              height: 70px;
              border: 1px solid #f5f5f5;
            }

            .info {
              width: 220px;
              text-align: left;
              padding: 0 10px;

              p {
                margin-bottom: 5px;

                &.name {
                  height: 38px;
                }

                &.attr {
                  color: #999;
                  font-size: 12px;

                  span {
                    margin-right: 5px;
                  }
                }
              }
            }

            .price {
              width: 100px;
            }

            .count {
              width: 80px;
            }
          }
        }
      }

      &.state {
        width: 120px;

        .green {
          color: $xtxColor;
        }
      }

      &.amount {
        width: 200px;

        .red {
          color: $priceColor;
        }
      }

      &.action {
        width: 140px;

        a {
          display: block;

          &:hover {
            color: $xtxColor;
          }
        }
      }
    }
  }
}
</style>