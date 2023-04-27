<template>
  <div class="md_view">
    <div>
      <gvb-nav></gvb-nav>
    </div>
    <div class="gvb_banner"></div>
    <div class="gvb_base_container">
      <div class="gvb_inner_container">
        <div style="position:absolute; left:7%; weight:86%">
          <div>
            <h2 style="text-align: center;">{{ blog.title }}</h2>
          </div>
          <div>
            <VueMarkDown :source="blog.content"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import VueMarkDown from 'vue-markdown'
import gvbNav from "@/views/components/gvbNav";

export default {
  components: {
    VueMarkDown,
    gvbNav
  },
  data() {
    return {
      id: 0,
      blog: {}
    }
  },
  methods: {
    getBlog() {
      axios.get('api/blog/id/' + this.id).then(res => {
        if (res.data.code === 0) {
          this.blog = res.data.data
        } else if (res.data.code === 1003) {
          this.$message({
            message: '文章不存在',
            type: 'error'
          })
        }
      }).catch(() => {
        this.$message({
          message: '服务器错误',
          type: 'error'
        })
      })
    }
  },
  mounted() {
    this.id = this.$route.query.id
    this.getBlog()
  }
}
</script>

<style lang="scss">

.md_view {
  img {
    max-width: 87%;
  }

  .gvb_banner {
    height: 50px;
    width: 100%;
    overflow: hidden;
  }

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