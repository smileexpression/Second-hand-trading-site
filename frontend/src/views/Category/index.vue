<script setup>
import GoodsItem from '../Home/components/GoodsItem.vue';
import { useBanner } from './composables/useBanner';
import { useCategory } from './composables/useCategory';

const { bannerList } = useBanner();
const { categoryData } = useCategory();
</script>

<template>
  <div class="container ">
    <!-- 面包屑 -->
    <div class="bread-container">
      <el-breadcrumb separator=">">
        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>{{ categoryData.name }}</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <!-- 轮播图 -->
    <div class="home-banner">
      <el-carousel height="500px">
        <el-carousel-item v-for="item in bannerList" :key="item.id">
          <img :src="item.imgUrl" alt="">
        </el-carousel-item>
      </el-carousel>
    </div>
    <div>
      <h3>- {{ categoryData.name }} -</h3>
    </div>
    <div class="sub-container">
      <!-- <el-tabs>
        <el-tab-pane label="推荐" name="orderNum"></el-tab-pane>
        <el-tab-pane label="最近发布" name="publishTime"></el-tab-pane>
      </el-tabs> -->
      <div class="body">
        <!-- 商品列表-->
        <GoodsItem v-for="goods in categoryData.goods" :goods="goods" :key="goods.id" />
      </div>
    </div>
  </div>
</template>


<style scoped lang="scss">
.bread-container {
  padding: 25px 0;
  color: #666;
}

h3 {
  font-size: 28px;
  color: #666;
  font-weight: normal;
  text-align: center;
  line-height: 100px;
}

.home-banner {
  width: 1240px;
  height: 500px;
  margin: 0 auto;

  img {
    width: 100%;
    height: 500px;
  }
}

.sub-container {
  padding: 20px 10px;
  background-color: #fff;

  .body {
    display: flex;
    flex-wrap: wrap;
    padding: 0 10px;
  }

  .goods-item {
    display: block;
    width: 220px;
    margin-right: 20px;
    padding: 20px 30px;
    text-align: center;

    img {
      width: 160px;
      height: 160px;
    }

    p {
      padding-top: 10px;
    }

    .name {
      font-size: 16px;
    }

    .desc {
      color: #999;
      height: 29px;
    }

    .price {
      color: $priceColor;
      font-size: 20px;
    }
  }

  .pagination-container {
    margin-top: 20px;
    display: flex;
    justify-content: center;
  }


}
</style>