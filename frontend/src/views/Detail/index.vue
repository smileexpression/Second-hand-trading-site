<script setup>
import { getDetail } from '@/apis/detail'
import { useRoute, useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useCartStore } from '@/stores/cartStore'
import { getImageUrl } from '@/apis/image'
import imageView from '@/components/imageView/index.vue';
import { addChatAPI } from '@/apis/chat'

const goods = ref({})
const pictures = ref({})
const user = ref({})
const category = ["手机严选", "保真奢品", "文玩珠宝", "限量潮品"]
const route = useRoute()
const router = useRouter()
const getGoods = async () => {
  const res = await getDetail(route.params.id)
  goods.value = res.result
  pictures.value = res.pictures
  user.value = res.user
  // console.log(pictures.value)
  // console.log(goods.value)
}
onMounted(() => getGoods())

const cartStore = useCartStore()
const addCart = () => {
  if (!goods.value.forsale) {
    //商品在售，可以加入购物车
    //console.log(goods.value.forSale)
    cartStore.addCart({
      id: String(goods.value.ID),
      name: goods.value.name,
      price: Number(goods.value.price),
      picture: goods.value.picture,
      selected: true
    })
  } else {
    //商品已卖出
    ElMessage.warning('来晚了，商品已卖出')
    //console.log(goods.value.forSale)
  }
}

// 聊天
const addChat = async () => {
  // console.log('aaaa', user.value.ID);
  const res = await addChatAPI({
    you: `${user.value.ID}`,
  })
  if (res.result === "bug") {
    ElMessage.warning('不能和自己聊天哦')
    return
  } else {
    ElMessage.success('已添加到聊天列表')
    router.replace({ path: '/chat' })
  }
  // console.log(res);
}

const onAddChat = () => {
  addChat()
  // router.replace('chat')
}
</script>

<template>
  <div class="xtx-goods-page">
    <div class="container" v-if="goods.ID">
      <div class="bread-container">
        <el-breadcrumb separator=">">
          <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
          <el-breadcrumb-item :to="{ path: `/category/${goods.Cate_Id}` }">{{ category[goods.Cate_Id - 1] }}
          </el-breadcrumb-item>
          <el-breadcrumb-item>{{ goods.name }}</el-breadcrumb-item>
        </el-breadcrumb>
      </div>
      <!-- 商品信息 -->
      <div class="info-container">
        <div>
          <div class="goods-info">
            <div class="media">
              <!-- 图片预览区 -->
              <imageView :imageList="pictures" />
            </div>
            <div class="spec">
              <!-- 商品信息区 -->
              <p class="g-name"> {{ goods.name }} </p>
              <p class="g-desc">{{ goods.desc }} </p>
              <p class="g-price">
                <span>{{ Number(goods.price).toFixed(2) }}</span>
              </p>
              <div class="g-service">
                <dl>
                  <dt>卖家</dt>
                  <img :src="getImageUrl(user.avatar)" alt="" />
                  <dd @click="onAddChat">{{ user.nickname }}</dd>
                </dl>
                <dl>
                  <dt>联系方式</dt>
                  <dd @click="onAddChat">{{ user.account }}</dd>
                  <el-button size="middle" round @click="onAddChat()">
                  加入聊天列表
                  </el-button>
                </dl>

              </div>
              <!-- 按钮组件 -->
              <div>
                <el-button type="primary" size="large" class="btn" @click="addCart()">
                  加入购物车
                </el-button>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>


<style scoped lang='scss'>
.xtx-goods-page {
  .goods-info {
    min-height: 600px;
    background: #fff;
    display: flex;

    .media {
      width: 580px;
      height: 600px;
      padding: 30px 50px;
    }

    .spec {
      flex: 1;
      padding: 30px 30px 30px 0;
    }
  }

  .g-name {
    font-size: 22px;
  }

  .g-desc {
    color: #999;
    margin-top: 10px;
  }

  .g-price {
    margin-top: 10px;

    span {
      &::before {
        content: "¥";
        font-size: 14px;
      }

      &:first-child {
        color: $priceColor;
        margin-right: 10px;
        font-size: 22px;
      }

      // &:last-child {
      //   color: #999;
      //   text-decoration: line-through;
      //   font-size: 16px;
      // }
    }
  }

  .g-service {
    background: #f5f5f5;
    width: 500px;
    padding: 20px 10px 0 10px;
    margin-top: 10px;

    dl {
      padding-bottom: 20px;
      display: flex;
      align-items: center;

      dt {
        width: 80px;
        color: #999;
      }

      img {
        width: 50px;
        height: 50px;
        border-radius: 100%;
      }

      dd {
        margin-left: 10px;
        color: #666;

        &:last-child {
          span {
            margin-right: 10px;

            &::before {
              content: "•";
              color: $xtxColor;
              margin-right: 2px;
            }
          }

          a {
            color: $xtxColor;
          }
        }
      }
    }
  }

}

.btn {
  margin-top: 20px;
}

.bread-container {
  padding: 25px 0;
}
</style>