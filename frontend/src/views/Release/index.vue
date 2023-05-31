<script setup>
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const UploadUrl = 'https://png.pngtree.com/png-vector/20191129/ourlarge/pngtree-image-upload-icon-photo-upload-icon-png-image_2047545.jpg'
const index = ref(0)

const open = () => {
  console.log("open")
  ElMessageBox.prompt('请输入图片url', '提示', {
    confirmButtonText: '完成',
    cancelButtonText: '取消',
    inputErrorMessage: 'Invalid Email',
  })
    .then(({ value }) => {
      index.value++
      form.picture.pop()
      form.picture.push(value)
      form.picture.push(UploadUrl)

      ElMessage({
        type: 'success',
        message: `Your email is:${value}`,
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Input canceled',
      })
    })
}
const del = (idx) => {
  console.log("del")
  form.picture.splice(idx, 1)
}

// do not use same name with ref
const form = reactive({
  name: '',
  cate_id: '',
  description: '',
  picture: [UploadUrl],
  price: '',
})

const onSubmit = () => {
  console.log('submit!')
}

</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">卖闲置</el-header>
      <el-main>
        <el-form :model="form">
          <el-form-item label="名称">
            <el-input v-model="form.name" />
          </el-form-item>
          <el-form-item label="描述">
            <el-input class="desc" v-model="form.desc" type="textarea" placeholder="描述一下宝贝的品牌型号、货品来源……" />
          </el-form-item>
          <el-form-item label="分类">
            <el-select v-model="form.region" placeholder="请选择宝贝的分类">
              <el-option label="手机严选" value="1" />
              <el-option label="保真奢品" value="beijing" />
            </el-select>
          </el-form-item>

          <div class="demo-image">
            <div v-for="(item, idx) in form.picture" class="block">
              <el-image v-if="idx < form.picture.length - 1" style="width: 100px; height: 100px" :src="item" :fit="fill"
                @click="del" />
              <el-image v-else-if="idx < 5" style="width: 100px; height: 100px" :src="item" :fit="fill" @click="open" />
            </div>
          </div>

          <el-form-item label="价格">
            <el-input v-model="form.price" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">发布</el-button>
            <el-button>取消</el-button>
          </el-form-item>
        </el-form>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped lang="scss">
.common-layout {
  position: relative;
  width: 60%;
  margin: auto;

  .header {
    text-align: center;
    font-size: 50px;
    font-weight: bold;
  }

  // .desc {
  //   height: 200;
  // }
}

.demo-image .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  display: inline-block;
  width: 20%;
  box-sizing: border-box;
  vertical-align: top;
}

.demo-image .block:last-child {
  border-right: none;
}

.demo-image .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  margin-bottom: 20px;
}
</style>