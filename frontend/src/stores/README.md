# stores

stores文件夹是一个用于存放前端状态管理的模块，包括Vuex、Redux、Pinia等状态管理工具。

## Pinia

Pinia是一个用于管理Vue.js应用程序状态的库，可以通过Pinia来实现状态的管理和共享等功能。

### 使用说明

1. 将Pinia文件复制到项目的相应位置。
2. 在Vue实例中注册Pinia，并配置Pinia Store。
3. 在需要使用Store的组件中，使用相应的Store实例。

### 示例

以下是一个Pinia Store的示例，用于实现购物车、分类、计数器、用户等状态的管理和共享：

javascript

Copy

```
// stores/cartStore.js
import { defineStore } from 'pinia';

export const useCartStore = defineStore('cart', {
  state: () => ({
    cartItems: []
  }),
  actions: {
    addToCart(item) {
      this.cartItems.push(item);
    }
  },
  getters: {
    cartItemsCount() {
      return this.cartItems.length;
    }
  }
});

// stores/category.js
import { defineStore } from 'pinia';

export const useCategoryStore = defineStore('category', {
  state: () => ({
    categories: []
  }),
  actions: {
    async fetchCategories() {
      const categories = await api.getCategories();
      this.categories = categories;
    }
  }
});

// stores/counter.js
import { defineStore } from 'pinia';

export const useCounterStore = defineStore('counter', {
  state: () => ({
    count: 0
  }),
  actions: {
    increment() {
      this.count++;
    }
  },
  getters: {
    doubleCount() {
      return this.count * 2;
    }
  }
});

// stores/user.js
import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null
  }),
  actions: {
    async login({ username, password }) {
      const user = await api.login(username, password);
      this.user = user;
    }
  },
  getters: {
    isLoggedIn() {
      return this.user !== null;
    }
  }
});

// 在Vue实例中注册Pinia Store
import { createPinia } from 'pinia';
import { useCartStore } from './stores/cartStore';
import { useCategoryStore } from './stores/categoryStore';
import { useCounterStore } from './stores/counterStore';
import { useUserStore } from './stores/userStore';

const pinia = createPinia();

pinia.useStore(useCartStore);
pinia.useStore(useCategoryStore);
pinia.useStore(useCounterStore);
pinia.useStore(useUserStore);

// 在模板中使用Pinia Store
<template>
  <div>{{ $store.counter.count }}</div>
  <div>{{ $store.cart.cartItemsCount }}</div>
  <button @click="$store.counter.increment()">增加计数器</button>
  <button @click="$store.cart.addToCart({ id: 1, name: '商品1' })">添加购物车</button>
</template>
```

## 注意事项

1. 在使用状态管理工具时，需要仔细阅读相应的文档，了解工具的使用方法和注意事项。
2. 在使用状态管理工具时，需要根据应用程序的需求和性能要求进行相应的设计和实现，以确保应用程序的正常运行和性能。
3. 在使用状态管理工具时，需要进行相应的错误处理和安全性检查，以确保应用程序的正常运行和安全性。
