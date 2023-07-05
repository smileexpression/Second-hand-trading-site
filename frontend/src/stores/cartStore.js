// 封装购物车模块
import { defineStore } from "pinia"
import { ref, computed } from "vue"
import { useUserStore } from "./user"
import { insertCartAPI, findNewCartListAPI, deleteCartAPI} from "@/apis/cart"
import {ElMessage} from 'element-plus'

export const useCartStore = defineStore('cart', () =>{
    const userStore = useUserStore()
    const isLogin = computed(() => userStore.userInfo.token)
    //定义state - cartList
    const cartList = ref([])
    //定义action 
    const addCart = async (goods) =>{
        // console.log(userStore.userInfo.account)
        // console.log(goods.owner)
        if(isLogin.value && goods.owner!=userStore.userInfo.account){
            //登录之后添加购物车
            const itemExisted = cartList.value.find((item) => goods.id === item.id)
            if(itemExisted){
                ElMessage('已经在购物车')
            }else{
                await insertCartAPI(goods.id)
                updateCart()
                ElMessage.success('成功加入购物车')
                // console.log(goods.id)
            }
        }else if(!isLogin.value)
        {//未登录则 提示无法加入购物车
            ElMessage('请先登录')
        }
        else if(goods.owner == userStore.userInfo.account)
        {
            ElMessage('不能购买自己发布的商品！')
        }


    }
    const delCart = async(goodsIdList) =>{
        if(isLogin.value){
            //登录之后删除购物车
            await deleteCartAPI(goodsIdList)
            updateCart()
        }else
        {
            //理论上不会执行这里，不登录时购物车为空
            ElMessage('请先登录')            
            // const i = cartList.value.findIndex((item) => goods.id === item.id)
            // cartList.value.splice(i, 1)
        }

    }
    const updateCart = async() =>{
        const request = await findNewCartListAPI()
        if(request.result){            
            cartList.value = request.result
            // console.log(cartList.value)
        }else{
            cartList.value = []
        }
    }
    const clearCart = () =>{
        cartList.value =[]
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
        updateCart,
        clearCart,
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