<template>
  <div class="md_view">
    <gvb-nav/>
    <gvb-banner/>
    <div class="blog_view_container">
      <div class="blog_view_inner_container">
        <div style="">
          <div>
            <h2 style="font-size: 34px;font-weight: 700;line-height: 1.3;">{{ blog.blog.title }}</h2>
            <div class="me-view-author">
              <a class="">
                <img class="me-view-picture" :src="blog.user.headshot">
              </a>
              <div class="me-view-info">
                <span>{{ blog.user.nickname }}</span>
                <div class="me-view-meta">
                  <span>{{ getFormatDate(blog.blog.createAt) }}</span>
                  <span>阅读   {{ blog.blog.views }}</span>
                  <span>评论   {{ blog.blog.commentCount }}</span>
                </div>
                <el-button
                    v-if="blog.blog.username == this.$store.state.username"
                    @click="editArticle()"
                    style="position: absolute;left: 60%;"
                    size="mini"
                    round
                    icon="el-icon-edit">编辑
                </el-button>
              </div>
            </div>
          </div>
          <div>
            <VueMarkDown :source="blog.blog.content"/>
          </div>
        </div>
        <div class="me-view-comment">
          <div class="me-view-comment-write">
            <el-row :gutter="20">
              <el-col :span="2">
                <a class="">
                  <img class="me-view-picture" :src="headshot">
                </a>
              </el-col>
              <el-col :span="22">
                <el-input
                    type="textarea"
                    :autosize="{ minRows: 2}"
                    placeholder="你的评论..."
                    class="me-view-comment-text"
                    v-model="myComment"
                    resize="none">
                </el-input>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="2" :offset="22">
                <el-button type="text" @click="publishComment()">评论</el-button>
              </el-col>
            </el-row>
          </div>

          <div class="me-view-comment-title">
            <span>{{ blog.blog.commentCount }} 条评论</span>
          </div>

          <comment-card
              v-for="(comment, index) in comments"
              :key="comment.comment.id"
              :blog-id="id"
              :index="index"
              :root-comment-counts="commentCount"
              :comment="comment"
          ></comment-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import VueMarkDown from 'vue-markdown'
import gvbNav from "@/views/components/gvbNav";
import GvbBanner from "@/views/components/gvbBanner";
import CommentCard from "@/views/components/index/CommentCard";
import Date from '@/utils/date'

export default {
  components: {
    GvbBanner,
    VueMarkDown,
    gvbNav,
    CommentCard
  },
  data() {
    return {
      headshot: localStorage.getItem("headshot"),
      id: 0,
      blog: {},
      myComment: "",
      comments: [],
      commentCount: 0,
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
    },
    publishComment() {
      if (!localStorage.getItem("loginStatus")) {
        this.$message({type: 'warning', message: '未登录', showClose: true})
        this.$router.push('login')
        return
      }
      const data = {
        blogId: Number(this.id),
        parentId: -1,
        replyId: -1,
        content: this.myComment,
      }
      axios.post('/api/comment/add', data).then(res => {
        if (res.data.code === 0) {
          this.$message({type: 'success', message: '评论成功', showClose: true})
          this.myComment = ''
          this.commentCount++
          this.comments.unshift(res.data.data)
          this.blog.blog.commentCount++
        } else if (res.data.code === 1006) {
          this.$message({type: 'warning', message: '未登录', showClose: true})
          this.$router.push('login')
        } else {
          this.$message({type: 'error', message: '评论失败', showClose: true})
        }
      }).catch(error => {
        if (error !== 'error') {
          this.$message({type: 'error', message: '评论失败', showClose: true})
        }
      })
    },
    getComment() {
      axios.get('/api/comment/' + this.id + '/10/0').then(res => {
        if (res.data.code === 0) {
          this.commentCount = res.data.data.count
          this.comments = res.data.data.comments
        } else {
          this.$message({type: 'error', message: '评论加载失败', showClose: true})
        }
      }).catch(error => {
        if (error !== 'error') {
          this.$message({type: 'error', message: '评论加载失败', showClose: true})
        }
      })
    },
    getFormatDate(date) {
      return Date.getFormatDate(date)
    }
  },
  mounted() {
    this.id = this.$route.query.id
    this.getBlog()
    this.getComment()
  }
}
</script>

<style lang="scss">

.md_view {
  img {
    max-width: 87%;
  }

  a {
    color: #2b3539;
    cursor: pointer;

    &:hover {
      color: cornflowerblue;
    }
  }

  background-color: #f0eeee;

  .blog_view_container {
    top: 120px;
    display: flex;
    justify-content: center;
    height: 100%;
    background-color: white;

    .blog_view_inner_container {
      width: 1200px;
      margin-top: 20px;
    }
  }

  .me-view-author {
    /*margin: 30px 0;*/
    margin-top: 30px;
    vertical-align: middle;
  }

  .me-view-picture {
    width: 40px;
    height: 40px;
    border: 1px solid #ddd;
    border-radius: 50%;
    vertical-align: middle;
    background-color: #5fb878;
  }

  .me-view-info {
    display: inline-block;
    vertical-align: middle;
    margin-left: 8px;
  }

  .me-view-meta {
    font-size: 12px;
    color: #969696;
  }

  .me-view-comment {
    margin-top: 60px;
  }

  .me-view-comment-title {
    font-weight: 600;
    border-bottom: 1px solid #f0f0f0;
    padding-bottom: 20px;
  }

  .me-view-comment-write {
    margin-top: 20px;
  }

  .me-view-comment-text {
    font-size: 16px;
  }

  .v-show-content {
    padding: 8px 25px 15px 0px !important;
  }

  .v-note-wrapper .v-note-panel {
    box-shadow: none !important;
  }

  .v-note-wrapper .v-note-panel .v-note-show .v-show-content, .v-note-wrapper .v-note-panel .v-note-show .v-show-content-html {
    background: #fff !important;
  }


}

</style>