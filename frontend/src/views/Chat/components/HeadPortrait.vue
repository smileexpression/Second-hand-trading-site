<script setup>
import { getImageUrl } from '@/apis/image'
const props = defineProps({
  imgUrl: {
    type: String,
  }
})
// console.log("head", props.imgUrl)
</script >

<template>
  <div class="head-portrait">
    <img :src="getImageUrl(props.imgUrl)" alt="">
  </div>
</template>

<style lang="scss" scoped>
.head-portrait {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  // border: 2px solid rgb(137,140,151);
  border: 2px solid rgb(0, 0, 0);
  position: relative;

  // &::before {
  //   content: '';
  //   width: 15px;
  //   height: 15px;
  //   z-index: 1;
  //   display: block;
  //   border-radius: 50%;
  //   background-color: rgb(144, 225, 80);
  //   position: absolute;
  //   right: 0;
  // }

  img {
    width: 45px;
    height: 45px;
    border-radius: 50%;
    // padding: 2px;
    box-sizing: border-box;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    vertical-align: middle;
  }
}
</style>