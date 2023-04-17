<template>
  <div class="index_view">
    <gvb-nav/>
<!--    <div class="gvb_banner">-->
<!--      <img src="https://featherwit-blog-img.oss-cn-shenzhen.aliyuncs.com/img/7c7dfa29baf2bd0f.jpg"/>-->
<!--    </div>-->
    <div class="gvb_base_container">
      <div class="gvb_inner_container">

      </div>
    </div>
    <div class="gvb_footer"></div>
  </div>
</template>

<script>
import axios from 'axios'
import config from '../../config/config'
import gvbNav from './components/gvbNav'
// import gvbBanner from "@/views/components/gvbBanner";

export default {
  components: {
    gvbNav,
    // gvbBanner
  },
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
    this.loginStatus = localStorage.getItem('loginStatus')
  },
}
</script>

<style lang="scss">
a {
  text-decoration: none;
}


.index_view {
  background-color: #f0eeee;

  .gvb_base_container {
    width: 1200px;
    display: flex;
    justify-content: center;
    .gvb_inner_container {
      background-color: white;
      min-height: 1000px;
      margin-top: 20px;
    }
  }

  .gvb_banner{
    height: 600px;
    width: 100%;
    background-color: darksalmon;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      display: block;
    }
  }
}

.user_show {
  cursor: pointer;

  &:hover {
    color: cornflowerblue;
  }
}
</style>