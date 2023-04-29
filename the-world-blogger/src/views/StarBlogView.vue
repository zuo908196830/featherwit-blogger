<template>
  <div class="star_blog_view">
    <gvb-nav/>
    <gvb-banner/>
    <div style="text-align: center; color: #555"><h2>收藏夹</h2></div>
    <div class="gvb_base_container">
      <div class="gvb_inner_container gvb_index_main">
        <gvb-blog-list star/>
      </div>
    </div>
  </div>
</template>

<script>
import gvbNav from "@/views/components/gvbNav";
import gvbBanner from "@/views/components/gvbBanner";
import axios from "axios";
import gvbBlogList from "@/views/components/index/gvbBlogList";

export default {
  name: "StarBlogView",
  components: {
    gvbNav,
    gvbBanner,
    gvbBlogList,
  },
  mounted() {
    if (localStorage.getItem('loginStatus') === 'false') {
      this.$message({
        message: '请先登录',
        type: 'warning'
      });
      this.$router.push('/login')
    }
    axios.defaults.headers.common['Authorization'] = localStorage.getItem("token")
  },
}
</script>

<style lang="scss">
.star_blog_view {
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
}
</style>