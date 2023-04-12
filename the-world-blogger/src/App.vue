<template>
  <div id="app">
    <el-header style="height:40px">
      <div>
        <el-input v-model="searchName" placeholder="搜索文章"
          style="position: absolute; left: 25%; width: 40%"></el-input>
        <el-button type="primary" icon="el-icon-search" @click=search
          style="position: absolute; left: 65.1%;">搜索</el-button>
      </div>
      <div style="position: absolute; top: 1px; right: 1px;">
        <nav>
          <router-link to="/login" v-if="!loginStatus" class="user_show">登录/注册</router-link>

          <!-- 登录后的用户展示界面 -->
          <el-button v-if="loginStatus" @click="drawer = true" type="primary" style="margin-left: 16px;"
            class="user_show">
            {{ username }}
          </el-button>

          <el-drawer title="我是标题" :visible.sync="drawer" :with-header="false">
            <div class="head_photo">
              <div class="block">
                <span class="demonstration"></span>
                <el-image>
                  <div slot="error" class="image-slot">
                    <i class="el-icon-picture-outline"></i>
                  </div>
                </el-image>
              </div>
            </div>
            <div>
              <router-link to="/user/data">
                <el-button style="width:100%; height:40px">个人详情</el-button>
              </router-link>
            </div>
            <div>
              <el-button style="width:100%; height:40px">收藏夹</el-button>
            </div>
            <div>
              <el-button style="width:100%; height:40px">历史记录</el-button>
            </div>
            <div>
              <el-button @click="logout" style="width:100%; height:40px">退出</el-button>
            </div>
          </el-drawer>

        </nav>
      </div>
    </el-header>
    <router-view />
  </div>
</template>

<script>
import axios from 'axios'
import config from '../config/config'

export default {
  data() {
    return {
      drawer: false,
      username: localStorage.getItem("user"),
      loginStatus: localStorage.getItem("loginStatus"),
      searchName: ""
    }
  },
  methods: {
    logout() {
      axios.get("/api/user/logout").then(res => {
        if (res.data.code === 0) {
          localStorage.removeItem("user")
          axios.defaults.headers.common['Authorization'] = ""
          this.username = ""
          this.loginStatus = false
          this.drawer = false
          localStorage.setItem("loginStatus", false)
          if (this.$route.path === "/user/data") {
            this.$router.push('/')
          }
        }
      })
    },
    search() {
    },
    // getLoginStatus() {
    //   axios.get("/api/user/status").then(res => {
    //     if (!res.data.data) {
    //       localStorage.setItem("loginStatus", false)
    //       localStorage.removeItem("user")
    //       axios.defaults.headers.common['Authorization'] = ""
    //       this.username = ""
    //       this.loginStatus = false
    //     }
    //   })
    // }
  },
  watch: {
    "$route.path": function () {      //监视每次router的变化
      this.username = localStorage.getItem("user")
      let lst = localStorage.getItem("loginStatus")
      if (lst === 'true') {
        this.loginStatus = true
      } else {
        this.loginStatus = false
      }
    }
  },
  created() {
    axios.defaults.baseURL = config.host
    if (localStorage.getItem("token")) {
      axios.defaults.headers.common['Authorization'] = localStorage.getItem("token")
      axios.get("/api/user/token/login").then(res => {
        if (res.data.code === 0) {
          if (res.data.data.nickname) {
            this.username = res.data.data.nickname
            localStorage.setItem("user", res.data.data.nickname)
          } else {
            this.username = res.data.data.username
            localStorage.setItem("user", res.data.data.username)
          }
          this.loginStatus = true
          localStorage.setItem("loginStatus", true)
        } else {
          localStorage.setItem("loginStatus", false)
          localStorage.removeItem("user")
          this.loginStatus = false
          this.username = false
          axios.defaults.headers.common['Authorization'] = ""
        }

      }).catch(() => {
        localStorage.setItem("loginStatus", false)
        localStorage.removeItem("user")
        this.loginStatus = false
        this.username = false
        axios.defaults.headers.common['Authorization'] = ""
      })
      // this.loginStatus()
    } else {
      localStorage.setItem("loginStatus", false)
      localStorage.removeItem("user")
      axios.defaults.headers.common['Authorization'] = ""
      this.username = ""
      this.loginStatus = false
    }
  },
}
</script>

<style>
.user_show {
  position: absolute;
  right: 0%;
  width: 100px;
  height: 40px;
}

.head_photo {
  height: 20%;
}

.block {
  width: 60%;
  height: 70%;
}

img {
  max-width: 86%;
}
</style>
