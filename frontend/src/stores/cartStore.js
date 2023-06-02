// 封装购物车模块
import { defineStore } from "pinia"
import { ref, computed } from "vue"
import { useUserStore } from "./user"
import { insertCartAPI, findNewCartLisrAPI} from "@/apis/cart"
export const useCartStore = defineStore('cart', () =>{
    const userStore = useUserStore()
    const isLogin = computed(() => userStore.userInfo.token)
    //定义state - cartList
    const cartList = ref([])
    //定义action - addCart
    const addCart = async (goods) =>{
        if(isLogin.value){
            //登录之后添加购物车
            await insertCartAPI(goods.id)
            const request = await findNewCartLisrAPI()
            cartList.value = request.result
        }else
        {//未登录则添加到本地购物车
            //添加购物车操作
            const itemExisted = cartList.value.find((item) => goods.id === item.id)
            if(itemExisted){
                //已经添加，提示
                ElMessage('已经在购物车')
            }else{
                //还没有添加，push并提示
                cartList.value.push(goods)
                ElMessage.success('成功加入购物车')
            }
        }


    }
    const delCart = (goods) =>{
        const i = cartList.value.findIndex((item) => goods.id === item.id)
        cartList.value.splice(i, 1)
    }
    const singleCheck = (id, selected) =>{
        const item = cartList.value.find((item) => item.id === id)
        item.selected = selected
    }
    const allCheck = (selected) =>{
        cartList.value.forEach(item => item.selected = selected)
    }

    const allCount = computed(()=> cartList.value.reduce((a) => a + 1, 0))
    const allPrice = computed(()=> cartList.value.reduce((a,c) => (a*100 + c.price*100)/100, 0))
    const allSelected = computed(()=> cartList.value.every((item) =>item.selected))
    const selectedCount = computed(()=> cartList.value.filter(item => item.selected).reduce((a) => a+1, 0))
    const selectedPrice = computed(()=> cartList.value.filter(item => item.selected).reduce((a,c) => (a*100 + c.price*100)/100, 0))
    
    return{
        cartList,
        addCart,
        delCart,
        singleCheck,
        allCheck,
        allCount,
        allPrice,
        allSelected,
        selectedCount,
        selectedPrice
    }
}, {
    persist: true,
})