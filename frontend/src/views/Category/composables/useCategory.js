// 封装分类
import { getCategoryAPI } from '@/apis/category';
import { useRoute } from 'vue-router';
import { onBeforeRouteUpdate } from 'vue-router';
import { onMounted, ref } from 'vue';

export function useCategory() {
  const categoryData = ref({})
  const route = useRoute()
  const getCategory = async (id = route.params.id) => {
    // 如何在setup中获取路由参数 useRoute() -> route 等价于this.$route
    const res = await getCategoryAPI(id)
    categoryData.value = res.result
  }

  onMounted(() => getCategory())

  // 目标：路由参数变化的时候 可以把分类数据接口重新发送
  onBeforeRouteUpdate((to) => {
    getCategory(to.params.id)
  })

  return {
    categoryData
  }
}