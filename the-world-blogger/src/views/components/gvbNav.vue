<template>
  <div :class="{
    gvb_nav: true,
  }">
    <div class="gvb_nav_container">
      <div class="logo">
        <div>这是一个博客</div>
      </div>
      <div class="left">
        <span><a href="/" class="router-link">首页</a></span>
        <span><a href="/about" class="router-link">关于</a></span>
        <span><a href="/search" class="router-link">文章搜索</a></span>
        <span><a href="/blog/content" class="router-link">写文章</a></span>
      </div>
      <div class="right">
        <span><a href="/login" v-if="!loginStatus">登录</a></span>
        <img v-if="loginStatus" class="me-view-picture" :src="headshot">
        <el-dropdown v-if="loginStatus">
          <span class="el-dropdown-link" v-if="loginStatus">
            {{ username }}
            <i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>
              <a href="/user/data">个人详情</a>
            </el-dropdown-item>
            <el-dropdown-item>
              <a href="/blog/star">收藏夹</a>
            </el-dropdown-item>
            <el-dropdown-item>
              <a>历史记录</a>
            </el-dropdown-item>
            <el-dropdown-item>
              <a @click="logout">退出</a>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import config from "../../../config/config";

export default {
  name: "gvbNav",
  data() {
    return {
      drawer: false,
      username: localStorage.getItem("user"),
      loginStatus: localStorage.getItem("loginStatus"),
      searchName: "",
      headshot: localStorage.getItem("headshot"),
    }
  },
  methods: {
    logout() {
      axios.get("/api/user/logout").then(() => {
        localStorage.removeItem("user")
        axios.defaults.headers.common['Authorization'] = ""
        this.username = ""
        this.loginStatus = false
        this.drawer = false
        localStorage.setItem("loginStatus", false)
        if (this.$route.path === "/user/data") {
          this.$router.push('/')
        }
      })
    },
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
    } else {
      localStorage.setItem("loginStatus", false)
      localStorage.removeItem("user")
      axios.defaults.headers.common['Authorization'] = ""
      this.username = ""
      this.loginStatus = false
    }
    this.loginStatus = localStorage.getItem('loginStatus')
  },
}
</script>

<style lang="scss">
.el-dropdown-link {
  cursor: pointer;
  color: #2b3539;
}

.el-icon-arrow-down {
  font-size: 12px;
}

a {
  color: #2b3539;
}

.gvb_nav {
  box-shadow: 1px 1px 5px #0003;
  width: 100%;
  position: fixed;
  top: 0;
  height: 60px;
  display: flex;
  justify-content: center;
  font-size: 16px;
  z-index: 100;
  color: #2b3539;
  background-color: white;

  .me-view-picture {
    width: 40px;
    height: 40px;
    border: 1px solid #ddd;
    border-radius: 50%;
    vertical-align: middle;
    background-color: #5fb878;
  }

  .gvb_nav_container {
    width: 1200px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .logo {
    width: 20%;
    font-size: 23px;
  }

  .left {
    width: 40%;
    display: flex;

    span {
      margin-right: 30px;
    }
  }

  a {
    color: #2b3539;

    &:hover {
      color: cornflowerblue;
    }
  }

  .right {
    width: 40%;
    text-align: right;
  }
}
</style>