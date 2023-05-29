<script setup>
import { getDetail } from '@/apis/detail'
import { useRoute } from 'vue-router'
import { ref, onMounted } from 'vue'
import {ElMessage} from 'element-plus'
import {useCartStore} from '@/stores/cartStore'
import imageView from '@/components/imageView/index.vue';

const goods = ref({})
const route = useRoute()
const getGoods = async () =>{
    const res = await getDetail(route.params.id)
    goods.value = res.result
}
onMounted(() => getGoods())

const cartStore = useCartStore()
const addCart = () =>{
  if(goods.value.forSale){
    //商品在售，可以加入购物车
    //console.log(goods.value.forSale)
    cartStore.addCart({
      id: goods.value.id,
      name: goods.value.name,
      price: goods.value.price,
      picture: goods.value.mainPictures[0],
      selected: true
    })
  }else{
    //商品已卖出
    ElMessage.warning('来晚了，商品已卖出')
    //console.log(goods.value.forSale)
  }
}
</script>

<template>
  <div class="xtx-goods-page">
    <div class="container" v-if="goods.category">
      <div class="bread-container">
        <el-breadcrumb separator=">">
          <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
          <el-breadcrumb-item :to="{ path: `/category/${goods.category.id}` }">{{goods.category.name}}
          </el-breadcrumb-item>          
          <el-breadcrumb-item>{{goods.name}}</el-breadcrumb-item>
        </el-breadcrumb>
      </div>
      <!-- 商品信息 -->
      <div class="info-container">
        <div>
          <div class="goods-info">
            <div class="media">
              <!-- 图片预览区 -->
              <imageView :imageList="goods.mainPictures"/>
            </div>
            <div class="spec">
              <!-- 商品信息区 -->
              <p class="g-name"> {{goods.name}} </p>
              <p class="g-desc">{{goods.desc}} </p>
              <p class="g-price">
                <span>{{goods.price}}</span>
              </p>
              <div class="g-service">
                <dl>                  
                  <dt>卖家</dt>                
                  <img :src = "goods.user.avator" alt="" /> 
                  <dd>{{goods.user.nickname}}</dd>
                </dl>
                <dl>
                  <dt>联系方式</dt>
                  <dd>{{goods.user.id}}</dd>
                </dl>
              </div>
              <!-- sku组件 -->

              <!-- 数据组件 -->

              <!-- 按钮组件 -->
              <div >
                <el-button size="large" class="btn" @click="addCart()">
                  加入购物车
                </el-button>
                <!-- <label class="g-label" v-show="!goods.forSale">
                  卖掉了
                </label> -->
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
// .g-label {
//   margin-top: 500px;
//   margin-left: 50px;
//     span {
//       font-size: 30px;
//   }
// }
.bread-container {
  padding: 25px 0;
}

</style>