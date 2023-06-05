<script setup>
import Nav from '@/views/Chat/components/Nav.vue'
import chathome from '@/views/Chat/components/chathome.vue'
</script>

<template>
  <div class="home">
    <el-container height="100%">
      <el-aside width="100px">
        <Nav></Nav>
      </el-aside>
      <el-main>
        <!-- <router-view></router-view> -->
        <chathome></chathome>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped lang="scss">
.home {
  display: flex;
  padding-top: 20px;
  // width: 90vw;
  // height: 90vh;
  background-color: rgb(39, 42, 55);
  border-radius: 15px;
  // position: relative;
  // left: 50%;
  // top: 50%;
  // transform: translate(-50%, -50%);
}
</style>