<template>
  <div class="blog_list_view">
    <div class="blog_list_body">
      <ul class="blog_list_ul">
        <li v-for="(blog, index) in this.blogs" :key="index">
          <div class="left">
            <img :src="blog.cover"/>
          </div>
          <div class="right">
            <h2>
              <a href="" class="" @click="toBlog(blog.id)">
                <div v-html="blog.title"></div>
              </a>
            </h2>
            <p>
              {{ blog.profile }}
            </p>
            <div class="blog_info">
              <span><i class="fa el-icon-pie-chart"></i>{{ getFormatDate(blog.updateAt) }}</span>
              <span><i class="fa el-icon-view"></i>{{ blog.views }}</span>
              <span class="netbook phone_550"><i class="fa el-icon-chat-round"></i>{{ blog.commentCount }}</span>
              <span><i class="fa el-icon-star-on"></i>{{ blog.likeCount }}</span>
            </div>
          </div>
        </li>
      </ul>
    </div>
    <div class="block">
      <el-pagination
          background
          layout="prev, pager, next"
          :total="total"
          :page-size="limit"
          :current-page.sync="page"
          @current-change="changePage"
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Date from '@/utils/date'

export default {
  name: "gvbBlogList",
  components: {},
  data() {
    return {
      blogs: Array,
      total: Number,
      limit: 8,
      offset: 0,
      page: 1,
      tagId: 0,
    }
  },
  props: {
    star: Boolean,
  },
  methods: {
    searchBlogs() {
      var url = ""
      if (this.$props.star) {
        url = "/api/blog/star?limit=" + this.limit + "&offset=" + this.offset
      } else {
        url = "/api/blog/search?order=1&limit=" + this.limit + "&offset=" + this.offset
        if (this.tagId !== 0) {
          url = url + "&tagId=" + this.tagId
        }
      }
      axios.get(url).then(res => {
        if (res.data.code === 0) {
          this.blogs = res.data.data.blogs
          this.total = res.data.data.total
        } else {
          this.$message({
            "message": "查询文章失败",
            "type": "warning"
          })
        }
      })
    },
    changePage() {
      this.offset = this.limit * (this.page - 1)
      this.searchBlogs()
    },
    toBlog(id) {
      this.$router.push('/blog?id=' + id)
    },
    getFormatDate(date) {
      return Date.getFormatDate(date)
    }
  },
  created() {
    this.searchBlogs()
  },

  watch: {
    '$store.state.tagId'(newTagId) {
      this.tagId = newTagId
      this.searchBlogs()
    }
  }
}
</script>

<style lang="scss">
.blog_list_body {
  .blog_list_ul {
    li {
      height: 200px;
      background-color: white;
      margin-bottom: 20px;
      display: flex;
      border-radius: 5px;
      transition: all 0.3s;

      img {
        width: 100%;
        height: 100%;
        transition: all 0.3s;
        object-fit: cover;
      }

      &:first-child {
        border-radius: 0 0 5px 5px;
        margin-top: 1px;
      }

      &:not(&:first-child):hover {
        transform: translateY(-10px);
        box-shadow: 0 0 10px #0000001a;
      }

      .left {
        width: 30%;
        padding: 20px 10px 20px 20px;

        > div {
          width: 100%;
          border-radius: 5px;
          height: 110px;
          overflow: hidden;
        }

        img {
          width: 100%;
          transition: all .3s;
          display: block;

          &:hover {
            transform: scale(1.05);
          }
        }
      }

      .right {
        width: 70%;
        padding: 20px 20px 20px 10px;
        display: flex;
        flex-direction: column;
        align-items: baseline;
        justify-content: space-between;

        h2 {
          font-size: 23px;
          font-weight: 600;

          a {
            color: #555;

          }
        }

        p {
          display: -webkit-box;
          -webkit-box-orient: vertical;
          -webkit-line-clamp: 3;
          overflow: hidden;
          text-overflow: ellipsis;
          margin-bottom: 10px;
          color: #555;
        }

        .blog_info {
          span {
            margin-right: 10px;
            display: inline-flex;
            align-items: center;

            i {
              margin-right: 5px;
              font-size: 17px;
            }
          }
        }
      }
    }
  }
}
</style>