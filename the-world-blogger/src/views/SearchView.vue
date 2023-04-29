<template>
  <div class="base_view">
    <gvb-nav/>
    <gvb-banner/>
    <div class="gvb_base_container">
      <div class="gvb_inner_container gvb_index_main">
        <div class="gvb_search_head">
          <div class="gvb_search_slogn">
            搜索文章
          </div>
          <div class="gvb_search_ipt">
            <el-input v-model="name" placeholder="请输入搜索内容"></el-input>
          </div>
          <div class="gvb_search_btn">
            <el-button type="primary" @click="search">搜索</el-button>
          </div>
        </div>
        <div class="gvb_search_action">
          <div class="gvb_search_order">
            <span :class="{active: order === 1}" @click="checkOrder(1)">最近发布</span>
            <span :class="{active: order === 2}" @click="checkOrder(2)">最多浏览</span>
            <span :class="{active: order === 3}" @click="checkOrder(3)">最多评论</span>
            <span :class="{active: order === 4}" @click="checkOrder(4)">最多收藏</span>
          </div>
          <div class="gvb_search_tags">
            <span :class="{active: 0 === tagId}" @click="checkTag(0)">全部标签</span>
            <span :class="{active: tag.id === tagId}" @click="checkTag(tag.id)" v-for="(tag, index) in tags" :key="index">{{ tag.name }}</span>
          </div>
        </div>
        <div class="gvb_search_result">
          <div class="gvb_search_result_item"
               style="cursor: pointer;"
               @click="toBlog(blog.id)"
               v-for="(blog, index) in blogs" :key="index">
            <a class="img">
              <img :src="blog.cover">
            </a>
            <div class="search_result_info">
              <div class="search_result_title">
                <a>{{ blog.title }}</a>
              </div>
            </div>
            <div class="search_result_icon">
              <span><i class="el-icon-view"></i>{{ blog.views }}</span>
              <span><i class="el-icon-pie-chart"></i>{{ getFormatDate(blog.createAt) }}</span>
            </div>
          </div>
        </div>
        <div class="gvb_search_page">
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
    </div>
    <gvb-footer></gvb-footer>
  </div>
</template>

<script>
import gvbNav from "@/views/components/gvbNav";
import gvbFooter from "@/views/components/gvbFooter";
import axios from "axios";
import Date from '@/utils/date'
import GvbBanner from "@/views/components/gvbBanner";

export default {
  name: "SearchView",
  components: {
    GvbBanner,
    gvbNav,
    gvbFooter,
  },
  data() {
    return {
      name: "",
      searchName: "",
      order: 1,
      tags: [],
      tagId: 0,
      total: 0,
      page: 1,
      limit: 15,
      offset: 0,
      blogs: [],
    }
  },
  methods: {
    checkOrder(order) {
      this.order = order
      this.searchBlogs()
    },
    checkTag(tagId){
      this.tagId = tagId
      this.searchBlogs()
    },
    searchTag() {
      axios.get('/api/tag/search').then(res => {
        if (res.data.code === 0) {
          this.tags = res.data.data.tag
        } else {
          this.$message({
            message: '获取标签失败',
            type: 'warning'
          })
        }
      })
    },
    searchBlogs() {
      var url = "/api/blog/search?limit=" + this.limit + "&offset=" + this.offset + "&order=" + this.order
      if (this.tagId !== 0) {
        url = url + "&tagId=" + this.tagId
      }
      if (this.searchName !== "") {
        url = url + "&name=" + this.searchName
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
    search() {
      this.searchName = this.name
      this.searchBlogs()
    },
    changePage() {
      this.offset = this.limit * (this.page - 1)
      this.searchBlogs()
    },
    toBlog(blogId) {
      this.$router.push('/blog?id=' + blogId)
    },
    getFormatDate(date) {
      return Date.getSimpleDate(date)
    },
  },
  created() {
    this.searchTag()
    this.searchBlogs()
  }
}
</script>

<style lang="scss">
.base_view {
  .gvb_search_head {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 20px;

    .gvb_search_slogn {
      width: 10%;
      font-size: 16px;
    }

    .gvb_search_ipt {
      width: 50%;
    }

    .gvb_search_btn {
      width: 10%;
    }
  }

  .gvb_search_action {
    border-top: 1px solid #e2e2e2;
    padding: 20px 0;
    border-bottom: 1px solid #e2e2e2;
    span {
      padding: 3px 6px;
      color: #555;
      cursor: pointer;
      font-size: 14px;
      margin-right: 20px ;
    }
    span.active {
      background-color: #2184fc;
      color: white;
      border: 5px;
    }
    .gvb_search_order {
      margin-bottom: 10px;
    }
  }

  .gvb_search_result {
    display: grid;
    grid-template-columns: repeat(5, 1fr);
    grid-column-gap: 20px;
    grid-row-gap: 20px;
    margin-bottom: 20px;
    margin-top: 20px;

    .gvb_search_result_item {
      border: 1px solid #e2e2e2;
      border-radius: 5px;
      overflow: hidden;
      margin: 0 30px 30px 0;
      transition: all 0.3s;
      &:hover{
        transform: translateY(-10px);
        .img img {
          transform: scale(1.1);
        }

        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
      }

      >a {
        display: block;
        width: 100%;
        height: 100px;
        overflow: hidden;

        img {
          width: 100%;
          height: 100%;
          transition: all 0.3s;
          object-fit: cover;
        }
      }

      .search_result_info {
        padding: 8px;
        color: #777;

        .search_result_title {
          margin-bottom: 5px;
          height: 43px;
        }
      }
      .search_result_icon {
        font-size: 13px;
        display: flex;
        justify-content: space-between;
      }
    }
  }
  .gvb_search_page {
    justify-content: space-between;
  }
}
</style>