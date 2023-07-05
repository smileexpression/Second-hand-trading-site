<script setup>
import { useUserStore } from '@/stores/user';
import { useRouter } from 'vue-router';

//获取userStore实例
const userStore = useUserStore()

const router = useRouter()
//执行退出登录功能
const logoutConfirm = () => {
  userStore.clearUserInfo()
  //跳转回登录页面
  router.push('/login')
}

</script>

<template>
  <nav class="app-topnav">
    <div class="container">
      <ul>
        <!-- 通过userStore是否携带token选择模板进行渲染 -->
        <template v-if="userStore.userInfo.token">
          <li><a href="javascript:;"><i class="iconfont icon-user"></i>{{ userStore.userInfo.nickname }}</a></li>
          <li>
            <el-popconfirm title="确认退出吗?" @confirm="logoutConfirm" confirm-button-text="确认" cancel-button-text="取消">
              <template #reference>
                <a href="javascript:;">退出登录</a>
              </template>
            </el-popconfirm>
          </li>
          <li><a href="javascript:;" @click="$router.push('/member/order')">我的订单</a></li>
          <li><a href="javascript:;" @click="$router.push('/member')">会员中心</a></li>
          <li>
            <RouterLink to="/release">发布闲置</RouterLink>
          </li>
          <li>
            <RouterLink to="/chat">我的消息</RouterLink>
          </li>
        </template>
        <template v-else>
          <li><a href="javascript:;" @click="$router.push('/login')">请先登录</a></li>
          <!-- <li><a href="javascript:;">帮助中心</a></li>
          <li><a href="javascript:;">关于我们</a></li> -->
        </template>
      </ul>
    </div>
  </nav>
</template>


<style scoped lang="scss">
.app-topnav {
  background: #333;

  ul {
    display: flex;
    height: 53px;
    justify-content: flex-end;
    align-items: center;

    li {
      a {
        padding: 0 15px;
        color: #cdcdcd;
        line-height: 1;
        display: inline-block;

        i {
          font-size: 14px;
          margin-right: 2px;
        }

        &:hover {
          color: $xtxColor;
        }
      }

      ~li {
        a {
          border-left: 2px solid #666;
        }
      }
    }
  }
}
</style>