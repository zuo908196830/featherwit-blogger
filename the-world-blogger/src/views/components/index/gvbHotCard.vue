<template>
  <gvb-card class="hot_card" title="热榜" link_name="查找更多" link="/hot" top20>
    <ui class="hot_ui">
      <li v-for="(blog, index) in this.hotBlogs" :key="index">
        <div class="hot_index">{{ index + 1 }}</div>
        <a @click="toHotBlog(blog.id)" href="">{{ blog.title }}</a>
      </li>
    </ui>
  </gvb-card>
</template>

<script>
import gvbCard from "@/views/components/gvbCard";
import axios from "axios";


export default {
  name: "gvbHotCard",
  components: {
    gvbCard
  },
  data() {
    return {
      hotBlogs: Array
    }
  },
  methods: {
    toHotBlog(blogId) {
      this.$router.push('/blog?id='+blogId)
    },
    getHotBlogs() {
      axios.get("/api/blog/search?order=2&limit=10&offset=0").then(res => {
        if(res.data.code === 0) {
          this.hotBlogs = res.data.data.blogs
        } else if (res.data.code === 1001) {
          this.$message({
            message: '无最热信息',
            type: 'warning'
          })
        }
      }).catch(() => {

      })
    }
  },
  created() {
    this.getHotBlogs()
  }
}
</script>

<style lang="scss">
.hot_ui {
  li {
    display: flex;
    flex-wrap: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    span {
      margin-right: 10px;
    }
  }

  .hot_index {
    width: 10%;
  }

  a {
    width: 70%;
    color: #2184fc;
  }
}
</style>