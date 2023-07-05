<script setup>
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router';
import 'element-plus/es/components/message-box/style/index'
import { releaseAPI } from '@/apis/release'
import { getImageUrl, uploadImageAPI, deleteImageAPI } from '@/apis/image'

const uploadUrl = 'http://localhost:8080/image/upload'
const images = ref([])
const form = reactive({
  name: '',
  cate_id: '',
  description: '',
  price: '',
})

const beforeUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  if (!isImage) {
    ElMessage.error('只能上传图片文件')
  }
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    ElMessage.error('上传图片文件大小不能超过 2MB')
  }
  return isImage && isLt2M
}

const handleUploadSuccess = (response, file) => {
  console.log(response)
  images.value.push({ id: response.imageIds[0] })
}

// 上传图片
const uploadImages = async () => {
  const formData = new FormData()
  for (const file of $refs.upload.uploadFiles) {
    formData.append('images', file.raw)
  }
  const response = await uploadImageAPI(formData)
  if (response.ok) {
    // 获取返回的图片id
    const data = await response.json()
    data.imageIds.forEach((id) => {
      images.value.push({ id })
    })
  } else {
    ElMessage.error('上传图片失败')
  }
}

const router = useRouter()
const release = async () => {
  const imageIds = images.value.map((images) => images.id.toString());
  const res = await releaseAPI({
    name: form.name,
    cate_id: form.cate_id,
    description: form.description,
    price: form.price.toString(),
    picture: imageIds,
  })
  console.log(res)
  if (res.result == 'Succeed') {
    ElMessage.success('发布成功')
    router.replace({ path: '/' })
  } else {
    ElMessage.error('发布失败')
  }
}

const onSubmit = () => {
  if (form.name === '') {
    ElMessage.error('请输入名称')
    return
  } else if (form.cate_id === '') {
    ElMessage.error('请选择分类')
    return
  } else if (form.description === '') {
    ElMessage.error('请输入描述')
    return
  } else if (form.price === '') {
    ElMessage.error('请输入价格')
    return
  } else if (images.value.length < 1 || images.value.length > 5) {
    ElMessage.error('请上传1-5张图片')
    return
  }
  release()
}

const onCancel = () => {
  router.replace({ path: '/' })
}

const deleteImage = async (id) => {
  console.log("deleteImage", id)
  const res = await deleteImageAPI(id.toString())
  console.log(res)
}
const del = async (index, id) => {
  // await deleteImageAPI(index)
  await deleteImage(id)
  images.value.splice(index, 1)
  console.log(images.value)
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
            <el-input class="desc" v-model="form.description" type="textarea" placeholder="描述一下宝贝的品牌型号、货品来源……"
              :rows="10" />
          </el-form-item>
          <el-form-item label="分类">
            <el-select v-model="form.cate_id" placeholder="请选择宝贝的分类">
              <el-option label="手机严选" value="1" />
              <el-option label="保真奢品" value="2" />
              <el-option label="文玩珠宝" value="3" />
              <el-option label="限量潮品" value="4" />
            </el-select>
          </el-form-item>
          <!-- 这里的uploadUrl必须再上面声明 -->
          <el-upload ref="upload" :action="uploadUrl" :show-file-list="false" :on-success="handleUploadSuccess"
            :before-upload="beforeUpload">
            <el-button style="margin-left: 40px;">上传图片</el-button>
          </el-upload>
          <div class="demo-image">
            <div class="block" v-for="(image, index) in images" :key="image.id">
              <!-- 通过调用getImageUrl(参数为图片id)接口将图片url绑定到 :src -->
              <el-image style="width: 100px; height: 100px" :src="getImageUrl(image.id)" :fit="fill"
                @click="del(index, image.id)" />
              <!-- <h1></h1> -->
            </div>
          </div>
          <el-form-item label="价格">
            <!-- <el-input v-model="form.price" style="width: 200px;" /> -->
            <el-input-number v-model="form.price" :precision="2" :step="0.01" :min="0" style="width: 200px;" />
          </el-form-item>
          <el-form-item class="footer">
            <el-button type="primary" @click="onSubmit">发布</el-button>
            <el-button @click="onCancel">取消</el-button>
          </el-form-item>
        </el-form>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped lang="scss">
.common-layout {
  position: relative;
  width: 40%;
  margin: auto;

  .header {
    text-align: center;
    font-size: 50px;
    font-weight: bold;
  }

  .footer {
    button {
      position: relative;
      margin: auto;
      width: 20%;
    }
  }
}

.demo-image .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  display: inline-block;
  width: 33%;
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