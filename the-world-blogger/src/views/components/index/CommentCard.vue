<template>
  <div class="me-view-comment-item">
    <div class="me-view-comment-author">
      <a class="">
        <img class="me-view-picture" :src="this.$props.comment.user.headshot">
      </a>
      <div class="me-view-info">
        <span class="me-view-nickname">{{ this.$props.comment.user.nickname }}</span>
        <div class="me-view-meta">
          <span>{{ this.$props.rootCommentCounts - this.$props.index }}楼</span>
          <span>{{ getFormatDate(this.$props.comment.comment.createAt) }}</span>
        </div>
      </div>
    </div>

    <div>
      <p class="me-view-comment-content">{{ this.$props.comment.comment.content }}</p>
      <div class="me-view-comment-tools">
        <a class="me-view-comment-tool"  @click="openComment" style="color: cornflowerblue">
          <i class="me-icon-comment"></i>&nbsp; 共{{this.$props.comment.childrenCount}}条评论
        </a>
        <a class="me-view-comment-tool"  @click="showComment">
          <i class="me-icon-comment"></i>&nbsp; 评论
        </a>
      </div>
      <div class="me-reply-list" v-if="open">
        <div class="me-reply-item" v-for="(c, index) in this.$props.comment.childrenComment" :key="index">
          <div style="font-size: 14px">
            <span class="me-reply-user">{{ c.user.nickname }}:&nbsp;&nbsp;</span>
            <span>{{ c.comment.content }}</span>
          </div>
          <div class="me-view-meta">
            <span style="padding-right: 10px">{{ getFormatDate(c.comment.createAt) }}</span>
          </div>

        </div>
      </div>
      <div class="me-view-comment-write" v-show="commentShow">
        <el-input
            v-model="reply.content"
            type="input"
            style="width: 90%"
            :placeholder="placeholder"
            class="me-view-comment-text"
            resize="none">
        </el-input>
        <el-button style="margin-left: 8px" @click="publishComment()" type="text">评论</el-button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Date from "@/utils/date";

export default {
  name: "CommentCard",

  props: {
    blogId: Number,
    comment: Object,
    index: Number,
    rootCommentCounts: Number,
  },
  data() {
    return {
      reply: this.getEmptyReply(),
      commentShow: false,
      open: false,
    }
  },
  methods: {
    openComment() {
      this.open = !this.open
    },
    getEmptyReply() {
      return {
        article: {
          id: this.blogId
        },
        parent: {
          id: this.comment.comment.id
        },
        content: ''
      }
    },
    showComment() {
      this.reply = this.getEmptyReply()
      this.commentShow = !this.commentShow
    },
    publishComment() {
      if (!localStorage.getItem("loginStatus")) {
        this.$message({type: 'warning', message: '未登录', showClose: true})
        this.$router.push('login')
        return
      }
      const data = {
        blogId: Number(this.blogId),
        parentId: Number(this.comment.comment.id),
        content: this.reply.content,
      }
      axios.post('/api/comment/add', data).then(res => {
        if (res.data.code === 0) {
          this.$message({type: 'success', message: '评论成功', showClose: true})
          this.reply = ''
          let newChildren = {}
          newChildren.user = res.data.data.user
          newChildren.comment = res.data.data.comment
          this.$props.comment.childrenComment.unshift(newChildren)
          this.$props.comment.childrenCount += 1
          this.reply = true
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
    getFormatDate(date) {
      return Date.getFormatDate(date)
    }
  },
}
</script>

<style lang="scss">

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

.me-view-comment-item {
  margin-top: 20px;
  padding-bottom: 16px;
  border-bottom: 1px solid #f0f0f0;
}

.me-view-comment-author {
  margin: 10px 0;
  vertical-align: middle;
}

.me-view-nickname {
  font-size: 14px;
}

.me-view-comment-content {
  line-height: 1.5;
}

.me-view-comment-tools {
  margin-top: 4px;
  margin-bottom: 10px;
}

.me-view-comment-tool {
  font-size: 13px;
  color: #a6a6a6;
  padding-right: 14px;
}

.v-note-wrapper .v-note-panel .v-note-show .v-show-content, .v-note-wrapper .v-note-panel .v-note-show .v-show-content-html {
  background: #fff !important;
}

.me-reply-list {
  padding-left: 16px;
  border-left: 4px solid #c5cac3;
}

.me-reply-item {
  margin-bottom: 8px;
  border-bottom: 1px solid #f0f0f0;
}

.me-reply-user {
  color: #78b6f7;
}
</style>