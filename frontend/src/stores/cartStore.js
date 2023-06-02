// 封装购物车模块
import { defineStore } from "pinia"
import { ref, computed } from "vue"
export const useCartStore = defineStore('cart', () =>{
    //定义state - cartList
    const cartList = ref([])
    //定义action - addCart
    const addCart = (goods) =>{
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
    const delCart = (goods) =>{
        const i = cartList.value.findIndex((item) => goods.id === item.id)
        cartList.value.splice(i, 1)
    }
    const allCount = computed(()=> cartList.value.reduce((a,c) => a + 1, 0))
    const allPrice = computed(()=> cartList.value.reduce((a,c) => (a*100 + c.price*100)/100, 0))

    return{
        cartList,
        addCart,
        delCart,
        allCount,
        allPrice
    }
}, {
    persist: true,
})