// 封装购物车模块
import { defineStore } from "pinia"
import { ref } from "vue"
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
    return{
        cartList,
        addCart
    }
}, {
    persist: true,
})