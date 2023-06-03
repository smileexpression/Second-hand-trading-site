<script setup>
import HeadPortrait from '@/views/Chat/components/HeadPortrait.vue'
import { useUserStore } from '@/stores/user'
const userStore = useUserStore()
</script>

<template>
  <div class="nav">
    <div class="own-pic">
      <HeadPortrait :imgUrl="userStore.userInfo?.avatar"></HeadPortrait>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.nav {
  width: 100%;
  height: 90vh;
  position: relative;
  border-radius: 20px 0 0 20px;

  .own-pic {
    position: absolute;
    bottom: 90%;
    margin-left: 25px;
  }
}
</style>