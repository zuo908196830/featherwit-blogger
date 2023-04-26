<template>
  <div class="base_view index_view">
    <gvb-nav/>
    <div class="gvb_banner">
    </div>
    <div class="gvb_base_container">
      <div class="gvb_inner_container gvb_index_main">
        <div class="left">
          <gvb-blog-list style="margin-top: 20px"></gvb-blog-list>
        </div>
        <div class="right">
          <gvb-tag-card :tag-nums="15"></gvb-tag-card>
          <gvb-hot-card style="margin-top: 20px"></gvb-hot-card>
        </div>
      </div>
    </div>
    <gvb-footer></gvb-footer>
  </div>
</template>

<script>
import axios from 'axios'
import gvbNav from './components/gvbNav'
import gvbTagCard from "@/views/components/index/gvbTagCard";
import gvbFooter from "@/views/components/gvbFooter";
import gvbHotCard from "@/views/components/index/gvbHotCard";
import gvbBlogList from "@/views/components/index/gvbBlogList";
import GvbBlogList from "@/views/components/index/gvbBlogList";

export default {
  components: {
    GvbBlogList,
    gvbNav,
    gvbTagCard,
    gvbFooter,
    gvbHotCard,
    gvbBlogList
  },
  data() {
    return {
    }
  },
  methods: {
  },
  created() {
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

li {
  list-style: none;
}

.base_view {
  background-color: #f0eeee;

  .gvb_base_container {
    top: 120px;
    display: flex;
    justify-content: center;

    .gvb_inner_container {
      width: 1200px;
      margin-top: 20px;
    }
  }

  .gvb_banner {
    height: 50px;
    width: 100%;
    overflow: hidden;
  }
}

.index_view {
  .gvb_inner_container {
    display: flex;
    justify-content: space-between;
  }

  .gvb_index_main {
    .left {
      width: calc(100% - 416px);
    }

    .right {
      width: 396px;
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